package sangvis

import "encoding/json/v2"

type SangvisAdvance struct {
	AdvanceArmor int64 `json:"advance_armor"`
	AdvanceDodge int64 `json:"advance_dodge"`
	AdvanceHit   int64 `json:"advance_hit"`
	AdvanceHp    int64 `json:"advance_hp"`
	AdvancePow   int64 `json:"advance_pow"`
	AdvanceRate  int64 `json:"advance_rate"`
	AdvanceRec   int64 `json:"advance_rec"`
	Lv           int64 `json:"lv"`
	UnlockLv     int64 `json:"unlock_lv"`
}

type SangvisAdvanceList []SangvisAdvance

func (s *SangvisAdvance) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
