package mission

import "encoding/json/v2"

type MissionKeyInfo struct {
	Id string `json:"id"`
}

type MissionKeyInfoList []MissionKeyInfo

func (s *MissionKeyInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
