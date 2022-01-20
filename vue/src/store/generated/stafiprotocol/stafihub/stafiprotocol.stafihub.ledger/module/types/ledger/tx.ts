/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.ledger'

export interface MsgAddNewPool {
  creator: string
  denom: string
  addr: string
}

export interface MsgAddNewPoolResponse {}

export interface MsgRemovePool {
  creator: string
  denom: string
  addr: string
}

export interface MsgRemovePoolResponse {}

export interface MsgSetEraUnbondLimit {
  creator: string
  denom: string
  limit: number
}

export interface MsgSetEraUnbondLimitResponse {}

export interface MsgSetInitBond {
  creator: string
  denom: string
  pool: string
  amount: string
  receiver: string
}

export interface MsgSetInitBondResponse {}

export interface MsgSetChainBondingDuration {
  creator: string
  denom: string
  era: number
}

export interface MsgSetChainBondingDurationResponse {}

export interface MsgSetPoolDetail {
  creator: string
  denom: string
  pool: string
  subAccounts: string[]
  threshold: number
}

export interface MsgSetPoolDetailResponse {}

export interface MsgSetLeastBond {
  creator: string
  denom: string
  amount: string
}

export interface MsgSetLeastBondResponse {}

export interface MsgClearCurrentEraSnapShots {
  creator: string
  denom: string
}

export interface MsgClearCurrentEraSnapShotsResponse {}

export interface MsgSetCommission {
  creator: string
  commission: string
}

export interface MsgSetCommissionResponse {}

export interface MsgSetReceiver {
  creator: string
  receiver: string
}

export interface MsgSetReceiverResponse {}

export interface MsgSetUnbondFee {
  creator: string
  value: string
}

export interface MsgSetUnbondFeeResponse {}

export interface MsgLiquidityUnbond {
  creator: string
  pool: string
  value: string
  recipient: string
}

export interface MsgLiquidityUnbondResponse {}

export interface MsgSetUnbondCommission {
  creator: string
  commission: string
}

export interface MsgSetUnbondCommissionResponse {}

const baseMsgAddNewPool: object = { creator: '', denom: '', addr: '' }

export const MsgAddNewPool = {
  encode(message: MsgAddNewPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.addr !== '') {
      writer.uint32(26).string(message.addr)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddNewPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddNewPool } as MsgAddNewPool
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.addr = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgAddNewPool {
    const message = { ...baseMsgAddNewPool } as MsgAddNewPool
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.addr !== undefined && object.addr !== null) {
      message.addr = String(object.addr)
    } else {
      message.addr = ''
    }
    return message
  },

  toJSON(message: MsgAddNewPool): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.addr !== undefined && (obj.addr = message.addr)
    return obj
  },

  fromPartial(object: DeepPartial<MsgAddNewPool>): MsgAddNewPool {
    const message = { ...baseMsgAddNewPool } as MsgAddNewPool
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.addr !== undefined && object.addr !== null) {
      message.addr = object.addr
    } else {
      message.addr = ''
    }
    return message
  }
}

const baseMsgAddNewPoolResponse: object = {}

export const MsgAddNewPoolResponse = {
  encode(_: MsgAddNewPoolResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddNewPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddNewPoolResponse } as MsgAddNewPoolResponse
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

  fromJSON(_: any): MsgAddNewPoolResponse {
    const message = { ...baseMsgAddNewPoolResponse } as MsgAddNewPoolResponse
    return message
  },

  toJSON(_: MsgAddNewPoolResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgAddNewPoolResponse>): MsgAddNewPoolResponse {
    const message = { ...baseMsgAddNewPoolResponse } as MsgAddNewPoolResponse
    return message
  }
}

const baseMsgRemovePool: object = { creator: '', denom: '', addr: '' }

export const MsgRemovePool = {
  encode(message: MsgRemovePool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.addr !== '') {
      writer.uint32(26).string(message.addr)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemovePool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRemovePool } as MsgRemovePool
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.addr = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgRemovePool {
    const message = { ...baseMsgRemovePool } as MsgRemovePool
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.addr !== undefined && object.addr !== null) {
      message.addr = String(object.addr)
    } else {
      message.addr = ''
    }
    return message
  },

  toJSON(message: MsgRemovePool): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.addr !== undefined && (obj.addr = message.addr)
    return obj
  },

  fromPartial(object: DeepPartial<MsgRemovePool>): MsgRemovePool {
    const message = { ...baseMsgRemovePool } as MsgRemovePool
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.addr !== undefined && object.addr !== null) {
      message.addr = object.addr
    } else {
      message.addr = ''
    }
    return message
  }
}

