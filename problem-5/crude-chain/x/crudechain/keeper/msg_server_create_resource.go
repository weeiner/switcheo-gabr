package keeper

import (
	"context"

	"crude-chain/x/crudechain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateResource(goCtx context.Context, msg *types.MsgCreateResource) (*types.MsgCreateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var resource = types.Resource{
		Creator:     msg.Creator,
		Name:        msg.Name,
		Description: msg.Description,
	}

	id := k.AppendResource(ctx, resource)

	return &types.MsgCreateResourceResponse{
		Id: id,
	}, nil
}
