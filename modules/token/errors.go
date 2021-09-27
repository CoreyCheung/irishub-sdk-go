package token

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrQueryAddress         = sdkerrors.Register(Codespace, 2, "query address error")
	ErrGenConn              = sdkerrors.Register(Codespace, 2, "generate conn error")
	ErrQueryParams          = sdkerrors.Register(Codespace, 2, "query params error")
	ErrValidateAccAddress  =  sdkerrors.Register(Codespace, 2, "ValidateAccAddress error ")
)
