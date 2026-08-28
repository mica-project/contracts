package chess

import "encoding/json/v2"

type ChessSeasonevent struct {
	EndTime   string `json:"end_time"`
	Id        int64  `json:"id"`
	MissionId int64  `json:"mission_id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
}

type ChessSeasoneventList []ChessSeasonevent

func (s *ChessSeasonevent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
