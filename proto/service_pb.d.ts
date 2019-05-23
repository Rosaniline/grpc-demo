// package: 
// file: service.proto

import * as jspb from "google-protobuf";
import * as person_pb from "./person_pb";

export class getPersonRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): getPersonRequest.AsObject;
  static toObject(includeInstance: boolean, msg: getPersonRequest): getPersonRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: getPersonRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): getPersonRequest;
  static deserializeBinaryFromReader(message: getPersonRequest, reader: jspb.BinaryReader): getPersonRequest;
}

export namespace getPersonRequest {
  export type AsObject = {
    id: number,
  }
}

