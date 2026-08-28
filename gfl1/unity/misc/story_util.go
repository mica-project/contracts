package misc

import "encoding/json/v2"

type StoryUtil struct {
	BackgroundCode     string   `json:"background_code"`
	Bgm                string   `json:"bgm"`
	Description        string   `json:"description"`
	End                string   `json:"end"`
	First              string   `json:"first"`
	Id                 int64    `json:"id"`
	IsUtil             int64    `json:"is_util"`
	MissionId          string   `json:"mission_id"`
	Scripts            []string `json:"scripts"`
	Start              string   `json:"start"`
	StoryPlaybackSubId int64    `json:"story_playback_sub_id"`
	Title              string   `json:"title"`
}

type StoryUtilList []StoryUtil

func (s *StoryUtil) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
