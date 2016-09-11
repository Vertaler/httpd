package main

import (
	"flag"
	"runtime"
	"./server"
)
func main() {
	 server := server.Server{}
	 server.CpuNumber = *flag.Int("c", runtime.NumCPU(), "-c NCPU")
	 server.Port = *flag.String("p", "8080", "-p PORT")
	 server.Root = *flag.String("r", "./", "-r ROOTDIR")
	 server.Host = *flag.String("h","localhost","-h HOST")
	 server.LogEnable = *flag.Bool("e", false, "-e")
	 server.LogPath = *flag.String("l","./","-l LOGPATH")
	 server.Start()
}
