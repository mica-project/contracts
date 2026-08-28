package misc

import "encoding/json/v2"

type GuildEmoji struct {
	EmojiCode string `json:"emoji_code"`
	Group     int64  `json:"group"`
	Id        int64  `json:"id"`
}

type GuildEmojiList []GuildEmoji

func (s *GuildEmoji) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
