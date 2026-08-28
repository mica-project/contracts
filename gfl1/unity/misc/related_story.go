package misc

import "encoding/json/v2"

type RelatedStory struct {
	Id                int64   `json:"id"`
	MindupdateStoryId int64   `json:"mindupdate_story_id"`
	SkinId            []int64 `json:"skin_id"`
	StoryCampaignId   []int64 `json:"story_campaign_id"`
}

type RelatedStoryList []RelatedStory

func (s *RelatedStory) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
