package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"crude-chain/x/crudechain/types"
)

// AppendUser stores a new user and returns the assigned ID
func (k Keeper) AppendUser(ctx sdk.Context, user types.User) uint64 {
	count := k.GetUserCount(ctx) // Get current user count
	user.Id = count              // ✅ Always set the `id` field

	fmt.Println("Storing User:", user) // Debug log

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))

	appendedValue := k.cdc.MustMarshal(&user)
	store.Set(GetUserIDBytes(user.Id), appendedValue)

	k.SetUserCount(ctx, count+1) // Increment and store the new user count
	return count
}

// GetUserCount retrieves the total number of users
func (k Keeper) GetUserCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.UserCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// Convert User ID to byte array for storage
func GetUserIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// SetUserCount updates the total user count
func (k Keeper) SetUserCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.UserCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetUser fetches a user by ID
func (k Keeper) GetUser(ctx sdk.Context, id uint64) (val types.User, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))
	b := store.Get(GetUserIDBytes(id))

	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
