/* eslint-disable */
import { BondAction, bondActionFromJSON, bondActionToJSON } from '../ledger/ledger'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'stafiprotocol.stafihub.ledger'

export interface SetChainEraProposal {
  proposer: string
  denom: string
  era: number
  propId: Uint8Array
}

export interface BondReportProposal {
  proposer: string
  denom: string
  shotId: Uint8Array
  action: BondAction
  propId: Uint8Array
}

export interface BondAndReportActiveProposal {
  proposer: string
  denom: string
  shotId: Uint8Array
  action: BondAction
  staked: string
  unstaked: string
  propId: Uint8Array
}

export interface ActiveReportProposal {
  proposer: string
  denom: string
  shotId: Uint8Array
  staked: string
  unstaked: string
  propId: Uint8Array
}

export interface WithdrawReportProposal {
  proposer: string
  denom: string
  shotId: Uint8Array
  propId: Uint8Array
}

export interface TransferReportProposal {
  proposer: string
  denom: string
  shotId: Uint8Array
  propId: Uint8Array
}

export interface ExecuteBondProposal {
  proposer: string
  denom: string
  bonder: string
  pool: string
  blockhash: string
  txhash: string
  amount: string
  propId: Uint8Array
}

const baseSetChainEraProposal: object = { proposer: '', denom: '', era: 0 }

export const SetChainEraProposal = {
  encode(message: SetChainEraProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.era !== 0) {
      writer.uint32(24).uint32(message.era)
    }
    if (message.propId.length !== 0) {
      writer.uint32(34).bytes(message.propId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): SetChainEraProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseSetChainEraProposal } as SetChainEraProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.era = reader.uint32()
          break
        case 4:
          message.propId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): SetChainEraProposal {
    const message = { ...baseSetChainEraProposal } as SetChainEraProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.era !== undefined && object.era !== null) {
      message.era = Number(object.era)
    } else {
      message.era = 0
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = bytesFromBase64(object.propId)
    }
    return message
  },

  toJSON(message: SetChainEraProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.denom !== undefined && (obj.denom = message.denom)
    message.era !== undefined && (obj.era = message.era)
    message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<SetChainEraProposal>): SetChainEraProposal {
    const message = { ...baseSetChainEraProposal } as SetChainEraProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.era !== undefined && object.era !== null) {
      message.era = object.era
    } else {
      message.era = 0
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = new Uint8Array()
    }
    return message
  }
}

const baseBondReportProposal: object = { proposer: '', denom: '', action: 0 }

export const BondReportProposal = {
  encode(message: BondReportProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.shotId.length !== 0) {
      writer.uint32(26).bytes(message.shotId)
    }
    if (message.action !== 0) {
      writer.uint32(32).int32(message.action)
    }
    if (message.propId.length !== 0) {
      writer.uint32(42).bytes(message.propId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): BondReportProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseBondReportProposal } as BondReportProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.shotId = reader.bytes()
          break
        case 4:
          message.action = reader.int32() as any
          break
        case 5:
          message.propId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): BondReportProposal {
    const message = { ...baseBondReportProposal } as BondReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = bytesFromBase64(object.shotId)
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = bondActionFromJSON(object.action)
    } else {
      message.action = 0
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = bytesFromBase64(object.propId)
    }
    return message
  },

  toJSON(message: BondReportProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.denom !== undefined && (obj.denom = message.denom)
    message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()))
    message.action !== undefined && (obj.action = bondActionToJSON(message.action))
    message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<BondReportProposal>): BondReportProposal {
    const message = { ...baseBondReportProposal } as BondReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = object.shotId
    } else {
      message.shotId = new Uint8Array()
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = object.action
    } else {
      message.action = 0
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = new Uint8Array()
    }
    return message
  }
}

const baseBondAndReportActiveProposal: object = { proposer: '', denom: '', action: 0, staked: '', unstaked: '' }

