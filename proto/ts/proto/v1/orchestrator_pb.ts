// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file proto/v1/orchestrator.proto (package proto.v1, syntax proto3)
/* eslint-disable */

import type {
  GenFile,
  GenMessage,
  GenService,
} from "@bufbuild/protobuf/codegenv1";
import {
  fileDesc,
  messageDesc,
  serviceDesc,
} from "@bufbuild/protobuf/codegenv1";
import type {
  Conclusion,
  ExecInfo,
  JobMeta,
  StartTaskData,
  TaskMeta,
  WorkflowAlert,
  WorkflowMeta,
} from "./shared_pb.ts";
import { file_proto_v1_shared } from "./shared_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file proto/v1/orchestrator.proto.
 */
export const file_proto_v1_orchestrator: GenFile = /*@__PURE__*/
  fileDesc(
    "Chtwcm90by92MS9vcmNoZXN0cmF0b3IucHJvdG8SCHByb3RvLnYxIloKKU9yY2hlc3RyYXRvclNlcnZpY2VXb3JrZmxvd1N0YXJ0ZWRSZXF1ZXN0Ei0KDXdvcmtmbG93X21ldGEYASABKAsyFi5wcm90by52MS5Xb3JrZmxvd01ldGEiLAoqT3JjaGVzdHJhdG9yU2VydmljZVdvcmtmbG93U3RhcnRlZFJlc3BvbnNlIpMBCiNPcmNoZXN0cmF0b3JTZXJ2aWNlQ3JlYXRlSm9iUmVxdWVzdBItCg13b3JrZmxvd19tZXRhGAEgASgLMhYucHJvdG8udjEuV29ya2Zsb3dNZXRhEgwKBG5hbWUYAiABKAkSEwoGcnVubmVyGAMgASgJSACIAQESDwoHc2tpcHBlZBgEIAEoCEIJCgdfcnVubmVyIksKJE9yY2hlc3RyYXRvclNlcnZpY2VDcmVhdGVKb2JSZXNwb25zZRIjCghqb2JfbWV0YRgBIAEoCzIRLnByb3RvLnYxLkpvYk1ldGEisgEKJE9yY2hlc3RyYXRvclNlcnZpY2VDcmVhdGVUYXNrUmVxdWVzdBItCg13b3JrZmxvd19tZXRhGAEgASgLMhYucHJvdG8udjEuV29ya2Zsb3dNZXRhEiMKCGpvYl9tZXRhGAIgASgLMhEucHJvdG8udjEuSm9iTWV0YRIlCgRkYXRhGAMgASgLMhcucHJvdG8udjEuU3RhcnRUYXNrRGF0YRIPCgdza2lwcGVkGAQgASgIIjMKJU9yY2hlc3RyYXRvclNlcnZpY2VDcmVhdGVUYXNrUmVzcG9uc2USCgoCaWQYASABKAkigwEKKE9yY2hlc3RyYXRvclNlcnZpY2VGaW5pc2hXb3JrZmxvd1JlcXVlc3QSLQoNd29ya2Zsb3dfbWV0YRgBIAEoCzIWLnByb3RvLnYxLldvcmtmbG93TWV0YRIoCgpjb25jbHVzaW9uGAIgASgOMhQucHJvdG8udjEuQ29uY2x1c2lvbiIrCilPcmNoZXN0cmF0b3JTZXJ2aWNlRmluaXNoV29ya2Zsb3dSZXNwb25zZSLLAQokT3JjaGVzdHJhdG9yU2VydmljZUZpbmlzaFRhc2tSZXF1ZXN0Ei0KDXdvcmtmbG93X21ldGEYASABKAsyFi5wcm90by52MS5Xb3JrZmxvd01ldGESIwoIam9iX21ldGEYAiABKAsyES5wcm90by52MS5Kb2JNZXRhEiUKCXRhc2tfbWV0YRgDIAEoCzISLnByb3RvLnYxLlRhc2tNZXRhEigKCmNvbmNsdXNpb24YBCABKA4yFC5wcm90by52MS5Db25jbHVzaW9uIicKJU9yY2hlc3RyYXRvclNlcnZpY2VGaW5pc2hUYXNrUmVzcG9uc2UiowEKI09yY2hlc3RyYXRvclNlcnZpY2VGaW5pc2hKb2JSZXF1ZXN0Ei0KDXdvcmtmbG93X21ldGEYASABKAsyFi5wcm90by52MS5Xb3JrZmxvd01ldGESIwoIam9iX21ldGEYAiABKAsyES5wcm90by52MS5Kb2JNZXRhEigKCmNvbmNsdXNpb24YAyABKA4yFC5wcm90by52MS5Db25jbHVzaW9uIiYKJE9yY2hlc3RyYXRvclNlcnZpY2VGaW5pc2hKb2JSZXNwb25zZSL4AQokT3JjaGVzdHJhdG9yU2VydmljZUNyZWF0ZVN0ZXBSZXF1ZXN0Ei0KDXdvcmtmbG93X21ldGEYASABKAsyFi5wcm90by52MS5Xb3JrZmxvd01ldGESIwoIam9iX21ldGEYAiABKAsyES5wcm90by52MS5Kb2JNZXRhEioKCXRhc2tfbWV0YRgDIAEoCzISLnByb3RvLnYxLlRhc2tNZXRhSAGIAQESDAoEbmFtZRgEIAEoCRInCglleGVjX2RhdGEYBSABKAsyEi5wcm90by52MS5FeGVjSW5mb0gAQgsKCXN0ZXBfZGF0YUIMCgpfdGFza19tZXRhIlEKJU9yY2hlc3RyYXRvclNlcnZpY2VDcmVhdGVTdGVwUmVzcG9uc2USCgoCaWQYASABKAkSHAoUcHJlc2lnbmVkX291dHB1dF91cmwYAiABKAki1wEKJE9yY2hlc3RyYXRvclNlcnZpY2VGaW5pc2hTdGVwUmVxdWVzdBItCg13b3JrZmxvd19tZXRhGAEgASgLMhYucHJvdG8udjEuV29ya2Zsb3dNZXRhEiMKCGpvYl9tZXRhGAIgASgLMhEucHJvdG8udjEuSm9iTWV0YRIlCgl0YXNrX21ldGEYAyABKAsyEi5wcm90by52MS5UYXNrTWV0YRIKCgJpZBgEIAEoCRIoCgpjb25jbHVzaW9uGAUgASgOMhQucHJvdG8udjEuQ29uY2x1c2lvbiInCiVPcmNoZXN0cmF0b3JTZXJ2aWNlRmluaXNoU3RlcFJlc3BvbnNlInoKJE9yY2hlc3RyYXRvclNlcnZpY2VKb2JTdGFydGVkUmVxdWVzdBItCg13b3JrZmxvd19tZXRhGAEgASgLMhYucHJvdG8udjEuV29ya2Zsb3dNZXRhEiMKCGpvYl9tZXRhGAIgASgLMhEucHJvdG8udjEuSm9iTWV0YSInCiVPcmNoZXN0cmF0b3JTZXJ2aWNlSm9iU3RhcnRlZFJlc3BvbnNlIoYBCi1PcmNoZXN0cmF0b3JTZXJ2aWNlQ3JlYXRlV29ya2Zsb3dBbGVydFJlcXVlc3QSLQoNd29ya2Zsb3dfbWV0YRgBIAEoCzIWLnByb3RvLnYxLldvcmtmbG93TWV0YRImCgVhbGVydBgCIAEoCzIXLnByb3RvLnYxLldvcmtmbG93QWxlcnQiMAouT3JjaGVzdHJhdG9yU2VydmljZUNyZWF0ZVdvcmtmbG93QWxlcnRSZXNwb25zZTKcCQoTT3JjaGVzdHJhdG9yU2VydmljZRJ8Cg9Xb3JrZmxvd1N0YXJ0ZWQSMy5wcm90by52MS5PcmNoZXN0cmF0b3JTZXJ2aWNlV29ya2Zsb3dTdGFydGVkUmVxdWVzdBo0LnByb3RvLnYxLk9yY2hlc3RyYXRvclNlcnZpY2VXb3JrZmxvd1N0YXJ0ZWRSZXNwb25zZRJ5Cg5GaW5pc2hXb3JrZmxvdxIyLnByb3RvLnYxLk9yY2hlc3RyYXRvclNlcnZpY2VGaW5pc2hXb3JrZmxvd1JlcXVlc3QaMy5wcm90by52MS5PcmNoZXN0cmF0b3JTZXJ2aWNlRmluaXNoV29ya2Zsb3dSZXNwb25zZRJqCglDcmVhdGVKb2ISLS5wcm90by52MS5PcmNoZXN0cmF0b3JTZXJ2aWNlQ3JlYXRlSm9iUmVxdWVzdBouLnByb3RvLnYxLk9yY2hlc3RyYXRvclNlcnZpY2VDcmVhdGVKb2JSZXNwb25zZRJqCglGaW5pc2hKb2ISLS5wcm90by52MS5PcmNoZXN0cmF0b3JTZXJ2aWNlRmluaXNoSm9iUmVxdWVzdBouLnByb3RvLnYxLk9yY2hlc3RyYXRvclNlcnZpY2VGaW5pc2hKb2JSZXNwb25zZRJtCgpKb2JTdGFydGVkEi4ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUpvYlN0YXJ0ZWRSZXF1ZXN0Gi8ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUpvYlN0YXJ0ZWRSZXNwb25zZRJtCgpDcmVhdGVUYXNrEi4ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUNyZWF0ZVRhc2tSZXF1ZXN0Gi8ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUNyZWF0ZVRhc2tSZXNwb25zZRJtCgpGaW5pc2hUYXNrEi4ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUZpbmlzaFRhc2tSZXF1ZXN0Gi8ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUZpbmlzaFRhc2tSZXNwb25zZRJtCgpDcmVhdGVTdGVwEi4ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUNyZWF0ZVN0ZXBSZXF1ZXN0Gi8ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUNyZWF0ZVN0ZXBSZXNwb25zZRJtCgpGaW5pc2hTdGVwEi4ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUZpbmlzaFN0ZXBSZXF1ZXN0Gi8ucHJvdG8udjEuT3JjaGVzdHJhdG9yU2VydmljZUZpbmlzaFN0ZXBSZXNwb25zZRKIAQoTQ3JlYXRlV29ya2Zsb3dBbGVydBI3LnByb3RvLnYxLk9yY2hlc3RyYXRvclNlcnZpY2VDcmVhdGVXb3JrZmxvd0FsZXJ0UmVxdWVzdBo4LnByb3RvLnYxLk9yY2hlc3RyYXRvclNlcnZpY2VDcmVhdGVXb3JrZmxvd0FsZXJ0UmVzcG9uc2VCDVoLcHJvdG8vZ28vdjFiBnByb3RvMw",
    [file_proto_v1_shared],
  );

