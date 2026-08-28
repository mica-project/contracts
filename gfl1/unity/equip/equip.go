package equip

import "encoding/json/v2"

type Equip struct {
	AutoSelectId      int64   `json:"auto_select_id"`
	Category          int64   `json:"category"`
	Code              string  `json:"code"`
	Company           string  `json:"company"`
	CriticalPercent   []int64 `json:"critical_percent"`
	Description       string  `json:"description"`
	DevelopDuration   int64   `json:"develop_duration"`
	EquipIntroduction string  `json:"equip_introduction"`
	ExclusiveRate     int64   `json:"exclusive_rate"`
	Id                int64   `json:"id"`
	IsShow            int64   `json:"is_show"`
	Name              string  `json:"name"`
	ObtainIds         []int64 `json:"obtain_ids"`
	Rank              int64   `json:"rank"`
	RetireMp          int64   `json:"retire_mp"`
	RetireMre         int64   `json:"retire_mre"`
	SpDescription     string  `json:"sp_description"`
	Type              int64   `json:"type"`
}

type EquipList []Equip

func (s *Equip) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
