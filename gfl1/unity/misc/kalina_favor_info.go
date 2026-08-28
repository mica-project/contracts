package misc

import "encoding/json/v2"

type KalinaFavorInfo struct {
	Level    string `json:"level"`
	MinFavor string `json:"min_favor"`
}

type KalinaFavorInfoList []KalinaFavorInfo

func (s *KalinaFavorInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
