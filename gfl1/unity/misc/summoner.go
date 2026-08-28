package misc

import "encoding/json/v2"

type Summoner struct {
	Angle          int64     `json:"angle"`
	ArmorPiercing  int64     `json:"armor_piercing"`
	Camp           int64     `json:"camp"`
	Code           string    `json:"code"`
	Def            int64     `json:"def"`
	Hit            int64     `json:"hit"`
	Hp             int64     `json:"hp"`
	Id             int64     `json:"id"`
	IsDamageShowed int64     `json:"is_damage_showed"`
	IsHpShowed     int64     `json:"is_hp_showed"`
	Level          int64     `json:"level"`
	Name           string    `json:"name"`
	NormalAttack   int64     `json:"normal_attack"`
	Number         int64     `json:"number"`
	PassiveSkill   []int64   `json:"passive_skill"`
	Pow            int64     `json:"pow"`
	Range          int64     `json:"range"`
	Rate           int64     `json:"rate"`
	Scale          []float64 `json:"scale"`
	Shield         int64     `json:"shield"`
	Speed          int64     `json:"speed"`
	StartOffset    []int64   `json:"start_offset"`
	UnableBuffId   []int64   `json:"unable_buff_id"`
}

type SummonerList []Summoner

func (s *Summoner) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