const baseMsgRemovePoolResponse: object = {}

export const MsgRemovePoolResponse = {
  encode(_: MsgRemovePoolResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemovePoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRemovePoolResponse } as MsgRemovePoolResponse
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

  fromJSON(_: any): MsgRemovePoolResponse {
    const message = { ...baseMsgRemovePoolResponse } as MsgRemovePoolResponse
    return message
  },

  toJSON(_: MsgRemovePoolResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgRemovePoolResponse>): MsgRemovePoolResponse {
    const message = { ...baseMsgRemovePoolResponse } as MsgRemovePoolResponse
    return message
  }
}

const baseMsgSetEraUnbondLimit: object = { creator: '', denom: '', limit: 0 }

export const MsgSetEraUnbondLimit = {
  encode(message: MsgSetEraUnbondLimit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.limit !== 0) {
      writer.uint32(24).uint32(message.limit)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetEraUnbondLimit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetEraUnbondLimit } as MsgSetEraUnbondLimit
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.limit = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetEraUnbondLimit {
    const message = { ...baseMsgSetEraUnbondLimit } as MsgSetEraUnbondLimit
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = Number(object.limit)
    } else {
      message.limit = 0
    }
    return message
  },

  toJSON(message: MsgSetEraUnbondLimit): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.limit !== undefined && (obj.limit = message.limit)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetEraUnbondLimit>): MsgSetEraUnbondLimit {
    const message = { ...baseMsgSetEraUnbondLimit } as MsgSetEraUnbondLimit
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = object.limit
    } else {
      message.limit = 0
    }
    return message
  }
}

const baseMsgSetEraUnbondLimitResponse: object = {}

export const MsgSetEraUnbondLimitResponse = {
  encode(_: MsgSetEraUnbondLimitResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetEraUnbondLimitResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetEraUnbondLimitResponse } as MsgSetEraUnbondLimitResponse
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

  fromJSON(_: any): MsgSetEraUnbondLimitResponse {
    const message = { ...baseMsgSetEraUnbondLimitResponse } as MsgSetEraUnbondLimitResponse
    return message
  },

  toJSON(_: MsgSetEraUnbondLimitResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetEraUnbondLimitResponse>): MsgSetEraUnbondLimitResponse {
    const message = { ...baseMsgSetEraUnbondLimitResponse } as MsgSetEraUnbondLimitResponse
    return message
  }
}

const baseMsgSetInitBond: object = { creator: '', denom: '', pool: '', amount: '', receiver: '' }

export const MsgSetInitBond = {
  encode(message: MsgSetInitBond, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.pool !== '') {
      writer.uint32(26).string(message.pool)
    }
    if (message.amount !== '') {
      writer.uint32(34).string(message.amount)
    }
    if (message.receiver !== '') {
      writer.uint32(42).string(message.receiver)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetInitBond {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetInitBond } as MsgSetInitBond
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.pool = reader.string()
          break
        case 4:
          message.amount = reader.string()
          break
        case 5:
          message.receiver = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetInitBond {
    const message = { ...baseMsgSetInitBond } as MsgSetInitBond
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
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
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount)
    } else {
      message.amount = ''
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver)
    } else {
      message.receiver = ''
    }
    return message
  },

  toJSON(message: MsgSetInitBond): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.pool !== undefined && (obj.pool = message.pool)
    message.amount !== undefined && (obj.amount = message.amount)
    message.receiver !== undefined && (obj.receiver = message.receiver)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetInitBond>): MsgSetInitBond {
    const message = { ...baseMsgSetInitBond } as MsgSetInitBond
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
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
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = ''
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver
    } else {
      message.receiver = ''
    }
    return message
  }
}

const baseMsgSetInitBondResponse: object = {}

