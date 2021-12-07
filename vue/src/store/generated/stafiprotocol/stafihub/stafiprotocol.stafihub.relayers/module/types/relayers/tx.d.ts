import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.relayers";
export interface MsgCreateRelayer {
    creator: string;
    denom: string;
    address: string;
}
export interface MsgCreateRelayerResponse {
}
export interface MsgDeleteRelayer {
    creator: string;
    denom: string;
    address: string;
}
export interface MsgDeleteRelayerResponse {
}
export interface MsgUpdateThreshold {
    creator: string;
    denom: string;
    value: number;
}
export interface MsgUpdateThresholdResponse {
}
export declare const MsgCreateRelayer: {
    encode(message: MsgCreateRelayer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateRelayer;
    fromJSON(object: any): MsgCreateRelayer;
    toJSON(message: MsgCreateRelayer): unknown;
    fromPartial(object: DeepPartial<MsgCreateRelayer>): MsgCreateRelayer;
};
export declare const MsgCreateRelayerResponse: {
    encode(_: MsgCreateRelayerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateRelayerResponse;
    fromJSON(_: any): MsgCreateRelayerResponse;
    toJSON(_: MsgCreateRelayerResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateRelayerResponse>): MsgCreateRelayerResponse;
};
export declare const MsgDeleteRelayer: {
    encode(message: MsgDeleteRelayer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteRelayer;
    fromJSON(object: any): MsgDeleteRelayer;
    toJSON(message: MsgDeleteRelayer): unknown;
    fromPartial(object: DeepPartial<MsgDeleteRelayer>): MsgDeleteRelayer;
};
export declare const MsgDeleteRelayerResponse: {
    encode(_: MsgDeleteRelayerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteRelayerResponse;
    fromJSON(_: any): MsgDeleteRelayerResponse;
    toJSON(_: MsgDeleteRelayerResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteRelayerResponse>): MsgDeleteRelayerResponse;
};
export declare const MsgUpdateThreshold: {
    encode(message: MsgUpdateThreshold, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateThreshold;
    fromJSON(object: any): MsgUpdateThreshold;
    toJSON(message: MsgUpdateThreshold): unknown;
    fromPartial(object: DeepPartial<MsgUpdateThreshold>): MsgUpdateThreshold;
};
export declare const MsgUpdateThresholdResponse: {
    encode(_: MsgUpdateThresholdResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateThresholdResponse;
    fromJSON(_: any): MsgUpdateThresholdResponse;
    toJSON(_: MsgUpdateThresholdResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateThresholdResponse>): MsgUpdateThresholdResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateRelayer(request: MsgCreateRelayer): Promise<MsgCreateRelayerResponse>;
    DeleteRelayer(request: MsgDeleteRelayer): Promise<MsgDeleteRelayerResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateThreshold(request: MsgUpdateThreshold): Promise<MsgUpdateThresholdResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateRelayer(request: MsgCreateRelayer): Promise<MsgCreateRelayerResponse>;
    DeleteRelayer(request: MsgDeleteRelayer): Promise<MsgDeleteRelayerResponse>;
    UpdateThreshold(request: MsgUpdateThreshold): Promise<MsgUpdateThresholdResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
