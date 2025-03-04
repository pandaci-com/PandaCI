import { Conclusion, WorkflowAlert_Type } from "@pandaci/proto";
import type { JobOptions, JobResult } from "./jobTypes.ts";
import { setImmediate } from "node:timers";
import { jobContext } from "../context.ts";
import { extendedEnv } from "../env.ts";
import { logger } from "../logger.ts";
import type { JobFn } from "./mod.ts";
import type { getWorkflowClient } from "../api.ts";

export class JobError extends Error {
  conclusion: Conclusion;
  isFailure: boolean;
  isSkipped: boolean;
  isSuccess: boolean;
  id: string;
  jobName: string;
  runner?: string;

  constructor(public result: JobResult) {
    super(`Job ${result.name} failed`);
    this.conclusion = result.conclusion;
    this.isFailure = result.isFailure;
    this.id = result.id;
    this.jobName = result.name;
    this.runner = result.runner;
    this.isSkipped = result.isSkipped;
    this.isSuccess = result.isSuccess;
  }
}

interface JobFunctionContext {
  getWorkflowClient: () => ReturnType<typeof getWorkflowClient>;
}

export class JobPromise extends Promise<JobResult> {
  private reject: (reason?: JobResult) => void;
  private resolve: (value: JobResult) => void;

  private name: string;
  private runner: string;

  private throws?: boolean;
  private skip?: boolean;

  constructor(
    ctx: JobFunctionContext,
    name: string,
    fn: (() => void) | (() => Promise<void>),
    options?: JobOptions,
  ) {
    let tempReject: (reason?: JobResult) => void = () => {};
    let tempResolve: (value: JobResult) => void = () => {};
    super((resolve, reject) => {
      tempReject = reject;
      tempResolve = resolve;

      // internally, .then() calls the executor function and requires it to resolve or reject.
      // We don't include this in our type definition because it'll just confuse users.
      if (typeof ctx === "function") (ctx as any)(resolve, reject);
    });

    this.reject = tempReject;
    this.resolve = tempResolve;

    this.name = name;
    this.runner = options?.runner ?? "ubuntu-4x";
    this.throws = options?.throws ?? true;
    this.skip = options?.skip ?? false;

    // If this is an internal .then() call then calling run will just
    if (typeof ctx === "function") return;

    // TODO - this can cause some issues since not all errors are JobError
    setImmediate(() => this.run(ctx, fn).catch(this.reject));
  }

  private async run(
    ctx: JobFunctionContext,
    fn: JobFn,
  ) {
    logger.info(`Starting job ${this.name}`);

    if (jobContext.getStore()) {
      await ctx.getWorkflowClient().createWorkflowAlert({
        workflowJwt: extendedEnv.PANDACI_WORKFLOW_JWT,
        alert: {
          type: WorkflowAlert_Type.ERROR,
          title: "Syntax error - Nested job detected",
          message: "Jobs cannot be nested inside other jobs or tasks",
        },
      });

      throw new Error("nested job");
    }

    const client = ctx.getWorkflowClient();

    const createJobResult = await client.startJob({
      name: this.name,
      runner: this.runner,
      workflowJwt: extendedEnv.PANDACI_WORKFLOW_JWT,
      skipped: this.skip,
    });

    if (this.skip) {
      logger.info(`Job ${this.name} skipped`);
      return this.resolve({
        conclusion: Conclusion.SKIPPED,
        id: createJobResult.jobMeta!.id,
        name: createJobResult.jobMeta!.name,
        isFailure: false,
        runner: createJobResult.jobMeta!.runner,
        isSkipped: true,
        isSuccess: false,
      });
    }

    const taskPromises: Promise<unknown>[] = [];

    const conclusion = await jobContext
      .run(
        {
          meta: createJobResult.jobMeta!,
          registerJobPromise: (promise) => {
            taskPromises.push(promise);
          },
        },
        async () => {
          try {
            await Promise.resolve(fn());
            await Promise.allSettled(taskPromises);
          } catch (err) {
            logger.error(`Job ${this.name} failed`, err);
            return Conclusion.FAILURE;
          }
          return Conclusion.UNSPECIFIED;
        },
      );

    const jobRes = await client.stopJob({
      $typeName: "proto.v1.WorkflowServiceStopJobRequest",
      jobMeta: createJobResult.jobMeta!,
      workflowJwt: extendedEnv.PANDACI_WORKFLOW_JWT,
      conclusion,
    });

    logger.info(`Job ${this.name} completed`);

    const res: JobResult = {
      conclusion: jobRes.conclusion,
      id: createJobResult.jobMeta!.id,
      name: this.name,
      runner: createJobResult.jobMeta!.runner,
      isFailure: jobRes.conclusion === Conclusion.FAILURE,
      isSkipped: jobRes.conclusion === Conclusion.SKIPPED,
      isSuccess: jobRes.conclusion === Conclusion.FAILURE,
    };

    if (this.throws && res.isFailure) {
      this.reject(new JobError(res));
    }

    this.resolve(res);
  }
}
