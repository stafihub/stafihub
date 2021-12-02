/* eslint-disable */
import { Relayer, Threshold } from '../relayers/relayer';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'stafiprotocol.stafihub.relayers';
const baseGenesisState = { admin: '' };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.admin !== '') {
            writer.uint32(10).string(message.admin);
        }
        for (const v of message.relayers) {
            Relayer.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.thresholds) {
            Threshold.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.relayers = [];
        message.thresholds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.admin = reader.string();
                    break;
                case 2:
                    message.relayers.push(Relayer.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.thresholds.push(Threshold.decode(reader, reader.uint32()));
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
        message.relayers = [];
        message.thresholds = [];
        if (object.admin !== undefined && object.admin !== null) {
            message.admin = String(object.admin);
        }
        else {
            message.admin = '';
        }
        if (object.relayers !== undefined && object.relayers !== null) {
            for (const e of object.relayers) {
                message.relayers.push(Relayer.fromJSON(e));
            }
        }
        if (object.thresholds !== undefined && object.thresholds !== null) {
            for (const e of object.thresholds) {
                message.thresholds.push(Threshold.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.admin !== undefined && (obj.admin = message.admin);
        if (message.relayers) {
            obj.relayers = message.relayers.map((e) => (e ? Relayer.toJSON(e) : undefined));
        }
        else {
            obj.relayers = [];
        }
        if (message.thresholds) {
            obj.thresholds = message.thresholds.map((e) => (e ? Threshold.toJSON(e) : undefined));
        }
        else {
            obj.thresholds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.relayers = [];
        message.thresholds = [];
        if (object.admin !== undefined && object.admin !== null) {
            message.admin = object.admin;
        }
        else {
            message.admin = '';
        }
        if (object.relayers !== undefined && object.relayers !== null) {
            for (const e of object.relayers) {
                message.relayers.push(Relayer.fromPartial(e));
            }
        }
        if (object.thresholds !== undefined && object.thresholds !== null) {
            for (const e of object.thresholds) {
                message.thresholds.push(Threshold.fromPartial(e));
            }
        }
        return message;
    }
};
