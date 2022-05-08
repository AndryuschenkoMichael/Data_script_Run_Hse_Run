package service

import "Data_script_Run_Hse_Run/model"

type CSVEncoder interface {
	GetEdgesFromCSV(path string) ([]model.Edge, error)
	GetVerticesFromCSV(path string) (model.Vertices, error)
	GetRoomsFromCSV(path string) (model.Rooms, error)
}

type CSVDecoder interface {
	CreateEdgesCSVFile(path string, edges []model.Edge) error
}

type DataHandler interface {
	ProcessData(edges []model.Edge, vertices model.Vertices) []model.Edge
}

type DBScriptGenerator interface {
	GenerateRoomsScript(csvPath, scriptPath string) error
	GenerateEdgeScript(csvPath, scriptPath string) error
}

type Service struct {
	CSVEncoder
	DataHandler
	CSVDecoder
	DBScriptGenerator
}

func NewService() *Service {
	return &Service{
		CSVEncoder:        NewCSVEncoderService(),
		DataHandler:       NewDataHandlerServices(),
		CSVDecoder:        NewCSVDecoderService(),
		DBScriptGenerator: NewDBScriptGeneratorService(),
	}
}