export const BondAndReportActiveProposal = {
  encode(message: BondAndReportActiveProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.shotId.length !== 0) {
      writer.uint32(26).bytes(message.shotId)
    }
    if (message.action !== 0) {
      writer.uint32(32).int32(message.action)
    }
    if (message.staked !== '') {
      writer.uint32(42).string(message.staked)
    }
    if (message.unstaked !== '') {
      writer.uint32(50).string(message.unstaked)
    }
    if (message.propId.length !== 0) {
      writer.uint32(58).bytes(message.propId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): BondAndReportActiveProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseBondAndReportActiveProposal } as BondAndReportActiveProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.shotId = reader.bytes()
          break
        case 4:
          message.action = reader.int32() as any
          break
        case 5:
          message.staked = reader.string()
          break
        case 6:
          message.unstaked = reader.string()
          break
        case 7:
          message.propId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): BondAndReportActiveProposal {
    const message = { ...baseBondAndReportActiveProposal } as BondAndReportActiveProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = bytesFromBase64(object.shotId)
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = bondActionFromJSON(object.action)
    } else {
      message.action = 0
    }
    if (object.staked !== undefined && object.staked !== null) {
      message.staked = String(object.staked)
    } else {
      message.staked = ''
    }
    if (object.unstaked !== undefined && object.unstaked !== null) {
      message.unstaked = String(object.unstaked)
    } else {
      message.unstaked = ''
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = bytesFromBase64(object.propId)
    }
    return message
  },

  toJSON(message: BondAndReportActiveProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.denom !== undefined && (obj.denom = message.denom)
    message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()))
    message.action !== undefined && (obj.action = bondActionToJSON(message.action))
    message.staked !== undefined && (obj.staked = message.staked)
    message.unstaked !== undefined && (obj.unstaked = message.unstaked)
    message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<BondAndReportActiveProposal>): BondAndReportActiveProposal {
    const message = { ...baseBondAndReportActiveProposal } as BondAndReportActiveProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = object.shotId
    } else {
      message.shotId = new Uint8Array()
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = object.action
    } else {
      message.action = 0
    }
    if (object.staked !== undefined && object.staked !== null) {
      message.staked = object.staked
    } else {
      message.staked = ''
    }
    if (object.unstaked !== undefined && object.unstaked !== null) {
      message.unstaked = object.unstaked
    } else {
      message.unstaked = ''
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = new Uint8Array()
    }
    return message
  }
}

const baseActiveReportProposal: object = { proposer: '', denom: '', staked: '', unstaked: '' }

export const ActiveReportProposal = {
  encode(message: ActiveReportProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.shotId.length !== 0) {
      writer.uint32(26).bytes(message.shotId)
    }
    if (message.staked !== '') {
      writer.uint32(34).string(message.staked)
    }
    if (message.unstaked !== '') {
      writer.uint32(42).string(message.unstaked)
    }
    if (message.propId.length !== 0) {
      writer.uint32(50).bytes(message.propId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ActiveReportProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseActiveReportProposal } as ActiveReportProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.shotId = reader.bytes()
          break
        case 4:
          message.staked = reader.string()
          break
        case 5:
          message.unstaked = reader.string()
          break
        case 6:
          message.propId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ActiveReportProposal {
    const message = { ...baseActiveReportProposal } as ActiveReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = bytesFromBase64(object.shotId)
    }
    if (object.staked !== undefined && object.staked !== null) {
      message.staked = String(object.staked)
    } else {
      message.staked = ''
    }
    if (object.unstaked !== undefined && object.unstaked !== null) {
      message.unstaked = String(object.unstaked)
    } else {
      message.unstaked = ''
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = bytesFromBase64(object.propId)
    }
    return message
  },

  toJSON(message: ActiveReportProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.denom !== undefined && (obj.denom = message.denom)
    message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()))
    message.staked !== undefined && (obj.staked = message.staked)
    message.unstaked !== undefined && (obj.unstaked = message.unstaked)
    message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<ActiveReportProposal>): ActiveReportProposal {
    const message = { ...baseActiveReportProposal } as ActiveReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = object.shotId
    } else {
      message.shotId = new Uint8Array()
    }
    if (object.staked !== undefined && object.staked !== null) {
      message.staked = object.staked
    } else {
      message.staked = ''
    }
    if (object.unstaked !== undefined && object.unstaked !== null) {
      message.unstaked = object.unstaked
    } else {
      message.unstaked = ''
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = new Uint8Array()
    }
    return message
  }
}

const baseWithdrawReportProposal: object = { proposer: '', denom: '' }

export const WithdrawReportProposal = {
  encode(message: WithdrawReportProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.shotId.length !== 0) {
      writer.uint32(26).bytes(message.shotId)
    }
    if (message.propId.length !== 0) {
      writer.uint32(34).bytes(message.propId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): WithdrawReportProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseWithdrawReportProposal } as WithdrawReportProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.shotId = reader.bytes()
          break
        case 4:
          message.propId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): WithdrawReportProposal {
    const message = { ...baseWithdrawReportProposal } as WithdrawReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = bytesFromBase64(object.shotId)
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = bytesFromBase64(object.propId)
    }
    return message
  },

  toJSON(message: WithdrawReportProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.denom !== undefined && (obj.denom = message.denom)
    message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()))
    message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<WithdrawReportProposal>): WithdrawReportProposal {
    const message = { ...baseWithdrawReportProposal } as WithdrawReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = object.shotId
    } else {
      message.shotId = new Uint8Array()
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = new Uint8Array()
    }
    return message
  }
}

const baseTransferReportProposal: object = { proposer: '', denom: '' }

