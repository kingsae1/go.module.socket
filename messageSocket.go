package messageSocket

import (
	socketio "github.com/googollee/go-socket.io"
)

var Server socketio.Server

func init() {
	Server := socketio.NewServer(nil)

	Server.Adapter(&socketio.RedisAdapterOptions{
		Host:   "127.0.0.1",
		Port:   "3010",
		Prefix: "socket.io",
	})

	// Server.OnConnect("/", func(s socketio.Conn) error {
	// 	s.SetContext("")
	// 	// fmt.Println("connected:", s.ID())
	// 	return nil
	// })

	// server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
	// 	s.SetContext(msg)
	// 	return "recv " + msg
	// })

	// server.OnEvent("/", "bye", func(s socketio.Conn) string {
	// 	last := s.Context().(string)
	// 	s.Emit("bye", last)
	// 	s.Close()
	// 	return last
	// })

	// server.OnError("/", func(s socketio.Conn, e error) {
	// 	fmt.Println("meet error:", e)
	// })

	// server.OnDisconnect("/", func(s socketio.Conn, reason string) {
	// 		fmt.Println("closed", reason)
	// })

	// go Server.Serve()
	// defer Server.Close()

	// http.Handle("/socket.io/", Server)
	// http.Handle("/", http.FileServer(http.Dir("./asset")))
	// log.Println("Serving at localhost:8000...")
	// log.Fatal(http.ListenAndServe(":8000", nil))
}

func GetServer() socketio.Server {
	return Server
}
