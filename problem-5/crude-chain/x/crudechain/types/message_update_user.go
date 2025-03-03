package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateUser{}

func NewMsgUpdateUser(creator string, id uint64, name string, email string, gender string, age uint64) *MsgUpdateUser {
	return &MsgUpdateUser{
		Creator: creator,
		Id:      id,
		Name:    name,
		Email:   email,
		Gender:  gender,
		Age:     age,
	}
}

func (msg *MsgUpdateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
