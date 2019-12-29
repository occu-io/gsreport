package utils

import (
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gopkg.in/ini.v1"
)

type Excel struct {
	File  *excelize.File
	Sheet string
}

type Report struct {
	Cl  *http.Client
	Cfg *ini.File
	Xls *Excel
}
