// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSetEraUnbondLimit } from "./types/ledger/tx";
import { MsgClearCurrentEraSnapShots } from "./types/ledger/tx";
import { MsgSetReceiver } from "./types/ledger/tx";
import { MsgRemovePool } from "./types/ledger/tx";
import { MsgSetPoolDetail } from "./types/ledger/tx";
import { MsgSetLeastBond } from "./types/ledger/tx";
import { MsgSetCommission } from "./types/ledger/tx";
import { MsgSetChainBondingDuration } from "./types/ledger/tx";
import { MsgAddNewPool } from "./types/ledger/tx";
import { MsgSetInitBond } from "./types/ledger/tx";
const types = [
    ["/stafiprotocol.stafihub.ledger.MsgSetEraUnbondLimit", MsgSetEraUnbondLimit],
    ["/stafiprotocol.stafihub.ledger.MsgClearCurrentEraSnapShots", MsgClearCurrentEraSnapShots],
    ["/stafiprotocol.stafihub.ledger.MsgSetReceiver", MsgSetReceiver],
    ["/stafiprotocol.stafihub.ledger.MsgRemovePool", MsgRemovePool],
    ["/stafiprotocol.stafihub.ledger.MsgSetPoolDetail", MsgSetPoolDetail],
    ["/stafiprotocol.stafihub.ledger.MsgSetLeastBond", MsgSetLeastBond],
    ["/stafiprotocol.stafihub.ledger.MsgSetCommission", MsgSetCommission],
    ["/stafiprotocol.stafihub.ledger.MsgSetChainBondingDuration", MsgSetChainBondingDuration],
    ["/stafiprotocol.stafihub.ledger.MsgAddNewPool", MsgAddNewPool],
    ["/stafiprotocol.stafihub.ledger.MsgSetInitBond", MsgSetInitBond],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSetEraUnbondLimit: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetEraUnbondLimit", value: data }),
        msgClearCurrentEraSnapShots: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgClearCurrentEraSnapShots", value: data }),
        msgSetReceiver: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetReceiver", value: data }),
        msgRemovePool: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgRemovePool", value: data }),
        msgSetPoolDetail: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetPoolDetail", value: data }),
        msgSetLeastBond: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetLeastBond", value: data }),
        msgSetCommission: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetCommission", value: data }),
        msgSetChainBondingDuration: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetChainBondingDuration", value: data }),
        msgAddNewPool: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgAddNewPool", value: data }),
        msgSetInitBond: (data) => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetInitBond", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };