/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.sudo'

export interface QueryAdminRequest {}

export interface QueryAdminResponse {
  address: string
}

const baseQueryAdminRequest: object = {}

export const QueryAdminRequest = {
  encode(_: QueryAdminRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAdminRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAdminRequest } as QueryAdminRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryAdminRequest {
    const message = { ...baseQueryAdminRequest } as QueryAdminRequest
    return message
  },

  toJSON(_: QueryAdminRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryAdminRequest>): QueryAdminRequest {
    const message = { ...baseQueryAdminRequest } as QueryAdminRequest
    return message
  }
}

const baseQueryAdminResponse: object = { address: '' }

export const QueryAdminResponse = {
  encode(message: QueryAdminResponse, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAdminResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAdminResponse } as QueryAdminResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAdminResponse {
    const message = { ...baseQueryAdminResponse } as QueryAdminResponse
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryAdminResponse): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAdminResponse>): QueryAdminResponse {
    const message = { ...baseQueryAdminResponse } as QueryAdminResponse
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a list of admin items. */
  Admin(request: QueryAdminRequest): Promise<QueryAdminResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Admin(request: QueryAdminRequest): Promise<QueryAdminResponse> {
    const data = QueryAdminRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.sudo.Query', 'Admin', data)
    return promise.then((data) => QueryAdminResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
