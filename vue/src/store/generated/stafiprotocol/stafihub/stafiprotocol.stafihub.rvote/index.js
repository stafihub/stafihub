import { txClient, queryClient, MissingWalletError } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { Proposal } from "./module/types/rvote/proposal";
export { Proposal };
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
        _Structure: {
            Proposal: getStructure(Proposal.fromPartial({})),
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
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: stafiprotocol.stafihub.rvote initialized!');
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
        async sendMsgSetProposalLife({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetProposalLife(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetProposalLife:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetProposalLife:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSubmitProposal({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSubmitProposal(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSubmitProposal:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSubmitProposal:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgSetProposalLife({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSetProposalLife(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSetProposalLife:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSetProposalLife:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSubmitProposal({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSubmitProposal(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSubmitProposal:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSubmitProposal:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
