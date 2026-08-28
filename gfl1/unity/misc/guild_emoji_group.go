package misc

import "encoding/json/v2"

type GuildEmojiGroup struct {
	IconEmojiCode string `json:"icon_emoji_code"`
	Id            int64  `json:"id"`
}

type GuildEmojiGroupList []GuildEmojiGroup

func (s *GuildEmojiGroup) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