export const MsgSetInitBondResponse = {
  encode(_: MsgSetInitBondResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetInitBondResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetInitBondResponse } as MsgSetInitBondResponse
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

  fromJSON(_: any): MsgSetInitBondResponse {
    const message = { ...baseMsgSetInitBondResponse } as MsgSetInitBondResponse
    return message
  },

  toJSON(_: MsgSetInitBondResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetInitBondResponse>): MsgSetInitBondResponse {
    const message = { ...baseMsgSetInitBondResponse } as MsgSetInitBondResponse
    return message
  }
}

const baseMsgSetChainBondingDuration: object = { creator: '', denom: '', era: 0 }

export const MsgSetChainBondingDuration = {
  encode(message: MsgSetChainBondingDuration, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.era !== 0) {
      writer.uint32(24).uint32(message.era)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetChainBondingDuration {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetChainBondingDuration } as MsgSetChainBondingDuration
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
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

  fromJSON(object: any): MsgSetChainBondingDuration {
    const message = { ...baseMsgSetChainBondingDuration } as MsgSetChainBondingDuration
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
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

  toJSON(message: MsgSetChainBondingDuration): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.era !== undefined && (obj.era = message.era)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetChainBondingDuration>): MsgSetChainBondingDuration {
    const message = { ...baseMsgSetChainBondingDuration } as MsgSetChainBondingDuration
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
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

const baseMsgSetChainBondingDurationResponse: object = {}

export const MsgSetChainBondingDurationResponse = {
  encode(_: MsgSetChainBondingDurationResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetChainBondingDurationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetChainBondingDurationResponse } as MsgSetChainBondingDurationResponse
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

  fromJSON(_: any): MsgSetChainBondingDurationResponse {
    const message = { ...baseMsgSetChainBondingDurationResponse } as MsgSetChainBondingDurationResponse
    return message
  },

  toJSON(_: MsgSetChainBondingDurationResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetChainBondingDurationResponse>): MsgSetChainBondingDurationResponse {
    const message = { ...baseMsgSetChainBondingDurationResponse } as MsgSetChainBondingDurationResponse
    return message
  }
}

const baseMsgSetPoolDetail: object = { creator: '', denom: '', pool: '', subAccounts: '', threshold: 0 }

export const MsgSetPoolDetail = {
  encode(message: MsgSetPoolDetail, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.pool !== '') {
      writer.uint32(26).string(message.pool)
    }
    for (const v of message.subAccounts) {
      writer.uint32(34).string(v!)
    }
    if (message.threshold !== 0) {
      writer.uint32(40).uint32(message.threshold)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetPoolDetail {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetPoolDetail } as MsgSetPoolDetail
    message.subAccounts = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.pool = reader.string()
          break
        case 4:
          message.subAccounts.push(reader.string())
          break
        case 5:
          message.threshold = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetPoolDetail {
    const message = { ...baseMsgSetPoolDetail } as MsgSetPoolDetail
    message.subAccounts = []
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
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
    if (object.subAccounts !== undefined && object.subAccounts !== null) {
      for (const e of object.subAccounts) {
        message.subAccounts.push(String(e))
      }
    }
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = Number(object.threshold)
    } else {
      message.threshold = 0
    }
    return message
  },

  toJSON(message: MsgSetPoolDetail): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.pool !== undefined && (obj.pool = message.pool)
    if (message.subAccounts) {
      obj.subAccounts = message.subAccounts.map((e) => e)
    } else {
      obj.subAccounts = []
    }
    message.threshold !== undefined && (obj.threshold = message.threshold)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetPoolDetail>): MsgSetPoolDetail {
    const message = { ...baseMsgSetPoolDetail } as MsgSetPoolDetail
    message.subAccounts = []
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
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
    if (object.subAccounts !== undefined && object.subAccounts !== null) {
      for (const e of object.subAccounts) {
        message.subAccounts.push(e)
      }
    }
    if (object.threshold !== undefined && object.threshold !== null) {
      message.threshold = object.threshold
    } else {
      message.threshold = 0
    }
    return message
  }
}

const baseMsgSetPoolDetailResponse: object = {}

export const MsgSetPoolDetailResponse = {
  encode(_: MsgSetPoolDetailResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetPoolDetailResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetPoolDetailResponse } as MsgSetPoolDetailResponse
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

  fromJSON(_: any): MsgSetPoolDetailResponse {
    const message = { ...baseMsgSetPoolDetailResponse } as MsgSetPoolDetailResponse
    return message
  },

  toJSON(_: MsgSetPoolDetailResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetPoolDetailResponse>): MsgSetPoolDetailResponse {
    const message = { ...baseMsgSetPoolDetailResponse } as MsgSetPoolDetailResponse
    return message
  }
}

const baseMsgSetLeastBond: object = { creator: '', denom: '', amount: '' }

export const MsgSetLeastBond = {
  encode(message: MsgSetLeastBond, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.amount !== '') {
      writer.uint32(26).string(message.amount)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetLeastBond {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetLeastBond } as MsgSetLeastBond
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.amount = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetLeastBond {
    const message = { ...baseMsgSetLeastBond } as MsgSetLeastBond
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount)
    } else {
      message.amount = ''
    }
    return message
  },

  toJSON(message: MsgSetLeastBond): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.amount !== undefined && (obj.amount = message.amount)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetLeastBond>): MsgSetLeastBond {
    const message = { ...baseMsgSetLeastBond } as MsgSetLeastBond
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = ''
    }
    return message
  }
}

