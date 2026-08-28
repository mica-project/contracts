package fairy

import "encoding/json/v2"

type FairyCharavoice struct {
	Code string `json:"code"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type FairyCharavoiceList []FairyCharavoice

func (s *FairyCharavoice) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
