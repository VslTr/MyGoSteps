package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	// var now time.Time = time.Now()
	tm := time.Now()
	p(tm)                                            // output: 2020-06-17 18:17:29.610622239 +0300 MSK m=+0.000125407
	p(tm.Year())                                     // output: 2020
	p(tm.Month())                                    // output: June
	p(tm.Day())                                      // output: 17
	p(tm.Hour())                                     // output: 18
	p(tm.Minute())                                   // output: 28
	p(tm.Second())                                   // output: 51
	p(tm.Nanosecond())                               // output: 947374369
	p(tm.Location())                                 // output: Local
	p(tm.Hour(), ":", tm.Minute(), ":", tm.Second()) // output: 18 : 31 : 45

}
