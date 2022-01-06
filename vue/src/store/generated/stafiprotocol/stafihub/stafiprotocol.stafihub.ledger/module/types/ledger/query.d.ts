import { Reader, Writer } from 'protobufjs/minimal';
import { ExchangeRate, EraExchangeRate } from '../ledger/ledger';
export declare const protobufPackage = "stafiprotocol.stafihub.ledger";
export interface QueryGetExchangeRateRequest {
    denom: string;
}
export interface QueryGetExchangeRateResponse {
    exchangeRate: ExchangeRate | undefined;
}
export interface QueryExchangeRateAllRequest {
}
export interface QueryExchangeRateAllResponse {
    exchangeRates: ExchangeRate[];
}
export interface QueryGetEraExchangeRateRequest {
    denom: string;
    era: number;
}
export interface QueryGetEraExchangeRateResponse {
    eraExchangeRate: EraExchangeRate | undefined;
}
export interface QueryEraExchangeRatesByDenomRequest {
    denom: string;
}
export interface QueryEraExchangeRatesByDenomResponse {
    eraExchangeRates: EraExchangeRate[];
}
export declare const QueryGetExchangeRateRequest: {
    encode(message: QueryGetExchangeRateRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetExchangeRateRequest;
    fromJSON(object: any): QueryGetExchangeRateRequest;
    toJSON(message: QueryGetExchangeRateRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetExchangeRateRequest>): QueryGetExchangeRateRequest;
};
export declare const QueryGetExchangeRateResponse: {
    encode(message: QueryGetExchangeRateResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetExchangeRateResponse;
    fromJSON(object: any): QueryGetExchangeRateResponse;
    toJSON(message: QueryGetExchangeRateResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetExchangeRateResponse>): QueryGetExchangeRateResponse;
};
export declare const QueryExchangeRateAllRequest: {
    encode(_: QueryExchangeRateAllRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryExchangeRateAllRequest;
    fromJSON(_: any): QueryExchangeRateAllRequest;
    toJSON(_: QueryExchangeRateAllRequest): unknown;
    fromPartial(_: DeepPartial<QueryExchangeRateAllRequest>): QueryExchangeRateAllRequest;
};
export declare const QueryExchangeRateAllResponse: {
    encode(message: QueryExchangeRateAllResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryExchangeRateAllResponse;
    fromJSON(object: any): QueryExchangeRateAllResponse;
    toJSON(message: QueryExchangeRateAllResponse): unknown;
    fromPartial(object: DeepPartial<QueryExchangeRateAllResponse>): QueryExchangeRateAllResponse;
};
export declare const QueryGetEraExchangeRateRequest: {
    encode(message: QueryGetEraExchangeRateRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetEraExchangeRateRequest;
    fromJSON(object: any): QueryGetEraExchangeRateRequest;
    toJSON(message: QueryGetEraExchangeRateRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetEraExchangeRateRequest>): QueryGetEraExchangeRateRequest;
};
export declare const QueryGetEraExchangeRateResponse: {
    encode(message: QueryGetEraExchangeRateResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetEraExchangeRateResponse;
    fromJSON(object: any): QueryGetEraExchangeRateResponse;
    toJSON(message: QueryGetEraExchangeRateResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetEraExchangeRateResponse>): QueryGetEraExchangeRateResponse;
};
export declare const QueryEraExchangeRatesByDenomRequest: {
    encode(message: QueryEraExchangeRatesByDenomRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryEraExchangeRatesByDenomRequest;
    fromJSON(object: any): QueryEraExchangeRatesByDenomRequest;
    toJSON(message: QueryEraExchangeRatesByDenomRequest): unknown;
    fromPartial(object: DeepPartial<QueryEraExchangeRatesByDenomRequest>): QueryEraExchangeRatesByDenomRequest;
};
export declare const QueryEraExchangeRatesByDenomResponse: {
    encode(message: QueryEraExchangeRatesByDenomResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryEraExchangeRatesByDenomResponse;
    fromJSON(object: any): QueryEraExchangeRatesByDenomResponse;
    toJSON(message: QueryEraExchangeRatesByDenomResponse): unknown;
    fromPartial(object: DeepPartial<QueryEraExchangeRatesByDenomResponse>): QueryEraExchangeRatesByDenomResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of getExchangeRate items. */
    GetExchangeRate(request: QueryGetExchangeRateRequest): Promise<QueryGetExchangeRateResponse>;
    /** Queries a list of exchangeRateAll items. */
    ExchangeRateAll(request: QueryExchangeRateAllRequest): Promise<QueryExchangeRateAllResponse>;
    /** Queries a list of getEraExchangeRate items. */
    GetEraExchangeRate(request: QueryGetEraExchangeRateRequest): Promise<QueryGetEraExchangeRateResponse>;
    /** Queries a list of eraExchangeRatesByDenom items. */
    EraExchangeRatesByDenom(request: QueryEraExchangeRatesByDenomRequest): Promise<QueryEraExchangeRatesByDenomResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    GetExchangeRate(request: QueryGetExchangeRateRequest): Promise<QueryGetExchangeRateResponse>;
    ExchangeRateAll(request: QueryExchangeRateAllRequest): Promise<QueryExchangeRateAllResponse>;
    GetEraExchangeRate(request: QueryGetEraExchangeRateRequest): Promise<QueryGetEraExchangeRateResponse>;
    EraExchangeRatesByDenom(request: QueryEraExchangeRatesByDenomRequest): Promise<QueryEraExchangeRatesByDenomResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
