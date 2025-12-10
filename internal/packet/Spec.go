package packet

// Inbound:

type Spawn struct {
	Nickname string
}

type Spectate struct{}

type MouseMove struct {
	X, Y int32
}

type Split struct{}

type SpectateFreeRoam struct{}  // ActionKeyPressed
type ActionKeyReleased struct{} // ActionKeyReleased

type EjectMass struct{}

// TODO: Outbound Packet to broadcast player actions (minion control, duet swap, etc)

// Outbound:

type UpdateNodes struct {
	EatRecordLen uint16
	EatRecords   []struct {
		HunterID uint32
		PreyID   uint32
	}
	RemoveRecordLen uint16
	RemoveRecords   []struct {
		CellID uint32
	}
}

type SpectatorUpdate struct {
	X, Y, ZoomFactor float32
}

type ClearAllNodes struct{}

type ClearMyNodes struct{}

// Drawn from all player cells to position
type DrawLine struct {
	X, Y uint16 // ?? why uint??
}

// Nodes added by this packet are centered on the client's camera.
type AddNode struct {
	NodeID uint32
}

type SetBorder struct {
	Left, Top, Right, Bottom float64
	Mode                     uint32
	ServerAlias              string
}

type StartingGame struct {
	TimeToStart uint32
}

type StartGame struct{}

type UpdateGame struct {
	AlivePlayerCount uint16
	Status           uint16
}

// All types but 0 have 1 field with a string of the killed nickname (type 0 has 1 more, killer nickname)
const (
	PlayerDeathTypePlayerAtePlayer      = 0
	PlayerDeathTypeVirusAtePlayer       = 1
	PlayerDeathTypePlayerCouldNotEscape = 2
	PlayerDeathTypePlayerDied           = 3
)

type PlayerDeath struct {
	Type uint8
}

type BattleResult struct {
	FinalPosition uint32
	TotalKills    uint32
	PlayerCount   uint16
	Players       []struct {
		Nickname string
		Position uint32
	}
}

type UpdateMiniMap struct {
}
