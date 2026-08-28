package misc

import "encoding/json/v2"

type NpcCharavoice struct {
	Code string `json:"code"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type NpcCharavoiceList []NpcCharavoice

func (s *NpcCharavoice) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