const baseMsgSetLeastBondResponse: object = {}

export const MsgSetLeastBondResponse = {
  encode(_: MsgSetLeastBondResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetLeastBondResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetLeastBondResponse } as MsgSetLeastBondResponse
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

  fromJSON(_: any): MsgSetLeastBondResponse {
    const message = { ...baseMsgSetLeastBondResponse } as MsgSetLeastBondResponse
    return message
  },

  toJSON(_: MsgSetLeastBondResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetLeastBondResponse>): MsgSetLeastBondResponse {
    const message = { ...baseMsgSetLeastBondResponse } as MsgSetLeastBondResponse
    return message
  }
}

const baseMsgClearCurrentEraSnapShots: object = { creator: '', denom: '' }

export const MsgClearCurrentEraSnapShots = {
  encode(message: MsgClearCurrentEraSnapShots, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClearCurrentEraSnapShots {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgClearCurrentEraSnapShots } as MsgClearCurrentEraSnapShots
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgClearCurrentEraSnapShots {
    const message = { ...baseMsgClearCurrentEraSnapShots } as MsgClearCurrentEraSnapShots
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    return message
  },

  toJSON(message: MsgClearCurrentEraSnapShots): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    return obj
  },

  fromPartial(object: DeepPartial<MsgClearCurrentEraSnapShots>): MsgClearCurrentEraSnapShots {
    const message = { ...baseMsgClearCurrentEraSnapShots } as MsgClearCurrentEraSnapShots
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    return message
  }
}

const baseMsgClearCurrentEraSnapShotsResponse: object = {}

export const MsgClearCurrentEraSnapShotsResponse = {
  encode(_: MsgClearCurrentEraSnapShotsResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClearCurrentEraSnapShotsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgClearCurrentEraSnapShotsResponse } as MsgClearCurrentEraSnapShotsResponse
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

  fromJSON(_: any): MsgClearCurrentEraSnapShotsResponse {
    const message = { ...baseMsgClearCurrentEraSnapShotsResponse } as MsgClearCurrentEraSnapShotsResponse
    return message
  },

  toJSON(_: MsgClearCurrentEraSnapShotsResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgClearCurrentEraSnapShotsResponse>): MsgClearCurrentEraSnapShotsResponse {
    const message = { ...baseMsgClearCurrentEraSnapShotsResponse } as MsgClearCurrentEraSnapShotsResponse
    return message
  }
}

const baseMsgSetCommission: object = { creator: '', commission: '' }

export const MsgSetCommission = {
  encode(message: MsgSetCommission, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.commission !== '') {
      writer.uint32(18).string(message.commission)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetCommission {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetCommission } as MsgSetCommission
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.commission = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetCommission {
    const message = { ...baseMsgSetCommission } as MsgSetCommission
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = String(object.commission)
    } else {
      message.commission = ''
    }
    return message
  },

  toJSON(message: MsgSetCommission): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.commission !== undefined && (obj.commission = message.commission)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetCommission>): MsgSetCommission {
    const message = { ...baseMsgSetCommission } as MsgSetCommission
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = object.commission
    } else {
      message.commission = ''
    }
    return message
  }
}

