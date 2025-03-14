export * from "./docker/mod.ts";
export * from "./job/mod.ts";
export { type Env, env } from "./env.ts";
export * from "./shell/mod.ts";
export { type Volume, volume } from "./volume.ts";
export type { ExecOptions, ExecResult } from "./shell/mod.ts";
export type { Conclusion } from "./types.ts";

/**
 * Pull request event type
 */
export type PullRequestEvent = "opened" | "closed" | "reopened" | "synchronize";

/**
 * Workflow config
 * @example export const config: Config = { name: "My Job", on: { push: { branches: ["main"] } } }
 */
export interface Config {
  /**
   * Job name, defaults to the filename and path.
   */
  name?: string;
  /**
   * Trigger configuration, defaults to running on every push.
   *
   * @default { push: { branches: ["*"] } }
   */
  on?: {
    /**
     * Push configuration, by default all branches are triggered.
     * If you want to trigger on specific branches, you need to specify the branches.
     */
    push?: {
      /**
       * Branches to trigger on, defaults to all branches.
       * You can use glob patterns to match multiple branches.
       * @example ["main", "feature/*"]
       */
      branches?: string[];
      /**
       * Branches to ignore, defaults to none.
       * You can use glob patterns to match multiple branches.
       * @example ["main", "feature/*"]
       */
      branchesIgnore?: string[];
    };
    /**
     * Pull request configuration, by default no pull requests are triggered.
     * If you want to trigger on pull requests, you need to specify the events.
     */
    pr?: {
      /**
       * Events to trigger on, defaults to none.
       * @example ["opened", "closed"]
       */
      events?: PullRequestEvent[];
      /**
       * Branches to trigger on, defaults to all branches.
       * You can use glob patterns to match multiple branches.
       * @example ["main", "feature/*"]
       */
      targetBranches?: string[];
      /**
       * Branches to ignore, defaults to none.
       * You can use glob patterns to match multiple branches.
       * @example ["main", "feature/*"]
       */
      targetBranchesIgnore?: string[];
    };
  };
}
