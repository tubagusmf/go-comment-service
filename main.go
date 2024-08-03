package main

import (
	"log"
	"net"

	grpcSvc "comment-service/internal/delivery/grpc"
	repository "comment-service/internal/repository"
	usecase "comment-service/internal/usecase"
	pb "comment-service/pb/comment"

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
