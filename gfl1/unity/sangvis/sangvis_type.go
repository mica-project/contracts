package sangvis

import "encoding/json/v2"

type SangvisType struct {
	AuthorSuccessr   int64   `json:"author_successr"`
	BasicArmor       float64 `json:"basic_armor"`
	BasicDodge       float64 `json:"basic_dodge"`
	BasicHit         float64 `json:"basic_hit"`
	BasicHp          float64 `json:"basic_hp"`
	BasicPow         float64 `json:"basic_pow"`
	BasicRate        float64 `json:"basic_rate"`
	BasicRec         float64 `json:"basic_rec"`
	BasicSpeed       int64   `json:"basic_speed"`
	DailySuccessr    int64   `json:"daily_successr"`
	DefaultAdvanceLv int64   `json:"default_advance_lv"`
	ExchangeNum      string  `json:"exchange_num"`
	FixTimeRatio     int64   `json:"fix_time_ratio"`
	Id               int64   `json:"id"`
	MpFixRatio       int64   `json:"mp_fix_ratio"`
	Name             string  `json:"name"`
	PartFixRatio     int64   `json:"part_fix_ratio"`
	PicAdvanceLv     int64   `json:"pic_advance_lv"`
	RepairCost       int64   `json:"repair_cost"`
	SkillAdvanceLv   int64   `json:"skill_advance_lv"`
	SkillsMaxLv      []int64 `json:"skills_max_lv"`
	TransNum         []int64 `json:"trans_num"`
}

type SangvisTypeList []SangvisType

func (s *SangvisType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
