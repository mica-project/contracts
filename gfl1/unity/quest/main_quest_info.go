package quest

import "encoding/json/v2"

type MainQuestInfo struct {
	Ammo     string `json:"ammo"`
	Content  string `json:"content"`
	Count    string `json:"count"`
	GunId    string `json:"gun_id"`
	Identity string `json:"identity"`
	Mp       string `json:"mp"`
	Mre      string `json:"mre"`
	Part     string `json:"part"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	UserExp  string `json:"user_exp"`
}

type MainQuestInfoList []MainQuestInfo

func (s *MainQuestInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
