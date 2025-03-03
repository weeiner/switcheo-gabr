package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crude-chain/x/crudechain/types"
)

func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Fetch the user
	val, found := k.GetUser(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("user ID %d doesn't exist", msg.Id))
	}

	// Ensure only the creator can delete their user
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Delete user from store
	k.RemoveUser(ctx, msg.Id)

	return &types.MsgDeleteUserResponse{}, nil
}
