package main

import (
	"flag"
	"runtime"
	"./server"
	"path/filepath"
	"fmt"
	"os"
)
func main() {
	 server := server.Server{}
	 server.CpuNumber = *flag.Int("c", runtime.NumCPU(), "-c NCPU")
	 server.Port = *flag.String("p", "80", "-p PORT")
	 root, err := filepath.Abs( *flag.String("r", "../../http-test-suite-master/httptest", "-r ROOTDIR") )
	 if err != nil{
		 fmt.Println("Some error in root")
		 os.Exit(1)
	 }
	 server.Root = root
	 server.Host = *flag.String("h","localhost","-h HOST")
	 server.LogEnable = *flag.Bool("e", false, "-e")
	 server.LogPath = *flag.String("l","./","-l LOGPATH")
	 server.Start()
}
