package repository

import "go-crud/entity"

type CommentRepository interface {
	FindAll() ([]entity.Comment, error)
	FindById(id int32) (entity.Comment, error)
	Create(comment entity.Comment) (entity.Comment, error)
	Update(comment entity.Comment) (entity.Comment, string, error)
	Delete(id int32) (string, error)
}
