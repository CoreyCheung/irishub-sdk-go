package keys

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace

var (
	ErrInsert  = sdkerrors.Register(Codespace, 2, "key-manager insert error")
	ErrRecover = sdkerrors.Register(Codespace, 2, "key-manager recover error")
	ErrImport  = sdkerrors.Register(Codespace, 2, "key-manager import error")
	ErrExport  = sdkerrors.Register(Codespace, 2, "key-manager export error")
	ErrDelete  = sdkerrors.Register(Codespace, 2, "key-manager delete error")
	ErrShow    = sdkerrors.Register(Codespace, 2, "key-manager show error")
)
