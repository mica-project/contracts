package enemy

import "encoding/json/v2"

type EnemyIllustration struct {
	Character  []int64 `json:"character"`
	Code       string  `json:"code"`
	Counter    string  `json:"counter"`
	Id         int64   `json:"id"`
	Introduce  string  `json:"introduce"`
	LaunchTime string  `json:"launch_time"`
	LifeRank   int64   `json:"life_rank"`
	Name       string  `json:"name"`
	SpineScale []int64 `json:"spine_scale"`
	SubId      int64   `json:"sub_id"`
	SubName    string  `json:"sub_name"`
	Type       int64   `json:"type"`
}

type EnemyIllustrationList []EnemyIllustration

func (s *EnemyIllustration) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