const baseMsgSetCommissionResponse: object = {}

export const MsgSetCommissionResponse = {
  encode(_: MsgSetCommissionResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetCommissionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetCommissionResponse } as MsgSetCommissionResponse
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

  fromJSON(_: any): MsgSetCommissionResponse {
    const message = { ...baseMsgSetCommissionResponse } as MsgSetCommissionResponse
    return message
  },

  toJSON(_: MsgSetCommissionResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetCommissionResponse>): MsgSetCommissionResponse {
    const message = { ...baseMsgSetCommissionResponse } as MsgSetCommissionResponse
    return message
  }
}

const baseMsgSetReceiver: object = { creator: '', receiver: '' }

export const MsgSetReceiver = {
  encode(message: MsgSetReceiver, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.receiver !== '') {
      writer.uint32(18).string(message.receiver)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetReceiver {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetReceiver } as MsgSetReceiver
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.receiver = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetReceiver {
    const message = { ...baseMsgSetReceiver } as MsgSetReceiver
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver)
    } else {
      message.receiver = ''
    }
    return message
  },

  toJSON(message: MsgSetReceiver): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.receiver !== undefined && (obj.receiver = message.receiver)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetReceiver>): MsgSetReceiver {
    const message = { ...baseMsgSetReceiver } as MsgSetReceiver
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver
    } else {
      message.receiver = ''
    }
    return message
  }
}

const baseMsgSetReceiverResponse: object = {}

export const MsgSetReceiverResponse = {
  encode(_: MsgSetReceiverResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetReceiverResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetReceiverResponse } as MsgSetReceiverResponse
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

  fromJSON(_: any): MsgSetReceiverResponse {
    const message = { ...baseMsgSetReceiverResponse } as MsgSetReceiverResponse
    return message
  },

  toJSON(_: MsgSetReceiverResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetReceiverResponse>): MsgSetReceiverResponse {
    const message = { ...baseMsgSetReceiverResponse } as MsgSetReceiverResponse
    return message
  }
}

const baseMsgSetUnbondFee: object = { creator: '', value: '' }

export const MsgSetUnbondFee = {
  encode(message: MsgSetUnbondFee, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.value !== '') {
      writer.uint32(18).string(message.value)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondFee {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetUnbondFee } as MsgSetUnbondFee
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.value = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetUnbondFee {
    const message = { ...baseMsgSetUnbondFee } as MsgSetUnbondFee
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value)
    } else {
      message.value = ''
    }
    return message
  },

  toJSON(message: MsgSetUnbondFee): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.value !== undefined && (obj.value = message.value)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetUnbondFee>): MsgSetUnbondFee {
    const message = { ...baseMsgSetUnbondFee } as MsgSetUnbondFee
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value
    } else {
      message.value = ''
    }
    return message
  }
}

const baseMsgSetUnbondFeeResponse: object = {}

export const MsgSetUnbondFeeResponse = {
  encode(_: MsgSetUnbondFeeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondFeeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetUnbondFeeResponse } as MsgSetUnbondFeeResponse
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

  fromJSON(_: any): MsgSetUnbondFeeResponse {
    const message = { ...baseMsgSetUnbondFeeResponse } as MsgSetUnbondFeeResponse
    return message
  },

  toJSON(_: MsgSetUnbondFeeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetUnbondFeeResponse>): MsgSetUnbondFeeResponse {
    const message = { ...baseMsgSetUnbondFeeResponse } as MsgSetUnbondFeeResponse
    return message
  }
}

const baseMsgLiquidityUnbond: object = { creator: '', pool: '', value: '', recipient: '' }