/**
 * @generated from message proto.v1.OrchestratorServiceWorkflowStartedRequest
 */
export type OrchestratorServiceWorkflowStartedRequest =
  & Message<"proto.v1.OrchestratorServiceWorkflowStartedRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceWorkflowStartedRequest.
 * Use `create(OrchestratorServiceWorkflowStartedRequestSchema)` to create a new message.
 */
export const OrchestratorServiceWorkflowStartedRequestSchema: GenMessage<
  OrchestratorServiceWorkflowStartedRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 0);

/**
 * @generated from message proto.v1.OrchestratorServiceWorkflowStartedResponse
 */
export type OrchestratorServiceWorkflowStartedResponse =
  & Message<"proto.v1.OrchestratorServiceWorkflowStartedResponse">
  & {};

/**
 * Describes the message proto.v1.OrchestratorServiceWorkflowStartedResponse.
 * Use `create(OrchestratorServiceWorkflowStartedResponseSchema)` to create a new message.
 */
export const OrchestratorServiceWorkflowStartedResponseSchema: GenMessage<
  OrchestratorServiceWorkflowStartedResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 1);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateJobRequest
 */
export type OrchestratorServiceCreateJobRequest =
  & Message<"proto.v1.OrchestratorServiceCreateJobRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: string name = 2;
     */
    name: string;

    /**
     * @generated from field: optional string runner = 3;
     */
    runner?: string;

    /**
     * @generated from field: bool skipped = 4;
     */
    skipped: boolean;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceCreateJobRequest.
 * Use `create(OrchestratorServiceCreateJobRequestSchema)` to create a new message.
 */
