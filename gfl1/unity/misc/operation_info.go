package misc

import "encoding/json/v2"

type OperationInfo struct {
	Ammo               string   `json:"ammo"`
	Description        []string `json:"description"`
	Duration           string   `json:"duration"`
	GunMin             string   `json:"gun_min"`
	Guntype3Min        string   `json:"guntype3_min"`
	Id                 string   `json:"id"`
	ItemPool           []int64  `json:"item_pool"`
	Mre                string   `json:"mre"`
	Name               string   `json:"name"`
	TeamLeaderMinLevel string   `json:"team_leader_min_level"`
}

type OperationInfoList []OperationInfo

func (s *OperationInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
