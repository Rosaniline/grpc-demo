// package: 
// file: service.proto

import * as service_pb from "./service_pb";
import * as person_pb from "./person_pb";
import {grpc} from "@improbable-eng/grpc-web";

type HuaFuGetPerson = {
  readonly methodName: string;
  readonly service: typeof HuaFu;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.getPersonRequest;
  readonly responseType: typeof person_pb.Person;
};

export class HuaFu {
  static readonly serviceName: string;
  static readonly GetPerson: HuaFuGetPerson;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class HuaFuClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getPerson(
    requestMessage: service_pb.getPersonRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: person_pb.Person|null) => void
  ): UnaryResponse;
  getPerson(
    requestMessage: service_pb.getPersonRequest,
    callback: (error: ServiceError|null, responseMessage: person_pb.Person|null) => void
  ): UnaryResponse;
}

