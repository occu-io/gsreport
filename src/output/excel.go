package output

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/occu-io/gsreport/src/utils"
)

func createFile(filename string) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

}

func NewExcel() *utils.Excel {
	return &utils.Excel{File: excelize.NewFile()}
}
