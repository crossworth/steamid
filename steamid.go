package steamid

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	UniverseInvalid = iota
	UniversePublic
	UniverseBeta
	UniverseInternal
	UniverseDev
)

const (
	TypeInvalid = iota
	TypeIndividual
	TypeMultiSeat
	TypeGameServer
	TypeAnonGameServer
	TypePending
	TypeContentServer
	TypeClan
	TypeChat
	TypeP2PSuperSeeder
	TypeAnonUser
)

const (
	InstanceAll = iota
	InstanceDesktop
	InstanceConsole
	InstanceWeb
)

const (
	AccountIDMask = 0xFFFFFFFF
	AccountInstanceMask = 0x000FFFFF
)

const (
	ChatInstanceFlagClan = (AccountInstanceMask + 1) >> 1
	ChatInstanceFlagLobby = (AccountInstanceMask + 1) >> 2
	ChatInstanceFlagMMSLobby = (AccountInstanceMask + 1) >> 3
)

type SteamID struct {
	Universe int
	Type int
	Instance int
	AccountID int
}

func NewSteamID(input ...interface{}) (SteamID, error) {
	steamID := SteamID{UniverseInvalid, TypeInvalid, InstanceAll, 0}

	if len(input) == 0 {
		return steamID, nil
	}

	idstr := fmt.Sprintf("%v", input[0])

	id2 := regexp.MustCompile("^STEAM_([0-5]):([0-1]):([0-9]+)$")

	if id2.MatchString(idstr) {
		sidid2 := id2.FindStringSubmatch(idstr)
		universe, _ := strconv.Atoi(sidid2[1])
		accountId, _ := strconv.Atoi(sidid2[3])
		authServer, _ := strconv.Atoi(sidid2[2])

		if universe == 0 {
			universe = UniversePublic
		}

		steamID.Universe = universe
		steamID.Type = TypeIndividual
		steamID.Instance = InstanceDesktop
		steamID.AccountID = accountId << 1 | authServer
		return steamID, nil
	}

	id3 := regexp.MustCompile(`^\[([a-zA-Z]):([0-5]):([0-9]+)(:[0-9]+)?\]$`)

	if id3.MatchString(idstr) {
		sidid3 := id3.FindStringSubmatch(idstr)
		universe, _ := strconv.Atoi(sidid3[2])
		accountId, _ := strconv.Atoi(sidid3[3])

		typeChar := sidid3[1]

		if sidid3[4] != "" {
			steamID.Instance, _ = strconv.Atoi(sidid3[4])
		} else if typeChar == "U" {
			steamID.Instance = InstanceDesktop
		}

		steamID.Universe = universe
		steamID.AccountID = accountId

		switch typeChar {
		case "c":
			steamID.Instance |= ChatInstanceFlagClan
			steamID.Type = TypeChat
		case "L":
			steamID.Instance |= ChatInstanceFlagLobby
			steamID.Type = TypeChat
		case "I":
			steamID.Type = TypeInvalid
		case "U":
			steamID.Type = TypeIndividual
		case "M":
			steamID.Type = TypeMultiSeat
		case "G":
			steamID.Type = TypeGameServer
		case "A":
			steamID.Type = TypeAnonGameServer
		case "P":
			steamID.Type = TypePending
		case "C":
			steamID.Type = TypeContentServer
		case "g":
			steamID.Type = TypeClan
		case "T":
			steamID.Type = TypeChat
		case "a":
			steamID.Type = TypeAnonUser
		default:
			steamID.Type = TypeInvalid
		}
		return steamID, nil
	}

	accountId, err := strconv.ParseInt(idstr, 10, 64)

	if err == nil {
		steamID.AccountID = int(accountId & 0xFFFFFFFF >> 0)
		steamID.Instance = int((accountId >> 32) & 0xFFFFF)
		steamID.Type = int((accountId >> 52) & 0xF)
		steamID.Universe = int((accountId >> 56) & 0xFF)
		return steamID, nil
	}

	return steamID, errors.New("invalid input")
}

func (sid *SteamID) IsValid() bool {

	if sid.Type <= TypeInvalid || sid.Type > TypeAnonUser {
		return false
	}

	if sid.Universe <= UniverseInvalid || sid.Universe > UniverseDev {
		return false
	}

	if sid.Type == TypeIndividual && (sid.AccountID == 0 || sid.Instance > InstanceWeb) {
		return false
	}

	if sid.Type == TypeClan && (sid.AccountID == 0 || sid.Instance != InstanceAll) {
		return false
	}

	if sid.Type == TypeGameServer && sid.AccountID == 0 {
		return false
	}

	return true
}

func (sid *SteamID) IsGroupChat() bool {
	return !!(sid.Type == TypeChat && (sid.Instance & ChatInstanceFlagClan) != 0)
}

func (sid *SteamID) IsLobby() bool {
	return !!(sid.Type == TypeChat &&
		((sid.Instance & ChatInstanceFlagLobby) != 0 ||
			(sid.Instance & ChatInstanceFlagMMSLobby) != 0))
}

func (sid *SteamID) GetSteam2RenderedID(newerFormat ...bool) string {
	universe := 0

	if len(newerFormat) > 0 && newerFormat[0] && universe == 0 {
		universe = 1
	}

	return "STEAM_" +  strconv.Itoa(universe) + ":" + strconv.Itoa(sid.AccountID & 1) + ":" + strconv.Itoa(sid.AccountID / 2)
}

func (sid *SteamID) GetSteam3RenderedID() string {
	typeChar := "i"

	switch sid.Type {
	case TypeInvalid:
		typeChar = "i"
	case TypeIndividual:
		typeChar = "U"
	case TypeMultiSeat:
		typeChar = "M"
	case TypeGameServer:
		typeChar = "G"
	case TypeAnonGameServer:
		typeChar = "A"
	case TypePending:
		typeChar = "P"
	case TypeContentServer:
		typeChar = "C"
	case TypeClan:
		typeChar = "g"
	case TypeChat:
		typeChar = "T"
	case TypeAnonUser:
		typeChar = "a"
	}

	if sid.Instance & ChatInstanceFlagClan != 0 {
		typeChar = "c"
	} else if sid.Instance & ChatInstanceFlagLobby != 0 {
		typeChar = "L"
	}

	renderInstance := sid.Type == TypeAnonGameServer ||
		sid.Type == TypeMultiSeat ||
		(sid.Type == TypeIndividual && sid.Instance != InstanceDesktop);

	instance := ""

	if renderInstance {
		instance = ":" + strconv.Itoa(sid.Instance)
	}

	return "[" + typeChar + ":" + strconv.Itoa(sid.Universe) + ":" + strconv.Itoa(sid.AccountID) + instance + "]"
}

func (sid *SteamID) GetSteamID64() uint64 {
	highPart := uint64((sid.Universe << 24) | (sid.Type << 20) | sid.Universe) << 32
	return highPart | uint64(sid.AccountID)
}

func (sid *SteamID) String() string {
	return strconv.FormatUint(sid.GetSteamID64(), 10)
}