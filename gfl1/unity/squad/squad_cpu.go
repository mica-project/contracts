package squad

import "encoding/json/v2"

type SquadCpu struct {
	Color    int64 `json:"color"`
	CpuBonus int64 `json:"cpu_bonus"`
	Grid1    int64 `json:"grid1"`
	Grid2    int64 `json:"grid2"`
	Grid3    int64 `json:"grid3"`
	Grid4    int64 `json:"grid4"`
	Grid5    int64 `json:"grid5"`
	Id       int64 `json:"id"`
}

type SquadCpuList []SquadCpu

func (s *SquadCpu) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
