package keeper_test

import (
    "strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

    keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
    "github.com/stafiprotocol/stafihub/x/relayers/keeper"
    "github.com/stafiprotocol/stafihub/x/relayers/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestThresholdMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RelayersKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateThreshold{Creator: creator,
		    Index: strconv.Itoa(i),
            
		}
		_, err := srv.CreateThreshold(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetThreshold(ctx,
		    expected.Index,
            
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestThresholdMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateThreshold
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateThreshold{Creator: creator,
			    Index: strconv.Itoa(0),
                
			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateThreshold{Creator: "B",
			    Index: strconv.Itoa(0),
                
			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdateThreshold{Creator: creator,
			    Index: strconv.Itoa(100000),
                
			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RelayersKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateThreshold{Creator: creator,
			    Index: strconv.Itoa(0),
                
			}
			_, err := srv.CreateThreshold(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateThreshold(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetThreshold(ctx,
				    expected.Index,
                    
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestThresholdMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteThreshold
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteThreshold{Creator: creator,
			    Index: strconv.Itoa(0),
                
			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteThreshold{Creator: "B",
			    Index: strconv.Itoa(0),
                
			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteThreshold{Creator: creator,
			    Index: strconv.Itoa(100000),
                
			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RelayersKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateThreshold(wctx, &types.MsgCreateThreshold{Creator: creator,
			    Index: strconv.Itoa(0),
                
			})
			require.NoError(t, err)
			_, err = srv.DeleteThreshold(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetThreshold(ctx,
				    tc.request.Index,
                    
				)
				require.False(t, found)
			}
		})
	}
}
