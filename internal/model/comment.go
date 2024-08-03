package model

type Comment struct {
	Id        string `json:"id"`
	ArticleId string `json:"article_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentUseCase interface {
	FindComments(articleId string, from string) ([]Comment, error)
}

type CommentRepository interface {
	FindComments(articleId string) ([]Comment, error)
	FindCommentFromDocuments(articleId string) ([]Comment, error)
}
