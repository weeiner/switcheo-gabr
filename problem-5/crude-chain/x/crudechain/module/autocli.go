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
				{
					RpcMethod:      "ShowUser",
					Use:            "show-user [id]",
					Short:          "Query show-user",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListUser",
					Use:            "list-user",
					Short:          "Query list-user",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod:      "CreateUser",
					Use:            "create-user [name] [email] [gender] [age]",
					Short:          "Send a create-user tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "email"}, {ProtoField: "gender"}, {ProtoField: "age"}},
				},
				{
					RpcMethod:      "UpdateUser",
					Use:            "update-user [id] [name] [email] [gender] [age]",
					Short:          "Send a update-user tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "email"}, {ProtoField: "gender"}, {ProtoField: "age"}},
				},
				{
					RpcMethod:      "DeleteUser",
					Use:            "delete-user [id]",
					Short:          "Send a delete-user tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
