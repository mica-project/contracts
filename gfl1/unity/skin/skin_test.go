package skin

import (
	"encoding/json/v2"
	"io"
	"net/http"
	"testing"
)

func TestUnmarshalSkin(t *testing.T) {
	url := "https://gfverse-assets.sfo3.cdn.digitaloceanspaces.com/gfl1/import/formatted/skin.json"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("fetch fixture: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("bad status: %d", resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	var v []Skin
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(v) == 0 {
		t.Fatalf("expected non-empty slice")
	}
}
