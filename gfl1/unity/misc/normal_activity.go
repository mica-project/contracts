package misc

import "encoding/json/v2"

type NormalActivity struct {
	BgCode      string   `json:"bg_code"`
	BgConfig    []string `json:"bg_config"`
	CampaignId  int64    `json:"campaign_id"`
	ChapterName string   `json:"chapter_name"`
	Desc        string   `json:"desc"`
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Order       int64    `json:"order"`
	Tag         string   `json:"tag"`
}

type NormalActivityList []NormalActivity

func (s *NormalActivity) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
