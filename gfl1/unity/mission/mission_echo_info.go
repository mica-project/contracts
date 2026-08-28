package mission

import "encoding/json/v2"

type MissionEchoInfo struct {
	Character string   `json:"character"`
	Content   []string `json:"content"`
	Id        int64    `json:"id"`
	Interval  int64    `json:"interval"`
	MissionId int64    `json:"mission_id"`
	Subtitle  []string `json:"subtitle"`
	Title     string   `json:"title"`
}

type MissionEchoInfoList []MissionEchoInfo

func (s *MissionEchoInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
