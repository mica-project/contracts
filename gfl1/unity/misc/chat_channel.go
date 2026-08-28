package misc

import "encoding/json/v2"

type ChatChannel struct {
	Id int64 `json:"id"`
}

type ChatChannelList []ChatChannel

func (s *ChatChannel) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
