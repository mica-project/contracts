package misc

import "encoding/json/v2"

type Spot struct {
	Belong       int64   `json:"belong"`
	CoordinatorX int64   `json:"coordinator_x"`
	CoordinatorY int64   `json:"coordinator_y"`
	CurveControl []int64 `json:"curve_control"`
	Id           int64   `json:"id"`
	MapCode      string  `json:"map_code"`
	MapRoute     string  `json:"map_route"`
	MapType      int64   `json:"map_type"`
	MissionId    int64   `json:"mission_id"`
	Route        string  `json:"route"`
	Type         int64   `json:"type"`
}

type SpotList []Spot

func (s *Spot) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
