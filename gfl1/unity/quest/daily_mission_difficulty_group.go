package quest

import "encoding/json/v2"

type DailyMissionDifficultyGroup struct {
	DefaultBg           string  `json:"default_bg"`
	Id                  int64   `json:"id"`
	Name                string  `json:"name"`
	RecommendFairies    []int64 `json:"recommend_fairies"`
	RecommendFairyLevel []int64 `json:"recommend_fairy_level"`
	RecommendFairyTip   string  `json:"recommend_fairy_tip"`
	RecommendGunLevel   int64   `json:"recommend_gun_level"`
	RecommendSquadTip   string  `json:"recommend_squad_tip"`
}

type DailyMissionDifficultyGroupList []DailyMissionDifficultyGroup

func (s *DailyMissionDifficultyGroup) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