export const OrchestratorServiceCreateJobRequestSchema: GenMessage<
  OrchestratorServiceCreateJobRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 2);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateJobResponse
 */
export type OrchestratorServiceCreateJobResponse =
  & Message<"proto.v1.OrchestratorServiceCreateJobResponse">
  & {
    /**
     * @generated from field: proto.v1.JobMeta job_meta = 1;
     */
    jobMeta?: JobMeta;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceCreateJobResponse.
 * Use `create(OrchestratorServiceCreateJobResponseSchema)` to create a new message.
 */
export const OrchestratorServiceCreateJobResponseSchema: GenMessage<
  OrchestratorServiceCreateJobResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 3);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateTaskRequest
 */
export type OrchestratorServiceCreateTaskRequest =
  & Message<"proto.v1.OrchestratorServiceCreateTaskRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.JobMeta job_meta = 2;
     */
    jobMeta?: JobMeta;

    /**
     * @generated from field: proto.v1.StartTaskData data = 3;
     */
    data?: StartTaskData;

    /**
     * @generated from field: bool skipped = 4;
     */
    skipped: boolean;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceCreateTaskRequest.
 * Use `create(OrchestratorServiceCreateTaskRequestSchema)` to create a new message.
 */
export const OrchestratorServiceCreateTaskRequestSchema: GenMessage<
  OrchestratorServiceCreateTaskRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 4);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateTaskResponse
 */
