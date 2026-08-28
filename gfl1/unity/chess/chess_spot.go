package chess

import "encoding/json/v2"

type ChessSpot struct {
	AxialCoordinatorQ int64    `json:"axial_coordinator_q"`
	AxialCoordinatorR int64    `json:"axial_coordinator_r"`
	ChessMissionId    int64    `json:"chess_mission_id"`
	Id                int64    `json:"id"`
	NegativeDirection []int64  `json:"negative_direction"`
	Neighbor          []string `json:"neighbor"`
	PlayerOrder       int64    `json:"player_order"`
	PositiveDirection []int64  `json:"positive_direction"`
	Type              int64    `json:"type"`
}

type ChessSpotList []ChessSpot

func (s *ChessSpot) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
