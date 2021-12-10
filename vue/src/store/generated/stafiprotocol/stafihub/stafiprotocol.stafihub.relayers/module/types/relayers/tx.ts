/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'stafiprotocol.stafihub.relayers'

export interface MsgCreateRelayer {
  creator: string
  denom: string
  address: string
}

export interface MsgCreateRelayerResponse {}

export interface MsgDeleteRelayer {
  creator: string
  denom: string
  address: string
}

export interface MsgDeleteRelayerResponse {}

export interface MsgUpdateThreshold {
  creator: string
  denom: string
  value: number
}

export interface MsgUpdateThresholdResponse {}

export interface MsgSetProposalLife {
  creator: string
  proposalLife: number
}

export interface MsgSetProposalLifeResponse {}

const baseMsgCreateRelayer: object = { creator: '', denom: '', address: '' }

export const MsgCreateRelayer = {
  encode(message: MsgCreateRelayer, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.address !== '') {
      writer.uint32(26).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateRelayer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateRelayer } as MsgCreateRelayer
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
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateRelayer {
    const message = { ...baseMsgCreateRelayer } as MsgCreateRelayer
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
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgCreateRelayer): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateRelayer>): MsgCreateRelayer {
    const message = { ...baseMsgCreateRelayer } as MsgCreateRelayer
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
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgCreateRelayerResponse: object = {}

export const MsgCreateRelayerResponse = {
  encode(_: MsgCreateRelayerResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateRelayerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateRelayerResponse } as MsgCreateRelayerResponse
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

  fromJSON(_: any): MsgCreateRelayerResponse {
    const message = { ...baseMsgCreateRelayerResponse } as MsgCreateRelayerResponse
    return message
  },

  toJSON(_: MsgCreateRelayerResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateRelayerResponse>): MsgCreateRelayerResponse {
    const message = { ...baseMsgCreateRelayerResponse } as MsgCreateRelayerResponse
    return message
  }
}

const baseMsgDeleteRelayer: object = { creator: '', denom: '', address: '' }

export const MsgDeleteRelayer = {
  encode(message: MsgDeleteRelayer, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.address !== '') {
      writer.uint32(26).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteRelayer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteRelayer } as MsgDeleteRelayer
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
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeleteRelayer {
    const message = { ...baseMsgDeleteRelayer } as MsgDeleteRelayer
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
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgDeleteRelayer): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteRelayer>): MsgDeleteRelayer {
    const message = { ...baseMsgDeleteRelayer } as MsgDeleteRelayer
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
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgDeleteRelayerResponse: object = {}

export const MsgDeleteRelayerResponse = {
  encode(_: MsgDeleteRelayerResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteRelayerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteRelayerResponse } as MsgDeleteRelayerResponse
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

  fromJSON(_: any): MsgDeleteRelayerResponse {
    const message = { ...baseMsgDeleteRelayerResponse } as MsgDeleteRelayerResponse
    return message
  },

  toJSON(_: MsgDeleteRelayerResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteRelayerResponse>): MsgDeleteRelayerResponse {
    const message = { ...baseMsgDeleteRelayerResponse } as MsgDeleteRelayerResponse
    return message
  }
}

const baseMsgUpdateThreshold: object = { creator: '', denom: '', value: 0 }

export const MsgUpdateThreshold = {
  encode(message: MsgUpdateThreshold, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.value !== 0) {
      writer.uint32(24).uint32(message.value)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateThreshold {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateThreshold } as MsgUpdateThreshold
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
          message.value = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateThreshold {
    const message = { ...baseMsgUpdateThreshold } as MsgUpdateThreshold
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
    if (object.value !== undefined && object.value !== null) {
      message.value = Number(object.value)
    } else {
      message.value = 0
    }
    return message
  },

  toJSON(message: MsgUpdateThreshold): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.value !== undefined && (obj.value = message.value)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateThreshold>): MsgUpdateThreshold {
    const message = { ...baseMsgUpdateThreshold } as MsgUpdateThreshold
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
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value
    } else {
      message.value = 0
    }
    return message
  }
}

const baseMsgUpdateThresholdResponse: object = {}

export const MsgUpdateThresholdResponse = {
  encode(_: MsgUpdateThresholdResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateThresholdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateThresholdResponse } as MsgUpdateThresholdResponse
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

  fromJSON(_: any): MsgUpdateThresholdResponse {
    const message = { ...baseMsgUpdateThresholdResponse } as MsgUpdateThresholdResponse
    return message
  },

  toJSON(_: MsgUpdateThresholdResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateThresholdResponse>): MsgUpdateThresholdResponse {
    const message = { ...baseMsgUpdateThresholdResponse } as MsgUpdateThresholdResponse
    return message
  }
}

const baseMsgSetProposalLife: object = { creator: '', proposalLife: 0 }

export const MsgSetProposalLife = {
  encode(message: MsgSetProposalLife, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.proposalLife !== 0) {
      writer.uint32(16).int64(message.proposalLife)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetProposalLife {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetProposalLife } as MsgSetProposalLife
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.proposalLife = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetProposalLife {
    const message = { ...baseMsgSetProposalLife } as MsgSetProposalLife
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.proposalLife !== undefined && object.proposalLife !== null) {
      message.proposalLife = Number(object.proposalLife)
    } else {
      message.proposalLife = 0
    }
    return message
  },

  toJSON(message: MsgSetProposalLife): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.proposalLife !== undefined && (obj.proposalLife = message.proposalLife)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetProposalLife>): MsgSetProposalLife {
    const message = { ...baseMsgSetProposalLife } as MsgSetProposalLife
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.proposalLife !== undefined && object.proposalLife !== null) {
      message.proposalLife = object.proposalLife
    } else {
      message.proposalLife = 0
    }
    return message
  }
}

