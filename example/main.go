package main

import (
	"fmt"
	"github.com/raspi/uint64timestamp/pkg/base10"
	"github.com/raspi/uint64timestamp/pkg/base2"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	bin, err := base2.TimeToUint64(now)
	if err != nil {
		panic(err)
	}

	binTime, err := base2.Uint64ToTime(bin, now.Location())
	if err != nil {
		panic(err)
	}

	hum, err := base10.TimeToUint64(now)
	if err != nil {
		panic(err)
	}

	humTime, err := base10.Uint64ToTime(hum, now.Location())
	if err != nil {
		panic(err)
	}

	const uint64pad = 20

	fmt.Printf(`   now: %s %v`+"\n", strings.Repeat(" ", uint64pad), now)
	fmt.Printf(` human: % *d %v`+"\n", uint64pad, hum, humTime)
	fmt.Printf(`   bin: % *d %v`+"\n", uint64pad, bin, binTime)
	fmt.Printf(`   max: 18446744073709551615` + "\n")

	fmt.Println()
	fmt.Printf(` human: % *b`+"\n", 64, hum)
	fmt.Printf(`   bin: % *b`+"\n", 64, bin)
	fmt.Printf(`   max: %s`+"\n", strings.Repeat("1", 64))

}
