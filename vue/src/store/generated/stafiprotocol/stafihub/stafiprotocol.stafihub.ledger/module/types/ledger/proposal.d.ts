import { BondAction } from '../ledger/ledger';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.ledger";
export interface SetChainEraProposal {
    proposer: string;
    denom: string;
    era: number;
    propId: Uint8Array;
}
export interface BondReportProposal {
    proposer: string;
    denom: string;
    shotId: Uint8Array;
    action: BondAction;
    propId: Uint8Array;
}
export interface BondAndReportActiveProposal {
    proposer: string;
    denom: string;
    shotId: Uint8Array;
    action: BondAction;
    staked: string;
    unstaked: string;
    propId: Uint8Array;
}
export interface ActiveReportProposal {
    proposer: string;
    denom: string;
    shotId: Uint8Array;
    staked: string;
    unstaked: string;
    propId: Uint8Array;
}
export interface WithdrawReportProposal {
    proposer: string;
    denom: string;
    shotId: Uint8Array;
    propId: Uint8Array;
}
export interface TransferReportProposal {
    proposer: string;
    denom: string;
    shotId: Uint8Array;
    propId: Uint8Array;
}
export declare const SetChainEraProposal: {
    encode(message: SetChainEraProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SetChainEraProposal;
    fromJSON(object: any): SetChainEraProposal;
    toJSON(message: SetChainEraProposal): unknown;
    fromPartial(object: DeepPartial<SetChainEraProposal>): SetChainEraProposal;
};
export declare const BondReportProposal: {
    encode(message: BondReportProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BondReportProposal;
    fromJSON(object: any): BondReportProposal;
    toJSON(message: BondReportProposal): unknown;
    fromPartial(object: DeepPartial<BondReportProposal>): BondReportProposal;
};
export declare const BondAndReportActiveProposal: {
    encode(message: BondAndReportActiveProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BondAndReportActiveProposal;
    fromJSON(object: any): BondAndReportActiveProposal;
    toJSON(message: BondAndReportActiveProposal): unknown;
    fromPartial(object: DeepPartial<BondAndReportActiveProposal>): BondAndReportActiveProposal;
};
export declare const ActiveReportProposal: {
    encode(message: ActiveReportProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ActiveReportProposal;
    fromJSON(object: any): ActiveReportProposal;
    toJSON(message: ActiveReportProposal): unknown;
    fromPartial(object: DeepPartial<ActiveReportProposal>): ActiveReportProposal;
};
export declare const WithdrawReportProposal: {
    encode(message: WithdrawReportProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): WithdrawReportProposal;
    fromJSON(object: any): WithdrawReportProposal;
    toJSON(message: WithdrawReportProposal): unknown;
    fromPartial(object: DeepPartial<WithdrawReportProposal>): WithdrawReportProposal;
};
export declare const TransferReportProposal: {
    encode(message: TransferReportProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): TransferReportProposal;
    fromJSON(object: any): TransferReportProposal;
    toJSON(message: TransferReportProposal): unknown;
    fromPartial(object: DeepPartial<TransferReportProposal>): TransferReportProposal;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
