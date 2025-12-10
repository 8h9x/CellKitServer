package main

import (
	"net"
	"net/http"

	"github.com/cellkit/server/internal/game"
	"github.com/cellkit/server/internal/geometry"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	SERVER_BOUND_WIDTH  = 8192
	SERVER_BOUND_HEIGHT = 8192
)

func main() {
	grid := geometry.NewSpatialGrid(SERVER_BOUND_WIDTH, SERVER_BOUND_HEIGHT, 64)

	gameInstance := game.Game{
		Grid:    grid,
		Sockets: make(map[net.Conn]*game.Socket),
	}

	gameInstance.StartLoop()

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go func() {
			defer conn.Close()

			gameInstance.OnWebsocketOpen(conn)

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				switch op {
				case ws.OpBinary:
					gameInstance.OnWebsocketMessage(conn, msg)
				case ws.OpClose:
					gameInstance.OnWebsocketClose(conn)
				default:
				}
				if err != nil {
					// log.Println("socket error", err)
					// handle error
				}
			}
		}()
	}))
}
