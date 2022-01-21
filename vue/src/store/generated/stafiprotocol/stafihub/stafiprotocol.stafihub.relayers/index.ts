import { txClient, queryClient, MissingWalletError } from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'

import { Relayer } from "./module/types/relayers/relayer"
import { Threshold } from "./module/types/relayers/relayer"
import { LastVoter } from "./module/types/relayers/relayer"


export { Relayer, Threshold, LastVoter };

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
				RelayerAll: {},
				RelayersByDenom: {},
				Threshold: {},
				ThresholdAll: {},
				
				_Structure: {
						Relayer: getStructure(Relayer.fromPartial({})),
						Threshold: getStructure(Threshold.fromPartial({})),
						LastVoter: getStructure(LastVoter.fromPartial({})),
						
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
				getRelayerAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RelayerAll[JSON.stringify(params)] ?? {}
		},
				getRelayersByDenom: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RelayersByDenom[JSON.stringify(params)] ?? {}
		},
				getThreshold: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Threshold[JSON.stringify(params)] ?? {}
		},
				getThresholdAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ThresholdAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: stafiprotocol.stafihub.relayers initialized!')
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
		
		
		
		 		
		
		
		async QueryRelayerAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRelayerAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.nextKey!=null) {
					let next_values=(await queryClient.queryRelayerAll({...query, 'pagination.key':(<any> value).pagination.nextKey})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RelayerAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRelayerAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRelayerAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryRelayerAll', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRelayersByDenom({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryRelayersByDenom( key.denom, query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.nextKey!=null) {
					let next_values=(await queryClient.queryRelayersByDenom( key.denom, {...query, 'pagination.key':(<any> value).pagination.nextKey})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RelayersByDenom', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRelayersByDenom', payload: { options: { all }, params: {...key},query }})
				return getters['getRelayersByDenom']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryRelayersByDenom', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryThreshold({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryThreshold( key.denom)).data
				
					
				commit('QUERY', { query: 'Threshold', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryThreshold', payload: { options: { all }, params: {...key},query }})
				return getters['getThreshold']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryThreshold', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryThresholdAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryThresholdAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.nextKey!=null) {
					let next_values=(await queryClient.queryThresholdAll({...query, 'pagination.key':(<any> value).pagination.nextKey})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ThresholdAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryThresholdAll', payload: { options: { all }, params: {...key},query }})
				return getters['getThresholdAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryThresholdAll', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateRelayer({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateRelayer(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateRelayer:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateRelayer:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteRelayer({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteRelayer(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDeleteRelayer:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteRelayer:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateThreshold({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateThreshold(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateThreshold:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateThreshold:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateRelayer({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateRelayer(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateRelayer:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateRelayer:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgDeleteRelayer({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteRelayer(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDeleteRelayer:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteRelayer:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgUpdateThreshold({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateThreshold(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateThreshold:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateThreshold:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		
	}
}
