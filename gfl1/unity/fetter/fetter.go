package fetter

import "encoding/json/v2"

type Fetter struct {
	Actor            string `json:"actor"`
	Code             string `json:"code"`
	Id               int64  `json:"id"`
	Milestone1       int64  `json:"milestone1"`
	Milestone1Reward string `json:"milestone1_reward"`
	Milestone2       int64  `json:"milestone2"`
	Milestone2Reward string `json:"milestone2_reward"`
	Milestone3       int64  `json:"milestone3"`
	Milestone3Reward string `json:"milestone3_reward"`
	Milestone4       int64  `json:"milestone4"`
	Milestone4Reward string `json:"milestone4_reward"`
	Milestone5       int64  `json:"milestone5"`
	Milestone5Reward string `json:"milestone5_reward"`
	Name             string `json:"name"`
}

type FetterList []Fetter

func (s *Fetter) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
