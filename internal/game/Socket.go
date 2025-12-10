package game

import (
	"fmt"
	"image/color"
	"net"
	"strings"
	"sync"

	"github.com/cellkit/server/internal/entity"
	"github.com/cellkit/server/internal/geometry"
	"github.com/cellkit/server/internal/packet"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	PLAYER_START_MASS            uint32 = 780
	MOCK_RANDOM_START_POSITION_X int32  = 43
	MOCK_RANDOM_START_POSITION_Y int32  = 456
)

type Socket struct {
	Conn     net.Conn
	Entities []entity.Entity
	Color    color.Color
	nickname string
	mu       sync.RWMutex
}

func NewSocket(conn net.Conn) *Socket {
	return &Socket{
		Conn:     conn,
		Color:    color.RGBA{125, 125, 125, 255},
		nickname: "An unnamed cell",
	}
}

func (s *Socket) SetNickname(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nickname = name
}

func (s *Socket) GetNickname() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.nickname
}

func (s *Socket) OnSpawn(nickname string) {
	if strings.TrimSpace(nickname) != "" {
		s.SetNickname(nickname)
	} else {
		s.SetNickname("An unnamed cell")
	}

	s.Color = randomColor()

	playerCell := &entity.PlayerCell{}
	playerCell.SetMass(PLAYER_START_MASS)
	playerCell.SetPosition(geometry.Vector2D{
		X: MOCK_RANDOM_START_POSITION_X,
		Y: MOCK_RANDOM_START_POSITION_Y,
	})
	playerCell.SetColor(s.Color)

	s.Entities = []entity.Entity{playerCell}
	// s.Entities = append(s.Entities, playerCell)

	fmt.Printf("Player spawn packet - posX: %d, posY: %d | mass: %d", playerCell.GetPosition().X, playerCell.GetPosition().Y, playerCell.GetMass())
}

func (s *Socket) SendPacket(id byte, p packet.Packet) error {
	buf := append([]byte{id}, p.Build()...)
	err := wsutil.WriteServerMessage(s.Conn, ws.OpBinary, buf)
	return err
}
