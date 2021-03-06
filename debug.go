// Copyright 2014 The DST Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package dst

import (
	"log"
	"os"
	"strings"
)

var (
	debugConnection bool
	debugMux        bool
	debugCC         bool
)

func init() {
	debug := make(map[string]bool)
	for _, s := range strings.Split(os.Getenv("DSTDEBUG"), ",") {
		debug[strings.TrimSpace(s)] = true
	}
	if len(debug) > 0 {
		log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	}
	debugConnection = debug["conn"]
	debugMux = debug["mux"]
	debugCC = debug["cc"]
}
