package vars

import "flag"

var (
	//Host             = flag.String("H", "0.0.0.0", "hostname to listen on.")
	//Port             = flag.Int("p", 9876, "port number to listen on.")
	//Timeout          = flag.Int("t", 10, "timeout for TCP/TLS connection.")
	//MaxFdNum         = flag.Uint64("r", 0, "maximum number of file descriptors (need root priviledge).")
	//Silent			 = flag.Bool("q", false, "silent mode.")
	//Version     	 = flag.Bool("v", false, "show version and exit.")
	ProcessNum       = flag.Int("n", 5, "Number of Workers")
	Timeout          = flag.Int("t", 10, "timeout for TCP/TLS connection.")
)