import { txClient, queryClient, MissingWalletError } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { ChainEra } from "./module/types/ledger/ledger";
import { ChainBondingDuration } from "./module/types/ledger/ledger";
import { Pool } from "./module/types/ledger/ledger";
import { TotalExpectedActive } from "./module/types/ledger/ledger";
import { BondPipeline } from "./module/types/ledger/ledger";
import { EraSnapShot } from "./module/types/ledger/ledger";
import { PoolUnbond } from "./module/types/ledger/ledger";
import { EraUnbondLimit } from "./module/types/ledger/ledger";
import { PoolDetail } from "./module/types/ledger/ledger";
import { LeastBond } from "./module/types/ledger/ledger";
import { LinkChunk } from "./module/types/ledger/ledger";
import { BondSnapshot } from "./module/types/ledger/ledger";
import { ExchangeRate } from "./module/types/ledger/ledger";
import { EraExchangeRate } from "./module/types/ledger/ledger";
import { UnbondFee } from "./module/types/ledger/ledger";
import { Unbonding } from "./module/types/ledger/ledger";
import { UserUnlockChunk } from "./module/types/ledger/ledger";
import { AccountUnbond } from "./module/types/ledger/ledger";
import { BondRecord } from "./module/types/ledger/ledger";
import { SetChainEraProposal } from "./module/types/ledger/proposal";
import { BondReportProposal } from "./module/types/ledger/proposal";
import { BondAndReportActiveProposal } from "./module/types/ledger/proposal";
import { ActiveReportProposal } from "./module/types/ledger/proposal";
import { WithdrawReportProposal } from "./module/types/ledger/proposal";
import { TransferReportProposal } from "./module/types/ledger/proposal";
import { ExecuteBondProposal } from "./module/types/ledger/proposal";
export { ChainEra, ChainBondingDuration, Pool, TotalExpectedActive, BondPipeline, EraSnapShot, PoolUnbond, EraUnbondLimit, PoolDetail, LeastBond, LinkChunk, BondSnapshot, ExchangeRate, EraExchangeRate, UnbondFee, Unbonding, UserUnlockChunk, AccountUnbond, BondRecord, SetChainEraProposal, BondReportProposal, BondAndReportActiveProposal, ActiveReportProposal, WithdrawReportProposal, TransferReportProposal, ExecuteBondProposal };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        GetExchangeRate: {},
        ExchangeRateAll: {},
        GetEraExchangeRate: {},
        EraExchangeRatesByDenom: {},
        _Structure: {
            ChainEra: getStructure(ChainEra.fromPartial({})),
            ChainBondingDuration: getStructure(ChainBondingDuration.fromPartial({})),
            Pool: getStructure(Pool.fromPartial({})),
            TotalExpectedActive: getStructure(TotalExpectedActive.fromPartial({})),
            BondPipeline: getStructure(BondPipeline.fromPartial({})),
            EraSnapShot: getStructure(EraSnapShot.fromPartial({})),
            PoolUnbond: getStructure(PoolUnbond.fromPartial({})),
            EraUnbondLimit: getStructure(EraUnbondLimit.fromPartial({})),
            PoolDetail: getStructure(PoolDetail.fromPartial({})),
            LeastBond: getStructure(LeastBond.fromPartial({})),
            LinkChunk: getStructure(LinkChunk.fromPartial({})),
            BondSnapshot: getStructure(BondSnapshot.fromPartial({})),
            ExchangeRate: getStructure(ExchangeRate.fromPartial({})),
            EraExchangeRate: getStructure(EraExchangeRate.fromPartial({})),
            UnbondFee: getStructure(UnbondFee.fromPartial({})),
            Unbonding: getStructure(Unbonding.fromPartial({})),
            UserUnlockChunk: getStructure(UserUnlockChunk.fromPartial({})),
            AccountUnbond: getStructure(AccountUnbond.fromPartial({})),
            BondRecord: getStructure(BondRecord.fromPartial({})),
            SetChainEraProposal: getStructure(SetChainEraProposal.fromPartial({})),
            BondReportProposal: getStructure(BondReportProposal.fromPartial({})),
            BondAndReportActiveProposal: getStructure(BondAndReportActiveProposal.fromPartial({})),
            ActiveReportProposal: getStructure(ActiveReportProposal.fromPartial({})),
            WithdrawReportProposal: getStructure(WithdrawReportProposal.fromPartial({})),
            TransferReportProposal: getStructure(TransferReportProposal.fromPartial({})),
            ExecuteBondProposal: getStructure(ExecuteBondProposal.fromPartial({})),
        },
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(subscription);
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(subscription);
        }
    },
    getters: {
        getGetExchangeRate: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.GetExchangeRate[JSON.stringify(params)] ?? {};
        },
        getExchangeRateAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ExchangeRateAll[JSON.stringify(params)] ?? {};
        },
        getGetEraExchangeRate: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.GetEraExchangeRate[JSON.stringify(params)] ?? {};
        },
        getEraExchangeRatesByDenom: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.EraExchangeRatesByDenom[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: stafiprotocol.stafihub.ledger initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    await dispatch(subscription.action, subscription.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryGetExchangeRate({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryGetExchangeRate(key.denom)).data;
                commit('QUERY', { query: 'GetExchangeRate', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryGetExchangeRate', payload: { options: { all }, params: { ...key }, query } });
                return getters['getGetExchangeRate']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryGetExchangeRate', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryExchangeRateAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryExchangeRateAll()).data;
                commit('QUERY', { query: 'ExchangeRateAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryExchangeRateAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getExchangeRateAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryExchangeRateAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryGetEraExchangeRate({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryGetEraExchangeRate(key.denom, key.era)).data;
                commit('QUERY', { query: 'GetEraExchangeRate', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryGetEraExchangeRate', payload: { options: { all }, params: { ...key }, query } });
                return getters['getGetEraExchangeRate']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryGetEraExchangeRate', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryEraExchangeRatesByDenom({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryEraExchangeRatesByDenom(key.denom)).data;
                commit('QUERY', { query: 'EraExchangeRatesByDenom', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryEraExchangeRatesByDenom', payload: { options: { all }, params: { ...key }, query } });
                return getters['getEraExchangeRatesByDenom']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryEraExchangeRatesByDenom', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgSetChainBondingDuration({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetChainBondingDuration(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgLiquidityUnbond({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgLiquidityUnbond(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgLiquidityUnbond:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgLiquidityUnbond:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetReceiver({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetReceiver(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetReceiver:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetReceiver:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetLeastBond({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetLeastBond(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetLeastBond:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetLeastBond:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetCommission({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetCommission(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetCommission:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetCommission:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgAddNewPool({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgAddNewPool(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgAddNewPool:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgAddNewPool:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetUnbondFee({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetUnbondFee(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetUnbondFee:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetUnbondFee:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgClearCurrentEraSnapShots({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgClearCurrentEraSnapShots(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetEraUnbondLimit({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetEraUnbondLimit(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgRemovePool({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRemovePool(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRemovePool:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRemovePool:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetUnbondCommission({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetUnbondCommission(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetUnbondCommission:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetUnbondCommission:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetPoolDetail({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetPoolDetail(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetPoolDetail:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetPoolDetail:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSetInitBond({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetInitBond(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetInitBond:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetInitBond:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgSetChainBondingDuration({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetChainBondingDuration(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgLiquidityUnbond({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgLiquidityUnbond(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgLiquidityUnbond:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgLiquidityUnbond:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetReceiver({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetReceiver(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetReceiver:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetReceiver:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetLeastBond({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetLeastBond(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetLeastBond:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetLeastBond:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetCommission({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetCommission(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetCommission:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetCommission:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgAddNewPool({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgAddNewPool(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgAddNewPool:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgAddNewPool:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetUnbondFee({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetUnbondFee(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetUnbondFee:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetUnbondFee:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgClearCurrentEraSnapShots({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgClearCurrentEraSnapShots(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetEraUnbondLimit({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetEraUnbondLimit(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgRemovePool({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRemovePool(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRemovePool:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRemovePool:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetUnbondCommission({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetUnbondCommission(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetUnbondCommission:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetUnbondCommission:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetPoolDetail({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetPoolDetail(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetPoolDetail:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetPoolDetail:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSetInitBond({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetInitBond(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetInitBond:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetInitBond:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
