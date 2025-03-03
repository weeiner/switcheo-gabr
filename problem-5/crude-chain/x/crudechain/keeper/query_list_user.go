package keeper

import (
	"context"
	"encoding/binary"
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"crude-chain/x/crudechain/types"
)

func (k Keeper) ListUser(ctx context.Context, req *types.QueryListUserRequest) (*types.QueryListUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))

	var users []types.User

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return err
		}

		// ✅ Ensure `id` is explicitly set even if `0`
		if user.Id == 0 {
			user.Id = binary.BigEndian.Uint64(key) // Force setting ID
		}

		fmt.Println("Retrieved User:", user) // Debugging log

		users = append(users, user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListUserResponse{User: users, Pagination: pageRes}, nil
}
