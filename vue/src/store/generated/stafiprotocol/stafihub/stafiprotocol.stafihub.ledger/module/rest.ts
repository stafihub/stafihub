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

export type LedgerMsgAddNewPoolResponse = object;

export type LedgerMsgClearCurrentEraSnapShotsResponse = object;

export type LedgerMsgRemovePoolResponse = object;

export type LedgerMsgSetChainBondingDurationResponse = object;

export type LedgerMsgSetCommissionResponse = object;

export type LedgerMsgSetEraUnbondLimitResponse = object;

export type LedgerMsgSetInitBondResponse = object;

export type LedgerMsgSetLeastBondResponse = object;

export type LedgerMsgSetPoolDetailResponse = object;

export type LedgerMsgSetReceiverResponse = object;

export interface LedgerQueryEraExchangeRatesByDenomResponse {
  eraExchangeRates?: LedgerEraExchangeRate[];
}

export interface LedgerQueryExchangeRateAllResponse {
  exchangeRates?: LedgerExchangeRate[];
}

export interface LedgerQueryGetEraExchangeRateResponse {
  eraExchangeRate?: LedgerEraExchangeRate;
}

export interface LedgerQueryGetExchangeRateResponse {
  exchangeRate?: LedgerExchangeRate;
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
}
