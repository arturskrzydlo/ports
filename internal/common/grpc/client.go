package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewClientConnectionContext(ctx context.Context, url string, keepaliveInSeconds int) (*grpc.ClientConn, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Duration(keepaliveInSeconds) * time.Second,
			PermitWithoutStream: true,
		}),
	}

	conn, err := grpc.DialContext(ctx, url, dialOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to grpc endpoint: %w", err)
	}
	return conn, nil
}
