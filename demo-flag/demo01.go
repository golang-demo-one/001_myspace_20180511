package main

import (
	"flag"
	"fmt"
)

/* https://studygolang.com/articles/4706 */
//go run pg_ctl_parse.go  -D /home/manu/DB_data/ -l /home/manu/DB_data/postgres_manu.log -W start

func main() {
	data_path := flag.String("D", "/home/menu/sample/", "DB data path")
	log_file := flag.String("l", "/home/menu/sample.log", "log file")
	nowait_flag := flag.Bool("W", false, "do not wait until operation complates")

	flag.Parse()

	var cmd string = flag.Arg(0)

	fmt.Printf("action:	%s\n", cmd)
	fmt.Printf("data path: %s\n", *data_path)
	fmt.Printf("log file: %s\n", *log_file)
	fmt.Printf("nowait: %v\n", *nowait_flag)

	fmt.Printf("-----------------------------------------\n")

	fmt.Printf("there are %d non-flag input params\n", flag.NArg())
	for i, param := range flag.Args() {
		fmt.Printf("#%d    :%s\n", i, param)
	}
}
