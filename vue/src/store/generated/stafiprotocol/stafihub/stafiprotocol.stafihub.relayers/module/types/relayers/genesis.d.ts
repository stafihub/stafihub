import { Relayer, Threshold } from '../relayers/relayer';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.relayers";
/** GenesisState defines the relayers module's genesis state. */
export interface GenesisState {
    /** admin as the only operator */
    admin: string;
    relayers: Relayer[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    thresholds: Threshold[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
