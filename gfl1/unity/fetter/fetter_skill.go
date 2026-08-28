package fetter

import "encoding/json/v2"

type FetterSkill struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Gun         int64   `json:"gun"`
	GunGroup    []int64 `json:"gun_group"`
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Skill1      int64   `json:"skill1"`
	Type        int64   `json:"type"`
}

type FetterSkillList []FetterSkill

func (s *FetterSkill) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
