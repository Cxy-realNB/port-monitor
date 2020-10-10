package argsparse

import "flag"

type InitArgs struct {
	IP      string
	Port    int
	LogPath string
	LogReq  bool
}

var Args *InitArgs

func ParseArgs() {
	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.Int("port", 8888, "HTTP server port")
	logPath := flag.String("logPath", "./stdout.log", "Log file path")
	logReq := flag.Bool("logReq", false, "Whether log request")
	flag.Parse()
	Args = &InitArgs{
		IP:      *ip,
		Port:    *port,
		LogPath: *logPath,
		LogReq:  *logReq,
	}
}
