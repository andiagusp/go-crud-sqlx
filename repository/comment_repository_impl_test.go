package repository

import (
	"fmt"
	gc "go-crud"
	"go-crud/entity"
	"testing"
)

// var schema string = "CREATE TABLE comment(id int PRIMARY KEY AUTO_INCREMENT, email VARCHAR(128) NOT NULL, comment VARCHAR(128) NOT NULL) engine=InnoDB;"

var lastCreateId int32

func TestCommentCreate(t *testing.T) {
	fmt.Println("Test Create")
	conn := gc.GetConnection()
	commentRepository := NewCommentRepository(conn)
	defer conn.Close()
	comment := entity.Comment{
		Email:   "rusdi@gmail.com",
		Comment: "New Comment",
	}

	result, err := commentRepository.Create(comment)
	if err != nil {
		panic(err)
	}

	lastCreateId = result.Id
	fmt.Println(result)
	fmt.Println("Done")
}

func TestCommentFindAll(t *testing.T) {
	fmt.Println("Test Find All")
	conn := gc.GetConnection()
	commentRepository := NewCommentRepository(conn)
	defer conn.Close()

	result, err := commentRepository.FindAll()
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println("Id:", comment.Id)
		fmt.Println("Email:", comment.Email)
		fmt.Println("Comment:", comment.Comment)
	}
	fmt.Println("Done")
}

func TestCommentFindById(t *testing.T) {
	fmt.Println("Test Find By Id")
	conn := gc.GetConnection()
	commentRepository := NewCommentRepository(conn)
	defer conn.Close()

	result, err := commentRepository.FindById(lastCreateId)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println("Done")
}

func TestCommentUpdate(t *testing.T) {
	fmt.Println("Test Update")
	conn := gc.GetConnection()
	commentRepository := NewCommentRepository(conn)
	defer conn.Close()
	comment := entity.Comment{
		Id:      lastCreateId,
		Email:   "rusdi@gmail.com",
		Comment: "Update Comment",
	}

	result, message, err := commentRepository.Update(comment)
	if err != nil {
		fmt.Println(message)
		panic(err)
	}

	fmt.Println(message)
	fmt.Println(result)
	fmt.Println("Done")
}

func TestCommentDelete(t *testing.T) {
	fmt.Println("Test Delete")
	conn := gc.GetConnection()
	commentRepository := NewCommentRepository(conn)
	defer conn.Close()

	message, err := commentRepository.Delete(lastCreateId)
	if err != nil {
		fmt.Println(message)
		panic(err)
	}

	fmt.Println(message)
	fmt.Println("Done")
}
