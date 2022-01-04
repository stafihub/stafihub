/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { Relayer, Threshold } from '../relayers/relayer'

export const protobufPackage = 'stafiprotocol.stafihub.relayers'

export interface QueryAllRelayerRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllRelayerResponse {
  relayers: Relayer[]
  pagination: PageResponse | undefined
}

export interface QueryRelayersByDenomRequest {
  denom: string
  pagination: PageRequest | undefined
}

export interface QueryRelayersByDenomResponse {
  relayers: Relayer[]
  pagination: PageResponse | undefined
}

export interface QueryGetThresholdRequest {
  denom: string
}

export interface QueryGetThresholdResponse {
  threshold: Threshold | undefined
}

export interface QueryAllThresholdRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllThresholdResponse {
  threshold: Threshold[]
  pagination: PageResponse | undefined
}

const baseQueryAllRelayerRequest: object = {}

export const QueryAllRelayerRequest = {
  encode(message: QueryAllRelayerRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRelayerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRelayerRequest } as QueryAllRelayerRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllRelayerRequest {
    const message = { ...baseQueryAllRelayerRequest } as QueryAllRelayerRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRelayerRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRelayerRequest>): QueryAllRelayerRequest {
    const message = { ...baseQueryAllRelayerRequest } as QueryAllRelayerRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllRelayerResponse: object = {}

export const QueryAllRelayerResponse = {
  encode(message: QueryAllRelayerResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.relayers) {
      Relayer.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRelayerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRelayerResponse } as QueryAllRelayerResponse
    message.relayers = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.relayers.push(Relayer.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllRelayerResponse {
    const message = { ...baseQueryAllRelayerResponse } as QueryAllRelayerResponse
    message.relayers = []
    if (object.relayers !== undefined && object.relayers !== null) {
      for (const e of object.relayers) {
        message.relayers.push(Relayer.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRelayerResponse): unknown {
    const obj: any = {}
    if (message.relayers) {
      obj.relayers = message.relayers.map((e) => (e ? Relayer.toJSON(e) : undefined))
    } else {
      obj.relayers = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRelayerResponse>): QueryAllRelayerResponse {
    const message = { ...baseQueryAllRelayerResponse } as QueryAllRelayerResponse
    message.relayers = []
    if (object.relayers !== undefined && object.relayers !== null) {
      for (const e of object.relayers) {
        message.relayers.push(Relayer.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryRelayersByDenomRequest: object = { denom: '' }

export const QueryRelayersByDenomRequest = {
  encode(message: QueryRelayersByDenomRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryRelayersByDenomRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryRelayersByDenomRequest } as QueryRelayersByDenomRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryRelayersByDenomRequest {
    const message = { ...baseQueryRelayersByDenomRequest } as QueryRelayersByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryRelayersByDenomRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryRelayersByDenomRequest>): QueryRelayersByDenomRequest {
    const message = { ...baseQueryRelayersByDenomRequest } as QueryRelayersByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryRelayersByDenomResponse: object = {}

export const QueryRelayersByDenomResponse = {
  encode(message: QueryRelayersByDenomResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.relayers) {
      Relayer.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryRelayersByDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryRelayersByDenomResponse } as QueryRelayersByDenomResponse
    message.relayers = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.relayers.push(Relayer.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryRelayersByDenomResponse {
    const message = { ...baseQueryRelayersByDenomResponse } as QueryRelayersByDenomResponse
    message.relayers = []
    if (object.relayers !== undefined && object.relayers !== null) {
      for (const e of object.relayers) {
        message.relayers.push(Relayer.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryRelayersByDenomResponse): unknown {
    const obj: any = {}
    if (message.relayers) {
      obj.relayers = message.relayers.map((e) => (e ? Relayer.toJSON(e) : undefined))
    } else {
      obj.relayers = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryRelayersByDenomResponse>): QueryRelayersByDenomResponse {
    const message = { ...baseQueryRelayersByDenomResponse } as QueryRelayersByDenomResponse
    message.relayers = []
    if (object.relayers !== undefined && object.relayers !== null) {
      for (const e of object.relayers) {
        message.relayers.push(Relayer.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetThresholdRequest: object = { denom: '' }

export const QueryGetThresholdRequest = {
  encode(message: QueryGetThresholdRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetThresholdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetThresholdRequest } as QueryGetThresholdRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetThresholdRequest {
    const message = { ...baseQueryGetThresholdRequest } as QueryGetThresholdRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryGetThresholdRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetThresholdRequest>): QueryGetThresholdRequest {
    const message = { ...baseQueryGetThresholdRequest } as QueryGetThresholdRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryGetThresholdResponse: object = {}

export const QueryGetThresholdResponse = {
  encode(message: QueryGetThresholdResponse, writer: Writer = Writer.create()): Writer {
    if (message.threshold !== undefined) {
      Threshold.encode(message.threshold, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetThresholdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetThresholdResponse } as QueryGetThresholdResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.threshold = Threshold.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetThresholdResponse {
    const message = { ...baseQueryGetThresholdResponse } as QueryGetThresholdResponse
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = Threshold.fromJSON(object.threshold)
    } else {
      message.threshold = undefined
    }
    return message
  },

  toJSON(message: QueryGetThresholdResponse): unknown {
    const obj: any = {}
    message.threshold !== undefined && (obj.threshold = message.threshold ? Threshold.toJSON(message.threshold) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetThresholdResponse>): QueryGetThresholdResponse {
    const message = { ...baseQueryGetThresholdResponse } as QueryGetThresholdResponse
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = Threshold.fromPartial(object.threshold)
    } else {
      message.threshold = undefined
    }
    return message
  }
}

const baseQueryAllThresholdRequest: object = {}

export const QueryAllThresholdRequest = {
  encode(message: QueryAllThresholdRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllThresholdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllThresholdRequest } as QueryAllThresholdRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllThresholdRequest {
    const message = { ...baseQueryAllThresholdRequest } as QueryAllThresholdRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllThresholdRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllThresholdRequest>): QueryAllThresholdRequest {
    const message = { ...baseQueryAllThresholdRequest } as QueryAllThresholdRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllThresholdResponse: object = {}

export const QueryAllThresholdResponse = {
  encode(message: QueryAllThresholdResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.threshold) {
      Threshold.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllThresholdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllThresholdResponse } as QueryAllThresholdResponse
    message.threshold = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.threshold.push(Threshold.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllThresholdResponse {
    const message = { ...baseQueryAllThresholdResponse } as QueryAllThresholdResponse
    message.threshold = []
    if (object.threshold !== undefined && object.threshold !== null) {
      for (const e of object.threshold) {
        message.threshold.push(Threshold.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllThresholdResponse): unknown {
    const obj: any = {}
    if (message.threshold) {
      obj.threshold = message.threshold.map((e) => (e ? Threshold.toJSON(e) : undefined))
    } else {
      obj.threshold = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllThresholdResponse>): QueryAllThresholdResponse {
    const message = { ...baseQueryAllThresholdResponse } as QueryAllThresholdResponse
    message.threshold = []
    if (object.threshold !== undefined && object.threshold !== null) {
      for (const e of object.threshold) {
        message.threshold.push(Threshold.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a list of relayer items. */
  RelayerAll(request: QueryAllRelayerRequest): Promise<QueryAllRelayerResponse>
  /** Queries a list of relayersByDenom items. */
  RelayersByDenom(request: QueryRelayersByDenomRequest): Promise<QueryRelayersByDenomResponse>
  /** Queries a threshold by denom. */
  Threshold(request: QueryGetThresholdRequest): Promise<QueryGetThresholdResponse>
  /** Queries a list of threshold items. */
  ThresholdAll(request: QueryAllThresholdRequest): Promise<QueryAllThresholdResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  RelayerAll(request: QueryAllRelayerRequest): Promise<QueryAllRelayerResponse> {
    const data = QueryAllRelayerRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'RelayerAll', data)
    return promise.then((data) => QueryAllRelayerResponse.decode(new Reader(data)))
  }

  RelayersByDenom(request: QueryRelayersByDenomRequest): Promise<QueryRelayersByDenomResponse> {
    const data = QueryRelayersByDenomRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'RelayersByDenom', data)
    return promise.then((data) => QueryRelayersByDenomResponse.decode(new Reader(data)))
  }

  Threshold(request: QueryGetThresholdRequest): Promise<QueryGetThresholdResponse> {
    const data = QueryGetThresholdRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'Threshold', data)
    return promise.then((data) => QueryGetThresholdResponse.decode(new Reader(data)))
  }

  ThresholdAll(request: QueryAllThresholdRequest): Promise<QueryAllThresholdResponse> {
    const data = QueryAllThresholdRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Query', 'ThresholdAll', data)
    return promise.then((data) => QueryAllThresholdResponse.decode(new Reader(data)))
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
