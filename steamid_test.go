package steamid_test

import (
	"fmt"
	"testing"

	"github.com/crossworth/steamid"
)

var validSteamIds = []interface{}{
	"STEAM_0:1:40225689",
	"[U:1:80451379]",
	"76561198040717107",
	76561198040717107,
}

var invalidSteamIds = []interface{}{
	"test",
	"1233.0f",
	"fff",
	"",
	"STEMM_0:1:40225689",
}

func TestCanCreateSteamID(t *testing.T) {

	for _, sid := range validSteamIds {
		_, err := steamid.New(sid)

		if err != nil {
			t.Fail()
		}
	}

	for _, sid := range invalidSteamIds {
		_, err := steamid.New(sid)

		if err == nil {
			t.Fail()
		}
	}
}

func TestCanValidateSteamID(t *testing.T) {
	for _, sid := range validSteamIds {
		s, _ := steamid.New(sid)

		if !s.IsValid() {
			t.Fail()
		}
	}
}

func TestCanGetSteam2RenderedIDWithOldFormat(t *testing.T) {
	sid, _ := steamid.New(76561198040717107)

	if sid.GetSteam2RenderedID() != "STEAM_0:1:40225689" {
		t.Errorf("Expected %v got %v\n", "STEAM_0:1:40225689", sid.GetSteam2RenderedID())
	}
}

func TestCanGetSteam2RenderedIDWithNewFormat(t *testing.T) {
	sid, _ := steamid.New(76561198040717107)

	if sid.GetSteam2RenderedID(true) != "STEAM_1:1:40225689" {
		t.Errorf("Expected %v got %v\n", "STEAM_1:1:40225689", sid.GetSteam2RenderedID())
	}
}

func TestCanGetSteam3RenderedID(t *testing.T) {
	sid, _ := steamid.New(76561198040717107)

	if sid.GetSteam3RenderedID() != "[U:1:80451379]" {
		t.Errorf("Expected %v got %v\n", "[U:1:80451379]", sid.GetSteam3RenderedID())
	}
}

func ExampleNew() {
	sid, _ := steamid.New(76561198040717107)
	fmt.Println(sid)
	// Output: {1 1 1 80451379}
}

func ExampleGetSteam2RenderedID() {
	sid, _ := steamid.New(76561198040717107)
	fmt.Println(sid.GetSteam2RenderedID(), sid.GetSteam2RenderedID(true))
	// Output: STEAM_0:1:40225689 STEAM_1:1:40225689
}

func ExampleCanGetSteam3RenderedID() {
	sid, _ := steamid.New(76561198040717107)
	fmt.Println(sid.GetSteam3RenderedID())
	// Output: [U:1:80451379]
}

func ExampleGetSteamID64() {
	sid, _ := steamid.New(76561198040717107)
	fmt.Println(sid.GetSteamID64())
	// Output: 76561198040717107
}
