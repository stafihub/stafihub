export interface LedgerAccountUnbond {
    unbonder?: string;
    denom?: string;
    chunks?: LedgerUserUnlockChunk[];
}
export interface LedgerBondPipeline {
    denom?: string;
    pool?: string;
    chunk?: LedgerLinkChunk;
}
export interface LedgerBondRecord {
    denom?: string;
    bonder?: string;
    pool?: string;
    blockhash?: string;
    txhash?: string;
    amount?: string;
}
export interface LedgerBondSnapshot {
    denom?: string;
    pool?: string;
    /** @format int64 */
    era?: number;
    chunk?: LedgerLinkChunk;
    lastVoter?: string;
    bondState?: LedgerPoolBondState;
}
export interface LedgerEraExchangeRate {
    denom?: string;
    /** @format int64 */
    era?: number;
    value?: string;
}
export interface LedgerExchangeRate {
    denom?: string;
    value?: string;
}
export interface LedgerLeastBond {
    denom?: string;
    amount?: string;
}
export interface LedgerLinkChunk {
    bond?: string;
    unbond?: string;
    active?: string;
}
export declare type LedgerMsgAddNewPoolResponse = object;
export declare type LedgerMsgClearCurrentEraSnapShotsResponse = object;
export declare type LedgerMsgLiquidityUnbondResponse = object;
export declare type LedgerMsgRemovePoolResponse = object;
export declare type LedgerMsgSetChainBondingDurationResponse = object;
export declare type LedgerMsgSetCommissionResponse = object;
export declare type LedgerMsgSetEraUnbondLimitResponse = object;
export declare type LedgerMsgSetInitBondResponse = object;
export declare type LedgerMsgSetLeastBondResponse = object;
export declare type LedgerMsgSetPoolDetailResponse = object;
export declare type LedgerMsgSetReceiverResponse = object;
export declare type LedgerMsgSetUnbondCommissionResponse = object;
export declare type LedgerMsgSetUnbondFeeResponse = object;
export declare enum LedgerPoolBondState {
    ERA_UPDATED = "ERA_UPDATED",
    BOND_REPORTED = "BOND_REPORTED",
    ACTIVE_REPORTED = "ACTIVE_REPORTED",
    WITHDRAW_SKIPPED = "WITHDRAW_SKIPPED",
    WITHDRAW_REPORTED = "WITHDRAW_REPORTED",
    TRANSFER_REPORTED = "TRANSFER_REPORTED"
}
export interface LedgerPoolDetail {
    denom?: string;
    pool?: string;
    subAccounts?: string[];
    /** @format int64 */
    threshold?: number;
}
export interface LedgerPoolUnbond {
    denom?: string;
    pool?: string;
    /** @format int64 */
    era?: number;
    unbondings?: LedgerUnbonding[];
}
export interface LedgerQueryBondedPoolsByDenomResponse {
    addrs?: string[];
}
export interface LedgerQueryEraExchangeRatesByDenomResponse {
    eraExchangeRates?: LedgerEraExchangeRate[];
}
export interface LedgerQueryExchangeRateAllResponse {
    exchangeRates?: LedgerExchangeRate[];
}
export interface LedgerQueryGetAccountUnbondResponse {
    unbond?: LedgerAccountUnbond;
}
export interface LedgerQueryGetBondPipeLineResponse {
    pipeline?: LedgerBondPipeline;
}
export interface LedgerQueryGetBondRecordResponse {
    bondRecord?: LedgerBondRecord;
}
export interface LedgerQueryGetChainBondingDurationResponse {
    /** @format int64 */
    era?: number;
}
export interface LedgerQueryGetChainEraResponse {
    /** @format int64 */
    era?: number;
}
export interface LedgerQueryGetCommissionResponse {
    commission?: string;
}
export interface LedgerQueryGetCurrentEraSnapshotResponse {
    shotIds?: string[];
}
export interface LedgerQueryGetEraExchangeRateResponse {
    eraExchangeRate?: LedgerEraExchangeRate;
}
export interface LedgerQueryGetEraSnapshotResponse {
    shotIds?: string[];
}
export interface LedgerQueryGetEraUnbondLimitResponse {
    /** @format int64 */
    limit?: number;
}
export interface LedgerQueryGetExchangeRateResponse {
    exchangeRate?: LedgerExchangeRate;
}
export interface LedgerQueryGetLeastBondResponse {
    leastBond?: LedgerLeastBond;
}
export interface LedgerQueryGetPoolDetailResponse {
    detail?: LedgerPoolDetail;
}
export interface LedgerQueryGetPoolUnbondResponse {
    unbond?: LedgerPoolUnbond;
}
export interface LedgerQueryGetReceiverResponse {
    receiver?: string;
}
export interface LedgerQueryGetSnapshotResponse {
    shot?: LedgerBondSnapshot;
}
export interface LedgerQueryGetTotalExpectedActiveResponse {
    active?: string;
}
export interface LedgerQueryGetUnbondCommissionResponse {
    commission?: string;
}
export interface LedgerQueryGetUnbondFeeResponse {
    fee?: LedgerUnbondFee;
}
export interface LedgerQueryPoolsByDenomResponse {
    addrs?: string[];
}
export interface LedgerUnbondFee {
    value?: string;
}
export interface LedgerUnbonding {
    unbonder?: string;
    amount?: string;
    recipient?: string;
}
export interface LedgerUserUnlockChunk {
    pool?: string;
    /** @format int64 */
    unlockEra?: number;
    value?: string;
    recipient?: string;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title ledger/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetEraExchangeRate
     * @summary Queries a list of getEraExchangeRate items.
     * @request GET:/stafiprotocol/stafihub/ledger/EraExchangeRate/{denom}/{era}
     */
    queryGetEraExchangeRate: (denom: string, era: number, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetEraExchangeRateResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryBondedPoolsByDenom
     * @summary Queries a list of bondedPoolsByDenom items.
     * @request GET:/stafiprotocol/stafihub/ledger/bondedPoolsByDenom/{denom}
     */
    queryBondedPoolsByDenom: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryBondedPoolsByDenomResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryEraExchangeRatesByDenom
     * @summary Queries a list of eraExchangeRatesByDenom items.
     * @request GET:/stafiprotocol/stafihub/ledger/eraExchangeRatesByDenom/{denom}
     */
    queryEraExchangeRatesByDenom: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryEraExchangeRatesByDenomResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryExchangeRateAll
     * @summary Queries a list of exchangeRateAll items.
     * @request GET:/stafiprotocol/stafihub/ledger/exchangeRateAll
     */
    queryExchangeRateAll: (params?: RequestParams) => Promise<HttpResponse<LedgerQueryExchangeRateAllResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetExchangeRate
     * @summary Queries a list of getExchangeRate items.
     * @request GET:/stafiprotocol/stafihub/ledger/exchangerate/{denom}
     */
    queryGetExchangeRate: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetExchangeRateResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetAccountUnbond
     * @summary Queries a list of getAccountUnbond items.
     * @request GET:/stafiprotocol/stafihub/ledger/getAccountUnbond/{denom}/{unbonder}
     */
    queryGetAccountUnbond: (denom: string, unbonder: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetAccountUnbondResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetBondPipeLine
     * @summary Queries a list of getBondPipeLine items.
     * @request GET:/stafiprotocol/stafihub/ledger/getBondPipeLine/{denom}/{pool}
     */
    queryGetBondPipeLine: (denom: string, pool: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetBondPipeLineResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetBondRecord
     * @summary Queries a list of getBondRecord items.
     * @request GET:/stafiprotocol/stafihub/ledger/getBondRecord/{denom}/{blockhash}/{txhash}
     */
    queryGetBondRecord: (denom: string, blockhash: string, txhash: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetBondRecordResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetChainBondingDuration
     * @summary Queries a list of getChainBondingDuration items.
     * @request GET:/stafiprotocol/stafihub/ledger/getChainBondingDuration/{denom}
     */
    queryGetChainBondingDuration: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetChainBondingDurationResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetChainEra
     * @summary Queries a list of getChainEra items.
     * @request GET:/stafiprotocol/stafihub/ledger/getChainEra/{denom}
     */
    queryGetChainEra: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetChainEraResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetCommission
     * @summary Queries a list of getCommission items.
     * @request GET:/stafiprotocol/stafihub/ledger/getCommission
     */
    queryGetCommission: (params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetCommissionResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetCurrentEraSnapshot
     * @summary Queries a list of getCurrentEraSnapshot items.
     * @request GET:/stafiprotocol/stafihub/ledger/getCurrentEraSnapshot/{denom}
     */
    queryGetCurrentEraSnapshot: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetCurrentEraSnapshotResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetEraSnapshot
     * @summary Queries a list of getEraSnapshot items.
     * @request GET:/stafiprotocol/stafihub/ledger/getEraSnapshot/{denom}/{era}
     */
    queryGetEraSnapshot: (denom: string, era: number, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetEraSnapshotResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetEraUnbondLimit
     * @summary Queries a list of getEraUnbondLimit items.
     * @request GET:/stafiprotocol/stafihub/ledger/getEraUnbondLimit/{denom}
     */
    queryGetEraUnbondLimit: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetEraUnbondLimitResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetLeastBond
     * @summary Queries a list of getLeastBond items.
     * @request GET:/stafiprotocol/stafihub/ledger/getLeastBond/{denom}
     */
    queryGetLeastBond: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetLeastBondResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetPoolDetail
     * @summary Queries a list of getPoolDetail items.
     * @request GET:/stafiprotocol/stafihub/ledger/getPoolDetail/{denom}/{pool}
     */
    queryGetPoolDetail: (denom: string, pool: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetPoolDetailResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetPoolUnbond
     * @summary Queries a list of getPoolUnbond items.
     * @request GET:/stafiprotocol/stafihub/ledger/getPoolUnbond/{denom}/{pool}/{era}
     */
    queryGetPoolUnbond: (denom: string, pool: string, era: number, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetPoolUnbondResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetReceiver
     * @summary Queries a list of getReceiver items.
     * @request GET:/stafiprotocol/stafihub/ledger/getReceiver
     */
    queryGetReceiver: (params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetReceiverResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetSnapshot
     * @summary Queries a list of getSnapshot items.
     * @request GET:/stafiprotocol/stafihub/ledger/getSnapshot/{shotId}
     */
    queryGetSnapshot: (shotId: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetSnapshotResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetTotalExpectedActive
     * @summary Queries a list of getTotalExpectedActive items.
     * @request GET:/stafiprotocol/stafihub/ledger/getTotalExpectedActive/{denom}/{era}
     */
    queryGetTotalExpectedActive: (denom: string, era: number, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetTotalExpectedActiveResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetUnbondCommission
     * @summary Queries a list of getUnbondCommission items.
     * @request GET:/stafiprotocol/stafihub/ledger/getUnbondCommission
     */
    queryGetUnbondCommission: (params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetUnbondCommissionResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetUnbondFee
     * @summary Queries a list of getUnbondFee items.
     * @request GET:/stafiprotocol/stafihub/ledger/getUnbondFee
     */
    queryGetUnbondFee: (params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetUnbondFeeResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryPoolsByDenom
     * @summary Queries a list of poolsByDenom items.
     * @request GET:/stafiprotocol/stafihub/ledger/poolsByDenom/{denom}
     */
    queryPoolsByDenom: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryPoolsByDenomResponse, RpcStatus>>;
}
export {};
