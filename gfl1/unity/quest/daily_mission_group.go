package quest

import "encoding/json/v2"

type DailyMissionGroup struct {
	DailyPt         int64   `json:"daily_pt"`
	Difficulty      int64   `json:"difficulty"`
	DifficultyGroup int64   `json:"difficulty_group"`
	Id              int64   `json:"id"`
	MissionGroup    []int64 `json:"mission_group"`
	StateType       int64   `json:"state_type"`
	Type            int64   `json:"type"`
}

type DailyMissionGroupList []DailyMissionGroup

func (s *DailyMissionGroup) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