export type OrchestratorServiceCreateTaskResponse =
  & Message<"proto.v1.OrchestratorServiceCreateTaskResponse">
  & {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceCreateTaskResponse.
 * Use `create(OrchestratorServiceCreateTaskResponseSchema)` to create a new message.
 */
export const OrchestratorServiceCreateTaskResponseSchema: GenMessage<
  OrchestratorServiceCreateTaskResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 5);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishWorkflowRequest
 */
export type OrchestratorServiceFinishWorkflowRequest =
  & Message<"proto.v1.OrchestratorServiceFinishWorkflowRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.Conclusion conclusion = 2;
     */
    conclusion: Conclusion;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceFinishWorkflowRequest.
 * Use `create(OrchestratorServiceFinishWorkflowRequestSchema)` to create a new message.
 */
export const OrchestratorServiceFinishWorkflowRequestSchema: GenMessage<
  OrchestratorServiceFinishWorkflowRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 6);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishWorkflowResponse
 */
export type OrchestratorServiceFinishWorkflowResponse =
  & Message<"proto.v1.OrchestratorServiceFinishWorkflowResponse">
  & {};

/**
 * Describes the message proto.v1.OrchestratorServiceFinishWorkflowResponse.
 * Use `create(OrchestratorServiceFinishWorkflowResponseSchema)` to create a new message.
 */
export const OrchestratorServiceFinishWorkflowResponseSchema: GenMessage<
  OrchestratorServiceFinishWorkflowResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 7);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishTaskRequest
 */
export type OrchestratorServiceFinishTaskRequest =
  & Message<"proto.v1.OrchestratorServiceFinishTaskRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.JobMeta job_meta = 2;
     */
    jobMeta?: JobMeta;

    /**
     * @generated from field: proto.v1.TaskMeta task_meta = 3;
     */
    taskMeta?: TaskMeta;

    /**
     * @generated from field: proto.v1.Conclusion conclusion = 4;
     */
    conclusion: Conclusion;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceFinishTaskRequest.
 * Use `create(OrchestratorServiceFinishTaskRequestSchema)` to create a new message.
 */
export const OrchestratorServiceFinishTaskRequestSchema: GenMessage<
  OrchestratorServiceFinishTaskRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 8);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishTaskResponse
 */
export type OrchestratorServiceFinishTaskResponse =
  & Message<"proto.v1.OrchestratorServiceFinishTaskResponse">
  & {};

/**
 * Describes the message proto.v1.OrchestratorServiceFinishTaskResponse.
 * Use `create(OrchestratorServiceFinishTaskResponseSchema)` to create a new message.
 */
export const OrchestratorServiceFinishTaskResponseSchema: GenMessage<
  OrchestratorServiceFinishTaskResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 9);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishJobRequest
 */
export type OrchestratorServiceFinishJobRequest =
  & Message<"proto.v1.OrchestratorServiceFinishJobRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.JobMeta job_meta = 2;
     */
    jobMeta?: JobMeta;

    /**
     * @generated from field: proto.v1.Conclusion conclusion = 3;
     */
    conclusion: Conclusion;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceFinishJobRequest.
 * Use `create(OrchestratorServiceFinishJobRequestSchema)` to create a new message.
 */
export const OrchestratorServiceFinishJobRequestSchema: GenMessage<
  OrchestratorServiceFinishJobRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 10);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishJobResponse
 */
