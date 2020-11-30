package gt_excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func Object(fileName string) (*excelize.File, error) {
	return excelize.OpenFile(fileName)
}

func Read_for_sheet_name(fileName string, sheetName string) ([][]string, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	return f.GetRows(sheetName)
}

func Get_sheet_list(fileName string) []string {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil
	}
	return f.GetSheetList()
}
