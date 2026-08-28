package mission

import "encoding/json/v2"

type MissionSkip struct {
	Content       string  `json:"content"`
	Id            int64   `json:"id"`
	Mission       []int64 `json:"mission"`
	MissionText   string  `json:"mission_text"`
	MissionUnlock []int64 `json:"mission_unlock"`
	Stage         int64   `json:"stage"`
	Title         string  `json:"title"`
}

type MissionSkipList []MissionSkip

func (s *MissionSkip) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
