// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateRelayer } from "./types/relayers/tx";
import { MsgDeleteRelayer } from "./types/relayers/tx";
import { MsgUpdateThreshold } from "./types/relayers/tx";
const types = [
    ["/stafiprotocol.stafihub.relayers.MsgCreateRelayer", MsgCreateRelayer],
    ["/stafiprotocol.stafihub.relayers.MsgDeleteRelayer", MsgDeleteRelayer],
    ["/stafiprotocol.stafihub.relayers.MsgUpdateThreshold", MsgUpdateThreshold],
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
        msgCreateRelayer: (data) => ({ typeUrl: "/stafiprotocol.stafihub.relayers.MsgCreateRelayer", value: data }),
        msgDeleteRelayer: (data) => ({ typeUrl: "/stafiprotocol.stafihub.relayers.MsgDeleteRelayer", value: data }),
        msgUpdateThreshold: (data) => ({ typeUrl: "/stafiprotocol.stafihub.relayers.MsgUpdateThreshold", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
