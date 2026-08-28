package contracts

import (
	"encoding/json/jsontext"
	"strconv"
	"strings"
)

type IntSet []int

type IntMap map[int]int

func (s *IntSet) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	tok, err := dec.ReadToken()

	if err != nil {
		return err
	}

	t := strings.SplitSeq(tok.String(), ",")

	for v := range t {
		i, _ := strconv.Atoi(v)
		*s = append(*s, i)
	}

	return err
}

//func (s *IntMap) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
//	tok, err := dec.ReadToken()
//
//	if err != nil {
//		return err
//	}
//
//	t := strings.SplitSeq(tok.String(), ";")
//
//	for v := range t {
//		i := strings.Split(v, ",")
//		key, _ := strconv.Atoi(i[0])
//		val, _ := strconv.Atoi(i[1])
//		*s[key] = val
//	}
//
//	return err
//}
