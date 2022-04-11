package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/types"
)

func (k Keeper) AddRelayer(ctx sdk.Context, arena, denom, addr string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RelayerStoreKey(arena, denom, addr), []byte{})
}

func (k Keeper) HasRelayer(ctx sdk.Context, arena, denom, addr string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RelayerStoreKey(arena, denom, addr))
}

func (k Keeper) RemoveRelayer(ctx sdk.Context, arena, denom, addr string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.RelayerStoreKey(arena, denom, addr))
}

func (k Keeper) GetRelayer(ctx sdk.Context, arena, denom string) []string {
	keyPrefix := make([]byte, 1+2+len(arena)+len(denom))
	keyPrefix[0] = types.RelayerPrefix[0]
	keyPrefix[1] = byte(len(arena))
	copy(keyPrefix[2:], arena)
	keyPrefix[2+len(arena)] = byte(len(denom))
	copy(keyPrefix[3+len(arena):], denom)

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, keyPrefix)
	defer iterator.Close()

	relayers := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		if len(key) <= len(keyPrefix)+1 {
			continue
		}
		relayers = append(relayers, string(key[len(keyPrefix)+1:]))
	}
	return relayers
}

// used for export genesis
func (k Keeper) GetAllRelayer(ctx sdk.Context) []types.Relayer {
	keyPrefix := make([]byte, 1)
	keyPrefix[0] = types.RelayerPrefix[0]

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, keyPrefix)
	defer iterator.Close()

	relayers := make([]types.Relayer, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		if len(key) <= len(keyPrefix) {
			continue
		}
		arenaLen := int(key[1])
		if len(key) <= len(keyPrefix)+1+arenaLen {
			continue
		}
		arena := key[2 : 2+arenaLen]

		denomLen := int(key[2+arenaLen])
		if len(key) <= len(keyPrefix)+1+arenaLen+1+denomLen {
			continue
		}
		denom := key[3+arenaLen : 3+arenaLen+denomLen]

		addrLen := int(key[3+arenaLen+denomLen])
		if len(key) < len(keyPrefix)+1+arenaLen+1+denomLen+1+addrLen {
			continue
		}
		addr := key[4+arenaLen+denomLen : 4+arenaLen+denomLen+addrLen]

		relayers = append(relayers, types.Relayer{
			Arena: string(arena),
			Denom: string(denom),
			Addrs: []string{string(addr)},
		})
	}
	return relayers
}
