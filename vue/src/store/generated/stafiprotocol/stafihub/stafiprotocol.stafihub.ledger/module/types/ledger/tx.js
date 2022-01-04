/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'stafiprotocol.stafihub.ledger';
const baseMsgAddNewPool = { creator: '', denom: '', addr: '' };
export const MsgAddNewPool = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.addr !== '') {
            writer.uint32(26).string(message.addr);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddNewPool };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.addr = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAddNewPool };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.addr !== undefined && object.addr !== null) {
            message.addr = String(object.addr);
        }
        else {
            message.addr = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        message.addr !== undefined && (obj.addr = message.addr);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddNewPool };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.addr !== undefined && object.addr !== null) {
            message.addr = object.addr;
        }
        else {
            message.addr = '';
        }
        return message;
    }
};
const baseMsgAddNewPoolResponse = {};
export const MsgAddNewPoolResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddNewPoolResponse };
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
        const message = { ...baseMsgAddNewPoolResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddNewPoolResponse };
        return message;
    }
};
const baseMsgRemovePool = { creator: '', denom: '', addr: '' };
export const MsgRemovePool = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.addr !== '') {
            writer.uint32(26).string(message.addr);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRemovePool };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.addr = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgRemovePool };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.addr !== undefined && object.addr !== null) {
            message.addr = String(object.addr);
        }
        else {
            message.addr = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        message.addr !== undefined && (obj.addr = message.addr);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRemovePool };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.addr !== undefined && object.addr !== null) {
            message.addr = object.addr;
        }
        else {
            message.addr = '';
        }
        return message;
    }
};
const baseMsgRemovePoolResponse = {};
export const MsgRemovePoolResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRemovePoolResponse };
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
        const message = { ...baseMsgRemovePoolResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRemovePoolResponse };
        return message;
    }
};
const baseMsgSetEraUnbondLimit = { creator: '', denom: '', limit: 0 };
export const MsgSetEraUnbondLimit = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.limit !== 0) {
            writer.uint32(24).uint32(message.limit);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetEraUnbondLimit };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
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
        const message = { ...baseMsgSetEraUnbondLimit };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        message.limit !== undefined && (obj.limit = message.limit);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetEraUnbondLimit };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.limit !== undefined && object.limit !== null) {
            message.limit = object.limit;
        }
        else {
            message.limit = 0;
        }
        return message;
    }
};
const baseMsgSetEraUnbondLimitResponse = {};
export const MsgSetEraUnbondLimitResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetEraUnbondLimitResponse };
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
        const message = { ...baseMsgSetEraUnbondLimitResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetEraUnbondLimitResponse };
        return message;
    }
};
const baseMsgSetInitBond = { creator: '', denom: '', pool: '', amount: '', receiver: '' };
export const MsgSetInitBond = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.pool !== '') {
            writer.uint32(26).string(message.pool);
        }
        if (message.amount !== '') {
            writer.uint32(34).string(message.amount);
        }
        if (message.receiver !== '') {
            writer.uint32(42).string(message.receiver);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetInitBond };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.pool = reader.string();
                    break;
                case 4:
                    message.amount = reader.string();
                    break;
                case 5:
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
        const message = { ...baseMsgSetInitBond };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = String(object.amount);
        }
        else {
            message.amount = '';
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        message.pool !== undefined && (obj.pool = message.pool);
        message.amount !== undefined && (obj.amount = message.amount);
        message.receiver !== undefined && (obj.receiver = message.receiver);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetInitBond };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
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
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = '';
        }
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = object.receiver;
        }
        else {
            message.receiver = '';
        }
        return message;
    }
};
const baseMsgSetInitBondResponse = {};
export const MsgSetInitBondResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetInitBondResponse };
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
        const message = { ...baseMsgSetInitBondResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetInitBondResponse };
        return message;
    }
};
const baseMsgSetChainBondingDuration = { creator: '', denom: '', era: 0 };
export const MsgSetChainBondingDuration = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.era !== 0) {
            writer.uint32(24).uint32(message.era);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetChainBondingDuration };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
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
        const message = { ...baseMsgSetChainBondingDuration };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        message.era !== undefined && (obj.era = message.era);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetChainBondingDuration };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
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
        return message;
    }
};
const baseMsgSetChainBondingDurationResponse = {};
export const MsgSetChainBondingDurationResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetChainBondingDurationResponse };
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
        const message = { ...baseMsgSetChainBondingDurationResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetChainBondingDurationResponse };
        return message;
    }
};
const baseMsgSetPoolDetail = { creator: '', denom: '', pool: '', subAccounts: '', threshold: 0 };
export const MsgSetPoolDetail = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.pool !== '') {
            writer.uint32(26).string(message.pool);
        }
        for (const v of message.subAccounts) {
            writer.uint32(34).string(v);
        }
        if (message.threshold !== 0) {
            writer.uint32(40).uint32(message.threshold);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetPoolDetail };
        message.subAccounts = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.pool = reader.string();
                    break;
                case 4:
                    message.subAccounts.push(reader.string());
                    break;
                case 5:
                    message.threshold = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSetPoolDetail };
        message.subAccounts = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        if (object.subAccounts !== undefined && object.subAccounts !== null) {
            for (const e of object.subAccounts) {
                message.subAccounts.push(String(e));
            }
        }
        if (object.threshold !== undefined && object.threshold !== null) {
            message.threshold = Number(object.threshold);
        }
        else {
            message.threshold = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        message.pool !== undefined && (obj.pool = message.pool);
        if (message.subAccounts) {
            obj.subAccounts = message.subAccounts.map((e) => e);
        }
        else {
            obj.subAccounts = [];
        }
        message.threshold !== undefined && (obj.threshold = message.threshold);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetPoolDetail };
        message.subAccounts = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
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
        if (object.subAccounts !== undefined && object.subAccounts !== null) {
            for (const e of object.subAccounts) {
                message.subAccounts.push(e);
            }
        }
        if (object.threshold !== undefined && object.threshold !== null) {
            message.threshold = object.threshold;
        }
        else {
            message.threshold = 0;
        }
        return message;
    }
};
const baseMsgSetPoolDetailResponse = {};
export const MsgSetPoolDetailResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetPoolDetailResponse };
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
        const message = { ...baseMsgSetPoolDetailResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetPoolDetailResponse };
        return message;
    }
};
const baseMsgSetLeastBond = { creator: '', denom: '', amount: '' };
export const MsgSetLeastBond = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        if (message.amount !== '') {
            writer.uint32(26).string(message.amount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetLeastBond };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.denom = reader.string();
                    break;
                case 3:
                    message.amount = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgSetLeastBond };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = String(object.amount);
        }
        else {
            message.amount = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        message.amount !== undefined && (obj.amount = message.amount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetLeastBond };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = '';
        }
        return message;
    }
};
const baseMsgSetLeastBondResponse = {};
export const MsgSetLeastBondResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetLeastBondResponse };
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
        const message = { ...baseMsgSetLeastBondResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetLeastBondResponse };
        return message;
    }
};
const baseMsgClearCurrentEraSnapShots = { creator: '', denom: '' };
export const MsgClearCurrentEraSnapShots = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.denom !== '') {
            writer.uint32(18).string(message.denom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgClearCurrentEraSnapShots };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        const message = { ...baseMsgClearCurrentEraSnapShots };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.denom !== undefined && (obj.denom = message.denom);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgClearCurrentEraSnapShots };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        return message;
    }
};
const baseMsgClearCurrentEraSnapShotsResponse = {};
export const MsgClearCurrentEraSnapShotsResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgClearCurrentEraSnapShotsResponse };
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
        const message = { ...baseMsgClearCurrentEraSnapShotsResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgClearCurrentEraSnapShotsResponse };
        return message;
    }
};
const baseMsgSetCommission = { creator: '', commission: '' };
export const MsgSetCommission = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.commission !== '') {
            writer.uint32(18).string(message.commission);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetCommission };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        const message = { ...baseMsgSetCommission };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.commission !== undefined && (obj.commission = message.commission);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetCommission };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.commission !== undefined && object.commission !== null) {
            message.commission = object.commission;
        }
        else {
            message.commission = '';
        }
        return message;
    }
};
const baseMsgSetCommissionResponse = {};
export const MsgSetCommissionResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetCommissionResponse };
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
        const message = { ...baseMsgSetCommissionResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetCommissionResponse };
        return message;
    }
};
const baseMsgSetReceiver = { creator: '', receiver: '' };
export const MsgSetReceiver = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.receiver !== '') {
            writer.uint32(18).string(message.receiver);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetReceiver };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
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
        const message = { ...baseMsgSetReceiver };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.receiver !== undefined && (obj.receiver = message.receiver);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgSetReceiver };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.receiver !== undefined && object.receiver !== null) {
            message.receiver = object.receiver;
        }
        else {
            message.receiver = '';
        }
        return message;
    }
};
const baseMsgSetReceiverResponse = {};
export const MsgSetReceiverResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgSetReceiverResponse };
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
        const message = { ...baseMsgSetReceiverResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgSetReceiverResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    AddNewPool(request) {
        const data = MsgAddNewPool.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'AddNewPool', data);
        return promise.then((data) => MsgAddNewPoolResponse.decode(new Reader(data)));
    }
    RemovePool(request) {
        const data = MsgRemovePool.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'RemovePool', data);
        return promise.then((data) => MsgRemovePoolResponse.decode(new Reader(data)));
    }
    SetEraUnbondLimit(request) {
        const data = MsgSetEraUnbondLimit.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetEraUnbondLimit', data);
        return promise.then((data) => MsgSetEraUnbondLimitResponse.decode(new Reader(data)));
    }
    SetInitBond(request) {
        const data = MsgSetInitBond.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetInitBond', data);
        return promise.then((data) => MsgSetInitBondResponse.decode(new Reader(data)));
    }
    SetChainBondingDuration(request) {
        const data = MsgSetChainBondingDuration.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetChainBondingDuration', data);
        return promise.then((data) => MsgSetChainBondingDurationResponse.decode(new Reader(data)));
    }
    SetPoolDetail(request) {
        const data = MsgSetPoolDetail.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetPoolDetail', data);
        return promise.then((data) => MsgSetPoolDetailResponse.decode(new Reader(data)));
    }
    SetLeastBond(request) {
        const data = MsgSetLeastBond.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetLeastBond', data);
        return promise.then((data) => MsgSetLeastBondResponse.decode(new Reader(data)));
    }
    ClearCurrentEraSnapShots(request) {
        const data = MsgClearCurrentEraSnapShots.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'ClearCurrentEraSnapShots', data);
        return promise.then((data) => MsgClearCurrentEraSnapShotsResponse.decode(new Reader(data)));
    }
    SetCommission(request) {
        const data = MsgSetCommission.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetCommission', data);
        return promise.then((data) => MsgSetCommissionResponse.decode(new Reader(data)));
    }
    SetReceiver(request) {
        const data = MsgSetReceiver.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.ledger.Msg', 'SetReceiver', data);
        return promise.then((data) => MsgSetReceiverResponse.decode(new Reader(data)));
    }
}
