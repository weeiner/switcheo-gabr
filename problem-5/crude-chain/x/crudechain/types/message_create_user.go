package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUser{}

func NewMsgCreateUser(creator string, name string, email string, gender string, age uint64) *MsgCreateUser {
	return &MsgCreateUser{
		Creator: creator,
		Name:    name,
		Email:   email,
		Gender:  gender,
		Age:     age,
	}
}

func (msg *MsgCreateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
