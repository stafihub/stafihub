/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'stafiprotocol.stafihub.sudo';
const baseSymbol = {};
export const Symbol = {
    encode(message, writer = Writer.create()) {
        Object.entries(message.denoms).forEach(([key, value]) => {
            Symbol_DenomsEntry.encode({ key: key, value }, writer.uint32(10).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSymbol };
        message.denoms = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    const entry1 = Symbol_DenomsEntry.decode(reader, reader.uint32());
                    if (entry1.value !== undefined) {
                        message.denoms[entry1.key] = entry1.value;
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSymbol };
        message.denoms = {};
        if (object.denoms !== undefined && object.denoms !== null) {
            Object.entries(object.denoms).forEach(([key, value]) => {
                message.denoms[key] = Boolean(value);
            });
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        obj.denoms = {};
        if (message.denoms) {
            Object.entries(message.denoms).forEach(([k, v]) => {
                obj.denoms[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSymbol };
        message.denoms = {};
        if (object.denoms !== undefined && object.denoms !== null) {
            Object.entries(object.denoms).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.denoms[key] = Boolean(value);
                }
            });
        }
        return message;
    }
};
const baseSymbol_DenomsEntry = { key: '', value: false };
export const Symbol_DenomsEntry = {
    encode(message, writer = Writer.create()) {
        if (message.key !== '') {
            writer.uint32(10).string(message.key);
        }
        if (message.value === true) {
            writer.uint32(16).bool(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSymbol_DenomsEntry };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.string();
                    break;
                case 2:
                    message.value = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSymbol_DenomsEntry };
        if (object.key !== undefined && object.key !== null) {
            message.key = String(object.key);
        }
        else {
            message.key = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = Boolean(object.value);
        }
        else {
            message.value = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.key !== undefined && (obj.key = message.key);
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSymbol_DenomsEntry };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = false;
        }
        return message;
    }
};
