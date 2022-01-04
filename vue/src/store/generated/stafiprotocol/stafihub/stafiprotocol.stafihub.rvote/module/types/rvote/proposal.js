/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
import { Any } from '../google/protobuf/any';
export const protobufPackage = 'stafiprotocol.stafihub.rvote';
/** ProposalStatus enumerates the valid statuses of a proposal. */
export var ProposalStatus;
(function (ProposalStatus) {
    ProposalStatus[ProposalStatus["PROPOSAL_STATUS_INITIATED"] = 0] = "PROPOSAL_STATUS_INITIATED";
    ProposalStatus[ProposalStatus["PROPOSAL_STATUS_APPROVED"] = 1] = "PROPOSAL_STATUS_APPROVED";
    ProposalStatus[ProposalStatus["PROPOSAL_STATUS_REJECTED"] = 2] = "PROPOSAL_STATUS_REJECTED";
    ProposalStatus[ProposalStatus["PROPOSAL_STATUS_EXPIRED"] = 3] = "PROPOSAL_STATUS_EXPIRED";
    ProposalStatus[ProposalStatus["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(ProposalStatus || (ProposalStatus = {}));
export function proposalStatusFromJSON(object) {
    switch (object) {
        case 0:
        case 'PROPOSAL_STATUS_INITIATED':
            return ProposalStatus.PROPOSAL_STATUS_INITIATED;
        case 1:
        case 'PROPOSAL_STATUS_APPROVED':
            return ProposalStatus.PROPOSAL_STATUS_APPROVED;
        case 2:
        case 'PROPOSAL_STATUS_REJECTED':
            return ProposalStatus.PROPOSAL_STATUS_REJECTED;
        case 3:
        case 'PROPOSAL_STATUS_EXPIRED':
            return ProposalStatus.PROPOSAL_STATUS_EXPIRED;
        case -1:
        case 'UNRECOGNIZED':
        default:
            return ProposalStatus.UNRECOGNIZED;
    }
}
export function proposalStatusToJSON(object) {
    switch (object) {
        case ProposalStatus.PROPOSAL_STATUS_INITIATED:
            return 'PROPOSAL_STATUS_INITIATED';
        case ProposalStatus.PROPOSAL_STATUS_APPROVED:
            return 'PROPOSAL_STATUS_APPROVED';
        case ProposalStatus.PROPOSAL_STATUS_REJECTED:
            return 'PROPOSAL_STATUS_REJECTED';
        case ProposalStatus.PROPOSAL_STATUS_EXPIRED:
            return 'PROPOSAL_STATUS_EXPIRED';
        default:
            return 'UNKNOWN';
    }
}
const baseProposal = { status: 0, votesFor: '', votesAgainst: '', startBlock: 0, expireBlock: 0 };
export const Proposal = {
    encode(message, writer = Writer.create()) {
        if (message.content !== undefined) {
            Any.encode(message.content, writer.uint32(10).fork()).ldelim();
        }
        if (message.status !== 0) {
            writer.uint32(16).int32(message.status);
        }
        for (const v of message.votesFor) {
            writer.uint32(26).string(v);
        }
        for (const v of message.votesAgainst) {
            writer.uint32(34).string(v);
        }
        if (message.startBlock !== 0) {
            writer.uint32(40).int64(message.startBlock);
        }
        if (message.expireBlock !== 0) {
            writer.uint32(48).int64(message.expireBlock);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseProposal };
        message.votesFor = [];
        message.votesAgainst = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.content = Any.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.status = reader.int32();
                    break;
                case 3:
                    message.votesFor.push(reader.string());
                    break;
                case 4:
                    message.votesAgainst.push(reader.string());
                    break;
                case 5:
                    message.startBlock = longToNumber(reader.int64());
                    break;
                case 6:
                    message.expireBlock = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseProposal };
        message.votesFor = [];
        message.votesAgainst = [];
        if (object.content !== undefined && object.content !== null) {
            message.content = Any.fromJSON(object.content);
        }
        else {
            message.content = undefined;
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = proposalStatusFromJSON(object.status);
        }
        else {
            message.status = 0;
        }
        if (object.votesFor !== undefined && object.votesFor !== null) {
            for (const e of object.votesFor) {
                message.votesFor.push(String(e));
            }
        }
        if (object.votesAgainst !== undefined && object.votesAgainst !== null) {
            for (const e of object.votesAgainst) {
                message.votesAgainst.push(String(e));
            }
        }
        if (object.startBlock !== undefined && object.startBlock !== null) {
            message.startBlock = Number(object.startBlock);
        }
        else {
            message.startBlock = 0;
        }
        if (object.expireBlock !== undefined && object.expireBlock !== null) {
            message.expireBlock = Number(object.expireBlock);
        }
        else {
            message.expireBlock = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.content !== undefined && (obj.content = message.content ? Any.toJSON(message.content) : undefined);
        message.status !== undefined && (obj.status = proposalStatusToJSON(message.status));
        if (message.votesFor) {
            obj.votesFor = message.votesFor.map((e) => e);
        }
        else {
            obj.votesFor = [];
        }
        if (message.votesAgainst) {
            obj.votesAgainst = message.votesAgainst.map((e) => e);
        }
        else {
            obj.votesAgainst = [];
        }
        message.startBlock !== undefined && (obj.startBlock = message.startBlock);
        message.expireBlock !== undefined && (obj.expireBlock = message.expireBlock);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseProposal };
        message.votesFor = [];
        message.votesAgainst = [];
        if (object.content !== undefined && object.content !== null) {
            message.content = Any.fromPartial(object.content);
        }
        else {
            message.content = undefined;
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = 0;
        }
        if (object.votesFor !== undefined && object.votesFor !== null) {
            for (const e of object.votesFor) {
                message.votesFor.push(e);
            }
        }
        if (object.votesAgainst !== undefined && object.votesAgainst !== null) {
            for (const e of object.votesAgainst) {
                message.votesAgainst.push(e);
            }
        }
        if (object.startBlock !== undefined && object.startBlock !== null) {
            message.startBlock = object.startBlock;
        }
        else {
            message.startBlock = 0;
        }
        if (object.expireBlock !== undefined && object.expireBlock !== null) {
            message.expireBlock = object.expireBlock;
        }
        else {
            message.expireBlock = 0;
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
