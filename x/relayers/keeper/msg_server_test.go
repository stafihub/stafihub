package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/relayers/keeper"
	"github.com/stafihub/stafihub/x/relayers/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RelayersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestRelayerMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RelayersKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	newRelayer := sample.AccAddress()
	creator := sample.TestAdmin

	expected := &types.MsgAddRelayer{
		Creator:   creator,
		Arena:     types.ModuleName,
		Denom:     sample.TestDenom,
		Addresses: []string{newRelayer},
	}
	_, err := srv.AddRelayer(wctx, expected)
	require.NoError(t, err)
	found := k.HasRelayer(ctx, expected.Arena, expected.Denom, expected.Addresses[0])
	require.True(t, found)
}

func TestRelayerMsgServerDelete(t *testing.T) {
	creator := sample.TestAdmin
	relayer := sample.AccAddress()

	k, ctx := keepertest.RelayersKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteRelayer
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRelayer{
				Creator: creator,
				Arena:   types.ModuleName,
				Denom:   sample.TestDenom,
				Address: relayer,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRelayer{
				Creator: "a",
				Arena:   types.ModuleName,
				Denom:   sample.TestDenom,
				Address: relayer,
			},
			err: sudotypes.ErrCreatorNotAdmin,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {

			_, err := srv.AddRelayer(wctx, &types.MsgAddRelayer{
				Creator:   creator,
				Arena:     tc.request.Arena,
				Denom:     tc.request.Denom,
				Addresses: []string{tc.request.Address},
			})
			require.NoError(t, err)

			found := k.HasRelayer(ctx, tc.request.Arena, tc.request.Denom, tc.request.Address)
			require.True(t, found)

			_, err = srv.DeleteRelayer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				found := k.HasRelayer(ctx, tc.request.Arena, tc.request.Denom, tc.request.Address)
				require.False(t, found)
			}
		})
	}
}
