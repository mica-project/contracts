package fairy

import "encoding/json/v2"

type FairySkinInfo struct {
	Ai           string `json:"ai"`
	GiftFairy    string `json:"gift_fairy"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	PicId        string `json:"pic_id"`
	SkinName     string `json:"skin_name"`
	StrengthenLv string `json:"strengthen_lv"`
}

type FairySkinInfoList []FairySkinInfo

func (s *FairySkinInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
