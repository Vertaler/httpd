package main

import (
	"flag"
	"runtime"
	"./server"
)
func main() {
	 server := server.Server{}
	 server.CpuNumber = *flag.Int("c", runtime.NumCPU(), "-c NCPU")
	 server.Port = *flag.Int("p", 8080, "-p PORT")
	 server.Root = *flag.String("r", "/", "-r ROOTDIR")
	 server.LogEnable = *flag.Bool("e", false, "-e")
	 server.LogPath = *flag.String("l","./","-l LOGPATH")
	 server.Start()
}
