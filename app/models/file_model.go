package models

import "time"

type FileInfo struct {
	Id           int       `db:"id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	Date         time.Time `db:"date"`
	Size         int       `db:"size"`
	Md5          string    `db:"md5"`
	CatalogID    int       `db:"catalog_id"`
	VendorID     int       `db:"vendor_id"`
	VerMajor     int       `db:"ver_major"`
	VerMinor     int       `db:"ver_minor"`
	VerRelease   int       `db:"ver_release"`
	VerBuild     int       `db:"ver_build"`
	VerRevision  string    `db:"ver_revision"`
	UploadUserID int       `db:"upload_user_id"`
	UploadDate   time.Time `db:"upload_date"`
}
