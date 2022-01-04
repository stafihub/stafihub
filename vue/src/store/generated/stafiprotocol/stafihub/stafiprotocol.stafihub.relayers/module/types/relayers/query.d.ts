import { Reader, Writer } from 'protobufjs/minimal';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { Relayer, Threshold } from '../relayers/relayer';
export declare const protobufPackage = "stafiprotocol.stafihub.relayers";
export interface QueryAllRelayerRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllRelayerResponse {
    relayers: Relayer[];
    pagination: PageResponse | undefined;
}
export interface QueryRelayersByDenomRequest {
    denom: string;
    pagination: PageRequest | undefined;
}
export interface QueryRelayersByDenomResponse {
    relayers: Relayer[];
    pagination: PageResponse | undefined;
}
export interface QueryGetThresholdRequest {
    denom: string;
}
export interface QueryGetThresholdResponse {
    threshold: Threshold | undefined;
}
export interface QueryAllThresholdRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllThresholdResponse {
    threshold: Threshold[];
    pagination: PageResponse | undefined;
}
export declare const QueryAllRelayerRequest: {
    encode(message: QueryAllRelayerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRelayerRequest;
    fromJSON(object: any): QueryAllRelayerRequest;
    toJSON(message: QueryAllRelayerRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRelayerRequest>): QueryAllRelayerRequest;
};
export declare const QueryAllRelayerResponse: {
    encode(message: QueryAllRelayerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRelayerResponse;
    fromJSON(object: any): QueryAllRelayerResponse;
    toJSON(message: QueryAllRelayerResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRelayerResponse>): QueryAllRelayerResponse;
};
export declare const QueryRelayersByDenomRequest: {
    encode(message: QueryRelayersByDenomRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryRelayersByDenomRequest;
    fromJSON(object: any): QueryRelayersByDenomRequest;
    toJSON(message: QueryRelayersByDenomRequest): unknown;
    fromPartial(object: DeepPartial<QueryRelayersByDenomRequest>): QueryRelayersByDenomRequest;
};
export declare const QueryRelayersByDenomResponse: {
    encode(message: QueryRelayersByDenomResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryRelayersByDenomResponse;
    fromJSON(object: any): QueryRelayersByDenomResponse;
    toJSON(message: QueryRelayersByDenomResponse): unknown;
    fromPartial(object: DeepPartial<QueryRelayersByDenomResponse>): QueryRelayersByDenomResponse;
};
export declare const QueryGetThresholdRequest: {
    encode(message: QueryGetThresholdRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetThresholdRequest;
    fromJSON(object: any): QueryGetThresholdRequest;
    toJSON(message: QueryGetThresholdRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetThresholdRequest>): QueryGetThresholdRequest;
};
export declare const QueryGetThresholdResponse: {
    encode(message: QueryGetThresholdResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetThresholdResponse;
    fromJSON(object: any): QueryGetThresholdResponse;
    toJSON(message: QueryGetThresholdResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetThresholdResponse>): QueryGetThresholdResponse;
};
export declare const QueryAllThresholdRequest: {
    encode(message: QueryAllThresholdRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllThresholdRequest;
    fromJSON(object: any): QueryAllThresholdRequest;
    toJSON(message: QueryAllThresholdRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllThresholdRequest>): QueryAllThresholdRequest;
};
export declare const QueryAllThresholdResponse: {
    encode(message: QueryAllThresholdResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllThresholdResponse;
    fromJSON(object: any): QueryAllThresholdResponse;
    toJSON(message: QueryAllThresholdResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllThresholdResponse>): QueryAllThresholdResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of relayer items. */
    RelayerAll(request: QueryAllRelayerRequest): Promise<QueryAllRelayerResponse>;
    /** Queries a list of relayersByDenom items. */
    RelayersByDenom(request: QueryRelayersByDenomRequest): Promise<QueryRelayersByDenomResponse>;
    /** Queries a threshold by denom. */
    Threshold(request: QueryGetThresholdRequest): Promise<QueryGetThresholdResponse>;
    /** Queries a list of threshold items. */
    ThresholdAll(request: QueryAllThresholdRequest): Promise<QueryAllThresholdResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    RelayerAll(request: QueryAllRelayerRequest): Promise<QueryAllRelayerResponse>;
    RelayersByDenom(request: QueryRelayersByDenomRequest): Promise<QueryRelayersByDenomResponse>;
    Threshold(request: QueryGetThresholdRequest): Promise<QueryGetThresholdResponse>;
    ThresholdAll(request: QueryAllThresholdRequest): Promise<QueryAllThresholdResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
