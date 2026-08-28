package sangvis

import "encoding/json/v2"

type SangvisResolution struct {
	EffectGridEffect []string `json:"effect_grid_effect"`
	GroupId          int64    `json:"group_id"`
	Id               int64    `json:"id"`
	ResolutionDodge  int64    `json:"resolution_dodge"`
	ResolutionHit    int64    `json:"resolution_hit"`
	ResolutionHp     int64    `json:"resolution_hp"`
	ResolutionNumber int64    `json:"resolution_number"`
	ResolutionPow    int64    `json:"resolution_pow"`
	ResolutionRate   int64    `json:"resolution_rate"`
	ResolutionRec    int64    `json:"resolution_rec"`
}

type SangvisResolutionList []SangvisResolution

func (s *SangvisResolution) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
