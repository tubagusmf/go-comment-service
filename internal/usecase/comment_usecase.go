package usecase

import "github.com/tubagusmf/go-comment-service/internal/model"

type commentUsecase struct {
	commentRepo model.CommentRepository
}

func NewCommentUsecase(cr model.CommentRepository) model.CommentUseCase {
	return &commentUsecase{
		commentRepo: cr,
	}
}

func (cu *commentUsecase) FindComments(articleId string, from string) ([]model.Comment, error) {

	if from == "elastic" {
		return cu.commentRepo.FindCommentFromDocuments(articleId)
	}

	return cu.commentRepo.FindComments(articleId)
}
