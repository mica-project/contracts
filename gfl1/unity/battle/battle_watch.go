package battle

import "encoding/json/v2"

type BattleWatch struct {
	ActivationSkills       string `json:"activation_skills"`
	ActivationWatchTrigger int64  `json:"activation_watch_trigger"`
	Description            string `json:"description"`
	Id                     int64  `json:"id"`
	NewRecordChara1        int64  `json:"new_record_chara1"`
	RecordChara1           int64  `json:"record_chara1"`
	WatchTriggerType       int64  `json:"watch_trigger_type"`
}

type BattleWatchList []BattleWatch

func (s *BattleWatch) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
