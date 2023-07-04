package repository

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const fileName = "app_user.csv"

// CSVRepository contains the methods to handle a user stored in a CSV file.
type CSVRepository struct {
	f *os.File
	r *csv.Reader
	w *csv.Writer
}

// NewCSVRepository creates a new *CSVRepository.
func NewCSVRepository() (*CSVRepository, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	w := csv.NewWriter(f)

	records, err := r.ReadAll()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to read file: %s", err))
	}

	repo := &CSVRepository{
		f: f,
		r: r,
		w: w,
	}

	if records == nil {
		err = repo.seed()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to populate file: %s", err))
		}
	}

	return repo, nil
}

// CreateUser executes the SQL statement to create a user.
func (cr *CSVRepository) CreateUser(ctx context.Context, name string, age int) error {
	_, err := cr.f.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to seek the start of the file: %s", err))
	}

	records, err := cr.r.ReadAll()
	if err != nil {
		return err
	}

	lastRecord := records[len(records)-1]
	lastRecordID, err := strconv.Atoi(lastRecord[0])
	if err != nil {
		return errors.New(fmt.Sprintf("failed to obtain ID from record: %s", err))
	}

	t := time.Now()

	newRecord := []string{
		strconv.Itoa(lastRecordID + 1),
		name, strconv.Itoa(age),
		t.Format(time.DateTime),
		t.Format(time.DateTime),
	}

	err = cr.w.Write(newRecord)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to write new record: %s", err))
	}

	cr.w.Flush()
	err = cr.w.Error()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to flush data: %s", err))
	}

	return nil
}

// DeleteUserByName deletes the first user that matches a specific name.
func (cr *CSVRepository) DeleteUserByName(ctx context.Context, name string) error {
	_, err := cr.f.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to seek the start of the file: %s", err))
	}

	fileData, err := cr.r.ReadAll()
	if err != nil {
		return err
	}
	headers := fileData[0]
	records := fileData[1:]

	for i, header := range headers {
		if header == "name" {
			for j, record := range records {
				if record[i] == name {
					records = append(records[:j], records[j+1:]...)
					fileData = append([][]string{headers}, records...)
					err = cr.writeData(fileData)
					if err != nil {
						return errors.New(fmt.Sprintf("failed to write data: %s", err))
					}
					break
				}
			}
			break
		}
	}

	cr.w.Flush()
	err = cr.w.Error()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to flush data: %s", err))
	}

	return nil
}

// GetUserByName executes the SQL statement to retrieve a user's data.
func (cr *CSVRepository) GetUserByName(ctx context.Context, name string) (*User, error) {
	_, err := cr.f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to seek the start of the file: %s", err))
	}

	fileData, err := cr.r.ReadAll()
	if err != nil {
		return nil, err
	}
	headers := fileData[0]
	records := fileData[1:]

	for i, header := range headers {
		if header == "Name" {
			for _, record := range records {
				if record[i] == name {
					return createUserFromRecord(record, headers)
				}
			}
		}
	}
	return nil, nil
}

// UpdateUser executes the SQL statement to update a user.
func (cr *CSVRepository) UpdateUser(ctx context.Context, name string, attributes map[string]interface{}) error {
	_, err := cr.f.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to seek the start of the file: %s", err))
	}

	fileData, err := cr.r.ReadAll()
	if err != nil {
		return err
	}

	headers := fileData[0]
	records := fileData[1:]

	for i, record := range records {
		if record[1] == name {
			for j, header := range headers {
				for key, value := range attributes {
					if key == strings.ToLower(header) {
						records[i][j] = fmt.Sprintf("%v", value)
					}
				}
			}
			fileData = append([][]string{headers}, records...)
			err = cr.writeData(fileData)
			if err != nil {
				return errors.New(fmt.Sprintf("failed to write data: %s", err))
			}
			break
		}
	}

	cr.w.Flush()
	err = cr.w.Error()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to flush data: %s", err))
	}

	return nil

}

func (cr *CSVRepository) seed() error {
	records := [][]string{
		{"ID", "Name", "Age", "CreatedAt", "UpdatedAt"},
		{"1", "Tom", "24", "2022-02-22 15:16:19", "2022-02-22 15:16:19"},
		{"2", "Lucy", "23", "2022-02-22 15:16:19", "2022-02-22 15:16:19"},
		{"3", "Jim", "33", "2022-02-22 15:16:20", "2022-02-22 15:16:20"},
		{"4", "Ben", "40", "2022-05-27 16:39:39", "2022-05-27 16:39:39"},
	}
	err := cr.w.WriteAll(records)
	if err != nil {
		return fmt.Errorf("failed to initialized CSV file: %+v", err)
	}
	return nil
}

func (cr *CSVRepository) writeData(fileData [][]string) error {
	tempFilename := fmt.Sprintf("%s_TEMP", fileName)
	f, err := os.OpenFile(tempFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	w := csv.NewWriter(f)

	err = w.WriteAll(fileData)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to write data: %s", err))
	}

	err = os.Remove(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to delete file: %s", err))
	}

	err = os.Rename(tempFilename, fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to rename file: %s", err))
	}

	cr.f = f
	cr.r = r
	cr.w = w

	return nil
}

func createUserFromRecord(record []string, headers []string) (*User, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to convert ID to int: %s", err))
	}

	age, err := strconv.Atoi(record[2])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to convert age to int: %s", err))
	}

	return &User{
		ID:   id,
		Name: record[1],
		Age:  age,
	}, nil
}
