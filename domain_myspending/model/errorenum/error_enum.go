package errorenum

import "your/path/project/shared/model/apperror"

const (
	SomethingError           apperror.ErrorType = "ER0000 something error"
	NilaiIsRequired          apperror.ErrorType = "ER0001 nilai is required"
	NilaiMustGreatenThanZero apperror.ErrorType = "ER0002 nilai must greaten than zero"
	DeskripsiIsRequired      apperror.ErrorType = "ER0003 deskripsi is required"
)
