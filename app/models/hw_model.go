package models

type HwDevice struct {
	ID           int    `db:"id"`
	ParentID     int    `db:"parent_id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	HardwareID   int    `db:"hardware_id"`
	EnumeratorID int    `db:"enumerator_id"`
}
