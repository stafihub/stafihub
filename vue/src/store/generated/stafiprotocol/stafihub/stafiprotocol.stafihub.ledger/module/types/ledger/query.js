/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ExchangeRate, EraExchangeRate } from '../ledger/ledger';
export const protobufPackage = 'stafiprotocol.stafihub.ledger';
const baseQueryGetExchangeRateRequest = { denom: '' };
export const QueryGetExchangeRateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetExchangeRateRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetExchangeRateRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetExchangeRateRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryGetExchangeRateResponse = {};
export const QueryGetExchangeRateResponse = {
    encode(message, writer = Writer.create()) {
        if (message.exchangeRate !== undefined) {
            ExchangeRate.encode(message.exchangeRate, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetExchangeRateResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.exchangeRate = ExchangeRate.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetExchangeRateResponse };
        if (object.exchangeRate !== undefined && object.exchangeRate !== null) {
            message.exchangeRate = ExchangeRate.fromJSON(object.exchangeRate);
        }
        else {
            message.exchangeRate = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.exchangeRate !== undefined && (obj.exchangeRate = message.exchangeRate ? ExchangeRate.toJSON(message.exchangeRate) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetExchangeRateResponse };
        if (object.exchangeRate !== undefined && object.exchangeRate !== null) {
            message.exchangeRate = ExchangeRate.fromPartial(object.exchangeRate);
        }
        else {
            message.exchangeRate = undefined;
        }
        return message;
    }
};
const baseQueryExchangeRateAllRequest = {};
export const QueryExchangeRateAllRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryExchangeRateAllRequest };
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
        const message = { ...baseQueryExchangeRateAllRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryExchangeRateAllRequest };
        return message;
    }
};
const baseQueryExchangeRateAllResponse = {};
export const QueryExchangeRateAllResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.exchangeRates) {
            ExchangeRate.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryExchangeRateAllResponse };
        message.exchangeRates = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.exchangeRates.push(ExchangeRate.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryExchangeRateAllResponse };
        message.exchangeRates = [];
        if (object.exchangeRates !== undefined && object.exchangeRates !== null) {
            for (const e of object.exchangeRates) {
                message.exchangeRates.push(ExchangeRate.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.exchangeRates) {
            obj.exchangeRates = message.exchangeRates.map((e) => (e ? ExchangeRate.toJSON(e) : undefined));
        }
        else {
            obj.exchangeRates = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryExchangeRateAllResponse };
        message.exchangeRates = [];
        if (object.exchangeRates !== undefined && object.exchangeRates !== null) {
            for (const e of object.exchangeRates) {
                message.exchangeRates.push(ExchangeRate.fromPartial(e));
            }
        }
        return message;
    }
};
const baseQueryGetEraExchangeRateRequest = { denom: '', era: 0 };
export const QueryGetEraExchangeRateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.era !== 0) {
            writer.uint32(16).uint32(message.era);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetEraExchangeRateRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.era = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetEraExchangeRateRequest };
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.era !== undefined && (obj.era = message.era);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetEraExchangeRateRequest };
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
        return message;
    }
};
const baseQueryGetEraExchangeRateResponse = {};
export const QueryGetEraExchangeRateResponse = {
    encode(message, writer = Writer.create()) {
        if (message.eraExchangeRate !== undefined) {
            EraExchangeRate.encode(message.eraExchangeRate, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetEraExchangeRateResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.eraExchangeRate = EraExchangeRate.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetEraExchangeRateResponse };
        if (object.eraExchangeRate !== undefined && object.eraExchangeRate !== null) {
            message.eraExchangeRate = EraExchangeRate.fromJSON(object.eraExchangeRate);
        }
        else {
            message.eraExchangeRate = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.eraExchangeRate !== undefined && (obj.eraExchangeRate = message.eraExchangeRate ? EraExchangeRate.toJSON(message.eraExchangeRate) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetEraExchangeRateResponse };
        if (object.eraExchangeRate !== undefined && object.eraExchangeRate !== null) {
            message.eraExchangeRate = EraExchangeRate.fromPartial(object.eraExchangeRate);
        }
        else {
            message.eraExchangeRate = undefined;
        }
        return message;
    }
};
const baseQueryEraExchangeRatesByDenomRequest = { denom: '' };
export const QueryEraExchangeRatesByDenomRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryEraExchangeRatesByDenomRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryEraExchangeRatesByDenomRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryEraExchangeRatesByDenomRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryEraExchangeRatesByDenomResponse = {};
export const QueryEraExchangeRatesByDenomResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.eraExchangeRates) {
            EraExchangeRate.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryEraExchangeRatesByDenomResponse };
        message.eraExchangeRates = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.eraExchangeRates.push(EraExchangeRate.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryEraExchangeRatesByDenomResponse };
        message.eraExchangeRates = [];
        if (object.eraExchangeRates !== undefined && object.eraExchangeRates !== null) {
            for (const e of object.eraExchangeRates) {
                message.eraExchangeRates.push(EraExchangeRate.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.eraExchangeRates) {
            obj.eraExchangeRates = message.eraExchangeRates.map((e) => (e ? EraExchangeRate.toJSON(e) : undefined));
        }
        else {
            obj.eraExchangeRates = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryEraExchangeRatesByDenomResponse };
        message.eraExchangeRates = [];
        if (object.eraExchangeRates !== undefined && object.eraExchangeRates !== null) {
            for (const e of object.eraExchangeRates) {
                message.eraExchangeRates.push(EraExchangeRate.fromPartial(e));
            }
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    GetExchangeRate(request) {
        const data = QueryGetExchangeRateRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetExchangeRate', data);
        return promise.then((data) => QueryGetExchangeRateResponse.decode(new Reader(data)));
    }
    ExchangeRateAll(request) {
        const data = QueryExchangeRateAllRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'ExchangeRateAll', data);
        return promise.then((data) => QueryExchangeRateAllResponse.decode(new Reader(data)));
    }
    GetEraExchangeRate(request) {
        const data = QueryGetEraExchangeRateRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetEraExchangeRate', data);
        return promise.then((data) => QueryGetEraExchangeRateResponse.decode(new Reader(data)));
    }
    EraExchangeRatesByDenom(request) {
        const data = QueryEraExchangeRatesByDenomRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'EraExchangeRatesByDenom', data);
        return promise.then((data) => QueryEraExchangeRatesByDenomResponse.decode(new Reader(data)));
    }
}
