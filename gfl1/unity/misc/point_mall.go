package misc

import "encoding/json/v2"

type PointMall struct {
	ClassificationId int64 `json:"classification_id"`
	DoubleGem        int64 `json:"double_gem"`
	DoubleGemPay     int64 `json:"double_gem_pay"`
	Gem              int64 `json:"gem"`
	GemPay           int64 `json:"gem_pay"`
	Id               int64 `json:"id"`
	Type             int64 `json:"type"`
}

type PointMallList []PointMall

func (s *PointMall) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
