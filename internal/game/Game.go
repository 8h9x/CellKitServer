package game

import (
	"fmt"
	"github.com/gobwas/ws"
	"image/color"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/cellkit/server/internal/entity"
	"github.com/cellkit/server/internal/geometry"
	"github.com/cellkit/server/internal/packet"
)

const TARGET_TPS = 25

type Game struct {
	Grid      *geometry.SpatialGrid
	Sockets   map[net.Conn]*Socket
	tickCount int
}

func (g *Game) Tick() {
	cell := entity.Cell{}
	cell.SetPosition(geometry.Vector2D{
		X: 107,
		Y: 304,
	})
	g.Grid.Insert(cell.GetPosition(), cell)

	for i := 0; i < 6; i++ {
		cell := entity.Cell{}
		cell.SetPosition(geometry.Vector2D{
			X: 9,
			Y: 18,
		})
		g.Grid.Insert(cell.GetPosition(), cell)
	}

	if (g.tickCount+7)%TARGET_TPS == 0 {
		records := make([]packet.LeaderboardRecord, 0)
		for _, socket := range g.Sockets {
			nickname := socket.GetNickname()
			log.Printf("Socket nickname: '%s' (len=%d, bytes=%v)",
				nickname, len(nickname), []byte(nickname))

			records = append(records, packet.LeaderboardRecord{
				Flags:    0,
				Nickname: nickname,
				ID:       uint32(len(records) + 1),
			})
		}

		g.BroadcastPacket(packet.OutboundPacketUpdateLeaderboardFFA, packet.UpdateLeaderboardFFA{
			Records: records,
		})
	}

	g.tickCount++
}

func (g *Game) StartLoop() {
	g.tickCount = 0
	ticker := time.NewTicker(time.Second / time.Duration(TARGET_TPS))

	go func() {
		for {
			select {
			case <-ticker.C:
				g.Tick()
				// log.Println("Tick at", t)
			}
		}
	}()
}

func (g *Game) BroadcastPacket(id byte, p packet.Packet) {
	for _, socket := range g.Sockets {
		socket.SendPacket(id, p)
	}
}

func (g *Game) OnWebsocketOpen(conn net.Conn) {
	g.Sockets[conn] = NewSocket(conn)
}

func (g *Game) OnWebsocketMessage(conn net.Conn, payload []byte) {
	socket, ok := g.Sockets[conn]
	if !ok {
		// WTF HOW MESSAGE BUT NOT JOINED????
		return
	}

	reader := packet.NewReader(payload)

	packetID := reader.ReadUint8()
	switch packetID {
	case packet.InboundPacketSpawn:
		nickname := reader.ReadString()
		g.Sockets[conn].OnSpawn(nickname)
	case packet.InboundPacketChatMessage:
		reader.ReadUint8()
		message := reader.ReadString()
		log.Println(socket.GetNickname(), message)
		g.BroadcastPacket(packet.InboundPacketChatMessage, &packet.SendChatMessage{
			Author: packet.ChatMessageAuthor{
				Nickname: socket.GetNickname(),
				Color:    socket.Color,
			},
			TextContent: message,
		})
	case packet.InboundPacketHandshakeStat:
		protocol := reader.ReadUint32()
		fmt.Printf("stat packet (254) received, protocol ver: %d\n", protocol)
		socket.SendPacket(packet.InboundPacketHandshakeStat, &MockPacket{
			Raw: payload,
		})
	case packet.InboundPacketHandshakeKey:
		protocol := reader.ReadUint32()
		fmt.Printf("protocol key packet (255) received, protocol key ver: %d\n", protocol)
	default:
		break
	}
}

func (g *Game) OnWebsocketClose(conn net.Conn, code ws.StatusCode) {
	delete(g.Sockets, conn)
}

func randomColor() color.Color {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	r := uint8(rng.Intn(256))
	g := uint8(rng.Intn(256))
	b := uint8(rng.Intn(256))

	return color.RGBA{R: r, G: g, B: b, A: 255}
}
