/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.relayers'

export interface MsgCreateRelayer {
  denom: string
  address: string
}

export interface MsgCreateRelayerResponse {}

export interface MsgDeleteRelayer {
  denom: string
  address: string
}

export interface MsgDeleteRelayerResponse {}

export interface MsgSetThreshold {
  denom: string
  value: string
}

export interface MsgSetThresholdResponse {}

const baseMsgCreateRelayer: object = { denom: '', address: '' }

export const MsgCreateRelayer = {
  encode(message: MsgCreateRelayer, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
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
          message.denom = reader.string()
          break
        case 2:
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
    message.denom !== undefined && (obj.denom = message.denom)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateRelayer>): MsgCreateRelayer {
    const message = { ...baseMsgCreateRelayer } as MsgCreateRelayer
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

const baseMsgDeleteRelayer: object = { denom: '', address: '' }

export const MsgDeleteRelayer = {
  encode(message: MsgDeleteRelayer, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
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
          message.denom = reader.string()
          break
        case 2:
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
    message.denom !== undefined && (obj.denom = message.denom)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteRelayer>): MsgDeleteRelayer {
    const message = { ...baseMsgDeleteRelayer } as MsgDeleteRelayer
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

const baseMsgSetThreshold: object = { denom: '', value: '' }

export const MsgSetThreshold = {
  encode(message: MsgSetThreshold, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.value !== '') {
      writer.uint32(18).string(message.value)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetThreshold {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetThreshold } as MsgSetThreshold
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
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

  fromJSON(object: any): MsgSetThreshold {
    const message = { ...baseMsgSetThreshold } as MsgSetThreshold
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value)
    } else {
      message.value = ''
    }
    return message
  },

  toJSON(message: MsgSetThreshold): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.value !== undefined && (obj.value = message.value)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSetThreshold>): MsgSetThreshold {
    const message = { ...baseMsgSetThreshold } as MsgSetThreshold
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value
    } else {
      message.value = ''
    }
    return message
  }
}

const baseMsgSetThresholdResponse: object = {}

export const MsgSetThresholdResponse = {
  encode(_: MsgSetThresholdResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetThresholdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSetThresholdResponse } as MsgSetThresholdResponse
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

  fromJSON(_: any): MsgSetThresholdResponse {
    const message = { ...baseMsgSetThresholdResponse } as MsgSetThresholdResponse
    return message
  },

  toJSON(_: MsgSetThresholdResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgSetThresholdResponse>): MsgSetThresholdResponse {
    const message = { ...baseMsgSetThresholdResponse } as MsgSetThresholdResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  CreateRelayer(request: MsgCreateRelayer): Promise<MsgCreateRelayerResponse>
  DeleteRelayer(request: MsgDeleteRelayer): Promise<MsgDeleteRelayerResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetThreshold(request: MsgSetThreshold): Promise<MsgSetThresholdResponse>
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

  SetThreshold(request: MsgSetThreshold): Promise<MsgSetThresholdResponse> {
    const data = MsgSetThreshold.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.relayers.Msg', 'SetThreshold', data)
    return promise.then((data) => MsgSetThresholdResponse.decode(new Reader(data)))
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
