import { txClient, queryClient, MissingWalletError } from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'

import { ChainEra } from "./module/types/ledger/ledger"
import { ChainBondingDuration } from "./module/types/ledger/ledger"
import { Pool } from "./module/types/ledger/ledger"
import { TotalExpectedActive } from "./module/types/ledger/ledger"
import { BondPipeline } from "./module/types/ledger/ledger"
import { EraSnapshot } from "./module/types/ledger/ledger"
import { PoolUnbond } from "./module/types/ledger/ledger"
import { EraUnbondLimit } from "./module/types/ledger/ledger"
import { PoolDetail } from "./module/types/ledger/ledger"
import { LeastBond } from "./module/types/ledger/ledger"
import { LinkChunk } from "./module/types/ledger/ledger"
import { BondSnapshot } from "./module/types/ledger/ledger"
import { ExchangeRate } from "./module/types/ledger/ledger"
import { EraExchangeRate } from "./module/types/ledger/ledger"
import { UnbondFee } from "./module/types/ledger/ledger"
import { Unbonding } from "./module/types/ledger/ledger"
import { UserUnlockChunk } from "./module/types/ledger/ledger"
import { AccountUnbond } from "./module/types/ledger/ledger"
import { BondRecord } from "./module/types/ledger/ledger"
import { SetChainEraProposal } from "./module/types/ledger/proposal"
import { BondReportProposal } from "./module/types/ledger/proposal"
import { BondAndReportActiveProposal } from "./module/types/ledger/proposal"
import { ActiveReportProposal } from "./module/types/ledger/proposal"
import { WithdrawReportProposal } from "./module/types/ledger/proposal"
import { TransferReportProposal } from "./module/types/ledger/proposal"
import { ExecuteBondProposal } from "./module/types/ledger/proposal"


