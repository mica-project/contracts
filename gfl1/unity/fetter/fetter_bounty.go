package fetter

import "encoding/json/v2"

type FetterBounty struct {
	Description string `json:"description"`
	FetterId    int64  `json:"fetter_id"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Point       int64  `json:"point"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}

type FetterBountyList []FetterBounty

func (s *FetterBounty) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
