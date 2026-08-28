package misc

import "encoding/json/v2"

type StoryPlayback struct {
	CampaignSubName string `json:"campaign_sub_name"`
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Order           int64  `json:"order"`
	Tag             string `json:"tag"`
	Type            int64  `json:"type"`
}

type StoryPlaybackList []StoryPlayback

func (s *StoryPlayback) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
