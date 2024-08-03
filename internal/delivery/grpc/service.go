package grpc

import (
	"comment-service/internal/model"
	pb "comment-service/pb/comment"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
	commentUseCase model.CommentUseCase
}

func NewCommentService(commentUseCase model.CommentUseCase) *CommentService {
	return &CommentService{
		commentUseCase: commentUseCase,
	}
}
