package mission

import "encoding/json/v2"

type MissionTargettrain struct {
	BattleTimelimit     int64    `json:"battle_timelimit"`
	Code                string   `json:"code"`
	DifficultLevel      int64    `json:"difficult_level"`
	DifficultLevelLimit int64    `json:"difficult_level_limit"`
	HpIsRe              int64    `json:"hp_is_re"`
	Id                  int64    `json:"id"`
	Name                string   `json:"name"`
	RecommendLevel      int64    `json:"recommend_level"`
	TargetDes           []string `json:"target_des"`
	TargetId            string   `json:"target_id"`
	TargetType          int64    `json:"target_type"`
}

type MissionTargettrainList []MissionTargettrain

func (s *MissionTargettrain) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
