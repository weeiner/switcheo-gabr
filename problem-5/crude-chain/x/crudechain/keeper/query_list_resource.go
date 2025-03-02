package keeper

import (
	"context"

	"crude-chain/x/crudechain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListResource(goCtx context.Context, req *types.QueryListResourceRequest) (*types.QueryListResourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ResourceKey))

	var resources []types.Resource // ✅ Change to `[]types.Resource`

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var resource types.Resource
		if err := k.cdc.Unmarshal(value, &resource); err != nil {
			return err
		}

		resources = append(resources, resource) // ✅ Append as value (not pointer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// ✅ Correct: Returning slice of `Resource` (not pointers)
	return &types.QueryListResourceResponse{Resource: resources, Pagination: pageRes}, nil
}
