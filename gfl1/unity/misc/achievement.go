package misc

import "encoding/json/v2"

type Achievement struct {
	Ammo      int64  `json:"ammo"`
	Content   string `json:"content"`
	Core      int64  `json:"core"`
	Count     int64  `json:"count"`
	IconCode  string `json:"icon_code"`
	Identity  int64  `json:"identity"`
	IfSteam   int64  `json:"if_steam"`
	ItemIds   string `json:"item_ids"`
	Mp        int64  `json:"mp"`
	Mre       int64  `json:"mre"`
	Part      int64  `json:"part"`
	Sort      int64  `json:"sort"`
	SteamCode string `json:"steam_code"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	TypeSort  int64  `json:"type_sort"`
	UserExp   int64  `json:"user_exp"`
}

type AchievementList []Achievement

func (s *Achievement) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
