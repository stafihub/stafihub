/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.relayers'

export interface Relayer {
  creator: string
  denom: string
  address: string
}

export interface Threshold {
  creator: string
  denom: string
  value: number
}

const baseRelayer: object = { creator: '', denom: '', address: '' }

export const Relayer = {
  encode(message: Relayer, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): Relayer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRelayer } as Relayer
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

  fromJSON(object: any): Relayer {
    const message = { ...baseRelayer } as Relayer
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

  toJSON(message: Relayer): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<Relayer>): Relayer {
    const message = { ...baseRelayer } as Relayer
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

const baseThreshold: object = { creator: '', denom: '', value: 0 }

export const Threshold = {
  encode(message: Threshold, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): Threshold {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseThreshold } as Threshold
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

  fromJSON(object: any): Threshold {
    const message = { ...baseThreshold } as Threshold
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

  toJSON(message: Threshold): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.denom !== undefined && (obj.denom = message.denom)
    message.value !== undefined && (obj.value = message.value)
    return obj
  },

  fromPartial(object: DeepPartial<Threshold>): Threshold {
    const message = { ...baseThreshold } as Threshold
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
