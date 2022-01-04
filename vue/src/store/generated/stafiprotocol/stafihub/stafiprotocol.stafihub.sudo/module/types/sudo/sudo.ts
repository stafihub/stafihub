/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.sudo'

export interface Symbol {
  denoms: { [key: string]: boolean }
}

export interface Symbol_DenomsEntry {
  key: string
  value: boolean
}

const baseSymbol: object = {}

export const Symbol = {
  encode(message: Symbol, writer: Writer = Writer.create()): Writer {
    Object.entries(message.denoms).forEach(([key, value]) => {
      Symbol_DenomsEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim()
    })
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Symbol {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseSymbol } as Symbol
    message.denoms = {}
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          const entry1 = Symbol_DenomsEntry.decode(reader, reader.uint32())
          if (entry1.value !== undefined) {
            message.denoms[entry1.key] = entry1.value
          }
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Symbol {
    const message = { ...baseSymbol } as Symbol
    message.denoms = {}
    if (object.denoms !== undefined && object.denoms !== null) {
      Object.entries(object.denoms).forEach(([key, value]) => {
        message.denoms[key] = Boolean(value)
      })
    }
    return message
  },

  toJSON(message: Symbol): unknown {
    const obj: any = {}
    obj.denoms = {}
    if (message.denoms) {
      Object.entries(message.denoms).forEach(([k, v]) => {
        obj.denoms[k] = v
      })
    }
    return obj
  },

  fromPartial(object: DeepPartial<Symbol>): Symbol {
    const message = { ...baseSymbol } as Symbol
    message.denoms = {}
    if (object.denoms !== undefined && object.denoms !== null) {
      Object.entries(object.denoms).forEach(([key, value]) => {
        if (value !== undefined) {
          message.denoms[key] = Boolean(value)
        }
      })
    }
    return message
  }
}

const baseSymbol_DenomsEntry: object = { key: '', value: false }

export const Symbol_DenomsEntry = {
  encode(message: Symbol_DenomsEntry, writer: Writer = Writer.create()): Writer {
    if (message.key !== '') {
      writer.uint32(10).string(message.key)
    }
    if (message.value === true) {
      writer.uint32(16).bool(message.value)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Symbol_DenomsEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseSymbol_DenomsEntry } as Symbol_DenomsEntry
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string()
          break
        case 2:
          message.value = reader.bool()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Symbol_DenomsEntry {
    const message = { ...baseSymbol_DenomsEntry } as Symbol_DenomsEntry
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key)
    } else {
      message.key = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = Boolean(object.value)
    } else {
      message.value = false
    }
    return message
  },

  toJSON(message: Symbol_DenomsEntry): unknown {
    const obj: any = {}
    message.key !== undefined && (obj.key = message.key)
    message.value !== undefined && (obj.value = message.value)
    return obj
  },

  fromPartial(object: DeepPartial<Symbol_DenomsEntry>): Symbol_DenomsEntry {
    const message = { ...baseSymbol_DenomsEntry } as Symbol_DenomsEntry
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key
    } else {
      message.key = ''
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value
    } else {
      message.value = false
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
