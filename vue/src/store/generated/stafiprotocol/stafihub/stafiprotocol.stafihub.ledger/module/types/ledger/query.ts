/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { ExchangeRate, EraExchangeRate } from '../ledger/ledger'

export const protobufPackage = 'stafiprotocol.stafihub.ledger'

export interface QueryGetExchangeRateRequest {
  denom: string
}

export interface QueryGetExchangeRateResponse {
  exchangeRate: ExchangeRate | undefined
}

export interface QueryExchangeRateAllRequest {}

export interface QueryExchangeRateAllResponse {
  exchangeRates: ExchangeRate[]
}

export interface QueryGetEraExchangeRateRequest {
  denom: string
  era: number
}

export interface QueryGetEraExchangeRateResponse {
  eraExchangeRate: EraExchangeRate | undefined
}

export interface QueryEraExchangeRatesByDenomRequest {
  denom: string
}

export interface QueryEraExchangeRatesByDenomResponse {
  eraExchangeRates: EraExchangeRate[]
}

const baseQueryGetExchangeRateRequest: object = { denom: '' }

export const QueryGetExchangeRateRequest = {
  encode(message: QueryGetExchangeRateRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetExchangeRateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetExchangeRateRequest } as QueryGetExchangeRateRequest
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

  fromJSON(object: any): QueryGetExchangeRateRequest {
    const message = { ...baseQueryGetExchangeRateRequest } as QueryGetExchangeRateRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryGetExchangeRateRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetExchangeRateRequest>): QueryGetExchangeRateRequest {
    const message = { ...baseQueryGetExchangeRateRequest } as QueryGetExchangeRateRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryGetExchangeRateResponse: object = {}

export const QueryGetExchangeRateResponse = {
  encode(message: QueryGetExchangeRateResponse, writer: Writer = Writer.create()): Writer {
    if (message.exchangeRate !== undefined) {
      ExchangeRate.encode(message.exchangeRate, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetExchangeRateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetExchangeRateResponse } as QueryGetExchangeRateResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.exchangeRate = ExchangeRate.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetExchangeRateResponse {
    const message = { ...baseQueryGetExchangeRateResponse } as QueryGetExchangeRateResponse
    if (object.exchangeRate !== undefined && object.exchangeRate !== null) {
      message.exchangeRate = ExchangeRate.fromJSON(object.exchangeRate)
    } else {
      message.exchangeRate = undefined
    }
    return message
  },

  toJSON(message: QueryGetExchangeRateResponse): unknown {
    const obj: any = {}
    message.exchangeRate !== undefined && (obj.exchangeRate = message.exchangeRate ? ExchangeRate.toJSON(message.exchangeRate) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetExchangeRateResponse>): QueryGetExchangeRateResponse {
    const message = { ...baseQueryGetExchangeRateResponse } as QueryGetExchangeRateResponse
    if (object.exchangeRate !== undefined && object.exchangeRate !== null) {
      message.exchangeRate = ExchangeRate.fromPartial(object.exchangeRate)
    } else {
      message.exchangeRate = undefined
    }
    return message
  }
}

const baseQueryExchangeRateAllRequest: object = {}

export const QueryExchangeRateAllRequest = {
  encode(_: QueryExchangeRateAllRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryExchangeRateAllRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryExchangeRateAllRequest } as QueryExchangeRateAllRequest
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

  fromJSON(_: any): QueryExchangeRateAllRequest {
    const message = { ...baseQueryExchangeRateAllRequest } as QueryExchangeRateAllRequest
    return message
  },

  toJSON(_: QueryExchangeRateAllRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryExchangeRateAllRequest>): QueryExchangeRateAllRequest {
    const message = { ...baseQueryExchangeRateAllRequest } as QueryExchangeRateAllRequest
    return message
  }
}

const baseQueryExchangeRateAllResponse: object = {}

export const QueryExchangeRateAllResponse = {
  encode(message: QueryExchangeRateAllResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.exchangeRates) {
      ExchangeRate.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryExchangeRateAllResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryExchangeRateAllResponse } as QueryExchangeRateAllResponse
    message.exchangeRates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.exchangeRates.push(ExchangeRate.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryExchangeRateAllResponse {
    const message = { ...baseQueryExchangeRateAllResponse } as QueryExchangeRateAllResponse
    message.exchangeRates = []
    if (object.exchangeRates !== undefined && object.exchangeRates !== null) {
      for (const e of object.exchangeRates) {
        message.exchangeRates.push(ExchangeRate.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: QueryExchangeRateAllResponse): unknown {
    const obj: any = {}
    if (message.exchangeRates) {
      obj.exchangeRates = message.exchangeRates.map((e) => (e ? ExchangeRate.toJSON(e) : undefined))
    } else {
      obj.exchangeRates = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<QueryExchangeRateAllResponse>): QueryExchangeRateAllResponse {
    const message = { ...baseQueryExchangeRateAllResponse } as QueryExchangeRateAllResponse
    message.exchangeRates = []
    if (object.exchangeRates !== undefined && object.exchangeRates !== null) {
      for (const e of object.exchangeRates) {
        message.exchangeRates.push(ExchangeRate.fromPartial(e))
      }
    }
    return message
  }
}

const baseQueryGetEraExchangeRateRequest: object = { denom: '', era: 0 }

export const QueryGetEraExchangeRateRequest = {
  encode(message: QueryGetEraExchangeRateRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.era !== 0) {
      writer.uint32(16).uint32(message.era)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEraExchangeRateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetEraExchangeRateRequest } as QueryGetEraExchangeRateRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.era = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetEraExchangeRateRequest {
    const message = { ...baseQueryGetEraExchangeRateRequest } as QueryGetEraExchangeRateRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.era !== undefined && object.era !== null) {
      message.era = Number(object.era)
    } else {
      message.era = 0
    }
    return message
  },

  toJSON(message: QueryGetEraExchangeRateRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.era !== undefined && (obj.era = message.era)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetEraExchangeRateRequest>): QueryGetEraExchangeRateRequest {
    const message = { ...baseQueryGetEraExchangeRateRequest } as QueryGetEraExchangeRateRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.era !== undefined && object.era !== null) {
      message.era = object.era
    } else {
      message.era = 0
    }
    return message
  }
}

const baseQueryGetEraExchangeRateResponse: object = {}

export const QueryGetEraExchangeRateResponse = {
  encode(message: QueryGetEraExchangeRateResponse, writer: Writer = Writer.create()): Writer {
    if (message.eraExchangeRate !== undefined) {
      EraExchangeRate.encode(message.eraExchangeRate, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEraExchangeRateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetEraExchangeRateResponse } as QueryGetEraExchangeRateResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.eraExchangeRate = EraExchangeRate.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetEraExchangeRateResponse {
    const message = { ...baseQueryGetEraExchangeRateResponse } as QueryGetEraExchangeRateResponse
    if (object.eraExchangeRate !== undefined && object.eraExchangeRate !== null) {
      message.eraExchangeRate = EraExchangeRate.fromJSON(object.eraExchangeRate)
    } else {
      message.eraExchangeRate = undefined
    }
    return message
  },

  toJSON(message: QueryGetEraExchangeRateResponse): unknown {
    const obj: any = {}
    message.eraExchangeRate !== undefined && (obj.eraExchangeRate = message.eraExchangeRate ? EraExchangeRate.toJSON(message.eraExchangeRate) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetEraExchangeRateResponse>): QueryGetEraExchangeRateResponse {
    const message = { ...baseQueryGetEraExchangeRateResponse } as QueryGetEraExchangeRateResponse
    if (object.eraExchangeRate !== undefined && object.eraExchangeRate !== null) {
      message.eraExchangeRate = EraExchangeRate.fromPartial(object.eraExchangeRate)
    } else {
      message.eraExchangeRate = undefined
    }
    return message
  }
}

const baseQueryEraExchangeRatesByDenomRequest: object = { denom: '' }

export const QueryEraExchangeRatesByDenomRequest = {
  encode(message: QueryEraExchangeRatesByDenomRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryEraExchangeRatesByDenomRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryEraExchangeRatesByDenomRequest } as QueryEraExchangeRatesByDenomRequest
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

  fromJSON(object: any): QueryEraExchangeRatesByDenomRequest {
    const message = { ...baseQueryEraExchangeRatesByDenomRequest } as QueryEraExchangeRatesByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryEraExchangeRatesByDenomRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryEraExchangeRatesByDenomRequest>): QueryEraExchangeRatesByDenomRequest {
    const message = { ...baseQueryEraExchangeRatesByDenomRequest } as QueryEraExchangeRatesByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryEraExchangeRatesByDenomResponse: object = {}

export const QueryEraExchangeRatesByDenomResponse = {
  encode(message: QueryEraExchangeRatesByDenomResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.eraExchangeRates) {
      EraExchangeRate.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryEraExchangeRatesByDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryEraExchangeRatesByDenomResponse } as QueryEraExchangeRatesByDenomResponse
    message.eraExchangeRates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.eraExchangeRates.push(EraExchangeRate.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryEraExchangeRatesByDenomResponse {
    const message = { ...baseQueryEraExchangeRatesByDenomResponse } as QueryEraExchangeRatesByDenomResponse
    message.eraExchangeRates = []
    if (object.eraExchangeRates !== undefined && object.eraExchangeRates !== null) {
      for (const e of object.eraExchangeRates) {
        message.eraExchangeRates.push(EraExchangeRate.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: QueryEraExchangeRatesByDenomResponse): unknown {
    const obj: any = {}
    if (message.eraExchangeRates) {
      obj.eraExchangeRates = message.eraExchangeRates.map((e) => (e ? EraExchangeRate.toJSON(e) : undefined))
    } else {
      obj.eraExchangeRates = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<QueryEraExchangeRatesByDenomResponse>): QueryEraExchangeRatesByDenomResponse {
    const message = { ...baseQueryEraExchangeRatesByDenomResponse } as QueryEraExchangeRatesByDenomResponse
    message.eraExchangeRates = []
    if (object.eraExchangeRates !== undefined && object.eraExchangeRates !== null) {
      for (const e of object.eraExchangeRates) {
        message.eraExchangeRates.push(EraExchangeRate.fromPartial(e))
      }
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a list of getExchangeRate items. */
  GetExchangeRate(request: QueryGetExchangeRateRequest): Promise<QueryGetExchangeRateResponse>
  /** Queries a list of exchangeRateAll items. */
  ExchangeRateAll(request: QueryExchangeRateAllRequest): Promise<QueryExchangeRateAllResponse>
  /** Queries a list of getEraExchangeRate items. */
  GetEraExchangeRate(request: QueryGetEraExchangeRateRequest): Promise<QueryGetEraExchangeRateResponse>
  /** Queries a list of eraExchangeRatesByDenom items. */
  EraExchangeRatesByDenom(request: QueryEraExchangeRatesByDenomRequest): Promise<QueryEraExchangeRatesByDenomResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  GetExchangeRate(request: QueryGetExchangeRateRequest): Promise<QueryGetExchangeRateResponse> {
    const data = QueryGetExchangeRateRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetExchangeRate', data)
    return promise.then((data) => QueryGetExchangeRateResponse.decode(new Reader(data)))
  }

  ExchangeRateAll(request: QueryExchangeRateAllRequest): Promise<QueryExchangeRateAllResponse> {
    const data = QueryExchangeRateAllRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'ExchangeRateAll', data)
    return promise.then((data) => QueryExchangeRateAllResponse.decode(new Reader(data)))
  }

  GetEraExchangeRate(request: QueryGetEraExchangeRateRequest): Promise<QueryGetEraExchangeRateResponse> {
    const data = QueryGetEraExchangeRateRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetEraExchangeRate', data)
    return promise.then((data) => QueryGetEraExchangeRateResponse.decode(new Reader(data)))
  }

  EraExchangeRatesByDenom(request: QueryEraExchangeRatesByDenomRequest): Promise<QueryEraExchangeRatesByDenomResponse> {
    const data = QueryEraExchangeRatesByDenomRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'EraExchangeRatesByDenom', data)
    return promise.then((data) => QueryEraExchangeRatesByDenomResponse.decode(new Reader(data)))
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
