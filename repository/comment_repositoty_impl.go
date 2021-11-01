package repository

import (
	"go-crud/entity"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type commentRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (commentRepository *commentRepositoryImpl) FindAll() ([]entity.Comment, error) {
	var comments []entity.Comment
	db := commentRepository.DB
	script := "SELECT id, email, comment FROM comment"
	err := db.Select(&comments, script)

	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (commentRepository *commentRepositoryImpl) FindById(id int32) (entity.Comment, error) {
	var comment entity.Comment
	db := commentRepository.DB
	script := "SELECT id, email, comment FROM comment WHERE id = ?"
	err := db.Get(&comment, script, id)

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (commentRepository *commentRepositoryImpl) Create(comment entity.Comment) (entity.Comment, error) {
	db := commentRepository.DB
	tx := db.MustBegin()
	script := "INSERT INTO comment(email, comment) VALUES(?, ?)"

	result := db.MustExec(script, comment.Email, comment.Comment)
	tx.Commit()

	value, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(value)
	return comment, nil
}

func (commentRepository *commentRepositoryImpl) Update(comment entity.Comment) (entity.Comment, string, error) {
	db := commentRepository.DB
	tx := db.MustBegin()
	script := "UPDATE comment SET email=?, comment=? WHERE id=?"

	result := db.MustExec(script, comment.Email, comment.Comment, comment.Id)
	tx.Commit()

	value, err := result.RowsAffected()
	if err != nil || value == 0 {
		return comment, "Update Failed", err
	}

	return comment, "Update Success", err
}

func (commentRepository *commentRepositoryImpl) Delete(id int32) (string, error) {
	db := commentRepository.DB
	tx := db.MustBegin()
	script := "DELETE FROM comment WHERE id=?"

	result := db.MustExec(script, id)
	tx.Commit()

	value, err := result.RowsAffected()
	if err != nil || value == 0 {
		return "Delete Failed Id: " + strconv.Itoa(int(id)), err
	}

	return "Delete Success Id: " + strconv.Itoa(int(id)), nil
}
