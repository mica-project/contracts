package carnival

import "encoding/json/v2"

type CarnivalGotoPageInfo struct {
	GotoPage string `json:"goto_page"`
	Id       string `json:"id"`
	Type     string `json:"type"`
}

type CarnivalGotoPageInfoList []CarnivalGotoPageInfo

func (s *CarnivalGotoPageInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
