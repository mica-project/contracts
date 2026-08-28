package misc

import "encoding/json/v2"

type MindupdateStoryInfo struct {
	GunId   string `json:"gun_id"`
	Id      string `json:"id"`
	PrizeId string `json:"prize_id"`
	Scripts string `json:"scripts"`
	StageId string `json:"stage_id"`
}

type MindupdateStoryInfoList []MindupdateStoryInfo

func (s *MindupdateStoryInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
