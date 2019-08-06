package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"taskulu/api/grpc"
	api "taskulu/api/proto/src"
	"taskulu/testkit"
	"testing"
	"time"
)

func TestExampleGrpc(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rsp, err := testkit.GetGrpcClient().RegisterExample(ctx, &api.ExampleRequest{AuthId: int64(10), UserId: int64(10)})
	assert.NotNil(t, err)
	assert.Nil(t, rsp)
	assert.Equal(t, err, grpc.Errors.Internal)

	rsp, err = testkit.GetGrpcClient().RegisterExample(ctx, &api.ExampleRequest{AuthId: int64(20), UserId: int64(20)})
	assert.NotNil(t, rsp)
	assert.Nil(t, err)

	assert.IsType(t, rsp, &api.ResponseVoid{})

}
