package keeper

import (
	"context"
	"strings"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"crude-chain/x/crudechain/types"
)

func (k Keeper) UsersByGender(goCtx context.Context, req *types.QueryUsersByGenderRequest) (*types.QueryUsersByGenderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKey))

	var users []*types.User

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var user types.User
		if err := k.cdc.Unmarshal(iterator.Value(), &user); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// Case insensitive match for gender
		if strings.EqualFold(user.Gender, req.Gender) {
			users = append(users, &user)
		}
	}

	return &types.QueryUsersByGenderResponse{User: users}, nil
}
