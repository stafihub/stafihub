import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSetProposalLife } from "./types/relayers/tx";
import { MsgUpdateThreshold } from "./types/relayers/tx";
import { MsgDeleteRelayer } from "./types/relayers/tx";
import { MsgCreateRelayer } from "./types/relayers/tx";
export declare const MissingWalletError: Error;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgSetProposalLife: (data: MsgSetProposalLife) => EncodeObject;
    msgUpdateThreshold: (data: MsgUpdateThreshold) => EncodeObject;
    msgDeleteRelayer: (data: MsgDeleteRelayer) => EncodeObject;
    msgCreateRelayer: (data: MsgCreateRelayer) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
