package skin

import "encoding/json/v2"

type SkinGroupInfo struct {
	Description []string `json:"description"`
	Icon        string   `json:"icon"`
	Id          string   `json:"id"`
	Order       string   `json:"order"`
	Skin        []int64  `json:"skin"`
	Theme       string   `json:"theme"`
	TitleCode   string   `json:"title_code"`
}

type SkinGroupInfoList []SkinGroupInfo

func (s *SkinGroupInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
