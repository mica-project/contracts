package misc

import "encoding/json/v2"

type DormEmojiTextInfo struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type DormEmojiTextInfoList []DormEmojiTextInfo

func (s *DormEmojiTextInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
