package repository

import "comment-service/internal/model"

type commentRepository struct {
}

func NewCommentRepository() model.CommentRepository {
	return &commentRepository{}
}

func (cr *commentRepository) FindComments(articleId string) ([]model.Comment, error) {
	var comments []model.Comment

	// select * from comments where article_id = '3'
	if articleId != "3" {
		return comments, nil
	}

	comments = getDataFromDB()
	return comments, nil
}

func (cr *commentRepository) FindCommentFromDocuments(articleId string) ([]model.Comment, error) {
	var comments []model.Comment

	// get from no sql, elastic search
	if articleId != "3" {
		return comments, nil
	}

	comments = getDataFromDB()
	return comments, nil
}

func getDataFromDB() []model.Comment {
	return []model.Comment{
		{
			Id:        "1",
			ArticleId: "1",
			Content:   "comment 1",
			CreatedAt: "2021-01-01T00:00:00Z",
			UpdatedAt: "2021-01-01T00:00:00Z",
		},
		{
			Id:        "2",
			ArticleId: "1",
			Content:   "comment 2",
			CreatedAt: "2021-01-01T00:00:00Z",
			UpdatedAt: "2021-01-01T00:00:00Z",
		},
	}
}
