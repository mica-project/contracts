package sangvis

import "encoding/json/v2"

type SangvisCharavoice struct {
	Code string `json:"code"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SangvisCharavoiceList []SangvisCharavoice

func (s *SangvisCharavoice) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
