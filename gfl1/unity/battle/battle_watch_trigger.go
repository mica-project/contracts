package battle

import "encoding/json/v2"

type BattleWatchTrigger struct {
	BuffJudge   int64  `json:"buff_judge"`
	Description string `json:"description"`
	Id          int64  `json:"id"`
}

type BattleWatchTriggerList []BattleWatchTrigger

func (s *BattleWatchTrigger) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
