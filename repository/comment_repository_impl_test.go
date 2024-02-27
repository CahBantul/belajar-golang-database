package repository

import (
	"testing"

	belajar_golang_database "github.com/CahBantul/belajar-golang-database"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
}
