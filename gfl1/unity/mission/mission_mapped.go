package mission

import "encoding/json/v2"

type MissionMapped struct {
	Id             int64   `json:"id"`
	MappedMissions []int64 `json:"mapped_missions"`
}

type MissionMappedList []MissionMapped

func (s *MissionMapped) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
