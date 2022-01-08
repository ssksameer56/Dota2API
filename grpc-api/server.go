package grpcapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/ssksameer56/Dota2API/handlers"
	"github.com/ssksameer56/Dota2API/models"
	dota2grpc "github.com/ssksameer56/Dota2API/models/grpc"
	"github.com/ssksameer56/Dota2API/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func StartGrpcServer(config models.Configuration, dataHandler *handlers.Dota2Handler, matchHandler *handlers.MatchDataHandler,
	wg sync.WaitGroup) {
	dotaserver := GrpcServer{
		Dota2Handler:     dataHandler,
		MatchDataHandler: matchHandler,
	}
	port := "8081"
	if config.GrpcAPIPort != "" {
		port = config.GrpcAPIPort
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		utils.LogFatal(err.Error(), "GRPC Server")
		wg.Done()
		return
	}

	opts := []grpc.ServerOption{

		grpc.UnaryInterceptor(UnaryLogger),
		grpc.StreamInterceptor(StreamLogger),
	}
	secureMode, _ := strconv.ParseBool(config.SecureModeGRPC)
	if secureMode {
		cert, err := tls.LoadX509KeyPair("./cert.pem", "key.pem")
		if err != nil {
			utils.LogFatal(err.Error(), "GRPC Server")
			wg.Done()
			return
		}
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))

	}
	grpcServer := grpc.NewServer(opts...)
	dota2grpc.RegisterDota2ServiceServer(grpcServer, &dotaserver)
	utils.LogInfo(fmt.Sprintf("Starting gRPC Server localhost:%s", port), "GRPC Server")
	err = grpcServer.Serve(lis)
	if err != nil {
		utils.LogFatal(err.Error(), "GRPC Server")
		wg.Done()
		return
	}
}

func UnaryLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	utils.LogInfo(fmt.Sprintf("got request at:%s", info.FullMethod), "GRPC Server")
	m, err := handler(ctx, req)
	return m, err

}

func StreamLogger(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	utils.LogInfo(fmt.Sprintf("got request at:%s", info.FullMethod), "GRPC Server")
	err := handler(srv, ss)
	return err
}
