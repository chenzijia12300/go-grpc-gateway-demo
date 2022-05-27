package core

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-demo/tools"
	"log"
)

type Auth struct {
}

func (auth *Auth) CheckAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("method: %s", info.FullMethod)
	if info.FullMethod == "/proto.UserService/Login" || info.FullMethod == "/proto.UserService/Register" {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("token error")
	}
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return nil, errors.New("token not found")
	}
	customClaims, err := tools.ParseToken(token[0])
	if err != nil {
		return nil, err
	}
	log.Printf("id:[%d]\tusername:[%s]", customClaims.ID, customClaims.Username)
	resp, err = handler(ctx, req)
	return resp, err
}