export const MsgLiquidityUnbond = {
  encode(message: MsgLiquidityUnbond, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.pool !== '') {
      writer.uint32(18).string(message.pool)
    }
    if (message.value !== '') {
      writer.uint32(26).string(message.value)
    }
    if (message.recipient !== '') {
      writer.uint32(34).string(message.recipient)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLiquidityUnbond {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgLiquidityUnbond } as MsgLiquidityUnbond
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.pool = reader.string()
          break
        case 3:
          message.value = reader.string()
          break
        case 4:
          message.recipient = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgLiquidityUnbond {
    const message = { ...baseMsgLiquidityUnbond } as MsgLiquidityUnbond
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = String(object.pool)
    } else {
      message.pool = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value)
    } else {
      message.value = ''
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = String(object.recipient)
    } else {
      message.recipient = ''
    }
    return message
  },

  toJSON(message: MsgLiquidityUnbond): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.pool !== undefined && (obj.pool = message.pool)
    message.value !== undefined && (obj.value = message.value)
    message.recipient !== undefined && (obj.recipient = message.recipient)
    return obj
  },

  fromPartial(object: DeepPartial<MsgLiquidityUnbond>): MsgLiquidityUnbond {
    const message = { ...baseMsgLiquidityUnbond } as MsgLiquidityUnbond
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = object.pool
    } else {
      message.pool = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value
    } else {
      message.value = ''
    }
    if (object.recipient !== undefined && object.recipient !== null) {
      message.recipient = object.recipient
    } else {
      message.recipient = ''
    }
    return message
  }
}

const baseMsgLiquidityUnbondResponse: object = {}

export const MsgLiquidityUnbondResponse = {
  encode(_: MsgLiquidityUnbondResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLiquidityUnbondResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgLiquidityUnbondResponse } as MsgLiquidityUnbondResponse
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

  fromJSON(_: any): MsgLiquidityUnbondResponse {
    const message = { ...baseMsgLiquidityUnbondResponse } as MsgLiquidityUnbondResponse
    return message
  },

  toJSON(_: MsgLiquidityUnbondResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgLiquidityUnbondResponse>): MsgLiquidityUnbondResponse {
    const message = { ...baseMsgLiquidityUnbondResponse } as MsgLiquidityUnbondResponse
    return message
  }
}

const baseMsgSetUnbondCommission: object = { creator: '', commission: '' }

export const MsgSetUnbondCommission = {
  encode(message: MsgSetUnbondCommission, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.commission !== '') {
      writer.uint32(18).string(message.commission)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondCommission {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetUnbondCommission } as MsgSetUnbondCommission
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.commission = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetUnbondCommission {
    const message = { ...baseMsgSetUnbondCommission } as MsgSetUnbondCommission
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = String(object.commission)
    } else {
      message.commission = ''
    }
    return message
  },

  toJSON(message: MsgSetUnbondCommission): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.commission !== undefined && (obj.commission = message.commission)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetUnbondCommission>): MsgSetUnbondCommission {
    const message = { ...baseMsgSetUnbondCommission } as MsgSetUnbondCommission
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = object.commission
    } else {
      message.commission = ''
    }
    return message
  }
}

const baseMsgSetUnbondCommissionResponse: object = {}

