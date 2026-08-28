package enemy

import "encoding/json/v2"

type EnemyVoiceInfo struct {
	Code string `json:"code"`
	Id   string `json:"id"`
}

type EnemyVoiceInfoList []EnemyVoiceInfo

func (s *EnemyVoiceInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
