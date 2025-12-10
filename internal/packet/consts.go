package packet

const (
	InboundPacketSpawn         byte = 0
	InboundPacketSpectate      byte = 1
	InboundPacketMouseMove     byte = 16
	InboundPacketSplit         byte = 17
	InboundPacketEjectMass     byte = 21
	InboundPacketChatMessage   byte = 99
	InboundPacketHandshakeStat byte = 254
	InboundPacketHandshakeKey  byte = 255
)

const (
	OutboundPacketUpdateLeaderboardFFA   = 49
	OutboundPacketUpdateLeaderboardTeams = 50
)
