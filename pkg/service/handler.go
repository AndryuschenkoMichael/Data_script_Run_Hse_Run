package service

import "Data_script_Run_Hse_Run/model"

type DataHandlerServices struct{}

func (d *DataHandlerServices) ProcessData(edges []model.Edge, vertices model.Vertices) []model.Edge {
	var listEdges []model.Edge
	graph := model.NewGraph(edges)

	for _, vertex := range vertices {
		ed := graph.CountDistances(vertex)
		for _, value := range ed {
			if value.StartVertex == value.EndVertex {
				continue
			}

			startVertex, ok := vertices[value.StartVertex]
			if !ok {
				continue
			}

			endVertex, ok := vertices[value.EndVertex]
			if !ok {
				continue
			}

			listEdges = append(listEdges, model.Edge{
				StartVertex: startVertex.RoomId,
				EndVertex:   endVertex.RoomId,
				Cost:        value.Cost,
				CampusId:    value.CampusId,
			})
		}
	}

	return listEdges
}

func NewDataHandlerServices() *DataHandlerServices {
	return &DataHandlerServices{}
}
