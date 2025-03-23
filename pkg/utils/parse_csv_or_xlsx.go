package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"server/internal/core/models"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ParseCsvOrXlsxResult struct {
	User     []models.BulkImportUser
	Failures []models.BulkImportFailure
}

func ParseCsv(file io.Reader) (*ParseCsvOrXlsxResult, error) {
	result := &ParseCsvOrXlsxResult{}
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.TrimLeadingSpace = true

	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	rowNumber := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		rowNumber++

		user, err := ParseRecord(record)
		if err != nil {
			result.Failures = append(result.Failures, models.BulkImportFailure{
				Name:            strings.TrimSpace(record[0]),
				Nickname:        strings.TrimSpace(record[1]),
				Email:           strings.TrimSpace(record[2]),
				SiteUserLevelId: strings.TrimSpace(record[3]),
				Message:         err.Error(),
			})
			continue
		}

		result.User = append(result.User, user)
		rowNumber++
	}

	return result, nil
}

func ParseXlsx(file io.Reader) (*ParseCsvOrXlsxResult, error) {
	result := &ParseCsvOrXlsxResult{}

	tempFile, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	f, err := excelize.OpenFile(tempFile.Name())
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		record := make([]string, len(row))
		copy(record, row)

		user, err := ParseRecord(record)
		if err != nil {
			result.Failures = append(result.Failures, models.BulkImportFailure{
				Name:            strings.TrimSpace(record[0]),
				Nickname:        strings.TrimSpace(record[1]),
				Email:           strings.TrimSpace(record[2]),
				SiteUserLevelId: strings.TrimSpace(record[3]),
				Message:         err.Error(),
			})
			continue
		}

		result.User = append(result.User, user)
	}

	return result, nil
}

func ParseRecord(record []string) (models.BulkImportUser, error) {
	requiredField := map[string]string{
		"Name":            record[0],
		"Nickname":        record[1],
		"Email":           record[2],
		"SiteUserLevelId": record[3],
	}

	for fieldName, fieldValue := range requiredField {
		if fieldValue == "" {
			return models.BulkImportUser{}, fmt.Errorf("field %s is required", fieldName)
		}
	}

	user := models.BulkImportUser{
		Name:            record[0],
		Nickname:        record[1],
		Email:           record[2],
		SiteUserLevelId: record[3],
	}

	return user, nil
}