export { ChainEra, ChainBondingDuration, Pool, TotalExpectedActive, BondPipeline, EraSnapshot, PoolUnbond, EraUnbondLimit, PoolDetail, LeastBond, LinkChunk, BondSnapshot, ExchangeRate, EraExchangeRate, UnbondFee, Unbonding, UserUnlockChunk, AccountUnbond, BondRecord, SetChainEraProposal, BondReportProposal, BondAndReportActiveProposal, ActiveReportProposal, WithdrawReportProposal, TransferReportProposal, ExecuteBondProposal };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				GetExchangeRate: {},
				ExchangeRateAll: {},
				GetEraExchangeRate: {},
				EraExchangeRatesByDenom: {},
				PoolsByDenom: {},
				BondedPoolsByDenom: {},
				GetPoolDetail: {},
				GetChainEra: {},
				GetCurrentEraSnapshot: {},
				GetReceiver: {},
				GetCommission: {},
				GetChainBondingDuration: {},
				GetUnbondFee: {},
				GetUnbondCommission: {},
				GetLeastBond: {},
				GetEraUnbondLimit: {},
				GetBondPipeLine: {},
				GetEraSnapshot: {},
				GetSnapshot: {},
				GetTotalExpectedActive: {},
				GetPoolUnbond: {},
				GetAccountUnbond: {},
				GetBondRecord: {},
				
				_Structure: {
						ChainEra: getStructure(ChainEra.fromPartial({})),
						ChainBondingDuration: getStructure(ChainBondingDuration.fromPartial({})),
						Pool: getStructure(Pool.fromPartial({})),
						TotalExpectedActive: getStructure(TotalExpectedActive.fromPartial({})),
						BondPipeline: getStructure(BondPipeline.fromPartial({})),
						EraSnapshot: getStructure(EraSnapshot.fromPartial({})),
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
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(subscription)
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(subscription)
		}
	},
	getters: {
				getGetExchangeRate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetExchangeRate[JSON.stringify(params)] ?? {}
		},
				getExchangeRateAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExchangeRateAll[JSON.stringify(params)] ?? {}
		},
				getGetEraExchangeRate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetEraExchangeRate[JSON.stringify(params)] ?? {}
		},
				getEraExchangeRatesByDenom: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EraExchangeRatesByDenom[JSON.stringify(params)] ?? {}
		},
				getPoolsByDenom: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PoolsByDenom[JSON.stringify(params)] ?? {}
		},
				getBondedPoolsByDenom: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BondedPoolsByDenom[JSON.stringify(params)] ?? {}
		},
				getGetPoolDetail: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetPoolDetail[JSON.stringify(params)] ?? {}
		},
				getGetChainEra: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetChainEra[JSON.stringify(params)] ?? {}
		},
				getGetCurrentEraSnapshot: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetCurrentEraSnapshot[JSON.stringify(params)] ?? {}
		},
				getGetReceiver: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetReceiver[JSON.stringify(params)] ?? {}
		},
				getGetCommission: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetCommission[JSON.stringify(params)] ?? {}
		},
				getGetChainBondingDuration: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetChainBondingDuration[JSON.stringify(params)] ?? {}
		},
				getGetUnbondFee: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetUnbondFee[JSON.stringify(params)] ?? {}
		},
				getGetUnbondCommission: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetUnbondCommission[JSON.stringify(params)] ?? {}
		},
				getGetLeastBond: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetLeastBond[JSON.stringify(params)] ?? {}
		},
				getGetEraUnbondLimit: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetEraUnbondLimit[JSON.stringify(params)] ?? {}
		},
				getGetBondPipeLine: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetBondPipeLine[JSON.stringify(params)] ?? {}
		},
				getGetEraSnapshot: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetEraSnapshot[JSON.stringify(params)] ?? {}
		},
				getGetSnapshot: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetSnapshot[JSON.stringify(params)] ?? {}
		},
				getGetTotalExpectedActive: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetTotalExpectedActive[JSON.stringify(params)] ?? {}
		},
				getGetPoolUnbond: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetPoolUnbond[JSON.stringify(params)] ?? {}
		},
				getGetAccountUnbond: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetAccountUnbond[JSON.stringify(params)] ?? {}
		},
				getGetBondRecord: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetBondRecord[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: stafiprotocol.stafihub.ledger initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					await dispatch(subscription.action, subscription.payload)
				}catch(e) {
					throw new SpVuexError('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryGetExchangeRate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetExchangeRate( key.denom)).data
				
					
				commit('QUERY', { query: 'GetExchangeRate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetExchangeRate', payload: { options: { all }, params: {...key},query }})
				return getters['getGetExchangeRate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetExchangeRate', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryExchangeRateAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryExchangeRateAll()).data
				
					
				commit('QUERY', { query: 'ExchangeRateAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryExchangeRateAll', payload: { options: { all }, params: {...key},query }})
				return getters['getExchangeRateAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryExchangeRateAll', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetEraExchangeRate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetEraExchangeRate( key.denom,  key.era)).data
				
					
				commit('QUERY', { query: 'GetEraExchangeRate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetEraExchangeRate', payload: { options: { all }, params: {...key},query }})
				return getters['getGetEraExchangeRate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetEraExchangeRate', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEraExchangeRatesByDenom({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryEraExchangeRatesByDenom( key.denom)).data
				
					
				commit('QUERY', { query: 'EraExchangeRatesByDenom', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEraExchangeRatesByDenom', payload: { options: { all }, params: {...key},query }})
				return getters['getEraExchangeRatesByDenom']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryEraExchangeRatesByDenom', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPoolsByDenom({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryPoolsByDenom( key.denom)).data
				
					
				commit('QUERY', { query: 'PoolsByDenom', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPoolsByDenom', payload: { options: { all }, params: {...key},query }})
				return getters['getPoolsByDenom']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryPoolsByDenom', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBondedPoolsByDenom({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryBondedPoolsByDenom( key.denom)).data
				
					
				commit('QUERY', { query: 'BondedPoolsByDenom', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBondedPoolsByDenom', payload: { options: { all }, params: {...key},query }})
				return getters['getBondedPoolsByDenom']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryBondedPoolsByDenom', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetPoolDetail({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetPoolDetail( key.denom,  key.pool)).data
				
					
				commit('QUERY', { query: 'GetPoolDetail', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetPoolDetail', payload: { options: { all }, params: {...key},query }})
				return getters['getGetPoolDetail']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetPoolDetail', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetChainEra({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetChainEra( key.denom)).data
				
					
				commit('QUERY', { query: 'GetChainEra', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetChainEra', payload: { options: { all }, params: {...key},query }})
				return getters['getGetChainEra']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetChainEra', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetCurrentEraSnapshot({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetCurrentEraSnapshot( key.denom)).data
				
					
				commit('QUERY', { query: 'GetCurrentEraSnapshot', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetCurrentEraSnapshot', payload: { options: { all }, params: {...key},query }})
				return getters['getGetCurrentEraSnapshot']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetCurrentEraSnapshot', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetReceiver({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetReceiver()).data
				
					
				commit('QUERY', { query: 'GetReceiver', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetReceiver', payload: { options: { all }, params: {...key},query }})
				return getters['getGetReceiver']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetReceiver', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetCommission({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetCommission()).data
				
					
				commit('QUERY', { query: 'GetCommission', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetCommission', payload: { options: { all }, params: {...key},query }})
				return getters['getGetCommission']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetCommission', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetChainBondingDuration({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetChainBondingDuration( key.denom)).data
				
					
				commit('QUERY', { query: 'GetChainBondingDuration', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetChainBondingDuration', payload: { options: { all }, params: {...key},query }})
				return getters['getGetChainBondingDuration']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetChainBondingDuration', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetUnbondFee({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetUnbondFee()).data
				
					
				commit('QUERY', { query: 'GetUnbondFee', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetUnbondFee', payload: { options: { all }, params: {...key},query }})
				return getters['getGetUnbondFee']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetUnbondFee', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetUnbondCommission({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetUnbondCommission()).data
				
					
				commit('QUERY', { query: 'GetUnbondCommission', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetUnbondCommission', payload: { options: { all }, params: {...key},query }})
				return getters['getGetUnbondCommission']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetUnbondCommission', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetLeastBond({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetLeastBond( key.denom)).data
				
					
				commit('QUERY', { query: 'GetLeastBond', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetLeastBond', payload: { options: { all }, params: {...key},query }})
				return getters['getGetLeastBond']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetLeastBond', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetEraUnbondLimit({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetEraUnbondLimit( key.denom)).data
				
					
				commit('QUERY', { query: 'GetEraUnbondLimit', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetEraUnbondLimit', payload: { options: { all }, params: {...key},query }})
				return getters['getGetEraUnbondLimit']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetEraUnbondLimit', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetBondPipeLine({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetBondPipeLine( key.denom,  key.pool)).data
				
					
				commit('QUERY', { query: 'GetBondPipeLine', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetBondPipeLine', payload: { options: { all }, params: {...key},query }})
				return getters['getGetBondPipeLine']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetBondPipeLine', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetEraSnapshot({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetEraSnapshot( key.denom,  key.era)).data
				
					
				commit('QUERY', { query: 'GetEraSnapshot', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetEraSnapshot', payload: { options: { all }, params: {...key},query }})
				return getters['getGetEraSnapshot']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetEraSnapshot', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetSnapshot({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetSnapshot( key.shotId)).data
				
					
				commit('QUERY', { query: 'GetSnapshot', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetSnapshot', payload: { options: { all }, params: {...key},query }})
				return getters['getGetSnapshot']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetSnapshot', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetTotalExpectedActive({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetTotalExpectedActive( key.denom,  key.era)).data
				
					
				commit('QUERY', { query: 'GetTotalExpectedActive', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetTotalExpectedActive', payload: { options: { all }, params: {...key},query }})
				return getters['getGetTotalExpectedActive']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetTotalExpectedActive', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetPoolUnbond({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetPoolUnbond( key.denom,  key.pool,  key.era)).data
				
					
				commit('QUERY', { query: 'GetPoolUnbond', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetPoolUnbond', payload: { options: { all }, params: {...key},query }})
				return getters['getGetPoolUnbond']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetPoolUnbond', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetAccountUnbond({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetAccountUnbond( key.denom,  key.unbonder)).data
				
					
				commit('QUERY', { query: 'GetAccountUnbond', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetAccountUnbond', payload: { options: { all }, params: {...key},query }})
				return getters['getGetAccountUnbond']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetAccountUnbond', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetBondRecord({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetBondRecord( key.denom,  key.blockhash,  key.txhash)).data
				
					
				commit('QUERY', { query: 'GetBondRecord', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetBondRecord', payload: { options: { all }, params: {...key},query }})
				return getters['getGetBondRecord']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryGetBondRecord', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSetInitBond({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetInitBond(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetInitBond:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetInitBond:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetEraUnbondLimit({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetEraUnbondLimit(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetUnbondCommission({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetUnbondCommission(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetUnbondCommission:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetUnbondCommission:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddNewPool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddNewPool(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgAddNewPool:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgAddNewPool:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetReceiver({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetReceiver(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetReceiver:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetReceiver:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetLeastBond({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetLeastBond(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetLeastBond:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetLeastBond:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgClearCurrentEraSnapShots({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgClearCurrentEraSnapShots(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetCommission({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetCommission(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetCommission:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetCommission:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetChainBondingDuration({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetChainBondingDuration(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRemovePool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRemovePool(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgRemovePool:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgRemovePool:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetPoolDetail({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetPoolDetail(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetPoolDetail:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetPoolDetail:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgLiquidityUnbond({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgLiquidityUnbond(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgLiquidityUnbond:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgLiquidityUnbond:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetUnbondFee({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetUnbondFee(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetUnbondFee:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetUnbondFee:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSetInitBond({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetInitBond(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetInitBond:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetInitBond:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetEraUnbondLimit({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetEraUnbondLimit(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetEraUnbondLimit:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetUnbondCommission({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetUnbondCommission(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetUnbondCommission:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetUnbondCommission:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgAddNewPool({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddNewPool(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgAddNewPool:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgAddNewPool:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetReceiver({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetReceiver(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetReceiver:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetReceiver:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetLeastBond({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetLeastBond(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetLeastBond:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetLeastBond:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgClearCurrentEraSnapShots({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgClearCurrentEraSnapShots(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgClearCurrentEraSnapShots:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetCommission({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetCommission(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetCommission:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetCommission:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetChainBondingDuration({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetChainBondingDuration(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetChainBondingDuration:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgRemovePool({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRemovePool(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgRemovePool:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgRemovePool:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetPoolDetail({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetPoolDetail(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetPoolDetail:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetPoolDetail:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgLiquidityUnbond({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgLiquidityUnbond(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgLiquidityUnbond:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgLiquidityUnbond:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSetUnbondFee({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetUnbondFee(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSetUnbondFee:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSetUnbondFee:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		
	}
}