export type OrchestratorServiceFinishJobResponse =
  & Message<"proto.v1.OrchestratorServiceFinishJobResponse">
  & {};

/**
 * Describes the message proto.v1.OrchestratorServiceFinishJobResponse.
 * Use `create(OrchestratorServiceFinishJobResponseSchema)` to create a new message.
 */
export const OrchestratorServiceFinishJobResponseSchema: GenMessage<
  OrchestratorServiceFinishJobResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 11);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateStepRequest
 */
export type OrchestratorServiceCreateStepRequest =
  & Message<"proto.v1.OrchestratorServiceCreateStepRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.JobMeta job_meta = 2;
     */
    jobMeta?: JobMeta;

    /**
     * @generated from field: optional proto.v1.TaskMeta task_meta = 3;
     */
    taskMeta?: TaskMeta;

    /**
     * @generated from field: string name = 4;
     */
    name: string;

    /**
     * @generated from oneof proto.v1.OrchestratorServiceCreateStepRequest.step_data
     */
    stepData: {
      /**
       * @generated from field: proto.v1.ExecInfo exec_data = 5;
       */
      value: ExecInfo;
      case: "execData";
    } | { case: undefined; value?: undefined };
  };

/**
 * Describes the message proto.v1.OrchestratorServiceCreateStepRequest.
 * Use `create(OrchestratorServiceCreateStepRequestSchema)` to create a new message.
 */
export const OrchestratorServiceCreateStepRequestSchema: GenMessage<
  OrchestratorServiceCreateStepRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 12);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateStepResponse
 */
export type OrchestratorServiceCreateStepResponse =
  & Message<"proto.v1.OrchestratorServiceCreateStepResponse">
  & {
    /**
     * @generated from field: string id = 1;
     */
    id: string;

    /**
     * @generated from field: string presigned_output_url = 2;
     */
    presignedOutputUrl: string;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceCreateStepResponse.
 * Use `create(OrchestratorServiceCreateStepResponseSchema)` to create a new message.
 */
export const OrchestratorServiceCreateStepResponseSchema: GenMessage<
  OrchestratorServiceCreateStepResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 13);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishStepRequest
 */
export type OrchestratorServiceFinishStepRequest =
  & Message<"proto.v1.OrchestratorServiceFinishStepRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.JobMeta job_meta = 2;
     */
    jobMeta?: JobMeta;

    /**
     * @generated from field: proto.v1.TaskMeta task_meta = 3;
     */
    taskMeta?: TaskMeta;

    /**
     * @generated from field: string id = 4;
     */
    id: string;

    /**
     * @generated from field: proto.v1.Conclusion conclusion = 5;
     */
    conclusion: Conclusion;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceFinishStepRequest.
 * Use `create(OrchestratorServiceFinishStepRequestSchema)` to create a new message.
 */
export const OrchestratorServiceFinishStepRequestSchema: GenMessage<
  OrchestratorServiceFinishStepRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 14);

/**
 * @generated from message proto.v1.OrchestratorServiceFinishStepResponse
 */
export type OrchestratorServiceFinishStepResponse =
  & Message<"proto.v1.OrchestratorServiceFinishStepResponse">
  & {};

/**
 * Describes the message proto.v1.OrchestratorServiceFinishStepResponse.
 * Use `create(OrchestratorServiceFinishStepResponseSchema)` to create a new message.
 */
export const OrchestratorServiceFinishStepResponseSchema: GenMessage<
  OrchestratorServiceFinishStepResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 15);

/**
 * @generated from message proto.v1.OrchestratorServiceJobStartedRequest
 */
export type OrchestratorServiceJobStartedRequest =
  & Message<"proto.v1.OrchestratorServiceJobStartedRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.JobMeta job_meta = 2;
     */
    jobMeta?: JobMeta;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceJobStartedRequest.
 * Use `create(OrchestratorServiceJobStartedRequestSchema)` to create a new message.
 */
export const OrchestratorServiceJobStartedRequestSchema: GenMessage<
  OrchestratorServiceJobStartedRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 16);

/**
 * @generated from message proto.v1.OrchestratorServiceJobStartedResponse
 */
export type OrchestratorServiceJobStartedResponse =
  & Message<"proto.v1.OrchestratorServiceJobStartedResponse">
  & {};

