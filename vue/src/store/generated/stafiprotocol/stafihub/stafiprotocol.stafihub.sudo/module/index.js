// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddDenom } from "./types/sudo/tx";
import { MsgUpdateAdmin } from "./types/sudo/tx";
const types = [
    ["/stafiprotocol.stafihub.sudo.MsgAddDenom", MsgAddDenom],
    ["/stafiprotocol.stafihub.sudo.MsgUpdateAdmin", MsgUpdateAdmin],
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
        msgAddDenom: (data) => ({ typeUrl: "/stafiprotocol.stafihub.sudo.MsgAddDenom", value: data }),
        msgUpdateAdmin: (data) => ({ typeUrl: "/stafiprotocol.stafihub.sudo.MsgUpdateAdmin", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
