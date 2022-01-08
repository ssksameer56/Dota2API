package grpcapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
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
	lis, err := net.Listen("tcp", ":"+config.GrpcAPIPort)
	if err != nil {
		utils.LogFatal(err.Error(), "GRPC Server")
		wg.Done()
		return
	}
	cert, err := tls.LoadX509KeyPair("./cert.pem", "key.pem")
	if err != nil {
		utils.LogFatal(err.Error(), "GRPC Server")
		wg.Done()
		return
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(UnaryLogger),
		grpc.StreamInterceptor(StreamLogger),
	}

	grpcServer := grpc.NewServer(opts...)
	dota2grpc.RegisterDota2ServiceServer(grpcServer, &dotaserver)
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
