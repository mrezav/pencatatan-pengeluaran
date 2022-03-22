package entity

import (
	"time"
	"your/path/project/domain_myspending/model/errorenum"
)

type Expense struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nilai     int    `json:"nilai"`
	Deskripsi string `json:"deskripsi"`
	Tanggal   string `json:"tanggal"`
}

type ExpenseRequest struct {
	Nilai              int
	Deskripsi, Tanggal string
}

func NewExpense(req ExpenseRequest) (*Expense, error) {

	// assign value here

	if req.Nilai <= 0 {
		return nil, errorenum.NilaiMustGreatenThanZero
	}

	if req.Deskripsi == "" {
		return nil, errorenum.DeskripsiIsRequired
	}

	var obj Expense
	obj.Nilai = req.Nilai
	obj.Deskripsi = req.Deskripsi
	obj.Tanggal = time.Now().String()

	return &obj, nil
}

func (r *Expense) Validate() error {
	return nil
}
