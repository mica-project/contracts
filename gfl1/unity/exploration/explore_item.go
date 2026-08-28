package exploration

import "encoding/json/v2"

type ExploreItem struct {
	ExploreTimeDown float64 `json:"explore_time_down"`
	Id              int64   `json:"id"`
	MidRewardUp     int64   `json:"mid_reward_up"`
	Reward1Up       int64   `json:"reward1_up"`
}

type ExploreItemList []ExploreItem

func (s *ExploreItem) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
