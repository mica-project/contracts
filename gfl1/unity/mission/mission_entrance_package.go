package mission

import "encoding/json/v2"

type MissionEntrancePackage struct {
	Id    int64    `json:"id"`
	Value []string `json:"value"`
}

type MissionEntrancePackageList []MissionEntrancePackage

func (s *MissionEntrancePackage) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
