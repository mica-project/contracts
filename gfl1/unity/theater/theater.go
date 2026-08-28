package theater

import "encoding/json/v2"

type Theater struct {
	Area                 []int64   `json:"area"`
	BackgroundMask       []float64 `json:"background_mask"`
	Bgm                  string    `json:"bgm"`
	Code                 string    `json:"code"`
	FormationLimit       int64     `json:"formation_limit"`
	Gauge                int64     `json:"gauge"`
	HocFormationNumber   int64     `json:"hoc_formation_number"`
	HocLimit             int64     `json:"hoc_limit"`
	Id                   int64     `json:"id"`
	MaskOffset           []int64   `json:"mask_offset"`
	Name                 string    `json:"name"`
	OccupiedPrize        int64     `json:"occupied_prize"`
	OccupiedPrizeDisplay []string  `json:"occupied_prize_display"`
	Rank                 int64     `json:"rank"`
	ReinforceCoef        int64     `json:"reinforce_coef"`
	SsocLimit            int64     `json:"ssoc_limit"`
	TheaterEventId       int64     `json:"theater_event_id"`
	Type                 int64     `json:"type"`
}

type TheaterList []Theater

func (s *Theater) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
