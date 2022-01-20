/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

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

export type LedgerMsgAddNewPoolResponse = object;

export type LedgerMsgClearCurrentEraSnapShotsResponse = object;

export type LedgerMsgLiquidityUnbondResponse = object;

export type LedgerMsgRemovePoolResponse = object;

export type LedgerMsgSetChainBondingDurationResponse = object;

export type LedgerMsgSetCommissionResponse = object;

export type LedgerMsgSetEraUnbondLimitResponse = object;

export type LedgerMsgSetInitBondResponse = object;

export type LedgerMsgSetLeastBondResponse = object;

export type LedgerMsgSetPoolDetailResponse = object;

export type LedgerMsgSetReceiverResponse = object;

export type LedgerMsgSetUnbondCommissionResponse = object;

export type LedgerMsgSetUnbondFeeResponse = object;

export enum LedgerPoolBondState {
  ERA_UPDATED = "ERA_UPDATED",
  BOND_REPORTED = "BOND_REPORTED",
  ACTIVE_REPORTED = "ACTIVE_REPORTED",
  WITHDRAW_SKIPPED = "WITHDRAW_SKIPPED",
  WITHDRAW_REPORTED = "WITHDRAW_REPORTED",
  TRANSFER_REPORTED = "TRANSFER_REPORTED",
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

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

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

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title ledger/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryGetEraExchangeRate
   * @summary Queries a list of getEraExchangeRate items.
   * @request GET:/stafiprotocol/stafihub/ledger/EraExchangeRate/{denom}/{era}
   */
  queryGetEraExchangeRate = (denom: string, era: number, params: RequestParams = {}) =>
    this.request<LedgerQueryGetEraExchangeRateResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/EraExchangeRate/${denom}/${era}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBondedPoolsByDenom
   * @summary Queries a list of bondedPoolsByDenom items.
   * @request GET:/stafiprotocol/stafihub/ledger/bondedPoolsByDenom/{denom}
   */
  queryBondedPoolsByDenom = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryBondedPoolsByDenomResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/bondedPoolsByDenom/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEraExchangeRatesByDenom
   * @summary Queries a list of eraExchangeRatesByDenom items.
   * @request GET:/stafiprotocol/stafihub/ledger/eraExchangeRatesByDenom/{denom}
   */
  queryEraExchangeRatesByDenom = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryEraExchangeRatesByDenomResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/eraExchangeRatesByDenom/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryExchangeRateAll
   * @summary Queries a list of exchangeRateAll items.
   * @request GET:/stafiprotocol/stafihub/ledger/exchangeRateAll
   */
  queryExchangeRateAll = (params: RequestParams = {}) =>
    this.request<LedgerQueryExchangeRateAllResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/exchangeRateAll`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetExchangeRate
   * @summary Queries a list of getExchangeRate items.
   * @request GET:/stafiprotocol/stafihub/ledger/exchangerate/{denom}
   */
  queryGetExchangeRate = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetExchangeRateResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/exchangerate/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetAccountUnbond
   * @summary Queries a list of getAccountUnbond items.
   * @request GET:/stafiprotocol/stafihub/ledger/getAccountUnbond/{denom}/{unbonder}
   */
  queryGetAccountUnbond = (denom: string, unbonder: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetAccountUnbondResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getAccountUnbond/${denom}/${unbonder}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetBondPipeLine
   * @summary Queries a list of getBondPipeLine items.
   * @request GET:/stafiprotocol/stafihub/ledger/getBondPipeLine/{denom}/{pool}
   */
  queryGetBondPipeLine = (denom: string, pool: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetBondPipeLineResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getBondPipeLine/${denom}/${pool}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetBondRecord
   * @summary Queries a list of getBondRecord items.
   * @request GET:/stafiprotocol/stafihub/ledger/getBondRecord/{denom}/{blockhash}/{txhash}
   */
  queryGetBondRecord = (denom: string, blockhash: string, txhash: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetBondRecordResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getBondRecord/${denom}/${blockhash}/${txhash}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetChainBondingDuration
   * @summary Queries a list of getChainBondingDuration items.
   * @request GET:/stafiprotocol/stafihub/ledger/getChainBondingDuration/{denom}
   */
  queryGetChainBondingDuration = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetChainBondingDurationResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getChainBondingDuration/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetChainEra
   * @summary Queries a list of getChainEra items.
   * @request GET:/stafiprotocol/stafihub/ledger/getChainEra/{denom}
   */
  queryGetChainEra = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetChainEraResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getChainEra/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetCommission
   * @summary Queries a list of getCommission items.
   * @request GET:/stafiprotocol/stafihub/ledger/getCommission
   */
  queryGetCommission = (params: RequestParams = {}) =>
    this.request<LedgerQueryGetCommissionResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getCommission`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetCurrentEraSnapshot
   * @summary Queries a list of getCurrentEraSnapshot items.
   * @request GET:/stafiprotocol/stafihub/ledger/getCurrentEraSnapshot/{denom}
   */
  queryGetCurrentEraSnapshot = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetCurrentEraSnapshotResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getCurrentEraSnapshot/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetEraSnapshot
   * @summary Queries a list of getEraSnapshot items.
   * @request GET:/stafiprotocol/stafihub/ledger/getEraSnapshot/{denom}/{era}
   */
  queryGetEraSnapshot = (denom: string, era: number, params: RequestParams = {}) =>
    this.request<LedgerQueryGetEraSnapshotResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getEraSnapshot/${denom}/${era}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetEraUnbondLimit
   * @summary Queries a list of getEraUnbondLimit items.
   * @request GET:/stafiprotocol/stafihub/ledger/getEraUnbondLimit/{denom}
   */
  queryGetEraUnbondLimit = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetEraUnbondLimitResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getEraUnbondLimit/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetLeastBond
   * @summary Queries a list of getLeastBond items.
   * @request GET:/stafiprotocol/stafihub/ledger/getLeastBond/{denom}
   */
  queryGetLeastBond = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetLeastBondResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getLeastBond/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetPoolDetail
   * @summary Queries a list of getPoolDetail items.
   * @request GET:/stafiprotocol/stafihub/ledger/getPoolDetail/{denom}/{pool}
   */
  queryGetPoolDetail = (denom: string, pool: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetPoolDetailResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getPoolDetail/${denom}/${pool}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetPoolUnbond
   * @summary Queries a list of getPoolUnbond items.
   * @request GET:/stafiprotocol/stafihub/ledger/getPoolUnbond/{denom}/{pool}/{era}
   */
  queryGetPoolUnbond = (denom: string, pool: string, era: number, params: RequestParams = {}) =>
    this.request<LedgerQueryGetPoolUnbondResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getPoolUnbond/${denom}/${pool}/${era}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetReceiver
   * @summary Queries a list of getReceiver items.
   * @request GET:/stafiprotocol/stafihub/ledger/getReceiver
   */
  queryGetReceiver = (params: RequestParams = {}) =>
    this.request<LedgerQueryGetReceiverResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getReceiver`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetSnapshot
   * @summary Queries a list of getSnapshot items.
   * @request GET:/stafiprotocol/stafihub/ledger/getSnapshot/{shotId}
   */
  queryGetSnapshot = (shotId: string, params: RequestParams = {}) =>
    this.request<LedgerQueryGetSnapshotResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getSnapshot/${shotId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetTotalExpectedActive
   * @summary Queries a list of getTotalExpectedActive items.
   * @request GET:/stafiprotocol/stafihub/ledger/getTotalExpectedActive/{denom}/{era}
   */
  queryGetTotalExpectedActive = (denom: string, era: number, params: RequestParams = {}) =>
    this.request<LedgerQueryGetTotalExpectedActiveResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getTotalExpectedActive/${denom}/${era}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetUnbondCommission
   * @summary Queries a list of getUnbondCommission items.
   * @request GET:/stafiprotocol/stafihub/ledger/getUnbondCommission
   */
  queryGetUnbondCommission = (params: RequestParams = {}) =>
    this.request<LedgerQueryGetUnbondCommissionResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getUnbondCommission`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetUnbondFee
   * @summary Queries a list of getUnbondFee items.
   * @request GET:/stafiprotocol/stafihub/ledger/getUnbondFee
   */
  queryGetUnbondFee = (params: RequestParams = {}) =>
    this.request<LedgerQueryGetUnbondFeeResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/getUnbondFee`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPoolsByDenom
   * @summary Queries a list of poolsByDenom items.
   * @request GET:/stafiprotocol/stafihub/ledger/poolsByDenom/{denom}
   */
  queryPoolsByDenom = (denom: string, params: RequestParams = {}) =>
    this.request<LedgerQueryPoolsByDenomResponse, RpcStatus>({
      path: `/stafiprotocol/stafihub/ledger/poolsByDenom/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
