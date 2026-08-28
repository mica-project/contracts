package mission

import "encoding/json/v2"

type MissionWinTypeConfig struct {
	Arguments   string `json:"arguments"`
	Id          int64  `json:"id"`
	IsShowCount string `json:"is_show_count"`
	Type        int64  `json:"type"`
}

type MissionWinTypeConfigList []MissionWinTypeConfig

func (s *MissionWinTypeConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
