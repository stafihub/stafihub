import { Reader, Writer } from 'protobufjs/minimal';
import { Metadata } from '../cosmos/bank/v1beta1/bank';
export declare const protobufPackage = "stafiprotocol.stafihub.sudo";
export interface MsgUpdateAdmin {
    creator: string;
    address: string;
}
export interface MsgUpdateAdminResponse {
}
export interface MsgAddDenom {
    creator: string;
    Metadata: Metadata | undefined;
}
export interface MsgAddDenomResponse {
}
export declare const MsgUpdateAdmin: {
    encode(message: MsgUpdateAdmin, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateAdmin;
    fromJSON(object: any): MsgUpdateAdmin;
    toJSON(message: MsgUpdateAdmin): unknown;
    fromPartial(object: DeepPartial<MsgUpdateAdmin>): MsgUpdateAdmin;
};
export declare const MsgUpdateAdminResponse: {
    encode(_: MsgUpdateAdminResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateAdminResponse;
    fromJSON(_: any): MsgUpdateAdminResponse;
    toJSON(_: MsgUpdateAdminResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateAdminResponse>): MsgUpdateAdminResponse;
};
export declare const MsgAddDenom: {
    encode(message: MsgAddDenom, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddDenom;
    fromJSON(object: any): MsgAddDenom;
    toJSON(message: MsgAddDenom): unknown;
    fromPartial(object: DeepPartial<MsgAddDenom>): MsgAddDenom;
};
export declare const MsgAddDenomResponse: {
    encode(_: MsgAddDenomResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddDenomResponse;
    fromJSON(_: any): MsgAddDenomResponse;
    toJSON(_: MsgAddDenomResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddDenomResponse>): MsgAddDenomResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    UpdateAdmin(request: MsgUpdateAdmin): Promise<MsgUpdateAdminResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    AddDenom(request: MsgAddDenom): Promise<MsgAddDenomResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    UpdateAdmin(request: MsgUpdateAdmin): Promise<MsgUpdateAdminResponse>;
    AddDenom(request: MsgAddDenom): Promise<MsgAddDenomResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
