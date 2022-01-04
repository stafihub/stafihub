import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.relayers";
export interface Relayer {
    denom: string;
    address: string;
}
export interface Threshold {
    denom: string;
    value: number;
}
export interface LastVoter {
    denom: string;
    voter: string;
}
export declare const Relayer: {
    encode(message: Relayer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Relayer;
    fromJSON(object: any): Relayer;
    toJSON(message: Relayer): unknown;
    fromPartial(object: DeepPartial<Relayer>): Relayer;
};
export declare const Threshold: {
    encode(message: Threshold, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Threshold;
    fromJSON(object: any): Threshold;
    toJSON(message: Threshold): unknown;
    fromPartial(object: DeepPartial<Threshold>): Threshold;
};
export declare const LastVoter: {
    encode(message: LastVoter, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): LastVoter;
    fromJSON(object: any): LastVoter;
    toJSON(message: LastVoter): unknown;
    fromPartial(object: DeepPartial<LastVoter>): LastVoter;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
