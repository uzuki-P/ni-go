package main

import (
	"flag"
	"fmt"

	"github.com/uzuki-P/ni-go/pkg"
)

func main() {
	pkg.GetPackageManager()
	// packageManager, err := pkg.GetPackageManager()
	// if err != nil {
	// 	return
	// }

	var global bool
	flag.BoolVar(&global, "g", false, "Whether use global install or not")

	flag.Parse()

	fmt.Printf("use global %t! ", global)

	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 0 {
		fmt.Println("\nNon-flag arguments:")
		for i, arg := range nonFlagArgs {
			fmt.Printf("  Arg %d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("\nNo non-flag arguments.")
	}
}
