// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteRelayer } from "./types/relayers/tx";
import { MsgCreateRelayer } from "./types/relayers/tx";
import { MsgUpdateThreshold } from "./types/relayers/tx";


const types = [
  ["/stafiprotocol.stafihub.relayers.MsgDeleteRelayer", MsgDeleteRelayer],
  ["/stafiprotocol.stafihub.relayers.MsgCreateRelayer", MsgCreateRelayer],
  ["/stafiprotocol.stafihub.relayers.MsgUpdateThreshold", MsgUpdateThreshold],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgDeleteRelayer: (data: MsgDeleteRelayer): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.relayers.MsgDeleteRelayer", value: data }),
    msgCreateRelayer: (data: MsgCreateRelayer): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.relayers.MsgCreateRelayer", value: data }),
    msgUpdateThreshold: (data: MsgUpdateThreshold): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.relayers.MsgUpdateThreshold", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
