import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.relayers";
export interface MsgCreateRelayer {
    denom: string;
    address: string;
}
export interface MsgCreateRelayerResponse {
}
export interface MsgDeleteRelayer {
    denom: string;
    address: string;
}
export interface MsgDeleteRelayerResponse {
}
export interface MsgSetThreshold {
    denom: string;
    value: string;
}
export interface MsgSetThresholdResponse {
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
export declare const MsgSetThreshold: {
    encode(message: MsgSetThreshold, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetThreshold;
    fromJSON(object: any): MsgSetThreshold;
    toJSON(message: MsgSetThreshold): unknown;
    fromPartial(object: DeepPartial<MsgSetThreshold>): MsgSetThreshold;
};
export declare const MsgSetThresholdResponse: {
    encode(_: MsgSetThresholdResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetThresholdResponse;
    fromJSON(_: any): MsgSetThresholdResponse;
    toJSON(_: MsgSetThresholdResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetThresholdResponse>): MsgSetThresholdResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateRelayer(request: MsgCreateRelayer): Promise<MsgCreateRelayerResponse>;
    DeleteRelayer(request: MsgDeleteRelayer): Promise<MsgDeleteRelayerResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SetThreshold(request: MsgSetThreshold): Promise<MsgSetThresholdResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateRelayer(request: MsgCreateRelayer): Promise<MsgCreateRelayerResponse>;
    DeleteRelayer(request: MsgDeleteRelayer): Promise<MsgDeleteRelayerResponse>;
    SetThreshold(request: MsgSetThreshold): Promise<MsgSetThresholdResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
