/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'stafiprotocol.stafihub.ledger';
export var PoolBondState;
(function (PoolBondState) {
    PoolBondState[PoolBondState["ERA_UPDATED"] = 0] = "ERA_UPDATED";
    PoolBondState[PoolBondState["BOND_REPORTED"] = 1] = "BOND_REPORTED";
    PoolBondState[PoolBondState["ACTIVE_REPORTED"] = 2] = "ACTIVE_REPORTED";
    PoolBondState[PoolBondState["WITHDRAW_SKIPPED"] = 3] = "WITHDRAW_SKIPPED";
    PoolBondState[PoolBondState["WITHDRAW_REPORTED"] = 4] = "WITHDRAW_REPORTED";
    PoolBondState[PoolBondState["TRANSFER_REPORTED"] = 5] = "TRANSFER_REPORTED";
    PoolBondState[PoolBondState["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(PoolBondState || (PoolBondState = {}));
export function poolBondStateFromJSON(object) {
    switch (object) {
        case 0:
        case 'ERA_UPDATED':
            return PoolBondState.ERA_UPDATED;
        case 1:
        case 'BOND_REPORTED':
            return PoolBondState.BOND_REPORTED;
        case 2:
        case 'ACTIVE_REPORTED':
            return PoolBondState.ACTIVE_REPORTED;
        case 3:
        case 'WITHDRAW_SKIPPED':
            return PoolBondState.WITHDRAW_SKIPPED;
        case 4:
        case 'WITHDRAW_REPORTED':
            return PoolBondState.WITHDRAW_REPORTED;
        case 5:
        case 'TRANSFER_REPORTED':
            return PoolBondState.TRANSFER_REPORTED;
        case -1:
        case 'UNRECOGNIZED':
        default:
            return PoolBondState.UNRECOGNIZED;
    }
}
export function poolBondStateToJSON(object) {
    switch (object) {
        case PoolBondState.ERA_UPDATED:
            return 'ERA_UPDATED';
        case PoolBondState.BOND_REPORTED:
            return 'BOND_REPORTED';
        case PoolBondState.ACTIVE_REPORTED:
            return 'ACTIVE_REPORTED';
        case PoolBondState.WITHDRAW_SKIPPED:
            return 'WITHDRAW_SKIPPED';
        case PoolBondState.WITHDRAW_REPORTED:
            return 'WITHDRAW_REPORTED';
        case PoolBondState.TRANSFER_REPORTED:
            return 'TRANSFER_REPORTED';
        default:
            return 'UNKNOWN';
    }
}
export var BondAction;
(function (BondAction) {
    BondAction[BondAction["BOND_ONLY"] = 0] = "BOND_ONLY";
    BondAction[BondAction["UNBOND_ONLY"] = 1] = "UNBOND_ONLY";
    BondAction[BondAction["BOTH_BOND_UNBOND"] = 2] = "BOTH_BOND_UNBOND";
    BondAction[BondAction["EITHER_BOND_UNBOND"] = 3] = "EITHER_BOND_UNBOND";
    BondAction[BondAction["INTER_DEDUCT"] = 4] = "INTER_DEDUCT";
    BondAction[BondAction["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(BondAction || (BondAction = {}));
export function bondActionFromJSON(object) {
    switch (object) {
        case 0:
        case 'BOND_ONLY':
            return BondAction.BOND_ONLY;
        case 1:
        case 'UNBOND_ONLY':
            return BondAction.UNBOND_ONLY;
        case 2:
        case 'BOTH_BOND_UNBOND':
            return BondAction.BOTH_BOND_UNBOND;
        case 3:
        case 'EITHER_BOND_UNBOND':
            return BondAction.EITHER_BOND_UNBOND;
        case 4:
        case 'INTER_DEDUCT':
            return BondAction.INTER_DEDUCT;
        case -1:
        case 'UNRECOGNIZED':
        default:
            return BondAction.UNRECOGNIZED;
    }
}
export function bondActionToJSON(object) {
    switch (object) {
        case BondAction.BOND_ONLY:
            return 'BOND_ONLY';
        case BondAction.UNBOND_ONLY:
            return 'UNBOND_ONLY';
        case BondAction.BOTH_BOND_UNBOND:
            return 'BOTH_BOND_UNBOND';
        case BondAction.EITHER_BOND_UNBOND:
            return 'EITHER_BOND_UNBOND';
        case BondAction.INTER_DEDUCT:
            return 'INTER_DEDUCT';
        default:
            return 'UNKNOWN';
    }
}
const baseChainEra = { denom: '', era: 0 };
export const ChainEra = {
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
        const message = { ...baseChainEra };
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
        const message = { ...baseChainEra };
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
        const message = { ...baseChainEra };
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
const baseChainBondingDuration = { denom: '', era: 0 };
export const ChainBondingDuration = {
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
        const message = { ...baseChainBondingDuration };
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
        const message = { ...baseChainBondingDuration };
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
        const message = { ...baseChainBondingDuration };
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
const basePool = { denom: '' };
export const Pool = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        Object.entries(message.addrs).forEach(([key, value]) => {
            Pool_AddrsEntry.encode({ key: key, value }, writer.uint32(18).fork()).ldelim();
        });
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePool };
        message.addrs = {};
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    const entry2 = Pool_AddrsEntry.decode(reader, reader.uint32());
                    if (entry2.value !== undefined) {
                        message.addrs[entry2.key] = entry2.value;
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
        const message = { ...basePool };
        message.addrs = {};
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.addrs !== undefined && object.addrs !== null) {
            Object.entries(object.addrs).forEach(([key, value]) => {
                message.addrs[key] = Boolean(value);
            });
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        obj.addrs = {};
        if (message.addrs) {
            Object.entries(message.addrs).forEach(([k, v]) => {
                obj.addrs[k] = v;
            });
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePool };
        message.addrs = {};
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.addrs !== undefined && object.addrs !== null) {
            Object.entries(object.addrs).forEach(([key, value]) => {
                if (value !== undefined) {
                    message.addrs[key] = Boolean(value);
                }
            });
        }
        return message;
    }
};
const basePool_AddrsEntry = { key: '', value: false };
export const Pool_AddrsEntry = {
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
        const message = { ...basePool_AddrsEntry };
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
        const message = { ...basePool_AddrsEntry };
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
        const message = { ...basePool_AddrsEntry };
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
const baseTotalExpectedActive = { denom: '', era: '', amount: '' };
export const TotalExpectedActive = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.era !== '') {
            writer.uint32(18).string(message.era);
        }
        if (message.amount !== '') {
            writer.uint32(26).string(message.amount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseTotalExpectedActive };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.era = reader.string();
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
        const message = { ...baseTotalExpectedActive };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.era !== undefined && object.era !== null) {
            message.era = String(object.era);
        }
        else {
            message.era = '';
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
        message.denom !== undefined && (obj.denom = message.denom);
        message.era !== undefined && (obj.era = message.era);
        message.amount !== undefined && (obj.amount = message.amount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseTotalExpectedActive };
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
            message.era = '';
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
const baseBondPipeline = { denom: '', pool: '' };
export const BondPipeline = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.pool !== '') {
            writer.uint32(18).string(message.pool);
        }
        if (message.chunk !== undefined) {
            LinkChunk.encode(message.chunk, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBondPipeline };
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
                    message.chunk = LinkChunk.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseBondPipeline };
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
        if (object.chunk !== undefined && object.chunk !== null) {
            message.chunk = LinkChunk.fromJSON(object.chunk);
        }
        else {
            message.chunk = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.pool !== undefined && (obj.pool = message.pool);
        message.chunk !== undefined && (obj.chunk = message.chunk ? LinkChunk.toJSON(message.chunk) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseBondPipeline };
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
        if (object.chunk !== undefined && object.chunk !== null) {
            message.chunk = LinkChunk.fromPartial(object.chunk);
        }
        else {
            message.chunk = undefined;
        }
        return message;
    }
};
const baseEraSnapShot = { denom: '' };
export const EraSnapShot = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        for (const v of message.shotIds) {
            writer.uint32(26).bytes(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEraSnapShot };
        message.shotIds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 3:
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
        const message = { ...baseEraSnapShot };
        message.shotIds = [];
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.shotIds !== undefined && object.shotIds !== null) {
            for (const e of object.shotIds) {
                message.shotIds.push(bytesFromBase64(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        if (message.shotIds) {
            obj.shotIds = message.shotIds.map((e) => base64FromBytes(e !== undefined ? e : new Uint8Array()));
        }
        else {
            obj.shotIds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEraSnapShot };
        message.shotIds = [];
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.shotIds !== undefined && object.shotIds !== null) {
            for (const e of object.shotIds) {
                message.shotIds.push(e);
            }
        }
        return message;
    }
};
const basePoolUnbond = { denom: '', pool: '', era: 0 };
export const PoolUnbond = {
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
        for (const v of message.unbondings) {
            Unbonding.encode(v, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePoolUnbond };
        message.unbondings = [];
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
                case 4:
                    message.unbondings.push(Unbonding.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...basePoolUnbond };
        message.unbondings = [];
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
        if (object.unbondings !== undefined && object.unbondings !== null) {
            for (const e of object.unbondings) {
                message.unbondings.push(Unbonding.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.pool !== undefined && (obj.pool = message.pool);
        message.era !== undefined && (obj.era = message.era);
        if (message.unbondings) {
            obj.unbondings = message.unbondings.map((e) => (e ? Unbonding.toJSON(e) : undefined));
        }
        else {
            obj.unbondings = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePoolUnbond };
        message.unbondings = [];
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
        if (object.unbondings !== undefined && object.unbondings !== null) {
            for (const e of object.unbondings) {
                message.unbondings.push(Unbonding.fromPartial(e));
            }
        }
        return message;
    }
};
const baseEraUnbondLimit = { denom: '', limit: 0 };
export const EraUnbondLimit = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.limit !== 0) {
            writer.uint32(16).uint32(message.limit);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEraUnbondLimit };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
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
        const message = { ...baseEraUnbondLimit };
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
        message.denom !== undefined && (obj.denom = message.denom);
        message.limit !== undefined && (obj.limit = message.limit);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEraUnbondLimit };
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
const basePoolDetail = { denom: '', pool: '', subAccounts: '', threshold: 0 };
export const PoolDetail = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.pool !== '') {
            writer.uint32(18).string(message.pool);
        }
        for (const v of message.subAccounts) {
            writer.uint32(26).string(v);
        }
        if (message.threshold !== 0) {
            writer.uint32(32).uint32(message.threshold);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePoolDetail };
        message.subAccounts = [];
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
                    message.subAccounts.push(reader.string());
                    break;
                case 4:
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
        const message = { ...basePoolDetail };
        message.subAccounts = [];
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
        const message = { ...basePoolDetail };
        message.subAccounts = [];
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
const baseLeastBond = { denom: '', amount: '' };
export const LeastBond = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.amount !== '') {
            writer.uint32(18).string(message.amount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseLeastBond };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
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
        const message = { ...baseLeastBond };
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
        message.denom !== undefined && (obj.denom = message.denom);
        message.amount !== undefined && (obj.amount = message.amount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseLeastBond };
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
const baseLinkChunk = { bond: '', unbond: '', active: '' };
export const LinkChunk = {
    encode(message, writer = Writer.create()) {
        if (message.bond !== '') {
            writer.uint32(10).string(message.bond);
        }
        if (message.unbond !== '') {
            writer.uint32(18).string(message.unbond);
        }
        if (message.active !== '') {
            writer.uint32(26).string(message.active);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseLinkChunk };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.bond = reader.string();
                    break;
                case 2:
                    message.unbond = reader.string();
                    break;
                case 3:
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
        const message = { ...baseLinkChunk };
        if (object.bond !== undefined && object.bond !== null) {
            message.bond = String(object.bond);
        }
        else {
            message.bond = '';
        }
        if (object.unbond !== undefined && object.unbond !== null) {
            message.unbond = String(object.unbond);
        }
        else {
            message.unbond = '';
        }
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
        message.bond !== undefined && (obj.bond = message.bond);
        message.unbond !== undefined && (obj.unbond = message.unbond);
        message.active !== undefined && (obj.active = message.active);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseLinkChunk };
        if (object.bond !== undefined && object.bond !== null) {
            message.bond = object.bond;
        }
        else {
            message.bond = '';
        }
        if (object.unbond !== undefined && object.unbond !== null) {
            message.unbond = object.unbond;
        }
        else {
            message.unbond = '';
        }
        if (object.active !== undefined && object.active !== null) {
            message.active = object.active;
        }
        else {
            message.active = '';
        }
        return message;
    }
};
const baseBondSnapshot = { denom: '', pool: '', era: 0, lastVoter: '', bondState: 0 };
export const BondSnapshot = {
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
        if (message.chunk !== undefined) {
            LinkChunk.encode(message.chunk, writer.uint32(34).fork()).ldelim();
        }
        if (message.lastVoter !== '') {
            writer.uint32(42).string(message.lastVoter);
        }
        if (message.bondState !== 0) {
            writer.uint32(48).int32(message.bondState);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBondSnapshot };
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
                case 4:
                    message.chunk = LinkChunk.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.lastVoter = reader.string();
                    break;
                case 6:
                    message.bondState = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseBondSnapshot };
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
        if (object.chunk !== undefined && object.chunk !== null) {
            message.chunk = LinkChunk.fromJSON(object.chunk);
        }
        else {
            message.chunk = undefined;
        }
        if (object.lastVoter !== undefined && object.lastVoter !== null) {
            message.lastVoter = String(object.lastVoter);
        }
        else {
            message.lastVoter = '';
        }
        if (object.bondState !== undefined && object.bondState !== null) {
            message.bondState = poolBondStateFromJSON(object.bondState);
        }
        else {
            message.bondState = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.pool !== undefined && (obj.pool = message.pool);
        message.era !== undefined && (obj.era = message.era);
        message.chunk !== undefined && (obj.chunk = message.chunk ? LinkChunk.toJSON(message.chunk) : undefined);
        message.lastVoter !== undefined && (obj.lastVoter = message.lastVoter);
        message.bondState !== undefined && (obj.bondState = poolBondStateToJSON(message.bondState));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseBondSnapshot };
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
        if (object.chunk !== undefined && object.chunk !== null) {
            message.chunk = LinkChunk.fromPartial(object.chunk);
        }
        else {
            message.chunk = undefined;
        }
        if (object.lastVoter !== undefined && object.lastVoter !== null) {
            message.lastVoter = object.lastVoter;
        }
        else {
            message.lastVoter = '';
        }
        if (object.bondState !== undefined && object.bondState !== null) {
            message.bondState = object.bondState;
        }
        else {
            message.bondState = 0;
        }
        return message;
    }
};
const baseUnbonding = { unbonder: '', amount: '', recipient: '' };
export const Unbonding = {
    encode(message, writer = Writer.create()) {
        if (message.unbonder !== '') {
            writer.uint32(10).string(message.unbonder);
        }
        if (message.amount !== '') {
            writer.uint32(18).string(message.amount);
        }
        if (message.recipient !== '') {
            writer.uint32(26).string(message.recipient);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseUnbonding };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.unbonder = reader.string();
                    break;
                case 2:
                    message.amount = reader.string();
                    break;
                case 3:
                    message.recipient = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseUnbonding };
        if (object.unbonder !== undefined && object.unbonder !== null) {
            message.unbonder = String(object.unbonder);
        }
        else {
            message.unbonder = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = String(object.amount);
        }
        else {
            message.amount = '';
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = String(object.recipient);
        }
        else {
            message.recipient = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.unbonder !== undefined && (obj.unbonder = message.unbonder);
        message.amount !== undefined && (obj.amount = message.amount);
        message.recipient !== undefined && (obj.recipient = message.recipient);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseUnbonding };
        if (object.unbonder !== undefined && object.unbonder !== null) {
            message.unbonder = object.unbonder;
        }
        else {
            message.unbonder = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = '';
        }
        if (object.recipient !== undefined && object.recipient !== null) {
            message.recipient = object.recipient;
        }
        else {
            message.recipient = '';
        }
        return message;
    }
};
const baseExchangeRate = { denom: '', value: '' };
export const ExchangeRate = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.value !== '') {
            writer.uint32(18).string(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseExchangeRate };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.value = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseExchangeRate };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = String(object.value);
        }
        else {
            message.value = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseExchangeRate };
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = '';
        }
        return message;
    }
};
const baseEraExchangeRate = { denom: '', era: 0, value: '' };
export const EraExchangeRate = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== '') {
            writer.uint32(10).string(message.denom);
        }
        if (message.era !== 0) {
            writer.uint32(16).uint32(message.era);
        }
        if (message.value !== '') {
            writer.uint32(26).string(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEraExchangeRate };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.era = reader.uint32();
                    break;
                case 3:
                    message.value = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseEraExchangeRate };
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
        if (object.value !== undefined && object.value !== null) {
            message.value = String(object.value);
        }
        else {
            message.value = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.era !== undefined && (obj.era = message.era);
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEraExchangeRate };
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
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = '';
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
