/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { Any } from '../google/protobuf/any'

export const protobufPackage = 'stafiprotocol.stafihub.rvote'

/** ProposalStatus enumerates the valid statuses of a proposal. */
export enum ProposalStatus {
  PROPOSAL_STATUS_INITIATED = 0,
  PROPOSAL_STATUS_APPROVED = 1,
  PROPOSAL_STATUS_EXPIRED = 2,
  UNRECOGNIZED = -1
}

export function proposalStatusFromJSON(object: any): ProposalStatus {
  switch (object) {
    case 0:
    case 'PROPOSAL_STATUS_INITIATED':
      return ProposalStatus.PROPOSAL_STATUS_INITIATED
    case 1:
    case 'PROPOSAL_STATUS_APPROVED':
      return ProposalStatus.PROPOSAL_STATUS_APPROVED
    case 2:
    case 'PROPOSAL_STATUS_EXPIRED':
      return ProposalStatus.PROPOSAL_STATUS_EXPIRED
    case -1:
    case 'UNRECOGNIZED':
    default:
      return ProposalStatus.UNRECOGNIZED
  }
}

export function proposalStatusToJSON(object: ProposalStatus): string {
  switch (object) {
    case ProposalStatus.PROPOSAL_STATUS_INITIATED:
      return 'PROPOSAL_STATUS_INITIATED'
    case ProposalStatus.PROPOSAL_STATUS_APPROVED:
      return 'PROPOSAL_STATUS_APPROVED'
    case ProposalStatus.PROPOSAL_STATUS_EXPIRED:
      return 'PROPOSAL_STATUS_EXPIRED'
    default:
      return 'UNKNOWN'
  }
}

export interface Proposal {
  content: Any | undefined
  status: ProposalStatus
  voted: string[]
  startBlock: number
  expireBlock: number
}

const baseProposal: object = { status: 0, voted: '', startBlock: 0, expireBlock: 0 }

export const Proposal = {
  encode(message: Proposal, writer: Writer = Writer.create()): Writer {
    if (message.content !== undefined) {
      Any.encode(message.content, writer.uint32(10).fork()).ldelim()
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status)
    }
    for (const v of message.voted) {
      writer.uint32(26).string(v!)
    }
    if (message.startBlock !== 0) {
      writer.uint32(32).int64(message.startBlock)
    }
    if (message.expireBlock !== 0) {
      writer.uint32(40).int64(message.expireBlock)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Proposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseProposal } as Proposal
    message.voted = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.content = Any.decode(reader, reader.uint32())
          break
        case 2:
          message.status = reader.int32() as any
          break
        case 3:
          message.voted.push(reader.string())
          break
        case 4:
          message.startBlock = longToNumber(reader.int64() as Long)
          break
        case 5:
          message.expireBlock = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Proposal {
    const message = { ...baseProposal } as Proposal
    message.voted = []
    if (object.content !== undefined && object.content !== null) {
      message.content = Any.fromJSON(object.content)
    } else {
      message.content = undefined
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = proposalStatusFromJSON(object.status)
    } else {
      message.status = 0
    }
    if (object.voted !== undefined && object.voted !== null) {
      for (const e of object.voted) {
        message.voted.push(String(e))
      }
    }
    if (object.startBlock !== undefined && object.startBlock !== null) {
      message.startBlock = Number(object.startBlock)
    } else {
      message.startBlock = 0
    }
    if (object.expireBlock !== undefined && object.expireBlock !== null) {
      message.expireBlock = Number(object.expireBlock)
    } else {
      message.expireBlock = 0
    }
    return message
  },

  toJSON(message: Proposal): unknown {
    const obj: any = {}
    message.content !== undefined && (obj.content = message.content ? Any.toJSON(message.content) : undefined)
    message.status !== undefined && (obj.status = proposalStatusToJSON(message.status))
    if (message.voted) {
      obj.voted = message.voted.map((e) => e)
    } else {
      obj.voted = []
    }
    message.startBlock !== undefined && (obj.startBlock = message.startBlock)
    message.expireBlock !== undefined && (obj.expireBlock = message.expireBlock)
    return obj
  },

  fromPartial(object: DeepPartial<Proposal>): Proposal {
    const message = { ...baseProposal } as Proposal
    message.voted = []
    if (object.content !== undefined && object.content !== null) {
      message.content = Any.fromPartial(object.content)
    } else {
      message.content = undefined
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status
    } else {
      message.status = 0
    }
    if (object.voted !== undefined && object.voted !== null) {
      for (const e of object.voted) {
        message.voted.push(e)
      }
    }
    if (object.startBlock !== undefined && object.startBlock !== null) {
      message.startBlock = object.startBlock
    } else {
      message.startBlock = 0
    }
    if (object.expireBlock !== undefined && object.expireBlock !== null) {
      message.expireBlock = object.expireBlock
    } else {
      message.expireBlock = 0
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