export const TransferReportProposal = {
  encode(message: TransferReportProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.shotId.length !== 0) {
      writer.uint32(26).bytes(message.shotId)
    }
    if (message.propId.length !== 0) {
      writer.uint32(34).bytes(message.propId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): TransferReportProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseTransferReportProposal } as TransferReportProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.shotId = reader.bytes()
          break
        case 4:
          message.propId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): TransferReportProposal {
    const message = { ...baseTransferReportProposal } as TransferReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = bytesFromBase64(object.shotId)
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = bytesFromBase64(object.propId)
    }
    return message
  },

  toJSON(message: TransferReportProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.denom !== undefined && (obj.denom = message.denom)
    message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()))
    message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<TransferReportProposal>): TransferReportProposal {
    const message = { ...baseTransferReportProposal } as TransferReportProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.shotId !== undefined && object.shotId !== null) {
      message.shotId = object.shotId
    } else {
      message.shotId = new Uint8Array()
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = new Uint8Array()
    }
    return message
  }
}

const baseExecuteBondProposal: object = { proposer: '', denom: '', bonder: '', pool: '', blockhash: '', txhash: '', amount: '' }

export const ExecuteBondProposal = {
  encode(message: ExecuteBondProposal, writer: Writer = Writer.create()): Writer {
    if (message.proposer !== '') {
      writer.uint32(10).string(message.proposer)
    }
    if (message.denom !== '') {
      writer.uint32(18).string(message.denom)
    }
    if (message.bonder !== '') {
      writer.uint32(26).string(message.bonder)
    }
    if (message.pool !== '') {
      writer.uint32(34).string(message.pool)
    }
    if (message.blockhash !== '') {
      writer.uint32(42).string(message.blockhash)
    }
    if (message.txhash !== '') {
      writer.uint32(50).string(message.txhash)
    }
    if (message.amount !== '') {
      writer.uint32(58).string(message.amount)
    }
    if (message.propId.length !== 0) {
      writer.uint32(66).bytes(message.propId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ExecuteBondProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseExecuteBondProposal } as ExecuteBondProposal
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposer = reader.string()
          break
        case 2:
          message.denom = reader.string()
          break
        case 3:
          message.bonder = reader.string()
          break
        case 4:
          message.pool = reader.string()
          break
        case 5:
          message.blockhash = reader.string()
          break
        case 6:
          message.txhash = reader.string()
          break
        case 7:
          message.amount = reader.string()
          break
        case 8:
          message.propId = reader.bytes()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ExecuteBondProposal {
    const message = { ...baseExecuteBondProposal } as ExecuteBondProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer)
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom)
    } else {
      message.denom = ''
    }
    if (object.bonder !== undefined && object.bonder !== null) {
      message.bonder = String(object.bonder)
    } else {
      message.bonder = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = String(object.pool)
    } else {
      message.pool = ''
    }
    if (object.blockhash !== undefined && object.blockhash !== null) {
      message.blockhash = String(object.blockhash)
    } else {
      message.blockhash = ''
    }
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = String(object.txhash)
    } else {
      message.txhash = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount)
    } else {
      message.amount = ''
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = bytesFromBase64(object.propId)
    }
    return message
  },

  toJSON(message: ExecuteBondProposal): unknown {
    const obj: any = {}
    message.proposer !== undefined && (obj.proposer = message.proposer)
    message.denom !== undefined && (obj.denom = message.denom)
    message.bonder !== undefined && (obj.bonder = message.bonder)
    message.pool !== undefined && (obj.pool = message.pool)
    message.blockhash !== undefined && (obj.blockhash = message.blockhash)
    message.txhash !== undefined && (obj.txhash = message.txhash)
    message.amount !== undefined && (obj.amount = message.amount)
    message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()))
    return obj
  },

  fromPartial(object: DeepPartial<ExecuteBondProposal>): ExecuteBondProposal {
    const message = { ...baseExecuteBondProposal } as ExecuteBondProposal
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer
    } else {
      message.proposer = ''
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom
    } else {
      message.denom = ''
    }
    if (object.bonder !== undefined && object.bonder !== null) {
      message.bonder = object.bonder
    } else {
      message.bonder = ''
    }
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = object.pool
    } else {
      message.pool = ''
    }
    if (object.blockhash !== undefined && object.blockhash !== null) {
      message.blockhash = object.blockhash
    } else {
      message.blockhash = ''
    }
    if (object.txhash !== undefined && object.txhash !== null) {
      message.txhash = object.txhash
    } else {
      message.txhash = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = ''
    }
    if (object.propId !== undefined && object.propId !== null) {
      message.propId = object.propId
    } else {
      message.propId = new Uint8Array()
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

const atob: (b64: string) => string = globalThis.atob || ((b64) => globalThis.Buffer.from(b64, 'base64').toString('binary'))
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64)
  const arr = new Uint8Array(bin.length)
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i)
  }
  return arr
}

const btoa: (bin: string) => string = globalThis.btoa || ((bin) => globalThis.Buffer.from(bin, 'binary').toString('base64'))
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = []
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]))
  }
  return btoa(bin.join(''))
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
