/* eslint-disable */
import { proposalStatusFromJSON, proposalStatusToJSON } from '../rvote/proposal';
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { Any } from '../google/protobuf/any';
export const protobufPackage = 'stafiprotocol.stafihub.rvote';
const baseMsgSetProposalLife = { creator: '', proposalLife: 0 };
export const MsgSetProposalLife = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.proposalLife !== 0) {
            writer.uint32(16).int64(message.proposalLife);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetProposalLife };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.proposalLife = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSetProposalLife };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.proposalLife !== undefined && object.proposalLife !== null) {
            message.proposalLife = Number(object.proposalLife);
        }
        else {
            message.proposalLife = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.proposalLife !== undefined && (obj.proposalLife = message.proposalLife);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetProposalLife };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.proposalLife !== undefined && object.proposalLife !== null) {
            message.proposalLife = object.proposalLife;
        }
        else {
            message.proposalLife = 0;
        }
        return message;
    }
};
const baseMsgSetProposalLifeResponse = {};
export const MsgSetProposalLifeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetProposalLifeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgSetProposalLifeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetProposalLifeResponse };
        return message;
    }
};
const baseMsgSubmitProposal = { proposer: '' };
export const MsgSubmitProposal = {
    encode(message, writer = Writer.create()) {
        if (message.proposer !== '') {
            writer.uint32(10).string(message.proposer);
        }
        if (message.content !== undefined) {
            Any.encode(message.content, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSubmitProposal };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposer = reader.string();
                    break;
                case 2:
                    message.content = Any.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSubmitProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = String(object.proposer);
        }
        else {
            message.proposer = '';
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = Any.fromJSON(object.content);
        }
        else {
            message.content = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposer !== undefined && (obj.proposer = message.proposer);
        message.content !== undefined && (obj.content = message.content ? Any.toJSON(message.content) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSubmitProposal };
        if (object.proposer !== undefined && object.proposer !== null) {
            message.proposer = object.proposer;
        }
        else {
            message.proposer = '';
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = Any.fromPartial(object.content);
        }
        else {
            message.content = undefined;
        }
        return message;
    }
};
const baseMsgSubmitProposalResponse = { propId: '', status: 0 };
export const MsgSubmitProposalResponse = {
    encode(message, writer = Writer.create()) {
        if (message.propId !== '') {
            writer.uint32(10).string(message.propId);
        }
        if (message.status !== 0) {
            writer.uint32(16).int32(message.status);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSubmitProposalResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.propId = reader.string();
                    break;
                case 2:
                    message.status = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSubmitProposalResponse };
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = String(object.propId);
        }
        else {
            message.propId = '';
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = proposalStatusFromJSON(object.status);
        }
        else {
            message.status = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.propId !== undefined && (obj.propId = message.propId);
        message.status !== undefined && (obj.status = proposalStatusToJSON(message.status));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSubmitProposalResponse };
        if (object.propId !== undefined && object.propId !== null) {
            message.propId = object.propId;
        }
        else {
            message.propId = '';
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = 0;
        }
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    SetProposalLife(request) {
        const data = MsgSetProposalLife.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.rvote.Msg', 'SetProposalLife', data);
        return promise.then((data) => MsgSetProposalLifeResponse.decode(new Reader(data)));
    }
    SubmitProposal(request) {
        const data = MsgSubmitProposal.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.rvote.Msg', 'SubmitProposal', data);
        return promise.then((data) => MsgSubmitProposalResponse.decode(new Reader(data)));
    }
}
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
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
