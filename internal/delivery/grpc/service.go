package grpc

import (
	"github.com/tubagusmf/go-comment-service/internal/model"
	pb "github.com/tubagusmf/go-comment-service/pb/comment"
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
