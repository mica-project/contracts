package exploration

import "encoding/json/v2"

type ExploreTimeType struct {
	Ammo           int64   `json:"ammo"`
	Duration       int64   `json:"duration"`
	Id             int64   `json:"id"`
	Mp             int64   `json:"mp"`
	Mre            int64   `json:"mre"`
	Part           int64   `json:"part"`
	RewardItemType []int64 `json:"reward_item_type"`
}

type ExploreTimeTypeList []ExploreTimeType

func (s *ExploreTimeType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
