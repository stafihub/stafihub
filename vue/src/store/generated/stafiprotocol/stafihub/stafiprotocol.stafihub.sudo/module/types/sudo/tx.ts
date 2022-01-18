/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Metadata } from '../cosmos/bank/v1beta1/bank'

export const protobufPackage = 'stafiprotocol.stafihub.sudo'

export interface MsgUpdateAdmin {
  creator: string
  address: string
}

export interface MsgUpdateAdminResponse {}

export interface MsgAddDenom {
  creator: string
  Metadata: Metadata | undefined
}

export interface MsgAddDenomResponse {}

const baseMsgUpdateAdmin: object = { creator: '', address: '' }

export const MsgUpdateAdmin = {
  encode(message: MsgUpdateAdmin, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateAdmin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateAdmin } as MsgUpdateAdmin
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
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

  fromJSON(object: any): MsgUpdateAdmin {
    const message = { ...baseMsgUpdateAdmin } as MsgUpdateAdmin
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgUpdateAdmin): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateAdmin>): MsgUpdateAdmin {
    const message = { ...baseMsgUpdateAdmin } as MsgUpdateAdmin
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgUpdateAdminResponse: object = {}

export const MsgUpdateAdminResponse = {
  encode(_: MsgUpdateAdminResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateAdminResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateAdminResponse } as MsgUpdateAdminResponse
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

  fromJSON(_: any): MsgUpdateAdminResponse {
    const message = { ...baseMsgUpdateAdminResponse } as MsgUpdateAdminResponse
    return message
  },

  toJSON(_: MsgUpdateAdminResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateAdminResponse>): MsgUpdateAdminResponse {
    const message = { ...baseMsgUpdateAdminResponse } as MsgUpdateAdminResponse
    return message
  }
}

const baseMsgAddDenom: object = { creator: '' }

export const MsgAddDenom = {
  encode(message: MsgAddDenom, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.Metadata !== undefined) {
      Metadata.encode(message.Metadata, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddDenom {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddDenom } as MsgAddDenom
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.Metadata = Metadata.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgAddDenom {
    const message = { ...baseMsgAddDenom } as MsgAddDenom
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.Metadata !== undefined && object.Metadata !== null) {
      message.Metadata = Metadata.fromJSON(object.Metadata)
    } else {
      message.Metadata = undefined
    }
    return message
  },

  toJSON(message: MsgAddDenom): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.Metadata !== undefined && (obj.Metadata = message.Metadata ? Metadata.toJSON(message.Metadata) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgAddDenom>): MsgAddDenom {
    const message = { ...baseMsgAddDenom } as MsgAddDenom
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.Metadata !== undefined && object.Metadata !== null) {
      message.Metadata = Metadata.fromPartial(object.Metadata)
    } else {
      message.Metadata = undefined
    }
    return message
  }
}

const baseMsgAddDenomResponse: object = {}

export const MsgAddDenomResponse = {
  encode(_: MsgAddDenomResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddDenomResponse } as MsgAddDenomResponse
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

  fromJSON(_: any): MsgAddDenomResponse {
    const message = { ...baseMsgAddDenomResponse } as MsgAddDenomResponse
    return message
  },

  toJSON(_: MsgAddDenomResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgAddDenomResponse>): MsgAddDenomResponse {
    const message = { ...baseMsgAddDenomResponse } as MsgAddDenomResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  UpdateAdmin(request: MsgUpdateAdmin): Promise<MsgUpdateAdminResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddDenom(request: MsgAddDenom): Promise<MsgAddDenomResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  UpdateAdmin(request: MsgUpdateAdmin): Promise<MsgUpdateAdminResponse> {
    const data = MsgUpdateAdmin.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.sudo.Msg', 'UpdateAdmin', data)
    return promise.then((data) => MsgUpdateAdminResponse.decode(new Reader(data)))
  }

  AddDenom(request: MsgAddDenom): Promise<MsgAddDenomResponse> {
    const data = MsgAddDenom.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.sudo.Msg', 'AddDenom', data)
    return promise.then((data) => MsgAddDenomResponse.decode(new Reader(data)))
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
