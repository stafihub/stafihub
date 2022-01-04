import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.sudo";
export interface QueryAdminRequest {
}
export interface QueryAdminResponse {
    address: string;
}
export interface QueryAllDenomsRequest {
}
export interface QueryAllDenomsResponse {
    denoms: string[];
}
export declare const QueryAdminRequest: {
    encode(_: QueryAdminRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAdminRequest;
    fromJSON(_: any): QueryAdminRequest;
    toJSON(_: QueryAdminRequest): unknown;
    fromPartial(_: DeepPartial<QueryAdminRequest>): QueryAdminRequest;
};
export declare const QueryAdminResponse: {
    encode(message: QueryAdminResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAdminResponse;
    fromJSON(object: any): QueryAdminResponse;
    toJSON(message: QueryAdminResponse): unknown;
    fromPartial(object: DeepPartial<QueryAdminResponse>): QueryAdminResponse;
};
export declare const QueryAllDenomsRequest: {
    encode(_: QueryAllDenomsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDenomsRequest;
    fromJSON(_: any): QueryAllDenomsRequest;
    toJSON(_: QueryAllDenomsRequest): unknown;
    fromPartial(_: DeepPartial<QueryAllDenomsRequest>): QueryAllDenomsRequest;
};
export declare const QueryAllDenomsResponse: {
    encode(message: QueryAllDenomsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDenomsResponse;
    fromJSON(object: any): QueryAllDenomsResponse;
    toJSON(message: QueryAllDenomsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllDenomsResponse>): QueryAllDenomsResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of admin items. */
    Admin(request: QueryAdminRequest): Promise<QueryAdminResponse>;
    /** Queries a list of allDenoms items. */
    AllDenoms(request: QueryAllDenomsRequest): Promise<QueryAllDenomsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Admin(request: QueryAdminRequest): Promise<QueryAdminResponse>;
    AllDenoms(request: QueryAllDenomsRequest): Promise<QueryAllDenomsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
