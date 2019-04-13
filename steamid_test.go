package steamid

import "testing"

func TestCanCreateSteamID2(t *testing.T) {
    _, err := NewSteamID("STEAM_0:1:40225689")

    if err != nil {
       t.Fail()
    }
}

func TestCanCreateSteamID3(t *testing.T) {
    _, err := NewSteamID("[U:1:80451379]")

    if err != nil {
       t.Fail()
    }
}

func TestCanCreateSteamID64WithString(t *testing.T) {
    _, err := NewSteamID("76561198040717107")

    if err != nil {
       t.Fail()
    }
}

func TestCanCreateSteamID64WithNumber(t *testing.T) {
    _, err := NewSteamID(76561198040717107)

    if err != nil {
       t.Fail()
    }
}

func TestCannotCreateSteamIDWithInvalidInput(t *testing.T) {
	values := []string {"test", "1233.0f", "fff", "", "STEMM_0:1:40225689"}

	for _, value := range values {
		_, err := NewSteamID(value)

		if err == nil {
		   t.Errorf("Input %v should not be valid", value)
		}
	}
}

func TestCanValidateSteamID2(t *testing.T) {
	sid, _ := NewSteamID("STEAM_0:1:40225689")

	if !sid.IsValid() {
		t.Fail()
	}
}

func TestCanValidateSteamID3(t *testing.T) {
	sid, _ := NewSteamID("[U:1:80451379]")

	if !sid.IsValid() {
		t.Fail()
	}
}

func TestCanValidateSteamID64(t *testing.T) {
	sid, _ := NewSteamID(76561198040717107)

	if !sid.IsValid() {
		t.Fail()
	}
}

func TestCanGetSteam2RenderedIDWithOldFormat(t *testing.T) {
	sid, _ := NewSteamID(76561198040717107)

	if sid.GetSteam2RenderedID() != "STEAM_0:1:40225689" {
		t.Errorf("Should have got %v got %v insted", "STEAM_0:1:40225689", sid.GetSteam2RenderedID())
	}
}

func TestCanGetSteam2RenderedIDWithNewFormat(t *testing.T) {
	sid, _ := NewSteamID(76561198040717107)

	if sid.GetSteam2RenderedID(true) != "STEAM_1:1:40225689" {
		t.Errorf("Should have got %v got %v insted", "STEAM_1:1:40225689", sid.GetSteam2RenderedID())
	}
}

func TestCanGetSteam3RenderedID(t *testing.T) {
	sid, _ := NewSteamID(76561198040717107)

	if sid.GetSteam3RenderedID() != "[U:1:80451379]" {
		t.Errorf("Should have got %v got %v insted", "[U:1:80451379]", sid.GetSteam3RenderedID())
	}
}