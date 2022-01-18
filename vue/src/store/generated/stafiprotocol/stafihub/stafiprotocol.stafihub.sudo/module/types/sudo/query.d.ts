import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.sudo";
export interface QueryAdminRequest {
}
export interface QueryAdminResponse {
    address: string;
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
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of admin items. */
    Admin(request: QueryAdminRequest): Promise<QueryAdminResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Admin(request: QueryAdminRequest): Promise<QueryAdminResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
