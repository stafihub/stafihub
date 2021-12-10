import { Writer, Reader } from 'protobufjs/minimal';
import { Relayer, Threshold } from '../relayers/relayer';
export declare const protobufPackage = "stafiprotocol.stafihub.relayers";
/** GenesisState defines the relayers module's genesis state. */
export interface GenesisState {
    relayers: Relayer[];
    thresholds: Threshold[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    proposalLife: number;
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
