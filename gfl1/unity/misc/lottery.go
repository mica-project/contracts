package misc

import "encoding/json/v2"

type Lottery struct {
	X any `json:"_"`
}

type LotteryList []Lottery

func (s *Lottery) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
