/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Metadata } from '../cosmos/bank/v1beta1/bank';
export const protobufPackage = 'stafiprotocol.stafihub.sudo';
const baseMsgUpdateAdmin = { creator: '', address: '' };
export const MsgUpdateAdmin = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateAdmin };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateAdmin };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateAdmin };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseMsgUpdateAdminResponse = {};
export const MsgUpdateAdminResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateAdminResponse };
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
        const message = { ...baseMsgUpdateAdminResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUpdateAdminResponse };
        return message;
    }
};
const baseMsgAddDenom = { creator: '' };
export const MsgAddDenom = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.Metadata !== undefined) {
            Metadata.encode(message.Metadata, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddDenom };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.Metadata = Metadata.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAddDenom };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.Metadata !== undefined && object.Metadata !== null) {
            message.Metadata = Metadata.fromJSON(object.Metadata);
        }
        else {
            message.Metadata = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.Metadata !== undefined && (obj.Metadata = message.Metadata ? Metadata.toJSON(message.Metadata) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddDenom };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.Metadata !== undefined && object.Metadata !== null) {
            message.Metadata = Metadata.fromPartial(object.Metadata);
        }
        else {
            message.Metadata = undefined;
        }
        return message;
    }
};
const baseMsgAddDenomResponse = {};
export const MsgAddDenomResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddDenomResponse };
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
        const message = { ...baseMsgAddDenomResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddDenomResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    UpdateAdmin(request) {
        const data = MsgUpdateAdmin.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.sudo.Msg', 'UpdateAdmin', data);
        return promise.then((data) => MsgUpdateAdminResponse.decode(new Reader(data)));
    }
    AddDenom(request) {
        const data = MsgAddDenom.encode(request).finish();
        const promise = this.rpc.request('stafiprotocol.stafihub.sudo.Msg', 'AddDenom', data);
        return promise.then((data) => MsgAddDenomResponse.decode(new Reader(data)));
    }
}
