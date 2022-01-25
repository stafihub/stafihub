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
/**
 * OriginalTxType enumerates the tx type of a signature.
 */
export var LedgerOriginalTxType;
(function (LedgerOriginalTxType) {
    LedgerOriginalTxType["ORIGINAL_TX_TYPE_TRANSFER"] = "ORIGINAL_TX_TYPE_TRANSFER";
    LedgerOriginalTxType["ORIGINAL_TX_TYPE_BOND"] = "ORIGINAL_TX_TYPE_BOND";
    LedgerOriginalTxType["ORIGINAL_TX_TYPE_UNBOND"] = "ORIGINAL_TX_TYPE_UNBOND";
    LedgerOriginalTxType["ORIGINAL_TX_TYPE_WITHDRAW"] = "ORIGINAL_TX_TYPE_WITHDRAW";
    LedgerOriginalTxType["ORIGINAL_TX_TYPE_CLAIM"] = "ORIGINAL_TX_TYPE_CLAIM";
})(LedgerOriginalTxType || (LedgerOriginalTxType = {}));
export var LedgerPoolBondState;
(function (LedgerPoolBondState) {
    LedgerPoolBondState["ERA_UPDATED"] = "ERA_UPDATED";
    LedgerPoolBondState["BOND_REPORTED"] = "BOND_REPORTED";
    LedgerPoolBondState["ACTIVE_REPORTED"] = "ACTIVE_REPORTED";
    LedgerPoolBondState["WITHDRAW_SKIPPED"] = "WITHDRAW_SKIPPED";
    LedgerPoolBondState["WITHDRAW_REPORTED"] = "WITHDRAW_REPORTED";
    LedgerPoolBondState["TRANSFER_REPORTED"] = "TRANSFER_REPORTED";
})(LedgerPoolBondState || (LedgerPoolBondState = {}));
export var ContentType;
(function (ContentType) {
    ContentType["Json"] = "application/json";
    ContentType["FormData"] = "multipart/form-data";
    ContentType["UrlEncoded"] = "application/x-www-form-urlencoded";
})(ContentType || (ContentType = {}));
export class HttpClient {
    constructor(apiConfig = {}) {
        this.baseUrl = "";
        this.securityData = null;
        this.securityWorker = null;
        this.abortControllers = new Map();
        this.baseApiParams = {
            credentials: "same-origin",
            headers: {},
            redirect: "follow",
            referrerPolicy: "no-referrer",
        };
        this.setSecurityData = (data) => {
            this.securityData = data;
        };
        this.contentFormatters = {
            [ContentType.Json]: (input) => input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
            [ContentType.FormData]: (input) => Object.keys(input || {}).reduce((data, key) => {
                data.append(key, input[key]);
                return data;
            }, new FormData()),
            [ContentType.UrlEncoded]: (input) => this.toQueryString(input),
        };
        this.createAbortSignal = (cancelToken) => {
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
        this.abortRequest = (cancelToken) => {
            const abortController = this.abortControllers.get(cancelToken);
            if (abortController) {
                abortController.abort();
                this.abortControllers.delete(cancelToken);
            }
        };
        this.request = ({ body, secure, path, type, query, format = "json", baseUrl, cancelToken, ...params }) => {
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
                const r = response;
                r.data = null;
                r.error = null;
                const data = await response[format]()
                    .then((data) => {
                    if (r.ok) {
                        r.data = data;
                    }
                    else {
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
                if (!response.ok)
                    throw data;
                return data;
            });
        };
        Object.assign(this, apiConfig);
    }
    addQueryParam(query, key) {
        const value = query[key];
        return (encodeURIComponent(key) +
            "=" +
            encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`));
    }
    toQueryString(rawQuery) {
        const query = rawQuery || {};
        const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
        return keys
            .map((key) => typeof query[key] === "object" && !Array.isArray(query[key])
            ? this.toQueryString(query[key])
            : this.addQueryParam(query, key))
            .join("&");
    }
    addQueryParams(rawQuery) {
        const queryString = this.toQueryString(rawQuery);
        return queryString ? `?${queryString}` : "";
    }
    mergeRequestParams(params1, params2) {
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
}
/**
 * @title ledger/genesis.proto
 * @version version not set
 */
export class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryGetEraExchangeRate
         * @summary Queries a list of getEraExchangeRate items.
         * @request GET:/stafiprotocol/stafihub/ledger/EraExchangeRate/{denom}/{era}
         */
        this.queryGetEraExchangeRate = (denom, era, params = {}) => this.request({
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
        this.queryBondedPoolsByDenom = (denom, params = {}) => this.request({
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
        this.queryEraExchangeRatesByDenom = (denom, params = {}) => this.request({
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
        this.queryExchangeRateAll = (params = {}) => this.request({
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
        this.queryGetExchangeRate = (denom, params = {}) => this.request({
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
        this.queryGetAccountUnbond = (denom, unbonder, params = {}) => this.request({
            path: `/stafiprotocol/stafihub/ledger/getAccountUnbond/${denom}/${unbonder}`,
            method: "GET",
            format: "json",
            ...params,
        });
        /**
         * No description
         *
         * @tags Query
         * @name QueryGetBondPipeline
         * @summary Queries a list of getBondPipeLine items.
         * @request GET:/stafiprotocol/stafihub/ledger/getBondPipeline/{denom}/{pool}
         */
        this.queryGetBondPipeline = (denom, pool, params = {}) => this.request({
            path: `/stafiprotocol/stafihub/ledger/getBondPipeline/${denom}/${pool}`,
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
        this.queryGetBondRecord = (denom, blockhash, txhash, params = {}) => this.request({
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
        this.queryGetChainBondingDuration = (denom, params = {}) => this.request({
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
        this.queryGetChainEra = (denom, params = {}) => this.request({
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
        this.queryGetCommission = (params = {}) => this.request({
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
        this.queryGetCurrentEraSnapshot = (denom, params = {}) => this.request({
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
        this.queryGetEraSnapshot = (denom, era, params = {}) => this.request({
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
        this.queryGetEraUnbondLimit = (denom, params = {}) => this.request({
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
        this.queryGetLeastBond = (denom, params = {}) => this.request({
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
        this.queryGetPoolDetail = (denom, pool, params = {}) => this.request({
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
        this.queryGetPoolUnbond = (denom, pool, era, params = {}) => this.request({
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
        this.queryGetReceiver = (params = {}) => this.request({
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
        this.queryGetSnapshot = (shotId, params = {}) => this.request({
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
        this.queryGetTotalExpectedActive = (denom, era, params = {}) => this.request({
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
        this.queryGetUnbondCommission = (params = {}) => this.request({
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
        this.queryGetUnbondFee = (params = {}) => this.request({
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
        this.queryPoolsByDenom = (denom, params = {}) => this.request({
            path: `/stafiprotocol/stafihub/ledger/poolsByDenom/${denom}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
