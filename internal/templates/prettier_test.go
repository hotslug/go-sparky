package templates

import (
	"encoding/json"
	"testing"
)

func TestPrettierConfig_IsValidJSON(t *testing.T) {
	cfg := PrettierConfig()
	var v any
	if err := json.Unmarshal([]byte(cfg), &v); err != nil {
		t.Fatalf("Prettier config is not valid JSON: %v", err)
	}
}
