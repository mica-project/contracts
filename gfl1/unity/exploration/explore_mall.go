package exploration

import "encoding/json/v2"

type ExploreMall struct {
	Cost    string `json:"cost"`
	Id      int64  `json:"id"`
	PrizeId int64  `json:"prize_id"`
	Sort    int64  `json:"sort"`
	Type    int64  `json:"type"`
	TypeId  int64  `json:"type_id"`
}

type ExploreMallList []ExploreMall

func (s *ExploreMall) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
