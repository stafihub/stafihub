/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.sudo'

/** GenesisState defines the sudo module's genesis state. */
export interface GenesisState {
  admin: string
  /** this line is used by starport scaffolding # genesis/proto/state */
  denoms: string[]
}

const baseGenesisState: object = { admin: '', denoms: '' }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.admin !== '') {
      writer.uint32(10).string(message.admin)
    }
    for (const v of message.denoms) {
      writer.uint32(18).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.denoms = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.admin = reader.string()
          break
        case 2:
          message.denoms.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.denoms = []
    if (object.admin !== undefined && object.admin !== null) {
      message.admin = String(object.admin)
    } else {
      message.admin = ''
    }
    if (object.denoms !== undefined && object.denoms !== null) {
      for (const e of object.denoms) {
        message.denoms.push(String(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    message.admin !== undefined && (obj.admin = message.admin)
    if (message.denoms) {
      obj.denoms = message.denoms.map((e) => e)
    } else {
      obj.denoms = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.denoms = []
    if (object.admin !== undefined && object.admin !== null) {
      message.admin = object.admin
    } else {
      message.admin = ''
    }
    if (object.denoms !== undefined && object.denoms !== null) {
      for (const e of object.denoms) {
        message.denoms.push(e)
      }
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
