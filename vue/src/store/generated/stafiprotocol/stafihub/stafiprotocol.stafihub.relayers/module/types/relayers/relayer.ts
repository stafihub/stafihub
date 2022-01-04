/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.relayers'

export interface Relayer {
  denom: string
  address: string
}

export interface Threshold {
  denom: string
  value: number
}

export interface LastVoter {
  denom: string
  voter: string
}

const baseRelayer: object = { denom: '', address: '' }

export const Relayer = {
  encode(message: Relayer, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Relayer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRelayer } as Relayer
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

  fromJSON(object: any): Relayer {
    const message = { ...baseRelayer } as Relayer
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

  toJSON(message: Relayer): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<Relayer>): Relayer {
    const message = { ...baseRelayer } as Relayer
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

const baseThreshold: object = { denom: '', value: 0 }

export const Threshold = {
  encode(message: Threshold, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.value !== 0) {
      writer.uint32(16).uint32(message.value)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Threshold {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseThreshold } as Threshold
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.value = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Threshold {
    const message = { ...baseThreshold } as Threshold
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

  toJSON(message: Threshold): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.value !== undefined && (obj.value = message.value)
    return obj
  },

  fromPartial(object: DeepPartial<Threshold>): Threshold {
    const message = { ...baseThreshold } as Threshold
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

const baseLastVoter: object = { denom: '', voter: '' }

export const LastVoter = {
  encode(message: LastVoter, writer: Writer = Writer.create()): Writer {
    if (message.denom !== '') {
      writer.uint32(10).string(message.denom)
    }
    if (message.voter !== '') {
      writer.uint32(18).string(message.voter)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): LastVoter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseLastVoter } as LastVoter
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string()
          break
        case 2:
          message.voter = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): LastVoter {
    const message = { ...baseLastVoter } as LastVoter
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter)
    } else {
      message.voter = ''
    }
    return message
  },

  toJSON(message: LastVoter): unknown {
    const obj: any = {}
    message.denom !== undefined && (obj.denom = message.denom)
    message.voter !== undefined && (obj.voter = message.voter)
    return obj
  },

  fromPartial(object: DeepPartial<LastVoter>): LastVoter {
    const message = { ...baseLastVoter } as LastVoter
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter
    } else {
      message.voter = ''
    }
    return message
  }
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
