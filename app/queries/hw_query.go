package queries

import (
	"github.com/jmoiron/sqlx"
	"go.fiber.sqlx/app/models"
)

type HwQuery struct {
	*sqlx.DB
}

func (q *HwQuery) GetHws() ([]models.HwDevice, error) {

	hws := []models.HwDevice{}

	query := `select * from tbl_hwdevice`

	err := q.Select(&hws, query)
	if err != nil {
		return hws, err
	}

	return hws, err
}

func (q *HwQuery) GetHw(id string) (models.HwDevice, error) {

	hw := models.HwDevice{}

	query := `select * from tbl_hwdevice where id=$1`

	err := q.Get(&hw, query, id)
	if err != nil {
		return hw, err
	}

	return hw, err
}
