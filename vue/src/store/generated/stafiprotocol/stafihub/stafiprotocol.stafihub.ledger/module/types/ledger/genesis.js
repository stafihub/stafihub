/* eslint-disable */
import { ExchangeRate } from '../ledger/ledger';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'stafiprotocol.stafihub.ledger';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.exchangeRateList) {
            ExchangeRate.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.exchangeRateList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.exchangeRateList.push(ExchangeRate.decode(reader, reader.uint32()));
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
        message.exchangeRateList = [];
        if (object.exchangeRateList !== undefined && object.exchangeRateList !== null) {
            for (const e of object.exchangeRateList) {
                message.exchangeRateList.push(ExchangeRate.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.exchangeRateList) {
            obj.exchangeRateList = message.exchangeRateList.map((e) => (e ? ExchangeRate.toJSON(e) : undefined));
        }
        else {
            obj.exchangeRateList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.exchangeRateList = [];
        if (object.exchangeRateList !== undefined && object.exchangeRateList !== null) {
            for (const e of object.exchangeRateList) {
                message.exchangeRateList.push(ExchangeRate.fromPartial(e));
            }
        }
        return message;
    }
};
