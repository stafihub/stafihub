import { Reader, Writer } from 'protobufjs/minimal';
import { ExchangeRate, EraExchangeRate, PoolDetail, UnbondFee, LeastBond, BondPipeline, BondSnapshot, PoolUnbond, AccountUnbond, BondRecord } from '../ledger/ledger';
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
export interface QueryPoolsByDenomRequest {
    denom: string;
}
export interface QueryPoolsByDenomResponse {
    addrs: string[];
}
export interface QueryBondedPoolsByDenomRequest {
    denom: string;
}
export interface QueryBondedPoolsByDenomResponse {
    addrs: string[];
}
export interface QueryGetPoolDetailRequest {
    denom: string;
    pool: string;
}
export interface QueryGetPoolDetailResponse {
    detail: PoolDetail | undefined;
}
export interface QueryGetChainEraRequest {
    denom: string;
}
export interface QueryGetChainEraResponse {
    era: number;
}
export interface QueryGetCurrentEraSnapshotRequest {
    denom: string;
}
export interface QueryGetCurrentEraSnapshotResponse {
    shotIds: Uint8Array[];
}
export interface QueryGetReceiverRequest {
}
export interface QueryGetReceiverResponse {
    receiver: string;
}
export interface QueryGetCommissionRequest {
}
export interface QueryGetCommissionResponse {
    commission: string;
}
export interface QueryGetChainBondingDurationRequest {
    denom: string;
}
export interface QueryGetChainBondingDurationResponse {
    era: number;
}
export interface QueryGetUnbondFeeRequest {
}
export interface QueryGetUnbondFeeResponse {
    fee: UnbondFee | undefined;
}
export interface QueryGetUnbondCommissionRequest {
}
export interface QueryGetUnbondCommissionResponse {
    commission: string;
}
export interface QueryGetLeastBondRequest {
    denom: string;
}
export interface QueryGetLeastBondResponse {
    leastBond: LeastBond | undefined;
}
export interface QueryGetEraUnbondLimitRequest {
    denom: string;
}
export interface QueryGetEraUnbondLimitResponse {
    limit: number;
}
export interface QueryGetBondPipeLineRequest {
    denom: string;
    pool: string;
}
export interface QueryGetBondPipeLineResponse {
    pipeline: BondPipeline | undefined;
}
export interface QueryGetEraSnapshotRequest {
    denom: string;
    era: number;
}
export interface QueryGetEraSnapshotResponse {
    shotIds: Uint8Array[];
}
export interface QueryGetSnapshotRequest {
    shotId: Uint8Array;
}
export interface QueryGetSnapshotResponse {
    shot: BondSnapshot | undefined;
}
export interface QueryGetTotalExpectedActiveRequest {
    denom: string;
    era: number;
}
export interface QueryGetTotalExpectedActiveResponse {
    active: string;
}
export interface QueryGetPoolUnbondRequest {
    denom: string;
    pool: string;
    era: number;
}
export interface QueryGetPoolUnbondResponse {
    unbond: PoolUnbond | undefined;
}
export interface QueryGetAccountUnbondRequest {
    denom: string;
    unbonder: string;
}
export interface QueryGetAccountUnbondResponse {
    unbond: AccountUnbond | undefined;
}
export interface QueryGetBondRecordRequest {
    denom: string;
    blockhash: string;
    txhash: string;
}
export interface QueryGetBondRecordResponse {
    bondRecord: BondRecord | undefined;
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
export declare const QueryPoolsByDenomRequest: {
    encode(message: QueryPoolsByDenomRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryPoolsByDenomRequest;
    fromJSON(object: any): QueryPoolsByDenomRequest;
    toJSON(message: QueryPoolsByDenomRequest): unknown;
    fromPartial(object: DeepPartial<QueryPoolsByDenomRequest>): QueryPoolsByDenomRequest;
};
export declare const QueryPoolsByDenomResponse: {
    encode(message: QueryPoolsByDenomResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryPoolsByDenomResponse;
    fromJSON(object: any): QueryPoolsByDenomResponse;
    toJSON(message: QueryPoolsByDenomResponse): unknown;
    fromPartial(object: DeepPartial<QueryPoolsByDenomResponse>): QueryPoolsByDenomResponse;
};
export declare const QueryBondedPoolsByDenomRequest: {
    encode(message: QueryBondedPoolsByDenomRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryBondedPoolsByDenomRequest;
    fromJSON(object: any): QueryBondedPoolsByDenomRequest;
    toJSON(message: QueryBondedPoolsByDenomRequest): unknown;
    fromPartial(object: DeepPartial<QueryBondedPoolsByDenomRequest>): QueryBondedPoolsByDenomRequest;
};
export declare const QueryBondedPoolsByDenomResponse: {
    encode(message: QueryBondedPoolsByDenomResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryBondedPoolsByDenomResponse;
    fromJSON(object: any): QueryBondedPoolsByDenomResponse;
    toJSON(message: QueryBondedPoolsByDenomResponse): unknown;
    fromPartial(object: DeepPartial<QueryBondedPoolsByDenomResponse>): QueryBondedPoolsByDenomResponse;
};
export declare const QueryGetPoolDetailRequest: {
    encode(message: QueryGetPoolDetailRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPoolDetailRequest;
    fromJSON(object: any): QueryGetPoolDetailRequest;
    toJSON(message: QueryGetPoolDetailRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPoolDetailRequest>): QueryGetPoolDetailRequest;
};
export declare const QueryGetPoolDetailResponse: {
    encode(message: QueryGetPoolDetailResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPoolDetailResponse;
    fromJSON(object: any): QueryGetPoolDetailResponse;
    toJSON(message: QueryGetPoolDetailResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPoolDetailResponse>): QueryGetPoolDetailResponse;
};
export declare const QueryGetChainEraRequest: {
    encode(message: QueryGetChainEraRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainEraRequest;
    fromJSON(object: any): QueryGetChainEraRequest;
    toJSON(message: QueryGetChainEraRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetChainEraRequest>): QueryGetChainEraRequest;
};
export declare const QueryGetChainEraResponse: {
    encode(message: QueryGetChainEraResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainEraResponse;
    fromJSON(object: any): QueryGetChainEraResponse;
    toJSON(message: QueryGetChainEraResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetChainEraResponse>): QueryGetChainEraResponse;
};
export declare const QueryGetCurrentEraSnapshotRequest: {
    encode(message: QueryGetCurrentEraSnapshotRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCurrentEraSnapshotRequest;
    fromJSON(object: any): QueryGetCurrentEraSnapshotRequest;
    toJSON(message: QueryGetCurrentEraSnapshotRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetCurrentEraSnapshotRequest>): QueryGetCurrentEraSnapshotRequest;
};
export declare const QueryGetCurrentEraSnapshotResponse: {
    encode(message: QueryGetCurrentEraSnapshotResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCurrentEraSnapshotResponse;
    fromJSON(object: any): QueryGetCurrentEraSnapshotResponse;
    toJSON(message: QueryGetCurrentEraSnapshotResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetCurrentEraSnapshotResponse>): QueryGetCurrentEraSnapshotResponse;
};
export declare const QueryGetReceiverRequest: {
    encode(_: QueryGetReceiverRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetReceiverRequest;
    fromJSON(_: any): QueryGetReceiverRequest;
    toJSON(_: QueryGetReceiverRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetReceiverRequest>): QueryGetReceiverRequest;
};
export declare const QueryGetReceiverResponse: {
    encode(message: QueryGetReceiverResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetReceiverResponse;
    fromJSON(object: any): QueryGetReceiverResponse;
    toJSON(message: QueryGetReceiverResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetReceiverResponse>): QueryGetReceiverResponse;
};
export declare const QueryGetCommissionRequest: {
    encode(_: QueryGetCommissionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCommissionRequest;
    fromJSON(_: any): QueryGetCommissionRequest;
    toJSON(_: QueryGetCommissionRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetCommissionRequest>): QueryGetCommissionRequest;
};
export declare const QueryGetCommissionResponse: {
    encode(message: QueryGetCommissionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCommissionResponse;
    fromJSON(object: any): QueryGetCommissionResponse;
    toJSON(message: QueryGetCommissionResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetCommissionResponse>): QueryGetCommissionResponse;
};
export declare const QueryGetChainBondingDurationRequest: {
    encode(message: QueryGetChainBondingDurationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainBondingDurationRequest;
    fromJSON(object: any): QueryGetChainBondingDurationRequest;
    toJSON(message: QueryGetChainBondingDurationRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetChainBondingDurationRequest>): QueryGetChainBondingDurationRequest;
};
export declare const QueryGetChainBondingDurationResponse: {
    encode(message: QueryGetChainBondingDurationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainBondingDurationResponse;
    fromJSON(object: any): QueryGetChainBondingDurationResponse;
    toJSON(message: QueryGetChainBondingDurationResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetChainBondingDurationResponse>): QueryGetChainBondingDurationResponse;
};
export declare const QueryGetUnbondFeeRequest: {
    encode(_: QueryGetUnbondFeeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondFeeRequest;
    fromJSON(_: any): QueryGetUnbondFeeRequest;
    toJSON(_: QueryGetUnbondFeeRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetUnbondFeeRequest>): QueryGetUnbondFeeRequest;
};
export declare const QueryGetUnbondFeeResponse: {
    encode(message: QueryGetUnbondFeeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondFeeResponse;
    fromJSON(object: any): QueryGetUnbondFeeResponse;
    toJSON(message: QueryGetUnbondFeeResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetUnbondFeeResponse>): QueryGetUnbondFeeResponse;
};
export declare const QueryGetUnbondCommissionRequest: {
    encode(_: QueryGetUnbondCommissionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondCommissionRequest;
    fromJSON(_: any): QueryGetUnbondCommissionRequest;
    toJSON(_: QueryGetUnbondCommissionRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetUnbondCommissionRequest>): QueryGetUnbondCommissionRequest;
};
export declare const QueryGetUnbondCommissionResponse: {
    encode(message: QueryGetUnbondCommissionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUnbondCommissionResponse;
    fromJSON(object: any): QueryGetUnbondCommissionResponse;
    toJSON(message: QueryGetUnbondCommissionResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetUnbondCommissionResponse>): QueryGetUnbondCommissionResponse;
};
export declare const QueryGetLeastBondRequest: {
    encode(message: QueryGetLeastBondRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLeastBondRequest;
    fromJSON(object: any): QueryGetLeastBondRequest;
    toJSON(message: QueryGetLeastBondRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetLeastBondRequest>): QueryGetLeastBondRequest;
};
export declare const QueryGetLeastBondResponse: {
    encode(message: QueryGetLeastBondResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLeastBondResponse;
    fromJSON(object: any): QueryGetLeastBondResponse;
    toJSON(message: QueryGetLeastBondResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetLeastBondResponse>): QueryGetLeastBondResponse;
};
export declare const QueryGetEraUnbondLimitRequest: {
    encode(message: QueryGetEraUnbondLimitRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetEraUnbondLimitRequest;
    fromJSON(object: any): QueryGetEraUnbondLimitRequest;
    toJSON(message: QueryGetEraUnbondLimitRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetEraUnbondLimitRequest>): QueryGetEraUnbondLimitRequest;
};
export declare const QueryGetEraUnbondLimitResponse: {
    encode(message: QueryGetEraUnbondLimitResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetEraUnbondLimitResponse;
    fromJSON(object: any): QueryGetEraUnbondLimitResponse;
    toJSON(message: QueryGetEraUnbondLimitResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetEraUnbondLimitResponse>): QueryGetEraUnbondLimitResponse;
};
export declare const QueryGetBondPipeLineRequest: {
    encode(message: QueryGetBondPipeLineRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBondPipeLineRequest;
    fromJSON(object: any): QueryGetBondPipeLineRequest;
    toJSON(message: QueryGetBondPipeLineRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetBondPipeLineRequest>): QueryGetBondPipeLineRequest;
};
export declare const QueryGetBondPipeLineResponse: {
    encode(message: QueryGetBondPipeLineResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBondPipeLineResponse;
    fromJSON(object: any): QueryGetBondPipeLineResponse;
    toJSON(message: QueryGetBondPipeLineResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetBondPipeLineResponse>): QueryGetBondPipeLineResponse;
};
export declare const QueryGetEraSnapshotRequest: {
    encode(message: QueryGetEraSnapshotRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetEraSnapshotRequest;
    fromJSON(object: any): QueryGetEraSnapshotRequest;
    toJSON(message: QueryGetEraSnapshotRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetEraSnapshotRequest>): QueryGetEraSnapshotRequest;
};
export declare const QueryGetEraSnapshotResponse: {
    encode(message: QueryGetEraSnapshotResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetEraSnapshotResponse;
    fromJSON(object: any): QueryGetEraSnapshotResponse;
    toJSON(message: QueryGetEraSnapshotResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetEraSnapshotResponse>): QueryGetEraSnapshotResponse;
};
export declare const QueryGetSnapshotRequest: {
    encode(message: QueryGetSnapshotRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSnapshotRequest;
    fromJSON(object: any): QueryGetSnapshotRequest;
    toJSON(message: QueryGetSnapshotRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSnapshotRequest>): QueryGetSnapshotRequest;
};
export declare const QueryGetSnapshotResponse: {
    encode(message: QueryGetSnapshotResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSnapshotResponse;
    fromJSON(object: any): QueryGetSnapshotResponse;
    toJSON(message: QueryGetSnapshotResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSnapshotResponse>): QueryGetSnapshotResponse;
};
export declare const QueryGetTotalExpectedActiveRequest: {
    encode(message: QueryGetTotalExpectedActiveRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTotalExpectedActiveRequest;
    fromJSON(object: any): QueryGetTotalExpectedActiveRequest;
    toJSON(message: QueryGetTotalExpectedActiveRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTotalExpectedActiveRequest>): QueryGetTotalExpectedActiveRequest;
};
export declare const QueryGetTotalExpectedActiveResponse: {
    encode(message: QueryGetTotalExpectedActiveResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTotalExpectedActiveResponse;
    fromJSON(object: any): QueryGetTotalExpectedActiveResponse;
    toJSON(message: QueryGetTotalExpectedActiveResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTotalExpectedActiveResponse>): QueryGetTotalExpectedActiveResponse;
};
export declare const QueryGetPoolUnbondRequest: {
    encode(message: QueryGetPoolUnbondRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPoolUnbondRequest;
    fromJSON(object: any): QueryGetPoolUnbondRequest;
    toJSON(message: QueryGetPoolUnbondRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPoolUnbondRequest>): QueryGetPoolUnbondRequest;
};
export declare const QueryGetPoolUnbondResponse: {
    encode(message: QueryGetPoolUnbondResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPoolUnbondResponse;
    fromJSON(object: any): QueryGetPoolUnbondResponse;
    toJSON(message: QueryGetPoolUnbondResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPoolUnbondResponse>): QueryGetPoolUnbondResponse;
};
export declare const QueryGetAccountUnbondRequest: {
    encode(message: QueryGetAccountUnbondRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountUnbondRequest;
    fromJSON(object: any): QueryGetAccountUnbondRequest;
    toJSON(message: QueryGetAccountUnbondRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAccountUnbondRequest>): QueryGetAccountUnbondRequest;
};
export declare const QueryGetAccountUnbondResponse: {
    encode(message: QueryGetAccountUnbondResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountUnbondResponse;
    fromJSON(object: any): QueryGetAccountUnbondResponse;
    toJSON(message: QueryGetAccountUnbondResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAccountUnbondResponse>): QueryGetAccountUnbondResponse;
};
export declare const QueryGetBondRecordRequest: {
    encode(message: QueryGetBondRecordRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBondRecordRequest;
    fromJSON(object: any): QueryGetBondRecordRequest;
    toJSON(message: QueryGetBondRecordRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetBondRecordRequest>): QueryGetBondRecordRequest;
};
export declare const QueryGetBondRecordResponse: {
    encode(message: QueryGetBondRecordResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetBondRecordResponse;
    fromJSON(object: any): QueryGetBondRecordResponse;
    toJSON(message: QueryGetBondRecordResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetBondRecordResponse>): QueryGetBondRecordResponse;
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
    /** Queries a list of poolsByDenom items. */
    PoolsByDenom(request: QueryPoolsByDenomRequest): Promise<QueryPoolsByDenomResponse>;
    /** Queries a list of bondedPoolsByDenom items. */
    BondedPoolsByDenom(request: QueryBondedPoolsByDenomRequest): Promise<QueryBondedPoolsByDenomResponse>;
    /** Queries a list of getPoolDetail items. */
    GetPoolDetail(request: QueryGetPoolDetailRequest): Promise<QueryGetPoolDetailResponse>;
    /** Queries a list of getChainEra items. */
    GetChainEra(request: QueryGetChainEraRequest): Promise<QueryGetChainEraResponse>;
    /** Queries a list of getCurrentEraSnapshot items. */
    GetCurrentEraSnapshot(request: QueryGetCurrentEraSnapshotRequest): Promise<QueryGetCurrentEraSnapshotResponse>;
    /** Queries a list of getReceiver items. */
    GetReceiver(request: QueryGetReceiverRequest): Promise<QueryGetReceiverResponse>;
    /** Queries a list of getCommission items. */
    GetCommission(request: QueryGetCommissionRequest): Promise<QueryGetCommissionResponse>;
    /** Queries a list of getChainBondingDuration items. */
    GetChainBondingDuration(request: QueryGetChainBondingDurationRequest): Promise<QueryGetChainBondingDurationResponse>;
    /** Queries a list of getUnbondFee items. */
    GetUnbondFee(request: QueryGetUnbondFeeRequest): Promise<QueryGetUnbondFeeResponse>;
    /** Queries a list of getUnbondCommission items. */
    GetUnbondCommission(request: QueryGetUnbondCommissionRequest): Promise<QueryGetUnbondCommissionResponse>;
    /** Queries a list of getLeastBond items. */
    GetLeastBond(request: QueryGetLeastBondRequest): Promise<QueryGetLeastBondResponse>;
    /** Queries a list of getEraUnbondLimit items. */
    GetEraUnbondLimit(request: QueryGetEraUnbondLimitRequest): Promise<QueryGetEraUnbondLimitResponse>;
    /** Queries a list of getBondPipeLine items. */
    GetBondPipeLine(request: QueryGetBondPipeLineRequest): Promise<QueryGetBondPipeLineResponse>;
    /** Queries a list of getEraSnapshot items. */
    GetEraSnapshot(request: QueryGetEraSnapshotRequest): Promise<QueryGetEraSnapshotResponse>;
    /** Queries a list of getSnapshot items. */
    GetSnapshot(request: QueryGetSnapshotRequest): Promise<QueryGetSnapshotResponse>;
    /** Queries a list of getTotalExpectedActive items. */
    GetTotalExpectedActive(request: QueryGetTotalExpectedActiveRequest): Promise<QueryGetTotalExpectedActiveResponse>;
    /** Queries a list of getPoolUnbond items. */
    GetPoolUnbond(request: QueryGetPoolUnbondRequest): Promise<QueryGetPoolUnbondResponse>;
    /** Queries a list of getAccountUnbond items. */
    GetAccountUnbond(request: QueryGetAccountUnbondRequest): Promise<QueryGetAccountUnbondResponse>;
    /** Queries a list of getBondRecord items. */
    GetBondRecord(request: QueryGetBondRecordRequest): Promise<QueryGetBondRecordResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    GetExchangeRate(request: QueryGetExchangeRateRequest): Promise<QueryGetExchangeRateResponse>;
    ExchangeRateAll(request: QueryExchangeRateAllRequest): Promise<QueryExchangeRateAllResponse>;
    GetEraExchangeRate(request: QueryGetEraExchangeRateRequest): Promise<QueryGetEraExchangeRateResponse>;
    EraExchangeRatesByDenom(request: QueryEraExchangeRatesByDenomRequest): Promise<QueryEraExchangeRatesByDenomResponse>;
    PoolsByDenom(request: QueryPoolsByDenomRequest): Promise<QueryPoolsByDenomResponse>;
    BondedPoolsByDenom(request: QueryBondedPoolsByDenomRequest): Promise<QueryBondedPoolsByDenomResponse>;
    GetPoolDetail(request: QueryGetPoolDetailRequest): Promise<QueryGetPoolDetailResponse>;
    GetChainEra(request: QueryGetChainEraRequest): Promise<QueryGetChainEraResponse>;
    GetCurrentEraSnapshot(request: QueryGetCurrentEraSnapshotRequest): Promise<QueryGetCurrentEraSnapshotResponse>;
    GetReceiver(request: QueryGetReceiverRequest): Promise<QueryGetReceiverResponse>;
    GetCommission(request: QueryGetCommissionRequest): Promise<QueryGetCommissionResponse>;
    GetChainBondingDuration(request: QueryGetChainBondingDurationRequest): Promise<QueryGetChainBondingDurationResponse>;
    GetUnbondFee(request: QueryGetUnbondFeeRequest): Promise<QueryGetUnbondFeeResponse>;
    GetUnbondCommission(request: QueryGetUnbondCommissionRequest): Promise<QueryGetUnbondCommissionResponse>;
    GetLeastBond(request: QueryGetLeastBondRequest): Promise<QueryGetLeastBondResponse>;
    GetEraUnbondLimit(request: QueryGetEraUnbondLimitRequest): Promise<QueryGetEraUnbondLimitResponse>;
    GetBondPipeLine(request: QueryGetBondPipeLineRequest): Promise<QueryGetBondPipeLineResponse>;
    GetEraSnapshot(request: QueryGetEraSnapshotRequest): Promise<QueryGetEraSnapshotResponse>;
    GetSnapshot(request: QueryGetSnapshotRequest): Promise<QueryGetSnapshotResponse>;
    GetTotalExpectedActive(request: QueryGetTotalExpectedActiveRequest): Promise<QueryGetTotalExpectedActiveResponse>;
    GetPoolUnbond(request: QueryGetPoolUnbondRequest): Promise<QueryGetPoolUnbondResponse>;
    GetAccountUnbond(request: QueryGetAccountUnbondRequest): Promise<QueryGetAccountUnbondResponse>;
    GetBondRecord(request: QueryGetBondRecordRequest): Promise<QueryGetBondRecordResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
