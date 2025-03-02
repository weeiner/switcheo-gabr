package crudechain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "crude-chain/api/crudechain/crudechain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateResource",
					Use:            "create-resource [title] [description]",
					Short:          "Send a create-resource tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "UpdateResource",
					Use:            "update-resource [title] [description] [id]",
					Short:          "Send a update-resource tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteResource",
					Use:            "delete-resource [id]",
					Short:          "Send a delete-resource tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
