# SteamID for GO

This library provide a easy way to handle SteamID on Go, it's based on the on the [NodeJS version](https://github.com/DoctorMcKay/node-steamid) and provide a very similar API.
 

# Installation

    $ go get github.com/crossworth/steamid


# Usage

You can parse the SteamID from a SteamID2, SteamID3 or SteamID64 (as uint or string).

## Steam2 ID

```go
sid, _ := steamid.NewSteamID("STEAM_0:1:40225689")
```

## Steam3 ID

```go
sid, _ := steamid.NewSteamID("[U:1:80451379]")
```

## SteamID64

```go
sid, _ := steamid.NewSteamID("76561198040717107")
```

or

```go
sid, _ := steamid.NewSteamID(76561198040717107)
```

# Using a SteamID

We provide the same methods as the NodeJS version.

## isValid() bool

Returns `true` if the object represents a valid SteamID, or `false` if not.

Example:

```go
sid, _ := steamid.NewSteamID(76561198040717107)
fmt.Println(sid.IsValid())
```

## isGroupChat() bool

Returns `true` if the `type` of this SteamID is `CHAT`, and it's associated with a Steam group's chat room.

Example:

```go
sid, _ := steamid.NewSteamID(76561198040717107)
fmt.Println(sid.IsGroupChat())
```

## isLobby() bool

Returns `true` if the `type` of thie SteamID is `CHAT`, and it's associated with a Steam lobby.

Example:

```go
sid, _ := steamid.NewSteamID(76561198040717107)
fmt.Println(sid.IsLobby())
```

## getSteam2RenderedID([newerFormat]) string


Returns the Steam2 rendered ID format for individual accounts. 

If you pass `true` for `newerFormat`, the first digit will be 1 instead of 0 for the public universe.


Example:

```go
sid, _ := steamid.NewSteamID(76561198040717107)
fmt.Println(sid.getSteam2RenderedID())
```

or


```go
sid, _ := steamid.NewSteamID(76561198040717107)
fmt.Println(sid.getSteam2RenderedID(true))
```

## getSteam3RenderedID() string

Returns the Steam3 rendered ID format.

Example:

```go
sid, _ := steamid.NewSteamID(76561198040717107)
fmt.Println(sid.getSteam3RenderedID())
```

## getSteamID64() uint64

Returns the uint64 representation of the SteamID.

Example:

```go
sid, _ := steamid.NewSteamID(76561198040717107)
fmt.Println(sid.getSteamID64())
```