/* eslint-disable */
import { ProposalStatus, proposalStatusFromJSON, proposalStatusToJSON } from '../rvote/proposal'
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { Any } from '../google/protobuf/any'

export const protobufPackage = 'stafiprotocol.stafihub.rvote'

export interface MsgSetProposalLife {
  creator: string
  proposalLife: number
}

export interface MsgSetProposalLifeResponse {}

export interface MsgSubmitProposal {
  proposer: string
  content: Any | undefined
}

export interface MsgSubmitProposalResponse {
  propId: string
  status: ProposalStatus
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

const baseMsgSubmitProposal: object = { proposer: '' }

export const MsgSubmitProposal = {
  encode(message: MsgSubmitProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.content !== undefined) {
      Any.encode(message.content, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSubmitProposal } as MsgSubmitProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.content = Any.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSubmitProposal {
    const message = { ...baseMsgSubmitProposal } as MsgSubmitProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = Any.fromJSON(object.content)
    } else {
      message.content = undefined
    }
    return message
  },

  toJSON(message: MsgSubmitProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.content !== undefined && (obj.content = message.content ? Any.toJSON(message.content) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgSubmitProposal>): MsgSubmitProposal {
    const message = { ...baseMsgSubmitProposal } as MsgSubmitProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = Any.fromPartial(object.content)
    } else {
      message.content = undefined
    }
    return message
  }
}

const baseMsgSubmitProposalResponse: object = { propId: '', status: 0 }

export const MsgSubmitProposalResponse = {
  encode(message: MsgSubmitProposalResponse, writer: Writer = Writer.create()): Writer {
    if (message.propId !== '') {
      writer.uint32(10).string(message.propId)
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitProposalResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgSubmitProposalResponse } as MsgSubmitProposalResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.propId = reader.string()
          break
        case 2:
          message.status = reader.int32() as any
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSubmitProposalResponse {
    const message = { ...baseMsgSubmitProposalResponse } as MsgSubmitProposalResponse
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = String(object.propId)
    } else {
      message.propId = ''
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = proposalStatusFromJSON(object.status)
    } else {
      message.status = 0
    }
    return message
  },

  toJSON(message: MsgSubmitProposalResponse): unknown {
    const obj: any = {}
    message.propId !== undefined && (obj.propId = message.propId)
    message.status !== undefined && (obj.status = proposalStatusToJSON(message.status))
    return obj
  },

  fromPartial(object: DeepPartial<MsgSubmitProposalResponse>): MsgSubmitProposalResponse {
    const message = { ...baseMsgSubmitProposalResponse } as MsgSubmitProposalResponse
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = ''
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status
    } else {
      message.status = 0
    }
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  SetProposalLife(request: MsgSetProposalLife): Promise<MsgSetProposalLifeResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitProposal(request: MsgSubmitProposal): Promise<MsgSubmitProposalResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  SetProposalLife(request: MsgSetProposalLife): Promise<MsgSetProposalLifeResponse> {
    const data = MsgSetProposalLife.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.rvote.Msg', 'SetProposalLife', data)
    return promise.then((data) => MsgSetProposalLifeResponse.decode(new Reader(data)))
  }

  SubmitProposal(request: MsgSubmitProposal): Promise<MsgSubmitProposalResponse> {
    const data = MsgSubmitProposal.encode(request).finish()
    const promise = this.rpc.request('stafiprotocol.stafihub.rvote.Msg', 'SubmitProposal', data)
    return promise.then((data) => MsgSubmitProposalResponse.decode(new Reader(data)))
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
