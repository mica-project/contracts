package battle

import "encoding/json/v2"

type BattleCreation struct {
	Code              string    `json:"code"`
	DestinationType   string    `json:"destination_type"`
	EffectArea        int64     `json:"effect_area"`
	HurtId            string    `json:"hurt_id"`
	Id                int64     `json:"id"`
	IsFormOffset      int64     `json:"is_form_offset"`
	IsFormOffsetStart int64     `json:"is_form_offset_start"`
	IsFormPlay        int64     `json:"is_form_play"`
	IsOrbEffect       int64     `json:"is_orb_effect"`
	Name              string    `json:"name"`
	RouteType         int64     `json:"route_type"`
	Scale             []float64 `json:"scale"`
	SoundOrder        string    `json:"sound_order"`
	Speed             int64     `json:"speed"`
	StartType         string    `json:"start_type"`
}

type BattleCreationList []BattleCreation

func (s *BattleCreation) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
