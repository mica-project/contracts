package squad

import "encoding/json/v2"

type SquadRank struct {
	ChipPoint int64  `json:"chip_point"`
	CostData  int64  `json:"cost_data"`
	CpuRate   int64  `json:"cpu_rate"`
	LvUnlock  string `json:"lv_unlock"`
	StarId    int64  `json:"star_id"`
}

type SquadRankList []SquadRank

func (s *SquadRank) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
