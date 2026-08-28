package sangvis

import "encoding/json/v2"

type SangvisGasha struct {
	AuthorPrice    string   `json:"author_price"`
	Banner         string   `json:"banner"`
	DailyPrice     string   `json:"daily_price"`
	EndTime        string   `json:"end_time"`
	GashaRewardIds []string `json:"gasha_reward_ids"`
	Name           string   `json:"name"`
	RefreshRate    int64    `json:"refresh_rate"`
	StartTime      string   `json:"start_time"`
	TabCode        string   `json:"tab_code"`
	Type           int64    `json:"type"`
}

type SangvisGashaList []SangvisGasha

func (s *SangvisGasha) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
