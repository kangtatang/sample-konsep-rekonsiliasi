package utils

import (
	"fmt"
	"go-big-internal/models"
	"time"

	"github.com/xuri/excelize/v2"
)

func ExportToExcel(data []models.ReconciliationResult) (*excelize.File, error) {
	f := excelize.NewFile()
	sheet := "Reconciliation"
	index, err := f.NewSheet(sheet)
	if err != nil {
		return nil, fmt.Errorf("failed to create new sheet: %v", err)
	}

	// Header
	f.SetCellValue(sheet, "A1", "Transaction ID")
	f.SetCellValue(sheet, "B1", "Status")
	f.SetCellValue(sheet, "C1", "Internal Amount")
	f.SetCellValue(sheet, "D1", "External Amount")
	f.SetCellValue(sheet, "E1", "Reconciled At")

	for i, row := range data {
		f.SetCellValue(sheet, "A"+toStr(i+2), row.TransactionID)
		f.SetCellValue(sheet, "B"+toStr(i+2), row.Status)
		f.SetCellValue(sheet, "C"+toStr(i+2), row.InternalAmount)
		f.SetCellValue(sheet, "D"+toStr(i+2), row.ExternalAmount)
		f.SetCellValue(sheet, "E"+toStr(i+2), row.ReconciledAt.Format(time.RFC3339))
	}

	f.SetActiveSheet(index)
	return f, nil
}

func toStr(n int) string {
	return fmt.Sprintf("%d", n)
}
