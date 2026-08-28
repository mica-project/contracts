package quest

import "encoding/json/v2"

type DailyMapLight struct {
	BrightContrastRatio       int64   `json:"bright_contrast_ratio"`
	DarkContrastRatio         float64 `json:"dark_contrast_ratio"`
	DirLightShadowColorLock   string  `json:"dir_light_shadow_colorLock"`
	DirLightShadowColorUnLock string  `json:"dir_light_shadow_colorUnLock"`
	Id                        int64   `json:"id"`
	LightAngle                int64   `json:"light_angle"`
	LightBrightness           float64 `json:"light_brightness"`
	LightColor                string  `json:"light_color"`
	LightDir                  string  `json:"light_dir"`
	LightPos                  string  `json:"light_pos"`
	LightTexBrightness        float64 `json:"light_tex_brightness"`
	ModelTexColor             string  `json:"model_tex_color"`
	SeasonCode                int64   `json:"season_code"`
	StateCode                 int64   `json:"state_code"`
}

type DailyMapLightList []DailyMapLight

func (s *DailyMapLight) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
