package gun

import "encoding/json/v2"

type GunInAlly struct {
	Dodge    int64 `json:"dodge"`
	EatLv    int64 `json:"eat_lv"`
	Equip1   int64 `json:"equip1"`
	Equip2   int64 `json:"equip2"`
	Equip3   int64 `json:"equip3"`
	GunId    int64 `json:"gun_id"`
	GunLevel int64 `json:"gun_level"`
	Hit      int64 `json:"hit"`
	Id       int64 `json:"id"`
	Life     int64 `json:"life"`
	Location int64 `json:"location"`
	Number   int64 `json:"number"`
	Position int64 `json:"position"`
	Pow      int64 `json:"pow"`
	Rate     int64 `json:"rate"`
	Rec      int64 `json:"rec"`
	Skill1   int64 `json:"skill1"`
}

type GunInAllyList []GunInAlly

func (s *GunInAlly) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
