package queries

import (
	"github.com/jmoiron/sqlx"
	"go.fiber.sqlx/app/models"
)

type FileQuery struct {
	*sqlx.DB
}

func (q *FileQuery) GetFiles() ([]models.FileInfo, error) {

	files := []models.FileInfo{}

	query := `select * from tbl_fileinfo`

	err := q.Select(&files, query)
	if err != nil {
		return files, err
	}

	return files, err
}

func (q *FileQuery) GetFile(id string) (models.FileInfo, error) {

	file := models.FileInfo{}

	query := `select * from tbl_fileinfo where id=$1`

	err := q.Get(&file, query, id)
	if err != nil {
		return file, err
	}

	return file, err
}
