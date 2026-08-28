package gun

import "encoding/json/v2"

type GunTypeInfo struct {
	BasicAttributeArmor string `json:"basic_attribute_armor"`
	BasicAttributeDodge string `json:"basic_attribute_dodge"`
	BasicAttributeHit   string `json:"basic_attribute_hit"`
	BasicAttributeLife  string `json:"basic_attribute_life"`
	BasicAttributePow   string `json:"basic_attribute_pow"`
	BasicAttributeRate  string `json:"basic_attribute_rate"`
	BasicAttributeRec   string `json:"basic_attribute_rec"`
	BasicAttributeSpeed string `json:"basic_attribute_speed"`
	FixTimeRatio        string `json:"fix_time_ratio"`
	Id                  string `json:"id"`
	MpFixRatio          string `json:"mp_fix_ratio"`
	Name                string `json:"name"`
	PartFixRatio        string `json:"part_fix_ratio"`
}

type GunTypeInfoList []GunTypeInfo

func (s *GunTypeInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
