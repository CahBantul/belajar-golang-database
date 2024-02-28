package repository

import (
	"context"
	"fmt"
	"testing"

	belajar "github.com/CahBantul/belajar-golang-database"
	"github.com/CahBantul/belajar-golang-database/entity"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "email@gmasil.com",
		Comment: "halo",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
func TestFindByIdSuccess(t *testing.T) {
	commentRepository := NewCommentRepository(belajar.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 40)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindByIdNotFound(t *testing.T) {
	commentRepository := NewCommentRepository(belajar.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 90)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for Index, commentValue := range comments {
		fmt.Println(Index, " : ", commentValue)
	}
}
