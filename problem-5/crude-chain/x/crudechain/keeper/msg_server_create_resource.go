package keeper

import (
	"context"

	"crude-chain/x/crudechain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateResource(goCtx context.Context, msg *types.MsgCreateResource) (*types.MsgCreateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var post = types.Resource{
		Title:       msg.Title,
		Description: msg.Description,
		Creator:     msg.Creator,
	}
	id := k.AppendResource(
		ctx,
		post,
	)
	return &types.MsgCreateResourceResponse{
		Id: id,
	}, nil
}
