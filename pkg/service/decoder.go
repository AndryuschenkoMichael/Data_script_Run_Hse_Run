package service

import (
	"Data_script_Run_Hse_Run/model"
	"encoding/csv"
	"os"
	"strconv"
)

type CSVDecoderService struct {}

func (C *CSVDecoderService) CreateEdgesCSVFile(path string, edges []model.Edge) error {
	csvFile, err := os.Create(path)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(csvFile)
	_ = csvWriter.Write([]string{"id", "start_room_id", "end_room_id", "cost", "campus"})
	for i, edge := range edges {
		_ = csvWriter.Write([]string{
			strconv.Itoa(i + 1),
			strconv.Itoa(edge.StartVertex),
			strconv.Itoa(edge.EndVertex),
			strconv.FormatFloat(edge.Cost, 'f', 4, 64),
			strconv.Itoa(edge.CampusId),
		})
	}

	csvWriter.Flush()
	_ = csvFile.Close()

	return nil
}

func NewCSVDecoderService() *CSVDecoderService {
	return &CSVDecoderService{}
}
