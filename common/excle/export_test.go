/*
 * @Author: 尤_Ta
 * @Date:  22:07
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  22:07
 */

package excle

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestCreateExcelFile(t *testing.T) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for idx, row := range [][]interface{}{
		//{nil, "Apple", "Orange", "Pear"},
		{"Small", 2, 3, 3},
		{"Normal", 5, 2, 4},
		{"Large", 6, 7, 8},
	} {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.SetSheetRow("Sheet1", cell, &row)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("../../storage/public/excle/Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func TestExport(t *testing.T) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("../../storage/public/excle/Book2.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func TestExport1(t *testing.T) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for idx, row := range [][]interface{}{
		{nil, "Apple", "Orange", "Pear"}, {"Small", 2, 3, 3},
		{"Normal", 5, 2, 4}, {"Large", 6, 7, 8},
	} {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetSheetRow("Sheet1", cell, &row)
	}
	if err := f.AddChart("Sheet1", "E1", &excelize.Chart{
		Type: excelize.Col3DClustered,
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$A$2",
				Categories: "Sheet1!$B$1:$D$1",
				Values:     "Sheet1!$B$2:$D$2",
			},
			{
				Name:       "Sheet1!$A$3",
				Categories: "Sheet1!$B$1:$D$1",
				Values:     "Sheet1!$B$3:$D$3",
			},
			{
				Name:       "Sheet1!$A$4",
				Categories: "Sheet1!$B$1:$D$1",
				Values:     "Sheet1!$B$4:$D$4",
			}},
		Title: []excelize.RichTextRun{
			{
				Text: "Fruit 3D Clustered Column Chart",
			},
		},
	}); err != nil {
		fmt.Println(err)
		return
	}
	// 根据指定路径保存文件
	if err := f.SaveAs("../../storage/public/excle/Book3.xlsx"); err != nil {
		fmt.Println(err)
	}
}
