package grpc

import (
	"context"
	"errors"

	pb "comment-service/pb/comment"
)

func (s *CommentService) FindComments(ctx context.Context, in *pb.CommentRequest) (*pb.CommentList, error) {
	if in.ArticleId == "" {
		return nil, errors.New("article id is required")
	}

	comments, err := s.commentUseCase.FindComments(in.ArticleId, "postgres")
	if err != nil {
		return nil, err
	}

	var commentPbs []*pb.Comment

	for _, comment := range comments {
		commentPb := &pb.Comment{
			Id:        comment.Id,
			ArticleId: comment.ArticleId,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		commentPbs = append(commentPbs, commentPb)
	}

	return &pb.CommentList{Comments: commentPbs}, nil
}
