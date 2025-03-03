package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"crude-chain/x/crudechain/types"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var user = types.User{
		Name:    msg.Name,
		Email:   msg.Email,
		Age:     msg.Age,
		Gender:  msg.Gender,
		Creator: msg.Creator,
	}

	id := k.AppendUser(ctx, user)

	return &types.MsgCreateUserResponse{
		Id: id,
	}, nil
}
