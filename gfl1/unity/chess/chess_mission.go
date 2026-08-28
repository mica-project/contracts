package chess

import "encoding/json/v2"

type ChessMission struct {
	CameraAngleH      int64     `json:"camera_angle_h"`
	CameraAngleL      int64     `json:"camera_angle_l"`
	CameraHeightRange []int64   `json:"camera_height_range"`
	ChessSpotids      []int64   `json:"chess_spotids"`
	GlobalLimit       []float64 `json:"global_limit"`
	GlobalPos         []float64 `json:"global_pos"`
	Id                int64     `json:"id"`
	MapId             []int64   `json:"map_id"`
	MapLimit          []int64   `json:"map_limit"`
	MissionIcon       string    `json:"mission_icon"`
	Name              string    `json:"name"`
	Rotation          int64     `json:"rotation"`
}

type ChessMissionList []ChessMission

func (s *ChessMission) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
