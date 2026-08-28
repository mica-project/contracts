package squad

import "encoding/json/v2"

type SquadChip struct {
	AssistDamage    int64  `json:"assist_damage"`
	AssistReload    int64  `json:"assist_reload"`
	Color           int64  `json:"color"`
	GridNumber      int64  `json:"grid_number"`
	Id              int64  `json:"id"`
	IsRandom        int64  `json:"is_random"`
	Name            string `json:"name"`
	Rank            int64  `json:"rank"`
	SquadId         int64  `json:"squad_id"`
	UnlockChipPoint int64  `json:"unlock_chip_point"`
}

type SquadChipList []SquadChip

func (s *SquadChip) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
