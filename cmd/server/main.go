package main

import (
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/cellkit/server/internal/game"
	"github.com/cellkit/server/internal/geometry"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	ServerBoundWidth  = 8192
	ServerBoundHeight = 8192
	ServerRegionSize  = 64
)

func main() {
	grid := geometry.NewSpatialGrid(ServerBoundWidth, ServerBoundHeight, ServerRegionSize)

	gameInstance := game.Game{
		Grid:    grid,
		Sockets: make(map[net.Conn]*game.Socket),
	}

	gameInstance.StartLoop()

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer conn.Close()

			gameInstance.OnWebsocketOpen(conn)

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					var closeErr wsutil.ClosedError
					if errors.As(err, &closeErr) {
						gameInstance.OnWebsocketClose(conn, closeErr.Code)
						return
					}
					log.Fatal(err)
				}

				switch op {
				case ws.OpBinary:
					gameInstance.OnWebsocketMessage(conn, msg)
				case ws.OpClose:
					// gameInstance.OnWebsocketClose(conn)
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
