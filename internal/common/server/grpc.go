package server

import (
	"net"

	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpctags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	grpclogrus.ReplaceGrpcLogger(logrus.NewEntry(logger))
}

func RunGRPCServer(serviceName string, registerServer func(server *grpc.Server)) {
	addr := viper.Sub(serviceName).GetString("grpc-addr")
	if addr == "" {
		// TODO: Warning log
		addr = viper.GetString("fallback.grpc-addr")
	}
	RunGRPCServerOnAddr(addr, registerServer)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpctags.UnaryServerInterceptor(grpctags.WithFieldExtractor(grpctags.CodeGenRequestFieldExtractor)),
			grpclogrus.UnaryServerInterceptor(logrusEntry),
			//srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			//logging.UnaryServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			//selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			//recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			grpctags.StreamServerInterceptor(grpctags.WithFieldExtractor(grpctags.CodeGenRequestFieldExtractor)),
			grpclogrus.StreamServerInterceptor(logrusEntry),
			//srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			//logging.StreamServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			//selector.StreamServerInterceptor(auth.StreamServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			//recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)
	registerServer(server)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Panic(err)
	}

	if err := server.Serve(listen); err != nil {
		logrus.Panic(err)
	}
}
