package chess

import "encoding/json/v2"

type ChessChip struct {
	ChipGroup         int64    `json:"chip_group"`
	ChipUpExp         int64    `json:"chip_up_exp"`
	Code              string   `json:"code"`
	Description       []string `json:"description"`
	Experience        int64    `json:"experience"`
	FitGunType        string   `json:"fit_gun_type"`
	Id                int64    `json:"id"`
	InitTime          int64    `json:"init_time"`
	Level             int64    `json:"level"`
	Name              string   `json:"name"`
	Number            int64    `json:"number"`
	Price             string   `json:"price"`
	Rank              int64    `json:"rank"`
	SellPrice         int64    `json:"sell_price"`
	SkillTargetFirst  string   `json:"skill_target_first"`
	TargetSelectFirst int64    `json:"target_select_first"`
	Type              int64    `json:"type"`
}

type ChessChipList []ChessChip

func (s *ChessChip) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
