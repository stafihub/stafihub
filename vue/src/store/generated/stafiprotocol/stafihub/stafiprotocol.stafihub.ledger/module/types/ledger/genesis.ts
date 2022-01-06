/* eslint-disable */
import { ExchangeRate } from '../ledger/ledger'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.ledger'

/** GenesisState defines the ledger module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  exchangeRateList: ExchangeRate[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.exchangeRateList) {
      ExchangeRate.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.exchangeRateList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.exchangeRateList.push(ExchangeRate.decode(reader, reader.uint32()))
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
    message.exchangeRateList = []
    if (object.exchangeRateList !== undefined && object.exchangeRateList !== null) {
      for (const e of object.exchangeRateList) {
        message.exchangeRateList.push(ExchangeRate.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.exchangeRateList) {
      obj.exchangeRateList = message.exchangeRateList.map((e) => (e ? ExchangeRate.toJSON(e) : undefined))
    } else {
      obj.exchangeRateList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.exchangeRateList = []
    if (object.exchangeRateList !== undefined && object.exchangeRateList !== null) {
      for (const e of object.exchangeRateList) {
        message.exchangeRateList.push(ExchangeRate.fromPartial(e))
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
