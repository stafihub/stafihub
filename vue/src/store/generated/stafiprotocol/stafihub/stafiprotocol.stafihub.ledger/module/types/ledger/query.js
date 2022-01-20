/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ExchangeRate, EraExchangeRate, PoolDetail, UnbondFee, LeastBond, BondPipeline, BondSnapshot, PoolUnbond, AccountUnbond, BondRecord } from '../ledger/ledger';
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
const baseQueryPoolsByDenomRequest = { denom: '' };
export const QueryPoolsByDenomRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryPoolsByDenomRequest };
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
        const message = { ...baseQueryPoolsByDenomRequest };
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
        const message = { ...baseQueryPoolsByDenomRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryPoolsByDenomResponse = { addrs: '' };
export const QueryPoolsByDenomResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.addrs) {
            writer.uint32(10).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryPoolsByDenomResponse };
        message.addrs = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.addrs.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryPoolsByDenomResponse };
        message.addrs = [];
        if (object.addrs !== undefined && object.addrs !== null) {
            for (const e of object.addrs) {
                message.addrs.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.addrs) {
            obj.addrs = message.addrs.map((e) => e);
        }
        else {
            obj.addrs = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryPoolsByDenomResponse };
        message.addrs = [];
        if (object.addrs !== undefined && object.addrs !== null) {
            for (const e of object.addrs) {
                message.addrs.push(e);
            }
        }
        return message;
    }
};
const baseQueryBondedPoolsByDenomRequest = { denom: '' };
export const QueryBondedPoolsByDenomRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryBondedPoolsByDenomRequest };
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
        const message = { ...baseQueryBondedPoolsByDenomRequest };
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
        const message = { ...baseQueryBondedPoolsByDenomRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryBondedPoolsByDenomResponse = { addrs: '' };
export const QueryBondedPoolsByDenomResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.addrs) {
            writer.uint32(10).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryBondedPoolsByDenomResponse };
        message.addrs = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.addrs.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryBondedPoolsByDenomResponse };
        message.addrs = [];
        if (object.addrs !== undefined && object.addrs !== null) {
            for (const e of object.addrs) {
                message.addrs.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.addrs) {
            obj.addrs = message.addrs.map((e) => e);
        }
        else {
            obj.addrs = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryBondedPoolsByDenomResponse };
        message.addrs = [];
        if (object.addrs !== undefined && object.addrs !== null) {
            for (const e of object.addrs) {
                message.addrs.push(e);
            }
        }
        return message;
    }
};
const baseQueryGetPoolDetailRequest = { denom: '', pool: '' };
export const QueryGetPoolDetailRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.pool !== '') {
            writer.uint32(18).string(message.pool);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPoolDetailRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.pool = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPoolDetailRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = String(object.pool);
        }
        else {
            message.pool = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.pool !== undefined && (obj.pool = message.pool);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPoolDetailRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = object.pool;
        }
        else {
            message.pool = '';
        }
        return message;
    }
};
const baseQueryGetPoolDetailResponse = {};
export const QueryGetPoolDetailResponse = {
    encode(message, writer = Writer.create()) {
        if (message.detail !== undefined) {
            PoolDetail.encode(message.detail, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPoolDetailResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.detail = PoolDetail.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPoolDetailResponse };
        if (object.detail !== undefined && object.detail !== null) {
            message.detail = PoolDetail.fromJSON(object.detail);
        }
        else {
            message.detail = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.detail !== undefined && (obj.detail = message.detail ? PoolDetail.toJSON(message.detail) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPoolDetailResponse };
        if (object.detail !== undefined && object.detail !== null) {
            message.detail = PoolDetail.fromPartial(object.detail);
        }
        else {
            message.detail = undefined;
        }
        return message;
    }
};
const baseQueryGetChainEraRequest = { denom: '' };
export const QueryGetChainEraRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainEraRequest };
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
        const message = { ...baseQueryGetChainEraRequest };
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
        const message = { ...baseQueryGetChainEraRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryGetChainEraResponse = { era: 0 };
export const QueryGetChainEraResponse = {
    encode(message, writer = Writer.create()) {
        if (message.era !== 0) {
            writer.uint32(8).uint32(message.era);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainEraResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseQueryGetChainEraResponse };
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
        message.era !== undefined && (obj.era = message.era);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChainEraResponse };
        if (object.era !== undefined && object.era !== null) {
            message.era = object.era;
        }
        else {
            message.era = 0;
        }
        return message;
    }
};
const baseQueryGetCurrentEraSnapshotRequest = { denom: '' };
export const QueryGetCurrentEraSnapshotRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCurrentEraSnapshotRequest };
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
        const message = { ...baseQueryGetCurrentEraSnapshotRequest };
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
        const message = { ...baseQueryGetCurrentEraSnapshotRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryGetCurrentEraSnapshotResponse = {};
export const QueryGetCurrentEraSnapshotResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.shotIds) {
            writer.uint32(10).bytes(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCurrentEraSnapshotResponse };
        message.shotIds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.shotIds.push(reader.bytes());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetCurrentEraSnapshotResponse };
        message.shotIds = [];
        if (object.shotIds !== undefined && object.shotIds !== null) {
            for (const e of object.shotIds) {
                message.shotIds.push(bytesFromBase64(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.shotIds) {
            obj.shotIds = message.shotIds.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
        }
        else {
            obj.shotIds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCurrentEraSnapshotResponse };
        message.shotIds = [];
        if (object.shotIds !== undefined && object.shotIds !== null) {
            for (const e of object.shotIds) {
                message.shotIds.push(e);
            }
        }
        return message;
    }
};
const baseQueryGetReceiverRequest = {};
export const QueryGetReceiverRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetReceiverRequest };
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
        const message = { ...baseQueryGetReceiverRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetReceiverRequest };
        return message;
    }
};
const baseQueryGetReceiverResponse = { receiver: '' };
export const QueryGetReceiverResponse = {
    encode(message, writer = Writer.create()) {
        if (message.receiver !== '') {
            writer.uint32(10).string(message.receiver);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetReceiverResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.receiver = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetReceiverResponse };
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = String(object.receiver);
        }
        else {
            message.receiver = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.receiver !== undefined && (obj.receiver = message.receiver);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetReceiverResponse };
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = object.receiver;
        }
        else {
            message.receiver = '';
        }
        return message;
    }
};
const baseQueryGetCommissionRequest = {};
export const QueryGetCommissionRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCommissionRequest };
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
        const message = { ...baseQueryGetCommissionRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetCommissionRequest };
        return message;
    }
};
const baseQueryGetCommissionResponse = { commission: '' };
export const QueryGetCommissionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.commission !== '') {
            writer.uint32(10).string(message.commission);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCommissionResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.commission = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetCommissionResponse };
        if (object.commission !== undefined && object.commission !== null) {
            message.commission = String(object.commission);
        }
        else {
            message.commission = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.commission !== undefined && (obj.commission = message.commission);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCommissionResponse };
        if (object.commission !== undefined && object.commission !== null) {
            message.commission = object.commission;
        }
        else {
            message.commission = '';
        }
        return message;
    }
};
const baseQueryGetChainBondingDurationRequest = { denom: '' };
export const QueryGetChainBondingDurationRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainBondingDurationRequest };
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
        const message = { ...baseQueryGetChainBondingDurationRequest };
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
        const message = { ...baseQueryGetChainBondingDurationRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryGetChainBondingDurationResponse = { era: 0 };
export const QueryGetChainBondingDurationResponse = {
    encode(message, writer = Writer.create()) {
        if (message.era !== 0) {
            writer.uint32(8).uint32(message.era);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainBondingDurationResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseQueryGetChainBondingDurationResponse };
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
        message.era !== undefined && (obj.era = message.era);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChainBondingDurationResponse };
        if (object.era !== undefined && object.era !== null) {
            message.era = object.era;
        }
        else {
            message.era = 0;
        }
        return message;
    }
};
const baseQueryGetUnbondFeeRequest = {};
export const QueryGetUnbondFeeRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetUnbondFeeRequest };
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
        const message = { ...baseQueryGetUnbondFeeRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetUnbondFeeRequest };
        return message;
    }
};
const baseQueryGetUnbondFeeResponse = {};
export const QueryGetUnbondFeeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.fee !== undefined) {
            UnbondFee.encode(message.fee, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetUnbondFeeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.fee = UnbondFee.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetUnbondFeeResponse };
        if (object.fee !== undefined && object.fee !== null) {
            message.fee = UnbondFee.fromJSON(object.fee);
        }
        else {
            message.fee = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.fee !== undefined && (obj.fee = message.fee ? UnbondFee.toJSON(message.fee) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetUnbondFeeResponse };
        if (object.fee !== undefined && object.fee !== null) {
            message.fee = UnbondFee.fromPartial(object.fee);
        }
        else {
            message.fee = undefined;
        }
        return message;
    }
};
const baseQueryGetUnbondCommissionRequest = {};
export const QueryGetUnbondCommissionRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetUnbondCommissionRequest };
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
        const message = { ...baseQueryGetUnbondCommissionRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetUnbondCommissionRequest };
        return message;
    }
};
const baseQueryGetUnbondCommissionResponse = { commission: '' };
export const QueryGetUnbondCommissionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.commission !== '') {
            writer.uint32(10).string(message.commission);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetUnbondCommissionResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.commission = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetUnbondCommissionResponse };
        if (object.commission !== undefined && object.commission !== null) {
            message.commission = String(object.commission);
        }
        else {
            message.commission = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.commission !== undefined && (obj.commission = message.commission);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetUnbondCommissionResponse };
        if (object.commission !== undefined && object.commission !== null) {
            message.commission = object.commission;
        }
        else {
            message.commission = '';
        }
        return message;
    }
};
const baseQueryGetLeastBondRequest = { denom: '' };
export const QueryGetLeastBondRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetLeastBondRequest };
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
        const message = { ...baseQueryGetLeastBondRequest };
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
        const message = { ...baseQueryGetLeastBondRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryGetLeastBondResponse = {};
export const QueryGetLeastBondResponse = {
    encode(message, writer = Writer.create()) {
        if (message.leastBond !== undefined) {
            LeastBond.encode(message.leastBond, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetLeastBondResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.leastBond = LeastBond.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetLeastBondResponse };
        if (object.leastBond !== undefined && object.leastBond !== null) {
            message.leastBond = LeastBond.fromJSON(object.leastBond);
        }
        else {
            message.leastBond = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.leastBond !== undefined && (obj.leastBond = message.leastBond ? LeastBond.toJSON(message.leastBond) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetLeastBondResponse };
        if (object.leastBond !== undefined && object.leastBond !== null) {
            message.leastBond = LeastBond.fromPartial(object.leastBond);
        }
        else {
            message.leastBond = undefined;
        }
        return message;
    }
};
const baseQueryGetEraUnbondLimitRequest = { denom: '' };
export const QueryGetEraUnbondLimitRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetEraUnbondLimitRequest };
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
        const message = { ...baseQueryGetEraUnbondLimitRequest };
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
        const message = { ...baseQueryGetEraUnbondLimitRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseQueryGetEraUnbondLimitResponse = { limit: 0 };
export const QueryGetEraUnbondLimitResponse = {
    encode(message, writer = Writer.create()) {
        if (message.limit !== 0) {
            writer.uint32(8).uint32(message.limit);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetEraUnbondLimitResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.limit = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetEraUnbondLimitResponse };
        if (object.limit !== undefined && object.limit !== null) {
            message.limit = Number(object.limit);
        }
        else {
            message.limit = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.limit !== undefined && (obj.limit = message.limit);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetEraUnbondLimitResponse };
        if (object.limit !== undefined && object.limit !== null) {
            message.limit = object.limit;
        }
        else {
            message.limit = 0;
        }
        return message;
    }
};
const baseQueryGetBondPipeLineRequest = { denom: '', pool: '' };
export const QueryGetBondPipeLineRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.pool !== '') {
            writer.uint32(18).string(message.pool);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetBondPipeLineRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.pool = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetBondPipeLineRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = String(object.pool);
        }
        else {
            message.pool = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.pool !== undefined && (obj.pool = message.pool);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetBondPipeLineRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = object.pool;
        }
        else {
            message.pool = '';
        }
        return message;
    }
};
const baseQueryGetBondPipeLineResponse = {};
export const QueryGetBondPipeLineResponse = {
    encode(message, writer = Writer.create()) {
        if (message.pipeline !== undefined) {
            BondPipeline.encode(message.pipeline, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetBondPipeLineResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pipeline = BondPipeline.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetBondPipeLineResponse };
        if (object.pipeline !== undefined && object.pipeline !== null) {
            message.pipeline = BondPipeline.fromJSON(object.pipeline);
        }
        else {
            message.pipeline = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pipeline !== undefined && (obj.pipeline = message.pipeline ? BondPipeline.toJSON(message.pipeline) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetBondPipeLineResponse };
        if (object.pipeline !== undefined && object.pipeline !== null) {
            message.pipeline = BondPipeline.fromPartial(object.pipeline);
        }
        else {
            message.pipeline = undefined;
        }
        return message;
    }
};
const baseQueryGetEraSnapshotRequest = { denom: '', era: 0 };
export const QueryGetEraSnapshotRequest = {
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
        const message = { ...baseQueryGetEraSnapshotRequest };
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
        const message = { ...baseQueryGetEraSnapshotRequest };
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
        const message = { ...baseQueryGetEraSnapshotRequest };
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
const baseQueryGetEraSnapshotResponse = {};
export const QueryGetEraSnapshotResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.shotIds) {
            writer.uint32(10).bytes(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetEraSnapshotResponse };
        message.shotIds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.shotIds.push(reader.bytes());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetEraSnapshotResponse };
        message.shotIds = [];
        if (object.shotIds !== undefined && object.shotIds !== null) {
            for (const e of object.shotIds) {
                message.shotIds.push(bytesFromBase64(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.shotIds) {
            obj.shotIds = message.shotIds.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
        }
        else {
            obj.shotIds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetEraSnapshotResponse };
        message.shotIds = [];
        if (object.shotIds !== undefined && object.shotIds !== null) {
            for (const e of object.shotIds) {
                message.shotIds.push(e);
            }
        }
        return message;
    }
};
const baseQueryGetSnapshotRequest = {};
export const QueryGetSnapshotRequest = {
    encode(message, writer = Writer.create()) {
        if (message.shotId.length !== 0) {
            writer.uint32(10).bytes(message.shotId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetSnapshotRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.shotId = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetSnapshotRequest };
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = bytesFromBase64(object.shotId);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.shotId !== undefined && (obj.shotId = base64FromBytes(message.shotId !== undefined ? message.shotId : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetSnapshotRequest };
        if (object.shotId !== undefined && object.shotId !== null) {
            message.shotId = object.shotId;
        }
        else {
            message.shotId = new Uint8Array();
        }
        return message;
    }
};
const baseQueryGetSnapshotResponse = {};
export const QueryGetSnapshotResponse = {
    encode(message, writer = Writer.create()) {
        if (message.shot !== undefined) {
            BondSnapshot.encode(message.shot, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetSnapshotResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.shot = BondSnapshot.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetSnapshotResponse };
        if (object.shot !== undefined && object.shot !== null) {
            message.shot = BondSnapshot.fromJSON(object.shot);
        }
        else {
            message.shot = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.shot !== undefined && (obj.shot = message.shot ? BondSnapshot.toJSON(message.shot) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetSnapshotResponse };
        if (object.shot !== undefined && object.shot !== null) {
            message.shot = BondSnapshot.fromPartial(object.shot);
        }
        else {
            message.shot = undefined;
        }
        return message;
    }
};
const baseQueryGetTotalExpectedActiveRequest = { denom: '', era: 0 };
export const QueryGetTotalExpectedActiveRequest = {
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
        const message = { ...baseQueryGetTotalExpectedActiveRequest };
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
        const message = { ...baseQueryGetTotalExpectedActiveRequest };
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
        const message = { ...baseQueryGetTotalExpectedActiveRequest };
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
const baseQueryGetTotalExpectedActiveResponse = { active: '' };
export const QueryGetTotalExpectedActiveResponse = {
    encode(message, writer = Writer.create()) {
        if (message.active !== '') {
            writer.uint32(10).string(message.active);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetTotalExpectedActiveResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.active = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetTotalExpectedActiveResponse };
        if (object.active !== undefined && object.active !== null) {
            message.active = String(object.active);
        }
        else {
            message.active = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.active !== undefined && (obj.active = message.active);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetTotalExpectedActiveResponse };
        if (object.active !== undefined && object.active !== null) {
            message.active = object.active;
        }
        else {
            message.active = '';
        }
        return message;
    }
};
const baseQueryGetPoolUnbondRequest = { denom: '', pool: '', era: 0 };
export const QueryGetPoolUnbondRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.pool !== '') {
            writer.uint32(18).string(message.pool);
        }
        if (message.era !== 0) {
            writer.uint32(24).uint32(message.era);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPoolUnbondRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.pool = reader.string();
                    break;
                case 3:
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
        const message = { ...baseQueryGetPoolUnbondRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = String(object.pool);
        }
        else {
            message.pool = '';
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
        message.pool !== undefined && (obj.pool = message.pool);
        message.era !== undefined && (obj.era = message.era);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPoolUnbondRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.pool !== undefined && object.pool !== null) {
            message.pool = object.pool;
        }
        else {
            message.pool = '';
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
const baseQueryGetPoolUnbondResponse = {};
export const QueryGetPoolUnbondResponse = {
    encode(message, writer = Writer.create()) {
        if (message.unbond !== undefined) {
            PoolUnbond.encode(message.unbond, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPoolUnbondResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.unbond = PoolUnbond.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPoolUnbondResponse };
        if (object.unbond !== undefined && object.unbond !== null) {
            message.unbond = PoolUnbond.fromJSON(object.unbond);
        }
        else {
            message.unbond = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.unbond !== undefined && (obj.unbond = message.unbond ? PoolUnbond.toJSON(message.unbond) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPoolUnbondResponse };
        if (object.unbond !== undefined && object.unbond !== null) {
            message.unbond = PoolUnbond.fromPartial(object.unbond);
        }
        else {
            message.unbond = undefined;
        }
        return message;
    }
};
const baseQueryGetAccountUnbondRequest = { denom: '', unbonder: '' };
export const QueryGetAccountUnbondRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.unbonder !== '') {
            writer.uint32(18).string(message.unbonder);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetAccountUnbondRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.unbonder = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetAccountUnbondRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.unbonder !== undefined && object.unbonder !== null) {
            message.unbonder = String(object.unbonder);
        }
        else {
            message.unbonder = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.unbonder !== undefined && (obj.unbonder = message.unbonder);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetAccountUnbondRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.unbonder !== undefined && object.unbonder !== null) {
            message.unbonder = object.unbonder;
        }
        else {
            message.unbonder = '';
        }
        return message;
    }
};
const baseQueryGetAccountUnbondResponse = {};
export const QueryGetAccountUnbondResponse = {
    encode(message, writer = Writer.create()) {
        if (message.unbond !== undefined) {
            AccountUnbond.encode(message.unbond, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetAccountUnbondResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.unbond = AccountUnbond.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetAccountUnbondResponse };
        if (object.unbond !== undefined && object.unbond !== null) {
            message.unbond = AccountUnbond.fromJSON(object.unbond);
        }
        else {
            message.unbond = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.unbond !== undefined && (obj.unbond = message.unbond ? AccountUnbond.toJSON(message.unbond) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetAccountUnbondResponse };
        if (object.unbond !== undefined && object.unbond !== null) {
            message.unbond = AccountUnbond.fromPartial(object.unbond);
        }
        else {
            message.unbond = undefined;
        }
        return message;
    }
};
const baseQueryGetBondRecordRequest = { denom: '', blockhash: '', txhash: '' };
export const QueryGetBondRecordRequest = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.blockhash !== '') {
            writer.uint32(18).string(message.blockhash);
        }
        if (message.txhash !== '') {
            writer.uint32(26).string(message.txhash);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetBondRecordRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.blockhash = reader.string();
                    break;
                case 3:
                    message.txhash = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetBondRecordRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.blockhash !== undefined && (obj.blockhash = message.blockhash);
        message.txhash !== undefined && (obj.txhash = message.txhash);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetBondRecordRequest };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
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
        return message;
    }
};
const baseQueryGetBondRecordResponse = {};
export const QueryGetBondRecordResponse = {
    encode(message, writer = Writer.create()) {
        if (message.bondRecord !== undefined) {
            BondRecord.encode(message.bondRecord, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetBondRecordResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.bondRecord = BondRecord.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetBondRecordResponse };
        if (object.bondRecord !== undefined && object.bondRecord !== null) {
            message.bondRecord = BondRecord.fromJSON(object.bondRecord);
        }
        else {
            message.bondRecord = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.bondRecord !== undefined && (obj.bondRecord = message.bondRecord ? BondRecord.toJSON(message.bondRecord) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetBondRecordResponse };
        if (object.bondRecord !== undefined && object.bondRecord !== null) {
            message.bondRecord = BondRecord.fromPartial(object.bondRecord);
        }
        else {
            message.bondRecord = undefined;
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
    PoolsByDenom(request) {
        const data = QueryPoolsByDenomRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'PoolsByDenom', data);
        return promise.then((data) => QueryPoolsByDenomResponse.decode(new Reader(data)));
    }
    BondedPoolsByDenom(request) {
        const data = QueryBondedPoolsByDenomRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'BondedPoolsByDenom', data);
        return promise.then((data) => QueryBondedPoolsByDenomResponse.decode(new Reader(data)));
    }
    GetPoolDetail(request) {
        const data = QueryGetPoolDetailRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetPoolDetail', data);
        return promise.then((data) => QueryGetPoolDetailResponse.decode(new Reader(data)));
    }
    GetChainEra(request) {
        const data = QueryGetChainEraRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetChainEra', data);
        return promise.then((data) => QueryGetChainEraResponse.decode(new Reader(data)));
    }
    GetCurrentEraSnapshot(request) {
        const data = QueryGetCurrentEraSnapshotRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetCurrentEraSnapshot', data);
        return promise.then((data) => QueryGetCurrentEraSnapshotResponse.decode(new Reader(data)));
    }
    GetReceiver(request) {
        const data = QueryGetReceiverRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetReceiver', data);
        return promise.then((data) => QueryGetReceiverResponse.decode(new Reader(data)));
    }
    GetCommission(request) {
        const data = QueryGetCommissionRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetCommission', data);
        return promise.then((data) => QueryGetCommissionResponse.decode(new Reader(data)));
    }
    GetChainBondingDuration(request) {
        const data = QueryGetChainBondingDurationRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetChainBondingDuration', data);
        return promise.then((data) => QueryGetChainBondingDurationResponse.decode(new Reader(data)));
    }
    GetUnbondFee(request) {
        const data = QueryGetUnbondFeeRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetUnbondFee', data);
        return promise.then((data) => QueryGetUnbondFeeResponse.decode(new Reader(data)));
    }
    GetUnbondCommission(request) {
        const data = QueryGetUnbondCommissionRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetUnbondCommission', data);
        return promise.then((data) => QueryGetUnbondCommissionResponse.decode(new Reader(data)));
    }
    GetLeastBond(request) {
        const data = QueryGetLeastBondRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetLeastBond', data);
        return promise.then((data) => QueryGetLeastBondResponse.decode(new Reader(data)));
    }
    GetEraUnbondLimit(request) {
        const data = QueryGetEraUnbondLimitRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetEraUnbondLimit', data);
        return promise.then((data) => QueryGetEraUnbondLimitResponse.decode(new Reader(data)));
    }
    GetBondPipeLine(request) {
        const data = QueryGetBondPipeLineRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetBondPipeLine', data);
        return promise.then((data) => QueryGetBondPipeLineResponse.decode(new Reader(data)));
    }
    GetEraSnapshot(request) {
        const data = QueryGetEraSnapshotRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetEraSnapshot', data);
        return promise.then((data) => QueryGetEraSnapshotResponse.decode(new Reader(data)));
    }
    GetSnapshot(request) {
        const data = QueryGetSnapshotRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetSnapshot', data);
        return promise.then((data) => QueryGetSnapshotResponse.decode(new Reader(data)));
    }
    GetTotalExpectedActive(request) {
        const data = QueryGetTotalExpectedActiveRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetTotalExpectedActive', data);
        return promise.then((data) => QueryGetTotalExpectedActiveResponse.decode(new Reader(data)));
    }
    GetPoolUnbond(request) {
        const data = QueryGetPoolUnbondRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetPoolUnbond', data);
        return promise.then((data) => QueryGetPoolUnbondResponse.decode(new Reader(data)));
    }
    GetAccountUnbond(request) {
        const data = QueryGetAccountUnbondRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetAccountUnbond', data);
        return promise.then((data) => QueryGetAccountUnbondResponse.decode(new Reader(data)));
    }
    GetBondRecord(request) {
        const data = QueryGetBondRecordRequest.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Query', 'GetBondRecord', data);
        return promise.then((data) => QueryGetBondRecordResponse.decode(new Reader(data)));
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
