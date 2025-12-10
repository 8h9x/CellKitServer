package packet

type LeaderboardRecord struct {
	Flags    uint8
	ID       uint32
	Nickname string
}

type UpdateLeaderboardFFA struct {
	Records []LeaderboardRecord
}

func (p UpdateLeaderboardFFA) Build() []byte {
	writer := NewWriter(make([]byte, 5678))

	writer.WriteUint32(uint32(len(p.Records)))

	for _, record := range p.Records {
		writer.WriteUint32(record.ID)
		writer.WriteStringUTF8(record.Nickname)
	}

	return writer.Buffer()
}

type UpdateLeaderboardTeams struct {
	Records []LeaderboardRecord
}
