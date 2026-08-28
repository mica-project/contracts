package fetter

import "encoding/json/v2"

type FetterStory struct {
	Actor       string `json:"actor"`
	Description string `json:"description"`
	FetterId    int64  `json:"fetter_id"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Reward      int64  `json:"reward"`
}

type FetterStoryList []FetterStory

func (s *FetterStory) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
