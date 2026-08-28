package gun

import "encoding/json/v2"

type GunCharavoice struct {
	Code string `json:"code"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GunCharavoiceList []GunCharavoice

func (s *GunCharavoice) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
