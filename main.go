package main

import (
	"log"
	"net"

	grpcSvc "github.com/tubagusmf/go-comment-service/internal/delivery/grpc"
	repository "github.com/tubagusmf/go-comment-service/internal/repository"
	usecase "github.com/tubagusmf/go-comment-service/internal/usecase"
	pb "github.com/tubagusmf/go-comment-service/pb/comment"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	commentRepo := repository.NewCommentRepository()
	commentUsecase := usecase.NewCommentUsecase(commentRepo)
	commentService := grpcSvc.NewCommentService(commentUsecase)

	pb.RegisterCommentServiceServer(s, commentService)

	listen, err := net.Listen("tcp", ":3100")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server is running on port 3100")

	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
