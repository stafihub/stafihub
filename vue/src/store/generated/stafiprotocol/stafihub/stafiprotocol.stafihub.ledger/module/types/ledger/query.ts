/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import {
  ExchangeRate,
  EraExchangeRate,
  PoolDetail,
  UnbondFee,
  LeastBond,
  BondPipeline,
  BondSnapshot,
  PoolUnbond,
  AccountUnbond,
  BondRecord
} from '../ledger/ledger'

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

export interface QueryPoolsByDenomRequest {
  denom: string
}

export interface QueryPoolsByDenomResponse {
  addrs: string[]
}

export interface QueryBondedPoolsByDenomRequest {
  denom: string
}

export interface QueryBondedPoolsByDenomResponse {
  addrs: string[]
}

export interface QueryGetPoolDetailRequest {
  denom: string
  pool: string
}

export interface QueryGetPoolDetailResponse {
  detail: PoolDetail | undefined
}

export interface QueryGetChainEraRequest {
  denom: string
}

export interface QueryGetChainEraResponse {
  era: number
}

export interface QueryGetCurrentEraSnapshotRequest {
  denom: string
}

export interface QueryGetCurrentEraSnapshotResponse {
  shotIds: Uint8Array[]
}

export interface QueryGetReceiverRequest {}

export interface QueryGetReceiverResponse {
  receiver: string
}

export interface QueryGetCommissionRequest {}

export interface QueryGetCommissionResponse {
  commission: string
}

export interface QueryGetChainBondingDurationRequest {
  denom: string
}

export interface QueryGetChainBondingDurationResponse {
  era: number
}

export interface QueryGetUnbondFeeRequest {}

export interface QueryGetUnbondFeeResponse {
  fee: UnbondFee | undefined
}

export interface QueryGetUnbondCommissionRequest {}

export interface QueryGetUnbondCommissionResponse {
  commission: string
}

export interface QueryGetLeastBondRequest {
  denom: string
}

export interface QueryGetLeastBondResponse {
  leastBond: LeastBond | undefined
}

export interface QueryGetEraUnbondLimitRequest {
  denom: string
}

export interface QueryGetEraUnbondLimitResponse {
  limit: number
}

export interface QueryGetBondPipeLineRequest {
  denom: string
  pool: string
}

export interface QueryGetBondPipeLineResponse {
  pipeline: BondPipeline | undefined
}

export interface QueryGetEraSnapshotRequest {
  denom: string
  era: number
}

export interface QueryGetEraSnapshotResponse {
  shotIds: Uint8Array[]
}

export interface QueryGetSnapshotRequest {
  shotId: Uint8Array
}

export interface QueryGetSnapshotResponse {
  shot: BondSnapshot | undefined
}

export interface QueryGetTotalExpectedActiveRequest {
  denom: string
  era: number
}

export interface QueryGetTotalExpectedActiveResponse {
  active: string
}

export interface QueryGetPoolUnbondRequest {
  denom: string
  pool: string
  era: number
}

export interface QueryGetPoolUnbondResponse {
  unbond: PoolUnbond | undefined
}

export interface QueryGetAccountUnbondRequest {
  denom: string
  unbonder: string
}

export interface QueryGetAccountUnbondResponse {
  unbond: AccountUnbond | undefined
}

export interface QueryGetBondRecordRequest {
  denom: string
  blockhash: string
  txhash: string
}

