package mission

import "encoding/json/v2"

type MissionGroupInfo struct {
	Campaign         string  `json:"campaign"`
	Id               string  `json:"id"`
	ResetDrawEvents  string  `json:"reset_draw_events"`
	ResetMissionKeys string  `json:"reset_mission_keys"`
	ResetMissions    []int64 `json:"reset_missions"`
}

type MissionGroupInfoList []MissionGroupInfo

func (s *MissionGroupInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
