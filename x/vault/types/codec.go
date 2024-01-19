package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateTokenAdmin{}, "vault/CreateTokenAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateTokenAdmin{}, "vault/UpdateTokenAdmin", nil)
	cdc.RegisterConcrete(&MsgCreateToken{}, "vault/CreateToken", nil)
	cdc.RegisterConcrete(&MsgUpdateToken{}, "vault/UpdateToken", nil)
	cdc.RegisterConcrete(&MsgAuditToken{}, "vault/AuditToken", nil)
	cdc.RegisterConcrete(&MsgConvertTokenToIns{}, "vault/ConvertTokenToIns", nil)
	cdc.RegisterConcrete(&MsgConvertInsToToken{}, "vault/ConvertInsToToken", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTokenAdmin{},
		&MsgUpdateTokenAdmin{},
		&MsgCreateToken{},
		&MsgUpdateToken{},
		&MsgAuditToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConvertTokenToIns{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConvertInsToToken{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
