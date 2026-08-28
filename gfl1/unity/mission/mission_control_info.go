package mission

import "encoding/json/v2"

type MissionControlInfo struct {
	BloodColor  string `json:"blood_color"`
	Id          string `json:"id"`
	SpineColor  string `json:"spine_color"`
	SpotColor   string `json:"spot_color"`
	TeamAiColor string `json:"team_ai_color"`
	TeamCode    string `json:"team_code"`
	TeamColor   string `json:"team_color"`
	TurnCode    string `json:"turn_code"`
	TurnColor   string `json:"turn_color"`
}

type MissionControlInfoList []MissionControlInfo

func (s *MissionControlInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
