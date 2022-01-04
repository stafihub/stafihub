import { ProposalStatus } from '../rvote/proposal';
import { Reader, Writer } from 'protobufjs/minimal';
import { Any } from '../google/protobuf/any';
export declare const protobufPackage = "stafiprotocol.stafihub.rvote";
export interface MsgSetProposalLife {
    creator: string;
    proposalLife: number;
}
export interface MsgSetProposalLifeResponse {
}
export interface MsgSubmitProposal {
    proposer: string;
    content: Any | undefined;
}
export interface MsgSubmitProposalResponse {
    propId: string;
    status: ProposalStatus;
}
export declare const MsgSetProposalLife: {
    encode(message: MsgSetProposalLife, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetProposalLife;
    fromJSON(object: any): MsgSetProposalLife;
    toJSON(message: MsgSetProposalLife): unknown;
    fromPartial(object: DeepPartial<MsgSetProposalLife>): MsgSetProposalLife;
};
export declare const MsgSetProposalLifeResponse: {
    encode(_: MsgSetProposalLifeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetProposalLifeResponse;
    fromJSON(_: any): MsgSetProposalLifeResponse;
    toJSON(_: MsgSetProposalLifeResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetProposalLifeResponse>): MsgSetProposalLifeResponse;
};
export declare const MsgSubmitProposal: {
    encode(message: MsgSubmitProposal, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSubmitProposal;
    fromJSON(object: any): MsgSubmitProposal;
    toJSON(message: MsgSubmitProposal): unknown;
    fromPartial(object: DeepPartial<MsgSubmitProposal>): MsgSubmitProposal;
};
export declare const MsgSubmitProposalResponse: {
    encode(message: MsgSubmitProposalResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSubmitProposalResponse;
    fromJSON(object: any): MsgSubmitProposalResponse;
    toJSON(message: MsgSubmitProposalResponse): unknown;
    fromPartial(object: DeepPartial<MsgSubmitProposalResponse>): MsgSubmitProposalResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    SetProposalLife(request: MsgSetProposalLife): Promise<MsgSetProposalLifeResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SubmitProposal(request: MsgSubmitProposal): Promise<MsgSubmitProposalResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    SetProposalLife(request: MsgSetProposalLife): Promise<MsgSetProposalLifeResponse>;
    SubmitProposal(request: MsgSubmitProposal): Promise<MsgSubmitProposalResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
