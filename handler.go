package schuko

import (
	"code.google.com/p/go.net/websocket"
	"github.com/garyburd/redigo/redis"
	"github.com/jcelliott/lumber"
)

var (
	RedisUrl               = ":6379"
	Log      lumber.Logger = lumber.NewConsoleLogger(lumber.INFO)
)

func NewHandler() websocket.Handler {
	return websocket.Handler(receiver)
}

func receiver(ws *websocket.Conn) {
	path := ws.Config().Location.Path
	Log.Info("WebSocket connection established. Origin: %s Path: %s", ws.Config().Origin, path)

	// connect to redis
	c, err := redis.Dial("tcp", RedisUrl)
	if err == nil {
		Log.Info("Connected to redis on %s", RedisUrl)
		defer c.Close()
	} else {
		Log.Error("ERROR: Failed to connect to redis on %s", RedisUrl)
	}

	// subscribe to channel based on path
	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe(path)

	// wait for message; return when client gone
	for {
		switch n := psc.Receive().(type) {
		case redis.Message:
			msg := string(n.Data)
			err := websocket.Message.Send(ws, msg)
			if err == nil {
				Log.Info("Message sent! %s", msg)
			} else {
				Log.Error("ERROR: %s", err)
				// error sending message - client gone
				// so return from goroutine
				return
			}
		case error:
			Log.Error("ERROR: Unknown")
			return
		}
	}

}
