package battle

import "encoding/json/v2"

type FightSuccessCondition struct {
	ConditionCoef string `json:"condition_coef"`
	Desc          string `json:"desc"`
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ScoreCoef     int64  `json:"score_coef"`
	Type          int64  `json:"type"`
}

type FightSuccessConditionList []FightSuccessCondition

func (s *FightSuccessCondition) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
