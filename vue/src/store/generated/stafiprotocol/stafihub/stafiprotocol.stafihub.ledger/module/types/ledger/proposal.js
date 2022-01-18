/* eslint-disable */
import { bondActionFromJSON, bondActionToJSON } from '../ledger/ledger';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'stafiprotocol.stafihub.ledger';
const baseSetChainEraProposal = { proposer: '', denom: '', era: 0 };
export const SetChainEraProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.era !== 0) {
            writer.uint32(24).uint32(message.era);
        }
        if (message.propId.length !== 0) {
            writer.uint32(34).bytes(message.propId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSetChainEraProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.era = reader.uint32();
                    break;
                case 4:
                    message.propId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSetChainEraProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.era !== undefined && object.era !== null) {
            message.era = Number(object.era);
        }
        else {
            message.era = 0;
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = bytesFromBase64(object.propId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.denom !== undefined && (obj.denom = message.denom);
        message.era !== undefined && (obj.era = message.era);
        message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSetChainEraProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.era !== undefined && object.era !== null) {
            message.era = object.era;
        }
        else {
            message.era = 0;
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = new Uint8Array();
        }
        return message;
    }
};
const baseBondReportProposal = { proposer: '', denom: '', action: 0 };
export const BondReportProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.shotId.length !== 0) {
            writer.uint32(26).bytes(message.shotId);
        }
        if (message.action !== 0) {
            writer.uint32(32).int32(message.action);
        }
        if (message.propId.length !== 0) {
            writer.uint32(42).bytes(message.propId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBondReportProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.shotId = reader.bytes();
                    break;
                case 4:
                    message.action = reader.int32();
                    break;
                case 5:
                    message.propId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseBondReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = bytesFromBase64(object.shotId);
        }
        if (object.action !== undefined && object.action !== null) {
            message.action = bondActionFromJSON(object.action);
        }
        else {
            message.action = 0;
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = bytesFromBase64(object.propId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.denom !== undefined && (obj.denom = message.denom);
        message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()));
        message.action !== undefined && (obj.action = bondActionToJSON(message.action));
        message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseBondReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = object.shotId;
        }
        else {
            message.shotId = new Uint8Array();
        }
        if (object.action !== undefined && object.action !== null) {
            message.action = object.action;
        }
        else {
            message.action = 0;
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = new Uint8Array();
        }
        return message;
    }
};
const baseBondAndReportActiveProposal = { proposer: '', denom: '', action: 0, staked: '', unstaked: '' };
export const BondAndReportActiveProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.shotId.length !== 0) {
            writer.uint32(26).bytes(message.shotId);
        }
        if (message.action !== 0) {
            writer.uint32(32).int32(message.action);
        }
        if (message.staked !== '') {
            writer.uint32(42).string(message.staked);
        }
        if (message.unstaked !== '') {
            writer.uint32(50).string(message.unstaked);
        }
        if (message.propId.length !== 0) {
            writer.uint32(58).bytes(message.propId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBondAndReportActiveProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.shotId = reader.bytes();
                    break;
                case 4:
                    message.action = reader.int32();
                    break;
                case 5:
                    message.staked = reader.string();
                    break;
                case 6:
                    message.unstaked = reader.string();
                    break;
                case 7:
                    message.propId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseBondAndReportActiveProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = bytesFromBase64(object.shotId);
        }
        if (object.action !== undefined && object.action !== null) {
            message.action = bondActionFromJSON(object.action);
        }
        else {
            message.action = 0;
        }
        if (object.staked !== undefined && object.staked !== null) {
            message.staked = String(object.staked);
        }
        else {
            message.staked = '';
        }
        if (object.unstaked !== undefined && object.unstaked !== null) {
            message.unstaked = String(object.unstaked);
        }
        else {
            message.unstaked = '';
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = bytesFromBase64(object.propId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.denom !== undefined && (obj.denom = message.denom);
        message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()));
        message.action !== undefined && (obj.action = bondActionToJSON(message.action));
        message.staked !== undefined && (obj.staked = message.staked);
        message.unstaked !== undefined && (obj.unstaked = message.unstaked);
        message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseBondAndReportActiveProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = object.shotId;
        }
        else {
            message.shotId = new Uint8Array();
        }
        if (object.action !== undefined && object.action !== null) {
            message.action = object.action;
        }
        else {
            message.action = 0;
        }
        if (object.staked !== undefined && object.staked !== null) {
            message.staked = object.staked;
        }
        else {
            message.staked = '';
        }
        if (object.unstaked !== undefined && object.unstaked !== null) {
            message.unstaked = object.unstaked;
        }
        else {
            message.unstaked = '';
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = new Uint8Array();
        }
        return message;
    }
};
const baseActiveReportProposal = { proposer: '', denom: '', staked: '', unstaked: '' };
export const ActiveReportProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.shotId.length !== 0) {
            writer.uint32(26).bytes(message.shotId);
        }
        if (message.staked !== '') {
            writer.uint32(34).string(message.staked);
        }
        if (message.unstaked !== '') {
            writer.uint32(42).string(message.unstaked);
        }
        if (message.propId.length !== 0) {
            writer.uint32(50).bytes(message.propId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseActiveReportProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.shotId = reader.bytes();
                    break;
                case 4:
                    message.staked = reader.string();
                    break;
                case 5:
                    message.unstaked = reader.string();
                    break;
                case 6:
                    message.propId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseActiveReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = bytesFromBase64(object.shotId);
        }
        if (object.staked !== undefined && object.staked !== null) {
            message.staked = String(object.staked);
        }
        else {
            message.staked = '';
        }
        if (object.unstaked !== undefined && object.unstaked !== null) {
            message.unstaked = String(object.unstaked);
        }
        else {
            message.unstaked = '';
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = bytesFromBase64(object.propId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.denom !== undefined && (obj.denom = message.denom);
        message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()));
        message.staked !== undefined && (obj.staked = message.staked);
        message.unstaked !== undefined && (obj.unstaked = message.unstaked);
        message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseActiveReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = object.shotId;
        }
        else {
            message.shotId = new Uint8Array();
        }
        if (object.staked !== undefined && object.staked !== null) {
            message.staked = object.staked;
        }
        else {
            message.staked = '';
        }
        if (object.unstaked !== undefined && object.unstaked !== null) {
            message.unstaked = object.unstaked;
        }
        else {
            message.unstaked = '';
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = new Uint8Array();
        }
        return message;
    }
};
const baseWithdrawReportProposal = { proposer: '', denom: '' };
export const WithdrawReportProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.shotId.length !== 0) {
            writer.uint32(26).bytes(message.shotId);
        }
        if (message.propId.length !== 0) {
            writer.uint32(34).bytes(message.propId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseWithdrawReportProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.shotId = reader.bytes();
                    break;
                case 4:
                    message.propId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseWithdrawReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = bytesFromBase64(object.shotId);
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = bytesFromBase64(object.propId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.denom !== undefined && (obj.denom = message.denom);
        message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()));
        message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseWithdrawReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = object.shotId;
        }
        else {
            message.shotId = new Uint8Array();
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = new Uint8Array();
        }
        return message;
    }
};
const baseTransferReportProposal = { proposer: '', denom: '' };
export const TransferReportProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.shotId.length !== 0) {
            writer.uint32(26).bytes(message.shotId);
        }
        if (message.propId.length !== 0) {
            writer.uint32(34).bytes(message.propId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseTransferReportProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.shotId = reader.bytes();
                    break;
                case 4:
                    message.propId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseTransferReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = bytesFromBase64(object.shotId);
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = bytesFromBase64(object.propId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.denom !== undefined && (obj.denom = message.denom);
        message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()));
        message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseTransferReportProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = object.shotId;
        }
        else {
            message.shotId = new Uint8Array();
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = new Uint8Array();
        }
        return message;
    }
};
const baseExecuteBondProposal = { proposer: '', denom: '', bonder: '', pool: '', blockhash: '', txhash: '', amount: '' };
export const ExecuteBondProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.bonder !== '') {
            writer.uint32(26).string(message.bonder);
        }
        if (message.pool !== '') {
            writer.uint32(34).string(message.pool);
        }
        if (message.blockhash !== '') {
            writer.uint32(42).string(message.blockhash);
        }
        if (message.txhash !== '') {
            writer.uint32(50).string(message.txhash);
        }
        if (message.amount !== '') {
            writer.uint32(58).string(message.amount);
        }
        if (message.propId.length !== 0) {
            writer.uint32(66).bytes(message.propId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseExecuteBondProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.bonder = reader.string();
                    break;
                case 4:
                    message.pool = reader.string();
                    break;
                case 5:
                    message.blockhash = reader.string();
                    break;
                case 6:
                    message.txhash = reader.string();
                    break;
                case 7:
                    message.amount = reader.string();
                    break;
                case 8:
                    message.propId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseExecuteBondProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.bonder !== undefined && object.bonder !== null) {
            message.bonder = String(object.bonder);
        }
        else {
            message.bonder = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = String(object.pool);
        }
        else {
            message.pool = '';
        }
        if (object.blockhash !== undefined && object.blockhash !== null) {
            message.blockhash = String(object.blockhash);
        }
        else {
            message.blockhash = '';
        }
        if (object.txhash !== undefined && object.txhash !== null) {
            message.txhash = String(object.txhash);
        }
        else {
            message.txhash = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = String(object.amount);
        }
        else {
            message.amount = '';
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = bytesFromBase64(object.propId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.denom !== undefined && (obj.denom = message.denom);
        message.bonder !== undefined && (obj.bonder = message.bonder);
        message.pool !== undefined && (obj.pool = message.pool);
        message.blockhash !== undefined && (obj.blockhash = message.blockhash);
        message.txhash !== undefined && (obj.txhash = message.txhash);
        message.amount !== undefined && (obj.amount = message.amount);
        message.propId !== undefined && (obj.propId = base64FromBytes(message.propId !== undefined ? message.propId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseExecuteBondProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.bonder !== undefined && object.bonder !== null) {
            message.bonder = object.bonder;
        }
        else {
            message.bonder = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = object.pool;
        }
        else {
            message.pool = '';
        }
        if (object.blockhash !== undefined && object.blockhash !== null) {
            message.blockhash = object.blockhash;
        }
        else {
            message.blockhash = '';
        }
        if (object.txhash !== undefined && object.txhash !== null) {
            message.txhash = object.txhash;
        }
        else {
            message.txhash = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = '';
        }
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = new Uint8Array();
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
const atob = globalThis.atob || ((b64) => globalThis.Buffer.from(b64, 'base64').toString('binary'));
function bytesFromBase64(b64) {
    const bin = atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
    }
    return arr;
}
const btoa = globalThis.btoa || ((bin) => globalThis.Buffer.from(bin, 'binary').toString('base64'));
function base64FromBytes(arr) {
    const bin = [];
    for (let i = 0; i < arr.byteLength; ++i) {
        bin.push(String.fromCharCode(arr[i]));
    }
    return btoa(bin.join(''));
}
