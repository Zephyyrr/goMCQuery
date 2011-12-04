/*
 C opyright 2011 Johan "Zephyyrr" *Fogelström
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
 http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.*/

package main

import (
	mcq "./goMcQuery"
	com "./commons"
	"flag"
	"net"
	"fmt"
	"time"
)

var (
	addr = flag.String("a", "localhost", "The adress you wish to connect to. (default = localhost)")
	port = flag.String("p", "25566", "The port the server is running on. (default = 25566")
	debug = flag.Int("d", 0, "The debug level, higher level mean more output. (default = 0)")
	
	id int32 = 1
	chall int32
)

func main() {
	flag.Parse()
	com.SetDebugLevel(*debug)
	mcq.SetDebugLevel(*debug)
	con := mcq.Connect(*addr + ":" + *port)
//	sendTest(con)
	mcq.Test(con, id)
	time.Sleep(2e9)
}

func sendTest(con net.Conn) {
	chall = mcq.GetChallengeCode(con, id)
	//ms := mcq.GetShortPack(con, id, chall)
	ml, pl := mcq.GetLongPack(con, id, chall)
//	fmt.Println(ms)
	fmt.Println(ml)
	fmt.Println(pl)
}