/**
 * Describes the message proto.v1.OrchestratorServiceJobStartedResponse.
 * Use `create(OrchestratorServiceJobStartedResponseSchema)` to create a new message.
 */
export const OrchestratorServiceJobStartedResponseSchema: GenMessage<
  OrchestratorServiceJobStartedResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 17);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateWorkflowAlertRequest
 */
export type OrchestratorServiceCreateWorkflowAlertRequest =
  & Message<"proto.v1.OrchestratorServiceCreateWorkflowAlertRequest">
  & {
    /**
     * @generated from field: proto.v1.WorkflowMeta workflow_meta = 1;
     */
    workflowMeta?: WorkflowMeta;

    /**
     * @generated from field: proto.v1.WorkflowAlert alert = 2;
     */
    alert?: WorkflowAlert;
  };

/**
 * Describes the message proto.v1.OrchestratorServiceCreateWorkflowAlertRequest.
 * Use `create(OrchestratorServiceCreateWorkflowAlertRequestSchema)` to create a new message.
 */
export const OrchestratorServiceCreateWorkflowAlertRequestSchema: GenMessage<
  OrchestratorServiceCreateWorkflowAlertRequest
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 18);

/**
 * @generated from message proto.v1.OrchestratorServiceCreateWorkflowAlertResponse
 */
export type OrchestratorServiceCreateWorkflowAlertResponse =
  & Message<"proto.v1.OrchestratorServiceCreateWorkflowAlertResponse">
  & {};

/**
 * Describes the message proto.v1.OrchestratorServiceCreateWorkflowAlertResponse.
 * Use `create(OrchestratorServiceCreateWorkflowAlertResponseSchema)` to create a new message.
 */
export const OrchestratorServiceCreateWorkflowAlertResponseSchema: GenMessage<
  OrchestratorServiceCreateWorkflowAlertResponse
> = /*@__PURE__*/
  messageDesc(file_proto_v1_orchestrator, 19);

/**
 * @generated from service proto.v1.OrchestratorService
 */
export const OrchestratorService: GenService<{
  /**
   * @generated from rpc proto.v1.OrchestratorService.WorkflowStarted
   */
  workflowStarted: {
    methodKind: "unary";
    input: typeof OrchestratorServiceWorkflowStartedRequestSchema;
    output: typeof OrchestratorServiceWorkflowStartedResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.FinishWorkflow
   */
  finishWorkflow: {
    methodKind: "unary";
    input: typeof OrchestratorServiceFinishWorkflowRequestSchema;
    output: typeof OrchestratorServiceFinishWorkflowResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.CreateJob
   */
  createJob: {
    methodKind: "unary";
    input: typeof OrchestratorServiceCreateJobRequestSchema;
    output: typeof OrchestratorServiceCreateJobResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.FinishJob
   */
  finishJob: {
    methodKind: "unary";
    input: typeof OrchestratorServiceFinishJobRequestSchema;
    output: typeof OrchestratorServiceFinishJobResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.JobStarted
   */
  jobStarted: {
    methodKind: "unary";
    input: typeof OrchestratorServiceJobStartedRequestSchema;
    output: typeof OrchestratorServiceJobStartedResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.CreateTask
   */
  createTask: {
    methodKind: "unary";
    input: typeof OrchestratorServiceCreateTaskRequestSchema;
    output: typeof OrchestratorServiceCreateTaskResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.FinishTask
   */
  finishTask: {
    methodKind: "unary";
    input: typeof OrchestratorServiceFinishTaskRequestSchema;
    output: typeof OrchestratorServiceFinishTaskResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.CreateStep
   */
  createStep: {
    methodKind: "unary";
    input: typeof OrchestratorServiceCreateStepRequestSchema;
    output: typeof OrchestratorServiceCreateStepResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.FinishStep
   */
  finishStep: {
    methodKind: "unary";
    input: typeof OrchestratorServiceFinishStepRequestSchema;
    output: typeof OrchestratorServiceFinishStepResponseSchema;
  };
  /**
   * @generated from rpc proto.v1.OrchestratorService.CreateWorkflowAlert
   */
  createWorkflowAlert: {
    methodKind: "unary";
    input: typeof OrchestratorServiceCreateWorkflowAlertRequestSchema;
    output: typeof OrchestratorServiceCreateWorkflowAlertResponseSchema;
  };
}> = /*@__PURE__*/
  serviceDesc(file_proto_v1_orchestrator, 0);
