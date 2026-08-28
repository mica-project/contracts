package battle

import "encoding/json/v2"

type BattleActionConfig struct {
	ActionOrder     string `json:"action_order"`
	ActionPlayspeed string `json:"action_playspeed"`
	CreationOrder   string `json:"creation_order"`
	Id              int64  `json:"id"`
	Name            string `json:"name"`
}

type BattleActionConfigList []BattleActionConfig

func (s *BattleActionConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
