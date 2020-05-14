package seed_test

import (
	"testing"

	"github.com/mrcook/svgtoolkit/geopattern/seed"
)

func TestSeed_ToInt(t *testing.T) {
	s := seed.New("Seed Test")
	result := s.ToInt(17, 1)
	if result != 9 {
		t.Fatalf("expected correct int value, got: %d", result)
	}

	result = s.ToInt(5, 3)
	if result != 3471 {
		t.Fatalf("expected correct int value, got: %d", result)
	}

	result = s.ToInt(2, 6)
	if result != 3603855 {
		t.Fatalf("expected correct int value, got: %d", result)
	}
}
