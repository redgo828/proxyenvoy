package main

import (
	"context"
	"net"
	"time"

	pb "backend/proto"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

var jwtSecret = []byte("secret123")

type server struct {
	pb.UnimplementedAuthServiceServer
	pb.UnimplementedHealthServiceServer
}

func issueJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "grpc-backend",
		"aud": "frontend",
		"sub": "user123",
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})
	s, _ := token.SignedString(jwtSecret)
	return s
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	jwt := issueJWT()

	md := metadata.Pairs(
		"set-cookie",
		"refresh_token22=ref123; HttpOnly; Path=/; SameSite=None",
	)
	grpc.SendHeader(ctx, md)

	return &pb.AuthResponse{AccessToken: jwt}, nil
}

func (s *server) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.AuthResponse, error) {
	jwt := issueJWT()

	md := metadata.Pairs(
		"set-cookie",
		"refresh_token22=ref456; HttpOnly; Path=/; SameSite=None",
	)
	grpc.SendHeader(ctx, md)

	return &pb.AuthResponse{AccessToken: jwt}, nil
}

func (s *server) Check(ctx context.Context, e *pb.Empty) (*pb.HealthReply, error) {
	return &pb.HealthReply{Status: "ok"}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	g := grpc.NewServer()

	pb.RegisterAuthServiceServer(g, &server{})
	pb.RegisterHealthServiceServer(g, &server{})
	reflection.Register(g)
	g.Serve(lis)
}
