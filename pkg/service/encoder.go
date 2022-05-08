package service

import (
	"Data_script_Run_Hse_Run/model"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

type CSVEncoderService struct{}

func parseEdgesLine(line []string) (model.Edge, error) {
	if len(line) != 4 {
		return model.Edge{}, errors.New("invalid csv")
	}

	startVertex, err := strconv.Atoi(line[0])
	if err != nil {
		return model.Edge{}, err
	}

	endVertex, err := strconv.Atoi(line[1])
	if err != nil {
		return model.Edge{}, err
	}

	cost, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		return model.Edge{}, err
	}

	campus, err := strconv.Atoi(line[3])
	if err != nil {
		return model.Edge{}, err
	}

	return model.Edge{
		StartVertex: startVertex,
		EndVertex:   endVertex,
		Cost:        cost,
		CampusId:    campus,
	}, nil
}

func parseVertexLine(line []string) (model.Vertex, error) {
	if len(line) != 4 {
		return model.Vertex{}, errors.New("invalid csv")
	}

	id, err := strconv.Atoi(line[0])
	if err != nil {
		return model.Vertex{}, err
	}

	typeVertex := line[1]

	roomId, err := strconv.Atoi(line[2])
	if err != nil {
		return model.Vertex{}, err
	}

	campus, err := strconv.Atoi(line[3])
	if err != nil {
		return model.Vertex{}, err
	}

	return model.Vertex{
		Id:       id,
		Type:     typeVertex,
		RoomId:   roomId,
		CampusId: campus,
	}, nil
}

func parseRoomLine(line []string) (model.Room, error) {
	if len(line) != 3 {
		return model.Room{}, errors.New("invalid csv")
	}

	id, err := strconv.Atoi(line[0])
	if err != nil {
		return model.Room{}, err
	}

	code := line[1]

	campus, err := strconv.Atoi(line[2])
	if err != nil {
		return model.Room{}, err
	}

	return model.Room{
		Id:       id,
		Code:     code,
		CampusId: campus,
	}, nil
}

func (e *CSVEncoderService) GetEdgesFromCSV(path string) ([]model.Edge, error) {
	csvFile, err := os.Open(path)
	defer csvFile.Close()

	if err != nil {
		return nil, err
	}

	var edges []model.Edge

	reader := csv.NewReader(csvFile)
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		edge, err := parseEdgesLine(line)
		if err != nil {
			return nil, err
		}
		edges = append(edges, edge)
	}

	return edges, nil
}

func (e *CSVEncoderService) GetVerticesFromCSV(path string) (model.Vertices, error) {
	csvFile, err := os.Open(path)
	defer csvFile.Close()

	vertices := model.NewVertices()

	if err != nil {
		return vertices, err
	}

	reader := csv.NewReader(csvFile)
	_, err = reader.Read()
	if err != nil {
		return vertices, err
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return vertices, err
		}

		vertex, err := parseVertexLine(line)
		if err != nil {
			return vertices, err
		}

		vertices[vertex.Id] = vertex
	}

	return vertices, nil
}

func (e *CSVEncoderService) GetRoomsFromCSV(path string) (model.Rooms, error) {
	csvFile, err := os.Open(path)
	defer csvFile.Close()

	rooms := model.NewRooms()

	if err != nil {
		return rooms, err
	}

	reader := csv.NewReader(csvFile)
	_, err = reader.Read()
	if err != nil {
		return rooms, err
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return rooms, err
		}

		room, err := parseRoomLine(line)
		if err != nil {
			return rooms, err
		}

		rooms[room.Id] = room
	}

	return rooms, nil
}

func NewCSVEncoderService() *CSVEncoderService {
	return &CSVEncoderService{}
}
