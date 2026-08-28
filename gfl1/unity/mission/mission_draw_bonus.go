package mission

import "encoding/json/v2"

type MissionDrawBonus struct {
	X any `json:"_"`
}

type MissionDrawBonusList []MissionDrawBonus

func (s *MissionDrawBonus) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
