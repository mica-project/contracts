package chess

import "encoding/json/v2"

type ChessVoice struct {
	Code      string `json:"code"`
	Id        int64  `json:"id"`
	IsShow    int64  `json:"is_show"`
	Situation string `json:"situation"`
}

type ChessVoiceList []ChessVoice

func (s *ChessVoice) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
