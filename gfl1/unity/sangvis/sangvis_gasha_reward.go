package sangvis

import "encoding/json/v2"

type SangvisGashaReward struct {
	Id         int64  `json:"id"`
	RewardType int64  `json:"reward_type"`
	SangvisId  string `json:"sangvis_id"`
}

type SangvisGashaRewardList []SangvisGashaReward

func (s *SangvisGashaReward) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
