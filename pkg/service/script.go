package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	edgesTable = "edges"
	roomsTable = "rooms"
)

type DBScriptGeneratorService struct{}

func (d *DBScriptGeneratorService) GenerateRoomsScript(csvPath, scriptPath string) error {
	csvFile, err := os.Open(csvPath)
	defer csvFile.Close()

	if err != nil {
		return err
	}

	outFile, err := os.Create(scriptPath)
	defer outFile.Close()

	if err != nil {
		return err
	}

	reader := csv.NewReader(csvFile)

	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if len(line) != 3 {
			return errors.New("can't read csv file")
		}

		outFile.WriteString(fmt.Sprintf("INSERT INTO %s values (%s, '%s', %s);", roomsTable, line[0], line[1], line[2]))
	}

	return nil
}

func (d *DBScriptGeneratorService) GenerateEdgeScript(csvPath, scriptPath string) error {
	csvFile, err := os.Open(csvPath)
	defer csvFile.Close()

	if err != nil {
		return err
	}

	outFile, err := os.Create(scriptPath)
	defer outFile.Close()

	if err != nil {
		return err
	}

	reader := csv.NewReader(csvFile)

	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if len(line) != 3 {
			return errors.New("can't read csv file")
		}

		outFile.WriteString(fmt.Sprintf("INSERT INTO %s values (%s, %s, %s, %s);", edgesTable, line[0], line[1], line[2], line[3]))
	}

	return nil
}

func NewDBScriptGeneratorService() *DBScriptGeneratorService {
	return &DBScriptGeneratorService{}
}
