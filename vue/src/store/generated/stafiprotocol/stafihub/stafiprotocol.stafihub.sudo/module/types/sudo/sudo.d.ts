import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.sudo";
export interface Symbol {
    denoms: {
        [key: string]: boolean;
    };
}
export interface Symbol_DenomsEntry {
    key: string;
    value: boolean;
}
export declare const Symbol: {
    encode(message: Symbol, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Symbol;
    fromJSON(object: any): Symbol;
    toJSON(message: Symbol): unknown;
    fromPartial(object: DeepPartial<Symbol>): Symbol;
};
export declare const Symbol_DenomsEntry: {
    encode(message: Symbol_DenomsEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Symbol_DenomsEntry;
    fromJSON(object: any): Symbol_DenomsEntry;
    toJSON(message: Symbol_DenomsEntry): unknown;
    fromPartial(object: DeepPartial<Symbol_DenomsEntry>): Symbol_DenomsEntry;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
