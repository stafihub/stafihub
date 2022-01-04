import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.ledger";
export declare enum PoolBondState {
    ERA_UPDATED = 0,
    BOND_REPORTED = 1,
    ACTIVE_REPORTED = 2,
    WITHDRAW_SKIPPED = 3,
    WITHDRAW_REPORTED = 4,
    TRANSFER_REPORTED = 5,
    UNRECOGNIZED = -1
}
export declare function poolBondStateFromJSON(object: any): PoolBondState;
export declare function poolBondStateToJSON(object: PoolBondState): string;
export declare enum BondAction {
    BOND_ONLY = 0,
    UNBOND_ONLY = 1,
    BOTH_BOND_UNBOND = 2,
    EITHER_BOND_UNBOND = 3,
    INTER_DEDUCT = 4,
    UNRECOGNIZED = -1
}
export declare function bondActionFromJSON(object: any): BondAction;
export declare function bondActionToJSON(object: BondAction): string;
export interface ChainEra {
    denom: string;
    era: number;
}
export interface ChainBondingDuration {
    denom: string;
    era: number;
}
export interface Pool {
    denom: string;
    addrs: {
        [key: string]: boolean;
    };
}
export interface Pool_AddrsEntry {
    key: string;
    value: boolean;
}
export interface TotalExpectedActive {
    denom: string;
    era: string;
    amount: string;
}
export interface BondPipeline {
    denom: string;
    pool: string;
    chunk: LinkChunk | undefined;
}
export interface EraSnapShot {
    denom: string;
    shotIds: Uint8Array[];
}
export interface PoolUnbond {
    denom: string;
    pool: string;
    era: number;
    unbondings: Unbonding[];
}
export interface EraUnbondLimit {
    denom: string;
    limit: number;
}
export interface PoolDetail {
    denom: string;
    pool: string;
    subAccounts: string[];
    threshold: number;
}
export interface LeastBond {
    denom: string;
    amount: string;
}
export interface LinkChunk {
    bond: string;
    unbond: string;
    active: string;
}
export interface BondSnapshot {
    denom: string;
    pool: string;
    era: number;
    chunk: LinkChunk | undefined;
    lastVoter: string;
    bondState: PoolBondState;
}
export interface Unbonding {
    unbonder: string;
    amount: string;
    recipient: string;
}
export declare const ChainEra: {
    encode(message: ChainEra, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ChainEra;
    fromJSON(object: any): ChainEra;
    toJSON(message: ChainEra): unknown;
    fromPartial(object: DeepPartial<ChainEra>): ChainEra;
};
export declare const ChainBondingDuration: {
    encode(message: ChainBondingDuration, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ChainBondingDuration;
    fromJSON(object: any): ChainBondingDuration;
    toJSON(message: ChainBondingDuration): unknown;
    fromPartial(object: DeepPartial<ChainBondingDuration>): ChainBondingDuration;
};
export declare const Pool: {
    encode(message: Pool, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Pool;
    fromJSON(object: any): Pool;
    toJSON(message: Pool): unknown;
    fromPartial(object: DeepPartial<Pool>): Pool;
};
export declare const Pool_AddrsEntry: {
    encode(message: Pool_AddrsEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Pool_AddrsEntry;
    fromJSON(object: any): Pool_AddrsEntry;
    toJSON(message: Pool_AddrsEntry): unknown;
    fromPartial(object: DeepPartial<Pool_AddrsEntry>): Pool_AddrsEntry;
};
export declare const TotalExpectedActive: {
    encode(message: TotalExpectedActive, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): TotalExpectedActive;
    fromJSON(object: any): TotalExpectedActive;
    toJSON(message: TotalExpectedActive): unknown;
    fromPartial(object: DeepPartial<TotalExpectedActive>): TotalExpectedActive;
};
export declare const BondPipeline: {
    encode(message: BondPipeline, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BondPipeline;
    fromJSON(object: any): BondPipeline;
    toJSON(message: BondPipeline): unknown;
    fromPartial(object: DeepPartial<BondPipeline>): BondPipeline;
};
export declare const EraSnapShot: {
    encode(message: EraSnapShot, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EraSnapShot;
    fromJSON(object: any): EraSnapShot;
    toJSON(message: EraSnapShot): unknown;
    fromPartial(object: DeepPartial<EraSnapShot>): EraSnapShot;
};
export declare const PoolUnbond: {
    encode(message: PoolUnbond, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PoolUnbond;
    fromJSON(object: any): PoolUnbond;
    toJSON(message: PoolUnbond): unknown;
    fromPartial(object: DeepPartial<PoolUnbond>): PoolUnbond;
};
export declare const EraUnbondLimit: {
    encode(message: EraUnbondLimit, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EraUnbondLimit;
    fromJSON(object: any): EraUnbondLimit;
    toJSON(message: EraUnbondLimit): unknown;
    fromPartial(object: DeepPartial<EraUnbondLimit>): EraUnbondLimit;
};
export declare const PoolDetail: {
    encode(message: PoolDetail, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PoolDetail;
    fromJSON(object: any): PoolDetail;
    toJSON(message: PoolDetail): unknown;
    fromPartial(object: DeepPartial<PoolDetail>): PoolDetail;
};
export declare const LeastBond: {
    encode(message: LeastBond, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): LeastBond;
    fromJSON(object: any): LeastBond;
    toJSON(message: LeastBond): unknown;
    fromPartial(object: DeepPartial<LeastBond>): LeastBond;
};
export declare const LinkChunk: {
    encode(message: LinkChunk, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): LinkChunk;
    fromJSON(object: any): LinkChunk;
    toJSON(message: LinkChunk): unknown;
    fromPartial(object: DeepPartial<LinkChunk>): LinkChunk;
};
export declare const BondSnapshot: {
    encode(message: BondSnapshot, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BondSnapshot;
    fromJSON(object: any): BondSnapshot;
    toJSON(message: BondSnapshot): unknown;
    fromPartial(object: DeepPartial<BondSnapshot>): BondSnapshot;
};
export declare const Unbonding: {
    encode(message: Unbonding, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Unbonding;
    fromJSON(object: any): Unbonding;
    toJSON(message: Unbonding): unknown;
    fromPartial(object: DeepPartial<Unbonding>): Unbonding;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
