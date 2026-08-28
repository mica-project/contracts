package quest

import "encoding/json/v2"

type DailyEventOption struct {
	BalanceOut              int64  `json:"balance_out"`
	BattleDuration          int64  `json:"battle_duration"`
	Description             string `json:"description"`
	FightEnvironmentSkillId int64  `json:"fight_environment_skill_id"`
	Id                      int64  `json:"id"`
	RoundDuration           int64  `json:"round_duration"`
}

type DailyEventOptionList []DailyEventOption

func (s *DailyEventOption) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
