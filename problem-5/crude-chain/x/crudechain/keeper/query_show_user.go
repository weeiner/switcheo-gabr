package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"crude-chain/x/crudechain/types"
)

func (k Keeper) ShowUser(goCtx context.Context, req *types.QueryShowUserRequest) (*types.QueryShowUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	user, found := k.GetUser(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &types.QueryShowUserResponse{User: user}, nil
}
