package main

import (
	"Data_script_Run_Hse_Run/pkg/service"
	"log"
)

const (
	edgesModifiedPath    = "modify_data/edges_modified.csv"
	roomsModifiedPath    = "hse_run_data/rooms_modified.csv"
	verticesModifiedPath = "modify_data/vertices_modified.csv"
	edgesPath            = "hse_run_data/edges_modified.csv"
	dbScriptEdgesPath    = "db_scripts/edges.sql"
	dbScriptRoomsPath    = "db_scripts/rooms.sql"
)

func CreateData(services *service.Service) {
	edges, err := services.GetEdgesFromCSV(edgesModifiedPath)
	if err != nil {
		log.Fatalf("can't read csv file: %s: %s", edgesModifiedPath, err.Error())
	}
	vertices, err := services.GetVerticesFromCSV(verticesModifiedPath)
	if err != nil {
		log.Fatalf("can't read csv file: %s: %s", verticesModifiedPath, err.Error())
	}
	_, err = services.GetRoomsFromCSV(roomsModifiedPath)
	if err != nil {
		log.Fatalf("can't read csv file: %s: %s", roomsModifiedPath, err.Error())
	}

	edgesInTable := services.ProcessData(edges, vertices)
	if err := services.CreateEdgesCSVFile(edgesPath, edgesInTable); err != nil {
		log.Fatalf("can't create csv file: %s: %s", edgesPath, err.Error())
	}
}

func CreateDBScripts(services *service.Service) {
	if err := services.GenerateRoomsScript(roomsModifiedPath, dbScriptRoomsPath); err != nil {
		log.Fatalf("can't open or create script: %s", err.Error())
	}

	if err := services.GenerateEdgeScript(edgesModifiedPath, dbScriptEdgesPath); err != nil {
		log.Fatalf("can't open or create script: %s", err.Error())
	}
}

func main() {
	services := service.NewService()
	CreateDBScripts(services)
}
