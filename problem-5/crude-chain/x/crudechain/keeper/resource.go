package keeper

import (
    "encoding/binary"

    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    sdk "github.com/cosmos/cosmos-sdk/types"

    "crude-chain/x/crudechain/types"
)

// AppendResource adds a new resource to the store and returns its unique ID.
func (k Keeper) AppendResource(ctx sdk.Context, resource types.Resource) uint64 {
    count := k.GetResourceCount(ctx) // Get the current count
    resource.Id = count              // Assign the new resource an ID

    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))

    appendedValue := k.cdc.MustMarshal(&resource)
    store.Set(GetResourceIDBytes(resource.Id), appendedValue) // Store the resource

    k.SetResourceCount(ctx, count+1) // Increment the count
    return count
}

// GetResourceCount retrieves the current number of stored resources.
func (k Keeper) GetResourceCount(ctx sdk.Context) uint64 {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, []byte{})
    byteKey := types.KeyPrefix(types.ResourceCountKey)

    bz := store.Get(byteKey)
    if bz == nil {
        return 0
    }
    return binary.BigEndian.Uint64(bz)
}

// SetResourceCount updates the resource count.
func (k Keeper) SetResourceCount(ctx sdk.Context, count uint64) {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, []byte{})
    byteKey := types.KeyPrefix(types.ResourceCountKey)

    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, count)
    store.Set(byteKey, bz)
}

// GetResource retrieves a specific resource from the store by ID.
func (k Keeper) GetResource(ctx sdk.Context, id uint64) (val types.Resource, found bool) {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))

    b := store.Get(GetResourceIDBytes(id))
    if b == nil {
        return val, false
    }
    k.cdc.MustUnmarshal(b, &val)
    return val, true
}
