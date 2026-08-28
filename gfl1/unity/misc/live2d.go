package misc

import "encoding/json/v2"

type Live2d struct {
	Code     string  `json:"code"`
	FitGun   int64   `json:"fit_gun"`
	Id       int64   `json:"id"`
	IsShow   int64   `json:"is_show"`
	Motions  []int64 `json:"motions"`
	SkinLogo int64   `json:"skinLogo"`
	SkinType int64   `json:"skinType"`
}

type Live2dList []Live2d

func (s *Live2d) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