export const MsgSetUnbondCommissionResponse = {
  encode(_: MsgSetUnbondCommissionResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondCommissionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetUnbondCommissionResponse } as MsgSetUnbondCommissionResponse
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

  fromJSON(_: any): MsgSetUnbondCommissionResponse {
    const message = { ...baseMsgSetUnbondCommissionResponse } as MsgSetUnbondCommissionResponse
    return message
  },

  toJSON(_: MsgSetUnbondCommissionResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetUnbondCommissionResponse>): MsgSetUnbondCommissionResponse {
    const message = { ...baseMsgSetUnbondCommissionResponse } as MsgSetUnbondCommissionResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  AddNewPool(request: MsgAddNewPool): Promise<MsgAddNewPoolResponse>
  RemovePool(request: MsgRemovePool): Promise<MsgRemovePoolResponse>
  SetEraUnbondLimit(request: MsgSetEraUnbondLimit): Promise<MsgSetEraUnbondLimitResponse>
  SetInitBond(request: MsgSetInitBond): Promise<MsgSetInitBondResponse>
  SetChainBondingDuration(request: MsgSetChainBondingDuration): Promise<MsgSetChainBondingDurationResponse>
  SetPoolDetail(request: MsgSetPoolDetail): Promise<MsgSetPoolDetailResponse>
  SetLeastBond(request: MsgSetLeastBond): Promise<MsgSetLeastBondResponse>
  ClearCurrentEraSnapShots(request: MsgClearCurrentEraSnapShots): Promise<MsgClearCurrentEraSnapShotsResponse>
  SetCommission(request: MsgSetCommission): Promise<MsgSetCommissionResponse>
  SetReceiver(request: MsgSetReceiver): Promise<MsgSetReceiverResponse>
  SetUnbondFee(request: MsgSetUnbondFee): Promise<MsgSetUnbondFeeResponse>
  LiquidityUnbond(request: MsgLiquidityUnbond): Promise<MsgLiquidityUnbondResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetUnbondCommission(request: MsgSetUnbondCommission): Promise<MsgSetUnbondCommissionResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  AddNewPool(request: MsgAddNewPool): Promise<MsgAddNewPoolResponse> {
    const data = MsgAddNewPool.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'AddNewPool', data)
    return promise.then((data) => MsgAddNewPoolResponse.decode(new Reader(data)))
  }

  RemovePool(request: MsgRemovePool): Promise<MsgRemovePoolResponse> {
    const data = MsgRemovePool.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'RemovePool', data)
    return promise.then((data) => MsgRemovePoolResponse.decode(new Reader(data)))
  }

  SetEraUnbondLimit(request: MsgSetEraUnbondLimit): Promise<MsgSetEraUnbondLimitResponse> {
    const data = MsgSetEraUnbondLimit.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetEraUnbondLimit', data)
    return promise.then((data) => MsgSetEraUnbondLimitResponse.decode(new Reader(data)))
  }

  SetInitBond(request: MsgSetInitBond): Promise<MsgSetInitBondResponse> {
    const data = MsgSetInitBond.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetInitBond', data)
    return promise.then((data) => MsgSetInitBondResponse.decode(new Reader(data)))
  }

  SetChainBondingDuration(request: MsgSetChainBondingDuration): Promise<MsgSetChainBondingDurationResponse> {
    const data = MsgSetChainBondingDuration.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetChainBondingDuration', data)
    return promise.then((data) => MsgSetChainBondingDurationResponse.decode(new Reader(data)))
  }

  SetPoolDetail(request: MsgSetPoolDetail): Promise<MsgSetPoolDetailResponse> {
    const data = MsgSetPoolDetail.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetPoolDetail', data)
    return promise.then((data) => MsgSetPoolDetailResponse.decode(new Reader(data)))
  }

  SetLeastBond(request: MsgSetLeastBond): Promise<MsgSetLeastBondResponse> {
    const data = MsgSetLeastBond.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetLeastBond', data)
    return promise.then((data) => MsgSetLeastBondResponse.decode(new Reader(data)))
  }

  ClearCurrentEraSnapShots(request: MsgClearCurrentEraSnapShots): Promise<MsgClearCurrentEraSnapShotsResponse> {
    const data = MsgClearCurrentEraSnapShots.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'ClearCurrentEraSnapShots', data)
    return promise.then((data) => MsgClearCurrentEraSnapShotsResponse.decode(new Reader(data)))
  }

  SetCommission(request: MsgSetCommission): Promise<MsgSetCommissionResponse> {
    const data = MsgSetCommission.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetCommission', data)
    return promise.then((data) => MsgSetCommissionResponse.decode(new Reader(data)))
  }

  SetReceiver(request: MsgSetReceiver): Promise<MsgSetReceiverResponse> {
    const data = MsgSetReceiver.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetReceiver', data)
    return promise.then((data) => MsgSetReceiverResponse.decode(new Reader(data)))
  }

  SetUnbondFee(request: MsgSetUnbondFee): Promise<MsgSetUnbondFeeResponse> {
    const data = MsgSetUnbondFee.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetUnbondFee', data)
    return promise.then((data) => MsgSetUnbondFeeResponse.decode(new Reader(data)))
  }

  LiquidityUnbond(request: MsgLiquidityUnbond): Promise<MsgLiquidityUnbondResponse> {
    const data = MsgLiquidityUnbond.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'LiquidityUnbond', data)
    return promise.then((data) => MsgLiquidityUnbondResponse.decode(new Reader(data)))
  }

  SetUnbondCommission(request: MsgSetUnbondCommission): Promise<MsgSetUnbondCommissionResponse> {
    const data = MsgSetUnbondCommission.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetUnbondCommission', data)
    return promise.then((data) => MsgSetUnbondCommissionResponse.decode(new Reader(data)))
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