const baseMsgSetProposalLifeResponse: object = {}

export const MsgSetProposalLifeResponse = {
  encode(_: MsgSetProposalLifeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetProposalLifeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetProposalLifeResponse } as MsgSetProposalLifeResponse
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

  fromJSON(_: any): MsgSetProposalLifeResponse {
    const message = { ...baseMsgSetProposalLifeResponse } as MsgSetProposalLifeResponse
    return message
  },

  toJSON(_: MsgSetProposalLifeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetProposalLifeResponse>): MsgSetProposalLifeResponse {
    const message = { ...baseMsgSetProposalLifeResponse } as MsgSetProposalLifeResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  CreateRelayer(request: MsgCreateRelayer): Promise<MsgCreateRelayerResponse>
  DeleteRelayer(request: MsgDeleteRelayer): Promise<MsgDeleteRelayerResponse>
  UpdateThreshold(request: MsgUpdateThreshold): Promise<MsgUpdateThresholdResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetProposalLife(request: MsgSetProposalLife): Promise<MsgSetProposalLifeResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateRelayer(request: MsgCreateRelayer): Promise<MsgCreateRelayerResponse> {
    const data = MsgCreateRelayer.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Msg', 'CreateRelayer', data)
    return promise.then((data) => MsgCreateRelayerResponse.decode(new Reader(data)))
  }

  DeleteRelayer(request: MsgDeleteRelayer): Promise<MsgDeleteRelayerResponse> {
    const data = MsgDeleteRelayer.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Msg', 'DeleteRelayer', data)
    return promise.then((data) => MsgDeleteRelayerResponse.decode(new Reader(data)))
  }

  UpdateThreshold(request: MsgUpdateThreshold): Promise<MsgUpdateThresholdResponse> {
    const data = MsgUpdateThreshold.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Msg', 'UpdateThreshold', data)
    return promise.then((data) => MsgUpdateThresholdResponse.decode(new Reader(data)))
  }

  SetProposalLife(request: MsgSetProposalLife): Promise<MsgSetProposalLifeResponse> {
    const data = MsgSetProposalLife.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Msg', 'SetProposalLife', data)
    return promise.then((data) => MsgSetProposalLifeResponse.decode(new Reader(data)))
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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
