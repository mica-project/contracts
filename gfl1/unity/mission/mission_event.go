package mission

import "encoding/json/v2"

type MissionEvent struct {
	EventName       string  `json:"event_name"`
	Id              int64   `json:"id"`
	InitMissionId   string  `json:"init_mission_id"`
	MissionCampaign []int64 `json:"mission_campaign"`
	Type            int64   `json:"type"`
}

type MissionEventList []MissionEvent

func (s *MissionEvent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
