package server

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-demo/core"
	"grpc-demo/proto"
	"grpc-demo/service"
	"grpc-demo/tools"
	"log"
	"net"
	"net/http"
)

var (
	Port        string
	CertName    string
	CertPemPath string
	CertKeyPath string
	SwaggerDir  string
)

func Server() (err error) {
	log.Printf("Port:%s\tCertName:%s\tCertPemPath:%s\tCertKeyPath:%s", Port, CertName, CertPemPath, CertKeyPath)
	address := ":" + Port
	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("TCP Listen err: %v\n", err)
		return err
	}
	tlsConfig := tools.GetTLSConfig(CertPemPath, CertKeyPath)
	server := createInternalServer(address, tlsConfig)
	log.Printf("gRPC and https listen on:%s\n", Port)
	newListener := tls.NewListener(conn, tlsConfig)
	if err = server.Serve(newListener); err != nil {
		log.Printf("Listen and Server err: %v\n", err)
		return err
	}
	return nil
}

/**
createInternalServer: 创建grpc服务
*/
func createInternalServer(address string, tlsConfig *tls.Config) *http.Server {
	var opts []grpc.ServerOption
	var auth core.Auth
	// grpc server
	cred, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials: %v\n", err)
	}
	opts = append(opts, grpc.Creds(cred))
	opts = append(opts, grpc.UnaryInterceptor(auth.CheckAuth))
	server := grpc.NewServer(opts...)
	proto.RegisterUserServiceServer(server, service.NewUserService())
	proto.RegisterGoodsServiceServer(server, service.NewGoodsServices())
	// gateway server
	ctx := context.Background()
	newCred, err := credentials.NewClientTLSFromFile(CertPemPath, CertName)
	if err != nil {
		log.Printf("Failed to create client TLS credentials: %v\n", err)
	}
	dialOpt := []grpc.DialOption{grpc.WithTransportCredentials(newCred)}
	gateWayMux := runtime.NewServeMux()

	// register grpc-gateway pb
	if err := proto.RegisterUserServiceHandlerFromEndpoint(ctx, gateWayMux, address, dialOpt); err != nil {
		log.Printf("Failed to register gateway server: %v\n", err)
	}
	if err := proto.RegisterGoodsServiceHandlerFromEndpoint(ctx, gateWayMux, address, dialOpt); err != nil {
		log.Printf("Failed to register gateway server: %v\n", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", gateWayMux)

	return &http.Server{
		Addr:      address,
		Handler:   tools.GrpcHandlerFunc(server, mux),
		TLSConfig: tlsConfig,
	}
}
