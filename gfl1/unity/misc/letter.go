package misc

import "encoding/json/v2"

type Letter struct {
	Code    string `json:"code"`
	Id      int64  `json:"id"`
	QuoteId int64  `json:"quote_id"`
	Type    int64  `json:"type"`
}

type LetterList []Letter

func (s *Letter) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
