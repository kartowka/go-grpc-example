package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"

	chatpb "github.com/antfley/go-grpc-example/chat"
	"github.com/antfley/go-grpc-example/config"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
	logger *slog.Logger
}
type ChatServiceServer struct {
	chatpb.UnimplementedChatServiceServer
}

func NewServer(config *config.Config, logger *slog.Logger) *Server {
	return &Server{config: config, logger: logger}
}
func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.config.PORT))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	gRPCServer := grpc.NewServer()
	chatpb.RegisterChatServiceServer(gRPCServer, &ChatServiceServer{})
	go func() {
		s.logger.Info("server started", "port", s.config.PORT)
		if err := gRPCServer.Serve(listener); err != nil {
			s.logger.Error("server error", "error", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Wait for the context to be done (signaling shutdown)
		<-ctx.Done()
		// Gracefully stop the gRPC server
		gRPCServer.GracefulStop()

		// Optionally, you can log here after server stop
		s.logger.Info("server gracefully stopped")
	}()

	// Wait until the context is done (shut down)
	wg.Wait()
	return nil
}
func (s *ChatServiceServer) SendMessage(ctx context.Context, in *chatpb.Message) (*chatpb.Message, error) {
	return &chatpb.Message{Body: "Hello, " + in.Body}, nil
}
