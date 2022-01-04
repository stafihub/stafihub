import { Writer, Reader } from 'protobufjs/minimal';
import { Any } from '../google/protobuf/any';
export declare const protobufPackage = "stafiprotocol.stafihub.rvote";
/** ProposalStatus enumerates the valid statuses of a proposal. */
export declare enum ProposalStatus {
    PROPOSAL_STATUS_INITIATED = 0,
    PROPOSAL_STATUS_APPROVED = 1,
    PROPOSAL_STATUS_REJECTED = 2,
    PROPOSAL_STATUS_EXPIRED = 3,
    UNRECOGNIZED = -1
}
export declare function proposalStatusFromJSON(object: any): ProposalStatus;
export declare function proposalStatusToJSON(object: ProposalStatus): string;
export interface Proposal {
    content: Any | undefined;
    status: ProposalStatus;
    votesFor: string[];
    votesAgainst: string[];
    startBlock: number;
    expireBlock: number;
}
export declare const Proposal: {
    encode(message: Proposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Proposal;
    fromJSON(object: any): Proposal;
    toJSON(message: Proposal): unknown;
    fromPartial(object: DeepPartial<Proposal>): Proposal;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
