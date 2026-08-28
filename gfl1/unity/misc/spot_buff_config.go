package misc

import "encoding/json/v2"

type SpotBuffConfig struct {
	Code          string `json:"code"`
	GrowEnemyPool int64  `json:"grow_enemy_pool"`
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Type          int64  `json:"type"`
	VisionBuff    string `json:"vision_buff"`
}

type SpotBuffConfigList []SpotBuffConfig

func (s *SpotBuffConfig) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
