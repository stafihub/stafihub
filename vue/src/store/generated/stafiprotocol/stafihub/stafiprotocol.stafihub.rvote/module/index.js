// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSetProposalLife } from "./types/rvote/tx";
import { MsgSubmitProposal } from "./types/rvote/tx";
const types = [
    ["/stafiprotocol.stafihub.rvote.MsgSetProposalLife", MsgSetProposalLife],
    ["/stafiprotocol.stafihub.rvote.MsgSubmitProposal", MsgSubmitProposal],
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
        msgSetProposalLife: (data) => ({ typeUrl: "/stafiprotocol.stafihub.rvote.MsgSetProposalLife", value: data }),
        msgSubmitProposal: (data) => ({ typeUrl: "/stafiprotocol.stafihub.rvote.MsgSubmitProposal", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
