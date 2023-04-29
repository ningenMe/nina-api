// @generated by protoc-gen-connect-es v0.8.6 with parameter "target=ts,import_extension=.ts"
// @generated from file v1/nina.proto (package nina.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { ContributionDeleteRequest, ContributionGetResponse, ContributionPostRequest, ContributionStatisticsGetRequest, ContributionStatisticsGetResponse } from "./nina_pb.ts";

/**
 * @generated from service nina.v1.NinaService
 */
export const NinaService = {
  typeName: "nina.v1.NinaService",
  methods: {
    /**
     * @generated from rpc nina.v1.NinaService.ContributionGet
     */
    contributionGet: {
      name: "ContributionGet",
      I: Empty,
      O: ContributionGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc nina.v1.NinaService.ContributionPost
     */
    contributionPost: {
      name: "ContributionPost",
      I: ContributionPostRequest,
      O: Empty,
      kind: MethodKind.ClientStreaming,
    },
    /**
     * @generated from rpc nina.v1.NinaService.ContributionDelete
     */
    contributionDelete: {
      name: "ContributionDelete",
      I: ContributionDeleteRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc nina.v1.NinaService.ContributionStatisticsGet
     */
    contributionStatisticsGet: {
      name: "ContributionStatisticsGet",
      I: ContributionStatisticsGetRequest,
      O: ContributionStatisticsGetResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
