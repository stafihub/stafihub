/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'stafiprotocol.stafihub.rvote';
const baseGenesisState = { proposalLife: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.proposalLife !== 0) {
            writer.uint32(8).int64(message.proposalLife);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseGenesisState };
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
        message.proposalLife !== undefined && (obj.proposalLife = message.proposalLife);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        if (object.proposalLife !== undefined && object.proposalLife !== null) {
            message.proposalLife = object.proposalLife;
        }
        else {
            message.proposalLife = 0;
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
