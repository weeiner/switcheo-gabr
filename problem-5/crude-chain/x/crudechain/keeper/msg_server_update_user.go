package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crude-chain/x/crudechain/types"
)

func (k msgServer) UpdateUser(goCtx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Fetch the existing user
	val, found := k.GetUser(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("user ID %d doesn't exist", msg.Id))
	}

	// Ensure only the original creator can update the user data
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Update only the fields that were provided
	val.Name = msg.Name
	val.Email = msg.Email
	val.Age = msg.Age
	val.Gender = msg.Gender

	// Save the updated user back to the store
	k.SetUser(ctx, val)

	return &types.MsgUpdateUserResponse{}, nil
}