export interface QueryGetBondRecordResponse {
  bondRecord: BondRecord | undefined
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

const baseQueryPoolsByDenomRequest: object = { denom: '' }

export const QueryPoolsByDenomRequest = {
  encode(message: QueryPoolsByDenomRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryPoolsByDenomRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryPoolsByDenomRequest } as QueryPoolsByDenomRequest
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

  fromJSON(object: any): QueryPoolsByDenomRequest {
    const message = { ...baseQueryPoolsByDenomRequest } as QueryPoolsByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryPoolsByDenomRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryPoolsByDenomRequest>): QueryPoolsByDenomRequest {
    const message = { ...baseQueryPoolsByDenomRequest } as QueryPoolsByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryPoolsByDenomResponse: object = { addrs: '' }

export const QueryPoolsByDenomResponse = {
  encode(message: QueryPoolsByDenomResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.addrs) {
      writer.uint32(10).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryPoolsByDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryPoolsByDenomResponse } as QueryPoolsByDenomResponse
    message.addrs = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.addrs.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryPoolsByDenomResponse {
    const message = { ...baseQueryPoolsByDenomResponse } as QueryPoolsByDenomResponse
    message.addrs = []
    if (object.addrs !== undefined && object.addrs !== null) {
      for (const e of object.addrs) {
        message.addrs.push(String(e))
      }
    }
    return message
  },

  toJSON(message: QueryPoolsByDenomResponse): unknown {
    const obj: any = {}
    if (message.addrs) {
      obj.addrs = message.addrs.map((e) => e)
    } else {
      obj.addrs = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<QueryPoolsByDenomResponse>): QueryPoolsByDenomResponse {
    const message = { ...baseQueryPoolsByDenomResponse } as QueryPoolsByDenomResponse
    message.addrs = []
    if (object.addrs !== undefined && object.addrs !== null) {
      for (const e of object.addrs) {
        message.addrs.push(e)
      }
    }
    return message
  }
}

const baseQueryBondedPoolsByDenomRequest: object = { denom: '' }

export const QueryBondedPoolsByDenomRequest = {
  encode(message: QueryBondedPoolsByDenomRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBondedPoolsByDenomRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryBondedPoolsByDenomRequest } as QueryBondedPoolsByDenomRequest
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

  fromJSON(object: any): QueryBondedPoolsByDenomRequest {
    const message = { ...baseQueryBondedPoolsByDenomRequest } as QueryBondedPoolsByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryBondedPoolsByDenomRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryBondedPoolsByDenomRequest>): QueryBondedPoolsByDenomRequest {
    const message = { ...baseQueryBondedPoolsByDenomRequest } as QueryBondedPoolsByDenomRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryBondedPoolsByDenomResponse: object = { addrs: '' }

export const QueryBondedPoolsByDenomResponse = {
  encode(message: QueryBondedPoolsByDenomResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.addrs) {
      writer.uint32(10).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBondedPoolsByDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryBondedPoolsByDenomResponse } as QueryBondedPoolsByDenomResponse
    message.addrs = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.addrs.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryBondedPoolsByDenomResponse {
    const message = { ...baseQueryBondedPoolsByDenomResponse } as QueryBondedPoolsByDenomResponse
    message.addrs = []
    if (object.addrs !== undefined && object.addrs !== null) {
      for (const e of object.addrs) {
        message.addrs.push(String(e))
      }
    }
    return message
  },

  toJSON(message: QueryBondedPoolsByDenomResponse): unknown {
    const obj: any = {}
    if (message.addrs) {
      obj.addrs = message.addrs.map((e) => e)
    } else {
      obj.addrs = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<QueryBondedPoolsByDenomResponse>): QueryBondedPoolsByDenomResponse {
    const message = { ...baseQueryBondedPoolsByDenomResponse } as QueryBondedPoolsByDenomResponse
    message.addrs = []
    if (object.addrs !== undefined && object.addrs !== null) {
      for (const e of object.addrs) {
        message.addrs.push(e)
      }
    }
    return message
  }
}

const baseQueryGetPoolDetailRequest: object = { denom: '', pool: '' }

export const QueryGetPoolDetailRequest = {
  encode(message: QueryGetPoolDetailRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.pool !== '') {
      writer.uint32(18).string(message.pool)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolDetailRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPoolDetailRequest } as QueryGetPoolDetailRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.pool = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPoolDetailRequest {
    const message = { ...baseQueryGetPoolDetailRequest } as QueryGetPoolDetailRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = String(object.pool)
    } else {
      message.pool = ''
    }
    return message
  },

  toJSON(message: QueryGetPoolDetailRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.pool !== undefined && (obj.pool = message.pool)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPoolDetailRequest>): QueryGetPoolDetailRequest {
    const message = { ...baseQueryGetPoolDetailRequest } as QueryGetPoolDetailRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = object.pool
    } else {
      message.pool = ''
    }
    return message
  }
}

const baseQueryGetPoolDetailResponse: object = {}

export const QueryGetPoolDetailResponse = {
  encode(message: QueryGetPoolDetailResponse, writer: Writer = Writer.create()): Writer {
    if (message.detail !== undefined) {
      PoolDetail.encode(message.detail, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolDetailResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPoolDetailResponse } as QueryGetPoolDetailResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.detail = PoolDetail.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPoolDetailResponse {
    const message = { ...baseQueryGetPoolDetailResponse } as QueryGetPoolDetailResponse
    if (object.detail !== undefined && object.detail !== null) {
      message.detail = PoolDetail.fromJSON(object.detail)
    } else {
      message.detail = undefined
    }
    return message
  },

  toJSON(message: QueryGetPoolDetailResponse): unknown {
    const obj: any = {}
    message.detail !== undefined && (obj.detail = message.detail ? PoolDetail.toJSON(message.detail) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPoolDetailResponse>): QueryGetPoolDetailResponse {
    const message = { ...baseQueryGetPoolDetailResponse } as QueryGetPoolDetailResponse
    if (object.detail !== undefined && object.detail !== null) {
      message.detail = PoolDetail.fromPartial(object.detail)
    } else {
      message.detail = undefined
    }
    return message
  }
}

const baseQueryGetChainEraRequest: object = { denom: '' }

export const QueryGetChainEraRequest = {
  encode(message: QueryGetChainEraRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainEraRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainEraRequest } as QueryGetChainEraRequest
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

  fromJSON(object: any): QueryGetChainEraRequest {
    const message = { ...baseQueryGetChainEraRequest } as QueryGetChainEraRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryGetChainEraRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainEraRequest>): QueryGetChainEraRequest {
    const message = { ...baseQueryGetChainEraRequest } as QueryGetChainEraRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryGetChainEraResponse: object = { era: 0 }

export const QueryGetChainEraResponse = {
  encode(message: QueryGetChainEraResponse, writer: Writer = Writer.create()): Writer {
    if (message.era !== 0) {
      writer.uint32(8).uint32(message.era)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainEraResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainEraResponse } as QueryGetChainEraResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.era = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetChainEraResponse {
    const message = { ...baseQueryGetChainEraResponse } as QueryGetChainEraResponse
    if (object.era !== undefined && object.era !== null) {
      message.era = Number(object.era)
    } else {
      message.era = 0
    }
    return message
  },

  toJSON(message: QueryGetChainEraResponse): unknown {
    const obj: any = {}
    message.era !== undefined && (obj.era = message.era)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainEraResponse>): QueryGetChainEraResponse {
    const message = { ...baseQueryGetChainEraResponse } as QueryGetChainEraResponse
    if (object.era !== undefined && object.era !== null) {
      message.era = object.era
    } else {
      message.era = 0
    }
    return message
  }
}

const baseQueryGetCurrentEraSnapshotRequest: object = { denom: '' }

export const QueryGetCurrentEraSnapshotRequest = {
  encode(message: QueryGetCurrentEraSnapshotRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCurrentEraSnapshotRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCurrentEraSnapshotRequest } as QueryGetCurrentEraSnapshotRequest
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

  fromJSON(object: any): QueryGetCurrentEraSnapshotRequest {
    const message = { ...baseQueryGetCurrentEraSnapshotRequest } as QueryGetCurrentEraSnapshotRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryGetCurrentEraSnapshotRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCurrentEraSnapshotRequest>): QueryGetCurrentEraSnapshotRequest {
    const message = { ...baseQueryGetCurrentEraSnapshotRequest } as QueryGetCurrentEraSnapshotRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryGetCurrentEraSnapshotResponse: object = {}

export const QueryGetCurrentEraSnapshotResponse = {
  encode(message: QueryGetCurrentEraSnapshotResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.shotIds) {
      writer.uint32(10).bytes(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCurrentEraSnapshotResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCurrentEraSnapshotResponse } as QueryGetCurrentEraSnapshotResponse
    message.shotIds = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.shotIds.push(reader.bytes())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetCurrentEraSnapshotResponse {
    const message = { ...baseQueryGetCurrentEraSnapshotResponse } as QueryGetCurrentEraSnapshotResponse
    message.shotIds = []
    if (object.shotIds !== undefined && object.shotIds !== null) {
      for (const e of object.shotIds) {
        message.shotIds.push(bytesFromBase64(e))
      }
    }
    return message
  },

  toJSON(message: QueryGetCurrentEraSnapshotResponse): unknown {
    const obj: any = {}
    if (message.shotIds) {
      obj.shotIds = message.shotIds.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()))
    } else {
      obj.shotIds = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCurrentEraSnapshotResponse>): QueryGetCurrentEraSnapshotResponse {
    const message = { ...baseQueryGetCurrentEraSnapshotResponse } as QueryGetCurrentEraSnapshotResponse
    message.shotIds = []
    if (object.shotIds !== undefined && object.shotIds !== null) {
      for (const e of object.shotIds) {
        message.shotIds.push(e)
      }
    }
    return message
  }
}

const baseQueryGetReceiverRequest: object = {}

export const QueryGetReceiverRequest = {
  encode(_: QueryGetReceiverRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetReceiverRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetReceiverRequest } as QueryGetReceiverRequest
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

  fromJSON(_: any): QueryGetReceiverRequest {
    const message = { ...baseQueryGetReceiverRequest } as QueryGetReceiverRequest
    return message
  },

  toJSON(_: QueryGetReceiverRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetReceiverRequest>): QueryGetReceiverRequest {
    const message = { ...baseQueryGetReceiverRequest } as QueryGetReceiverRequest
    return message
  }
}

const baseQueryGetReceiverResponse: object = { receiver: '' }

export const QueryGetReceiverResponse = {
  encode(message: QueryGetReceiverResponse, writer: Writer = Writer.create()): Writer {
    if (message.receiver !== '') {
      writer.uint32(10).string(message.receiver)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetReceiverResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetReceiverResponse } as QueryGetReceiverResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.receiver = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetReceiverResponse {
    const message = { ...baseQueryGetReceiverResponse } as QueryGetReceiverResponse
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver)
    } else {
      message.receiver = ''
    }
    return message
  },

  toJSON(message: QueryGetReceiverResponse): unknown {
    const obj: any = {}
    message.receiver !== undefined && (obj.receiver = message.receiver)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetReceiverResponse>): QueryGetReceiverResponse {
    const message = { ...baseQueryGetReceiverResponse } as QueryGetReceiverResponse
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver
    } else {
      message.receiver = ''
    }
    return message
  }
}

const baseQueryGetCommissionRequest: object = {}

export const QueryGetCommissionRequest = {
  encode(_: QueryGetCommissionRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCommissionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCommissionRequest } as QueryGetCommissionRequest
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

  fromJSON(_: any): QueryGetCommissionRequest {
    const message = { ...baseQueryGetCommissionRequest } as QueryGetCommissionRequest
    return message
  },

  toJSON(_: QueryGetCommissionRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetCommissionRequest>): QueryGetCommissionRequest {
    const message = { ...baseQueryGetCommissionRequest } as QueryGetCommissionRequest
    return message
  }
}

const baseQueryGetCommissionResponse: object = { commission: '' }

export const QueryGetCommissionResponse = {
  encode(message: QueryGetCommissionResponse, writer: Writer = Writer.create()): Writer {
    if (message.commission !== '') {
      writer.uint32(10).string(message.commission)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCommissionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCommissionResponse } as QueryGetCommissionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.commission = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetCommissionResponse {
    const message = { ...baseQueryGetCommissionResponse } as QueryGetCommissionResponse
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = String(object.commission)
    } else {
      message.commission = ''
    }
    return message
  },

  toJSON(message: QueryGetCommissionResponse): unknown {
    const obj: any = {}
    message.commission !== undefined && (obj.commission = message.commission)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCommissionResponse>): QueryGetCommissionResponse {
    const message = { ...baseQueryGetCommissionResponse } as QueryGetCommissionResponse
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = object.commission
    } else {
      message.commission = ''
    }
    return message
  }
}

const baseQueryGetChainBondingDurationRequest: object = { denom: '' }

export const QueryGetChainBondingDurationRequest = {
  encode(message: QueryGetChainBondingDurationRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainBondingDurationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainBondingDurationRequest } as QueryGetChainBondingDurationRequest
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

  fromJSON(object: any): QueryGetChainBondingDurationRequest {
    const message = { ...baseQueryGetChainBondingDurationRequest } as QueryGetChainBondingDurationRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryGetChainBondingDurationRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainBondingDurationRequest>): QueryGetChainBondingDurationRequest {
    const message = { ...baseQueryGetChainBondingDurationRequest } as QueryGetChainBondingDurationRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryGetChainBondingDurationResponse: object = { era: 0 }

export const QueryGetChainBondingDurationResponse = {
  encode(message: QueryGetChainBondingDurationResponse, writer: Writer = Writer.create()): Writer {
    if (message.era !== 0) {
      writer.uint32(8).uint32(message.era)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainBondingDurationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainBondingDurationResponse } as QueryGetChainBondingDurationResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.era = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetChainBondingDurationResponse {
    const message = { ...baseQueryGetChainBondingDurationResponse } as QueryGetChainBondingDurationResponse
    if (object.era !== undefined && object.era !== null) {
      message.era = Number(object.era)
    } else {
      message.era = 0
    }
    return message
  },

  toJSON(message: QueryGetChainBondingDurationResponse): unknown {
    const obj: any = {}
    message.era !== undefined && (obj.era = message.era)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainBondingDurationResponse>): QueryGetChainBondingDurationResponse {
    const message = { ...baseQueryGetChainBondingDurationResponse } as QueryGetChainBondingDurationResponse
    if (object.era !== undefined && object.era !== null) {
      message.era = object.era
    } else {
      message.era = 0
    }
    return message
  }
}

const baseQueryGetUnbondFeeRequest: object = {}

export const QueryGetUnbondFeeRequest = {
  encode(_: QueryGetUnbondFeeRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondFeeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUnbondFeeRequest } as QueryGetUnbondFeeRequest
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

  fromJSON(_: any): QueryGetUnbondFeeRequest {
    const message = { ...baseQueryGetUnbondFeeRequest } as QueryGetUnbondFeeRequest
    return message
  },

  toJSON(_: QueryGetUnbondFeeRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetUnbondFeeRequest>): QueryGetUnbondFeeRequest {
    const message = { ...baseQueryGetUnbondFeeRequest } as QueryGetUnbondFeeRequest
    return message
  }
}

const baseQueryGetUnbondFeeResponse: object = {}

export const QueryGetUnbondFeeResponse = {
  encode(message: QueryGetUnbondFeeResponse, writer: Writer = Writer.create()): Writer {
    if (message.fee !== undefined) {
      UnbondFee.encode(message.fee, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondFeeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUnbondFeeResponse } as QueryGetUnbondFeeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.fee = UnbondFee.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetUnbondFeeResponse {
    const message = { ...baseQueryGetUnbondFeeResponse } as QueryGetUnbondFeeResponse
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = UnbondFee.fromJSON(object.fee)
    } else {
      message.fee = undefined
    }
    return message
  },

  toJSON(message: QueryGetUnbondFeeResponse): unknown {
    const obj: any = {}
    message.fee !== undefined && (obj.fee = message.fee ? UnbondFee.toJSON(message.fee) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetUnbondFeeResponse>): QueryGetUnbondFeeResponse {
    const message = { ...baseQueryGetUnbondFeeResponse } as QueryGetUnbondFeeResponse
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = UnbondFee.fromPartial(object.fee)
    } else {
      message.fee = undefined
    }
    return message
  }
}

const baseQueryGetUnbondCommissionRequest: object = {}

export const QueryGetUnbondCommissionRequest = {
  encode(_: QueryGetUnbondCommissionRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondCommissionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUnbondCommissionRequest } as QueryGetUnbondCommissionRequest
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

  fromJSON(_: any): QueryGetUnbondCommissionRequest {
    const message = { ...baseQueryGetUnbondCommissionRequest } as QueryGetUnbondCommissionRequest
    return message
  },

  toJSON(_: QueryGetUnbondCommissionRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetUnbondCommissionRequest>): QueryGetUnbondCommissionRequest {
    const message = { ...baseQueryGetUnbondCommissionRequest } as QueryGetUnbondCommissionRequest
    return message
  }
}

const baseQueryGetUnbondCommissionResponse: object = { commission: '' }

export const QueryGetUnbondCommissionResponse = {
  encode(message: QueryGetUnbondCommissionResponse, writer: Writer = Writer.create()): Writer {
    if (message.commission !== '') {
      writer.uint32(10).string(message.commission)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondCommissionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetUnbondCommissionResponse } as QueryGetUnbondCommissionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.commission = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetUnbondCommissionResponse {
    const message = { ...baseQueryGetUnbondCommissionResponse } as QueryGetUnbondCommissionResponse
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = String(object.commission)
    } else {
      message.commission = ''
    }
    return message
  },

  toJSON(message: QueryGetUnbondCommissionResponse): unknown {
    const obj: any = {}
    message.commission !== undefined && (obj.commission = message.commission)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetUnbondCommissionResponse>): QueryGetUnbondCommissionResponse {
    const message = { ...baseQueryGetUnbondCommissionResponse } as QueryGetUnbondCommissionResponse
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = object.commission
    } else {
      message.commission = ''
    }
    return message
  }
}

const baseQueryGetLeastBondRequest: object = { denom: '' }

export const QueryGetLeastBondRequest = {
  encode(message: QueryGetLeastBondRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLeastBondRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetLeastBondRequest } as QueryGetLeastBondRequest
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

  fromJSON(object: any): QueryGetLeastBondRequest {
    const message = { ...baseQueryGetLeastBondRequest } as QueryGetLeastBondRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryGetLeastBondRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetLeastBondRequest>): QueryGetLeastBondRequest {
    const message = { ...baseQueryGetLeastBondRequest } as QueryGetLeastBondRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryGetLeastBondResponse: object = {}

export const QueryGetLeastBondResponse = {
  encode(message: QueryGetLeastBondResponse, writer: Writer = Writer.create()): Writer {
    if (message.leastBond !== undefined) {
      LeastBond.encode(message.leastBond, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLeastBondResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetLeastBondResponse } as QueryGetLeastBondResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.leastBond = LeastBond.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetLeastBondResponse {
    const message = { ...baseQueryGetLeastBondResponse } as QueryGetLeastBondResponse
    if (object.leastBond !== undefined && object.leastBond !== null) {
      message.leastBond = LeastBond.fromJSON(object.leastBond)
    } else {
      message.leastBond = undefined
    }
    return message
  },

  toJSON(message: QueryGetLeastBondResponse): unknown {
    const obj: any = {}
    message.leastBond !== undefined && (obj.leastBond = message.leastBond ? LeastBond.toJSON(message.leastBond) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetLeastBondResponse>): QueryGetLeastBondResponse {
    const message = { ...baseQueryGetLeastBondResponse } as QueryGetLeastBondResponse
    if (object.leastBond !== undefined && object.leastBond !== null) {
      message.leastBond = LeastBond.fromPartial(object.leastBond)
    } else {
      message.leastBond = undefined
    }
    return message
  }
}

const baseQueryGetEraUnbondLimitRequest: object = { denom: '' }

export const QueryGetEraUnbondLimitRequest = {
  encode(message: QueryGetEraUnbondLimitRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEraUnbondLimitRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetEraUnbondLimitRequest } as QueryGetEraUnbondLimitRequest
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

  fromJSON(object: any): QueryGetEraUnbondLimitRequest {
    const message = { ...baseQueryGetEraUnbondLimitRequest } as QueryGetEraUnbondLimitRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: QueryGetEraUnbondLimitRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetEraUnbondLimitRequest>): QueryGetEraUnbondLimitRequest {
    const message = { ...baseQueryGetEraUnbondLimitRequest } as QueryGetEraUnbondLimitRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseQueryGetEraUnbondLimitResponse: object = { limit: 0 }

export const QueryGetEraUnbondLimitResponse = {
  encode(message: QueryGetEraUnbondLimitResponse, writer: Writer = Writer.create()): Writer {
    if (message.limit !== 0) {
      writer.uint32(8).uint32(message.limit)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEraUnbondLimitResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetEraUnbondLimitResponse } as QueryGetEraUnbondLimitResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.limit = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetEraUnbondLimitResponse {
    const message = { ...baseQueryGetEraUnbondLimitResponse } as QueryGetEraUnbondLimitResponse
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = Number(object.limit)
    } else {
      message.limit = 0
    }
    return message
  },

  toJSON(message: QueryGetEraUnbondLimitResponse): unknown {
    const obj: any = {}
    message.limit !== undefined && (obj.limit = message.limit)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetEraUnbondLimitResponse>): QueryGetEraUnbondLimitResponse {
    const message = { ...baseQueryGetEraUnbondLimitResponse } as QueryGetEraUnbondLimitResponse
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = object.limit
    } else {
      message.limit = 0
    }
    return message
  }
}

const baseQueryGetBondPipeLineRequest: object = { denom: '', pool: '' }

export const QueryGetBondPipeLineRequest = {
  encode(message: QueryGetBondPipeLineRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.pool !== '') {
      writer.uint32(18).string(message.pool)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBondPipeLineRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetBondPipeLineRequest } as QueryGetBondPipeLineRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.pool = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetBondPipeLineRequest {
    const message = { ...baseQueryGetBondPipeLineRequest } as QueryGetBondPipeLineRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = String(object.pool)
    } else {
      message.pool = ''
    }
    return message
  },

  toJSON(message: QueryGetBondPipeLineRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.pool !== undefined && (obj.pool = message.pool)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetBondPipeLineRequest>): QueryGetBondPipeLineRequest {
    const message = { ...baseQueryGetBondPipeLineRequest } as QueryGetBondPipeLineRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = object.pool
    } else {
      message.pool = ''
    }
    return message
  }
}

const baseQueryGetBondPipeLineResponse: object = {}

export const QueryGetBondPipeLineResponse = {
  encode(message: QueryGetBondPipeLineResponse, writer: Writer = Writer.create()): Writer {
    if (message.pipeline !== undefined) {
      BondPipeline.encode(message.pipeline, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBondPipeLineResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetBondPipeLineResponse } as QueryGetBondPipeLineResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pipeline = BondPipeline.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetBondPipeLineResponse {
    const message = { ...baseQueryGetBondPipeLineResponse } as QueryGetBondPipeLineResponse
    if (object.pipeline !== undefined && object.pipeline !== null) {
      message.pipeline = BondPipeline.fromJSON(object.pipeline)
    } else {
      message.pipeline = undefined
    }
    return message
  },

  toJSON(message: QueryGetBondPipeLineResponse): unknown {
    const obj: any = {}
    message.pipeline !== undefined && (obj.pipeline = message.pipeline ? BondPipeline.toJSON(message.pipeline) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetBondPipeLineResponse>): QueryGetBondPipeLineResponse {
    const message = { ...baseQueryGetBondPipeLineResponse } as QueryGetBondPipeLineResponse
    if (object.pipeline !== undefined && object.pipeline !== null) {
      message.pipeline = BondPipeline.fromPartial(object.pipeline)
    } else {
      message.pipeline = undefined
    }
    return message
  }
}

const baseQueryGetEraSnapshotRequest: object = { denom: '', era: 0 }

export const QueryGetEraSnapshotRequest = {
  encode(message: QueryGetEraSnapshotRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.era !== 0) {
      writer.uint32(16).uint32(message.era)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEraSnapshotRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetEraSnapshotRequest } as QueryGetEraSnapshotRequest
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

  fromJSON(object: any): QueryGetEraSnapshotRequest {
    const message = { ...baseQueryGetEraSnapshotRequest } as QueryGetEraSnapshotRequest
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

  toJSON(message: QueryGetEraSnapshotRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.era !== undefined && (obj.era = message.era)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetEraSnapshotRequest>): QueryGetEraSnapshotRequest {
    const message = { ...baseQueryGetEraSnapshotRequest } as QueryGetEraSnapshotRequest
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

const baseQueryGetEraSnapshotResponse: object = {}

export const QueryGetEraSnapshotResponse = {
  encode(message: QueryGetEraSnapshotResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.shotIds) {
      writer.uint32(10).bytes(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEraSnapshotResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetEraSnapshotResponse } as QueryGetEraSnapshotResponse
    message.shotIds = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.shotIds.push(reader.bytes())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetEraSnapshotResponse {
    const message = { ...baseQueryGetEraSnapshotResponse } as QueryGetEraSnapshotResponse
    message.shotIds = []
    if (object.shotIds !== undefined && object.shotIds !== null) {
      for (const e of object.shotIds) {
        message.shotIds.push(bytesFromBase64(e))
      }
    }
    return message
  },

  toJSON(message: QueryGetEraSnapshotResponse): unknown {
    const obj: any = {}
    if (message.shotIds) {
      obj.shotIds = message.shotIds.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()))
    } else {
      obj.shotIds = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetEraSnapshotResponse>): QueryGetEraSnapshotResponse {
    const message = { ...baseQueryGetEraSnapshotResponse } as QueryGetEraSnapshotResponse
    message.shotIds = []
    if (object.shotIds !== undefined && object.shotIds !== null) {
      for (const e of object.shotIds) {
        message.shotIds.push(e)
      }
    }
    return message
  }
}

const baseQueryGetSnapshotRequest: object = {}

export const QueryGetSnapshotRequest = {
  encode(message: QueryGetSnapshotRequest, writer: Writer = Writer.create()): Writer {
    if (message.shotId.length !== 0) {
      writer.uint32(10).bytes(message.shotId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetSnapshotRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetSnapshotRequest } as QueryGetSnapshotRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.shotId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetSnapshotRequest {
    const message = { ...baseQueryGetSnapshotRequest } as QueryGetSnapshotRequest
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = bytesFromBase64(object.shotId)
    }
    return message
  },

  toJSON(message: QueryGetSnapshotRequest): unknown {
    const obj: any = {}
    message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetSnapshotRequest>): QueryGetSnapshotRequest {
    const message = { ...baseQueryGetSnapshotRequest } as QueryGetSnapshotRequest
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = object.shotId
    } else {
      message.shotId = new Uint8Array()
    }
    return message
  }
}

const baseQueryGetSnapshotResponse: object = {}

export const QueryGetSnapshotResponse = {
  encode(message: QueryGetSnapshotResponse, writer: Writer = Writer.create()): Writer {
    if (message.shot !== undefined) {
      BondSnapshot.encode(message.shot, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetSnapshotResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetSnapshotResponse } as QueryGetSnapshotResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.shot = BondSnapshot.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetSnapshotResponse {
    const message = { ...baseQueryGetSnapshotResponse } as QueryGetSnapshotResponse
    if (object.shot !== undefined && object.shot !== null) {
      message.shot = BondSnapshot.fromJSON(object.shot)
    } else {
      message.shot = undefined
    }
    return message
  },

  toJSON(message: QueryGetSnapshotResponse): unknown {
    const obj: any = {}
    message.shot !== undefined && (obj.shot = message.shot ? BondSnapshot.toJSON(message.shot) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetSnapshotResponse>): QueryGetSnapshotResponse {
    const message = { ...baseQueryGetSnapshotResponse } as QueryGetSnapshotResponse
    if (object.shot !== undefined && object.shot !== null) {
      message.shot = BondSnapshot.fromPartial(object.shot)
    } else {
      message.shot = undefined
    }
    return message
  }
}

const baseQueryGetTotalExpectedActiveRequest: object = { denom: '', era: 0 }

export const QueryGetTotalExpectedActiveRequest = {
  encode(message: QueryGetTotalExpectedActiveRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.era !== 0) {
      writer.uint32(16).uint32(message.era)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTotalExpectedActiveRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTotalExpectedActiveRequest } as QueryGetTotalExpectedActiveRequest
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

  fromJSON(object: any): QueryGetTotalExpectedActiveRequest {
    const message = { ...baseQueryGetTotalExpectedActiveRequest } as QueryGetTotalExpectedActiveRequest
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

  toJSON(message: QueryGetTotalExpectedActiveRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.era !== undefined && (obj.era = message.era)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetTotalExpectedActiveRequest>): QueryGetTotalExpectedActiveRequest {
    const message = { ...baseQueryGetTotalExpectedActiveRequest } as QueryGetTotalExpectedActiveRequest
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

const baseQueryGetTotalExpectedActiveResponse: object = { active: '' }

export const QueryGetTotalExpectedActiveResponse = {
  encode(message: QueryGetTotalExpectedActiveResponse, writer: Writer = Writer.create()): Writer {
    if (message.active !== '') {
      writer.uint32(10).string(message.active)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTotalExpectedActiveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTotalExpectedActiveResponse } as QueryGetTotalExpectedActiveResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.active = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetTotalExpectedActiveResponse {
    const message = { ...baseQueryGetTotalExpectedActiveResponse } as QueryGetTotalExpectedActiveResponse
    if (object.active !== undefined && object.active !== null) {
      message.active = String(object.active)
    } else {
      message.active = ''
    }
    return message
  },

  toJSON(message: QueryGetTotalExpectedActiveResponse): unknown {
    const obj: any = {}
    message.active !== undefined && (obj.active = message.active)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetTotalExpectedActiveResponse>): QueryGetTotalExpectedActiveResponse {
    const message = { ...baseQueryGetTotalExpectedActiveResponse } as QueryGetTotalExpectedActiveResponse
    if (object.active !== undefined && object.active !== null) {
      message.active = object.active
    } else {
      message.active = ''
    }
    return message
  }
}

const baseQueryGetPoolUnbondRequest: object = { denom: '', pool: '', era: 0 }

export const QueryGetPoolUnbondRequest = {
  encode(message: QueryGetPoolUnbondRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.pool !== '') {
      writer.uint32(18).string(message.pool)
    }
    if (message.era !== 0) {
      writer.uint32(24).uint32(message.era)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolUnbondRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPoolUnbondRequest } as QueryGetPoolUnbondRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.pool = reader.string()
          break
        case 3:
          message.era = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPoolUnbondRequest {
    const message = { ...baseQueryGetPoolUnbondRequest } as QueryGetPoolUnbondRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = String(object.pool)
    } else {
      message.pool = ''
    }
    if (object.era !== undefined && object.era !== null) {
      message.era = Number(object.era)
    } else {
      message.era = 0
    }
    return message
  },

  toJSON(message: QueryGetPoolUnbondRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.pool !== undefined && (obj.pool = message.pool)
    message.era !== undefined && (obj.era = message.era)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPoolUnbondRequest>): QueryGetPoolUnbondRequest {
    const message = { ...baseQueryGetPoolUnbondRequest } as QueryGetPoolUnbondRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = object.pool
    } else {
      message.pool = ''
    }
    if (object.era !== undefined && object.era !== null) {
      message.era = object.era
    } else {
      message.era = 0
    }
    return message
  }
}

const baseQueryGetPoolUnbondResponse: object = {}

export const QueryGetPoolUnbondResponse = {
  encode(message: QueryGetPoolUnbondResponse, writer: Writer = Writer.create()): Writer {
    if (message.unbond !== undefined) {
      PoolUnbond.encode(message.unbond, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolUnbondResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPoolUnbondResponse } as QueryGetPoolUnbondResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.unbond = PoolUnbond.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPoolUnbondResponse {
    const message = { ...baseQueryGetPoolUnbondResponse } as QueryGetPoolUnbondResponse
    if (object.unbond !== undefined && object.unbond !== null) {
      message.unbond = PoolUnbond.fromJSON(object.unbond)
    } else {
      message.unbond = undefined
    }
    return message
  },

  toJSON(message: QueryGetPoolUnbondResponse): unknown {
    const obj: any = {}
    message.unbond !== undefined && (obj.unbond = message.unbond ? PoolUnbond.toJSON(message.unbond) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPoolUnbondResponse>): QueryGetPoolUnbondResponse {
    const message = { ...baseQueryGetPoolUnbondResponse } as QueryGetPoolUnbondResponse
    if (object.unbond !== undefined && object.unbond !== null) {
      message.unbond = PoolUnbond.fromPartial(object.unbond)
    } else {
      message.unbond = undefined
    }
    return message
  }
}

const baseQueryGetAccountUnbondRequest: object = { denom: '', unbonder: '' }

export const QueryGetAccountUnbondRequest = {
  encode(message: QueryGetAccountUnbondRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.unbonder !== '') {
      writer.uint32(18).string(message.unbonder)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAccountUnbondRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAccountUnbondRequest } as QueryGetAccountUnbondRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.unbonder = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetAccountUnbondRequest {
    const message = { ...baseQueryGetAccountUnbondRequest } as QueryGetAccountUnbondRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.unbonder !== undefined && object.unbonder !== null) {
      message.unbonder = String(object.unbonder)
    } else {
      message.unbonder = ''
    }
    return message
  },

  toJSON(message: QueryGetAccountUnbondRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.unbonder !== undefined && (obj.unbonder = message.unbonder)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetAccountUnbondRequest>): QueryGetAccountUnbondRequest {
    const message = { ...baseQueryGetAccountUnbondRequest } as QueryGetAccountUnbondRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.unbonder !== undefined && object.unbonder !== null) {
      message.unbonder = object.unbonder
    } else {
      message.unbonder = ''
    }
    return message
  }
}

const baseQueryGetAccountUnbondResponse: object = {}

export const QueryGetAccountUnbondResponse = {
  encode(message: QueryGetAccountUnbondResponse, writer: Writer = Writer.create()): Writer {
    if (message.unbond !== undefined) {
      AccountUnbond.encode(message.unbond, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAccountUnbondResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAccountUnbondResponse } as QueryGetAccountUnbondResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.unbond = AccountUnbond.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetAccountUnbondResponse {
    const message = { ...baseQueryGetAccountUnbondResponse } as QueryGetAccountUnbondResponse
    if (object.unbond !== undefined && object.unbond !== null) {
      message.unbond = AccountUnbond.fromJSON(object.unbond)
    } else {
      message.unbond = undefined
    }
    return message
  },

  toJSON(message: QueryGetAccountUnbondResponse): unknown {
    const obj: any = {}
    message.unbond !== undefined && (obj.unbond = message.unbond ? AccountUnbond.toJSON(message.unbond) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetAccountUnbondResponse>): QueryGetAccountUnbondResponse {
    const message = { ...baseQueryGetAccountUnbondResponse } as QueryGetAccountUnbondResponse
    if (object.unbond !== undefined && object.unbond !== null) {
      message.unbond = AccountUnbond.fromPartial(object.unbond)
    } else {
      message.unbond = undefined
    }
    return message
  }
}

const baseQueryGetBondRecordRequest: object = { denom: '', blockhash: '', txhash: '' }

export const QueryGetBondRecordRequest = {
  encode(message: QueryGetBondRecordRequest, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.blockhash !== '') {
      writer.uint32(18).string(message.blockhash)
    }
    if (message.txhash !== '') {
      writer.uint32(26).string(message.txhash)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBondRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetBondRecordRequest } as QueryGetBondRecordRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.blockhash = reader.string()
          break
        case 3:
          message.txhash = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetBondRecordRequest {
    const message = { ...baseQueryGetBondRecordRequest } as QueryGetBondRecordRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.blockhash !== undefined && object.blockhash !== null) {
      message.blockhash = String(object.blockhash)
    } else {
      message.blockhash = ''
    }
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = String(object.txhash)
    } else {
      message.txhash = ''
    }
    return message
  },

  toJSON(message: QueryGetBondRecordRequest): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.blockhash !== undefined && (obj.blockhash = message.blockhash)
    message.txhash !== undefined && (obj.txhash = message.txhash)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetBondRecordRequest>): QueryGetBondRecordRequest {
    const message = { ...baseQueryGetBondRecordRequest } as QueryGetBondRecordRequest
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.blockhash !== undefined && object.blockhash !== null) {
      message.blockhash = object.blockhash
    } else {
      message.blockhash = ''
    }
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = object.txhash
    } else {
      message.txhash = ''
    }
    return message
  }
}

const baseQueryGetBondRecordResponse: object = {}

export const QueryGetBondRecordResponse = {
  encode(message: QueryGetBondRecordResponse, writer: Writer = Writer.create()): Writer {
    if (message.bondRecord !== undefined) {
      BondRecord.encode(message.bondRecord, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBondRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetBondRecordResponse } as QueryGetBondRecordResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.bondRecord = BondRecord.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetBondRecordResponse {
    const message = { ...baseQueryGetBondRecordResponse } as QueryGetBondRecordResponse
    if (object.bondRecord !== undefined && object.bondRecord !== null) {
      message.bondRecord = BondRecord.fromJSON(object.bondRecord)
    } else {
      message.bondRecord = undefined
    }
    return message
  },

  toJSON(message: QueryGetBondRecordResponse): unknown {
    const obj: any = {}
    message.bondRecord !== undefined && (obj.bondRecord = message.bondRecord ? BondRecord.toJSON(message.bondRecord) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetBondRecordResponse>): QueryGetBondRecordResponse {
    const message = { ...baseQueryGetBondRecordResponse } as QueryGetBondRecordResponse
    if (object.bondRecord !== undefined && object.bondRecord !== null) {
      message.bondRecord = BondRecord.fromPartial(object.bondRecord)
    } else {
      message.bondRecord = undefined
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
  /** Queries a list of poolsByDenom items. */
  PoolsByDenom(request: QueryPoolsByDenomRequest): Promise<QueryPoolsByDenomResponse>
  /** Queries a list of bondedPoolsByDenom items. */
  BondedPoolsByDenom(request: QueryBondedPoolsByDenomRequest): Promise<QueryBondedPoolsByDenomResponse>
  /** Queries a list of getPoolDetail items. */
  GetPoolDetail(request: QueryGetPoolDetailRequest): Promise<QueryGetPoolDetailResponse>
  /** Queries a list of getChainEra items. */
  GetChainEra(request: QueryGetChainEraRequest): Promise<QueryGetChainEraResponse>
  /** Queries a list of getCurrentEraSnapshot items. */
  GetCurrentEraSnapshot(request: QueryGetCurrentEraSnapshotRequest): Promise<QueryGetCurrentEraSnapshotResponse>
  /** Queries a list of getReceiver items. */
  GetReceiver(request: QueryGetReceiverRequest): Promise<QueryGetReceiverResponse>
  /** Queries a list of getCommission items. */
  GetCommission(request: QueryGetCommissionRequest): Promise<QueryGetCommissionResponse>
  /** Queries a list of getChainBondingDuration items. */
  GetChainBondingDuration(request: QueryGetChainBondingDurationRequest): Promise<QueryGetChainBondingDurationResponse>
  /** Queries a list of getUnbondFee items. */
  GetUnbondFee(request: QueryGetUnbondFeeRequest): Promise<QueryGetUnbondFeeResponse>
  /** Queries a list of getUnbondCommission items. */
  GetUnbondCommission(request: QueryGetUnbondCommissionRequest): Promise<QueryGetUnbondCommissionResponse>
  /** Queries a list of getLeastBond items. */
  GetLeastBond(request: QueryGetLeastBondRequest): Promise<QueryGetLeastBondResponse>
  /** Queries a list of getEraUnbondLimit items. */
  GetEraUnbondLimit(request: QueryGetEraUnbondLimitRequest): Promise<QueryGetEraUnbondLimitResponse>
  /** Queries a list of getBondPipeLine items. */
  GetBondPipeLine(request: QueryGetBondPipeLineRequest): Promise<QueryGetBondPipeLineResponse>
  /** Queries a list of getEraSnapshot items. */
  GetEraSnapshot(request: QueryGetEraSnapshotRequest): Promise<QueryGetEraSnapshotResponse>
  /** Queries a list of getSnapshot items. */
  GetSnapshot(request: QueryGetSnapshotRequest): Promise<QueryGetSnapshotResponse>
  /** Queries a list of getTotalExpectedActive items. */
  GetTotalExpectedActive(request: QueryGetTotalExpectedActiveRequest): Promise<QueryGetTotalExpectedActiveResponse>
  /** Queries a list of getPoolUnbond items. */
  GetPoolUnbond(request: QueryGetPoolUnbondRequest): Promise<QueryGetPoolUnbondResponse>
  /** Queries a list of getAccountUnbond items. */
  GetAccountUnbond(request: QueryGetAccountUnbondRequest): Promise<QueryGetAccountUnbondResponse>
  /** Queries a list of getBondRecord items. */
  GetBondRecord(request: QueryGetBondRecordRequest): Promise<QueryGetBondRecordResponse>
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

  PoolsByDenom(request: QueryPoolsByDenomRequest): Promise<QueryPoolsByDenomResponse> {
    const data = QueryPoolsByDenomRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'PoolsByDenom', data)
    return promise.then((data) => QueryPoolsByDenomResponse.decode(new Reader(data)))
  }

  BondedPoolsByDenom(request: QueryBondedPoolsByDenomRequest): Promise<QueryBondedPoolsByDenomResponse> {
    const data = QueryBondedPoolsByDenomRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'BondedPoolsByDenom', data)
    return promise.then((data) => QueryBondedPoolsByDenomResponse.decode(new Reader(data)))
  }

  GetPoolDetail(request: QueryGetPoolDetailRequest): Promise<QueryGetPoolDetailResponse> {
    const data = QueryGetPoolDetailRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetPoolDetail', data)
    return promise.then((data) => QueryGetPoolDetailResponse.decode(new Reader(data)))
  }

  GetChainEra(request: QueryGetChainEraRequest): Promise<QueryGetChainEraResponse> {
    const data = QueryGetChainEraRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetChainEra', data)
    return promise.then((data) => QueryGetChainEraResponse.decode(new Reader(data)))
  }

  GetCurrentEraSnapshot(request: QueryGetCurrentEraSnapshotRequest): Promise<QueryGetCurrentEraSnapshotResponse> {
    const data = QueryGetCurrentEraSnapshotRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetCurrentEraSnapshot', data)
    return promise.then((data) => QueryGetCurrentEraSnapshotResponse.decode(new Reader(data)))
  }

  GetReceiver(request: QueryGetReceiverRequest): Promise<QueryGetReceiverResponse> {
    const data = QueryGetReceiverRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetReceiver', data)
    return promise.then((data) => QueryGetReceiverResponse.decode(new Reader(data)))
  }

  GetCommission(request: QueryGetCommissionRequest): Promise<QueryGetCommissionResponse> {
    const data = QueryGetCommissionRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetCommission', data)
    return promise.then((data) => QueryGetCommissionResponse.decode(new Reader(data)))
  }

  GetChainBondingDuration(request: QueryGetChainBondingDurationRequest): Promise<QueryGetChainBondingDurationResponse> {
    const data = QueryGetChainBondingDurationRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetChainBondingDuration', data)
    return promise.then((data) => QueryGetChainBondingDurationResponse.decode(new Reader(data)))
  }

  GetUnbondFee(request: QueryGetUnbondFeeRequest): Promise<QueryGetUnbondFeeResponse> {
    const data = QueryGetUnbondFeeRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetUnbondFee', data)
    return promise.then((data) => QueryGetUnbondFeeResponse.decode(new Reader(data)))
  }

  GetUnbondCommission(request: QueryGetUnbondCommissionRequest): Promise<QueryGetUnbondCommissionResponse> {
    const data = QueryGetUnbondCommissionRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetUnbondCommission', data)
    return promise.then((data) => QueryGetUnbondCommissionResponse.decode(new Reader(data)))
  }

  GetLeastBond(request: QueryGetLeastBondRequest): Promise<QueryGetLeastBondResponse> {
    const data = QueryGetLeastBondRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetLeastBond', data)
    return promise.then((data) => QueryGetLeastBondResponse.decode(new Reader(data)))
  }

  GetEraUnbondLimit(request: QueryGetEraUnbondLimitRequest): Promise<QueryGetEraUnbondLimitResponse> {
    const data = QueryGetEraUnbondLimitRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetEraUnbondLimit', data)
    return promise.then((data) => QueryGetEraUnbondLimitResponse.decode(new Reader(data)))
  }

  GetBondPipeLine(request: QueryGetBondPipeLineRequest): Promise<QueryGetBondPipeLineResponse> {
    const data = QueryGetBondPipeLineRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetBondPipeLine', data)
    return promise.then((data) => QueryGetBondPipeLineResponse.decode(new Reader(data)))
  }

  GetEraSnapshot(request: QueryGetEraSnapshotRequest): Promise<QueryGetEraSnapshotResponse> {
    const data = QueryGetEraSnapshotRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetEraSnapshot', data)
    return promise.then((data) => QueryGetEraSnapshotResponse.decode(new Reader(data)))
  }

  GetSnapshot(request: QueryGetSnapshotRequest): Promise<QueryGetSnapshotResponse> {
    const data = QueryGetSnapshotRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetSnapshot', data)
    return promise.then((data) => QueryGetSnapshotResponse.decode(new Reader(data)))
  }

  GetTotalExpectedActive(request: QueryGetTotalExpectedActiveRequest): Promise<QueryGetTotalExpectedActiveResponse> {
    const data = QueryGetTotalExpectedActiveRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetTotalExpectedActive', data)
    return promise.then((data) => QueryGetTotalExpectedActiveResponse.decode(new Reader(data)))
  }

  GetPoolUnbond(request: QueryGetPoolUnbondRequest): Promise<QueryGetPoolUnbondResponse> {
    const data = QueryGetPoolUnbondRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetPoolUnbond', data)
    return promise.then((data) => QueryGetPoolUnbondResponse.decode(new Reader(data)))
  }

  GetAccountUnbond(request: QueryGetAccountUnbondRequest): Promise<QueryGetAccountUnbondResponse> {
    const data = QueryGetAccountUnbondRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetAccountUnbond', data)
    return promise.then((data) => QueryGetAccountUnbondResponse.decode(new Reader(data)))
  }

  GetBondRecord(request: QueryGetBondRecordRequest): Promise<QueryGetBondRecordResponse> {
    const data = QueryGetBondRecordRequest.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetBondRecord', data)
    return promise.then((data) => QueryGetBondRecordResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

const atob: (b64: string) => string = globalThis.atob || ((b64) => globalThis.Buffer.from(b64, 'base64').toString('binary'))
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64)
  const arr = new Uint8Array(bin.length)
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i)
  }
  return arr
}

const btoa: (bin: string) => string = globalThis.btoa || ((bin) => globalThis.Buffer.from(bin, 'binary').toString('base64'))
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = []
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]))
  }
  return btoa(bin.join(''))
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
