/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { Relayer, Threshold } from '../relayers/relayer'

export const protobufPackage = 'stafiprotocol.stafihub.relayers'

/** GenesisState defines the relayers module's genesis state. */
export interface GenesisState {
  relayers: Relayer[]
  thresholds: Threshold[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  proposalLife: number
}

const baseGenesisState: object = { proposalLife: 0 }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.relayers) {
      Relayer.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.thresholds) {
      Threshold.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    if (message.proposalLife !== 0) {
      writer.uint32(24).int64(message.proposalLife)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.relayers = []
    message.thresholds = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.relayers.push(Relayer.decode(reader, reader.uint32()))
          break
        case 2:
          message.thresholds.push(Threshold.decode(reader, reader.uint32()))
          break
        case 3:
          message.proposalLife = longToNumber(reader.int64() as Long)
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
    message.relayers = []
    message.thresholds = []
    if (object.relayers !== undefined && object.relayers !== null) {
      for (const e of object.relayers) {
        message.relayers.push(Relayer.fromJSON(e))
      }
    }
    if (object.thresholds !== undefined && object.thresholds !== null) {
      for (const e of object.thresholds) {
        message.thresholds.push(Threshold.fromJSON(e))
      }
    }
    if (object.proposalLife !== undefined && object.proposalLife !== null) {
      message.proposalLife = Number(object.proposalLife)
    } else {
      message.proposalLife = 0
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.relayers) {
      obj.relayers = message.relayers.map((e) => (e ? Relayer.toJSON(e) : undefined))
    } else {
      obj.relayers = []
    }
    if (message.thresholds) {
      obj.thresholds = message.thresholds.map((e) => (e ? Threshold.toJSON(e) : undefined))
    } else {
      obj.thresholds = []
    }
    message.proposalLife !== undefined && (obj.proposalLife = message.proposalLife)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.relayers = []
    message.thresholds = []
    if (object.relayers !== undefined && object.relayers !== null) {
      for (const e of object.relayers) {
        message.relayers.push(Relayer.fromPartial(e))
      }
    }
    if (object.thresholds !== undefined && object.thresholds !== null) {
      for (const e of object.thresholds) {
        message.thresholds.push(Threshold.fromPartial(e))
      }
    }
    if (object.proposalLife !== undefined && object.proposalLife !== null) {
      message.proposalLife = object.proposalLife
    } else {
      message.proposalLife = 0
    }
    return message
  }
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
