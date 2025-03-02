package keeper

import (
    "context"
    "fmt"

    "hello/x/hello/types"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (k Keeper) SayHello(ctx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    // Validation and Context unwrapping
    sdkCtx := sdk.UnwrapSDKContext(ctx)

    _ = sdkCtx
    // Custom Response
    return &types.QuerySayHelloResponse{Name: fmt.Sprintf("Hello %s!", req.Name)}, nil
}