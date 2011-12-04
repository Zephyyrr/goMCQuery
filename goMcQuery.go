package mcQuery

import (
	"net"
	"fmt"
//	"os"
	"strings"
	com "./commons"
)

const (
	Magic1 byte = 0xFE
	Magic2 byte = 0xFD
	Challenge byte = 0x09
	Request byte = 0x00
)

var(
	debug = 0
)

/**
* Sätter debug variabeln för loggning.
*/
func SetDebugLevel(level int) {
  debug = level
}

func Test(con net.Conn, id int32) {
	chall := GetChallengeCode(con, id)
	ms := GetShortPack(con, id, chall)
	fmt.Println(ms)
	fmt.Println("\n\n")
	ml, pl := GetLongPack(con, id, chall)
	fmt.Println(ml)
	fmt.Println("\n\n")
	fmt.Println(pl)
}

/**
 * This function connects to a given adress
 */
func Connect(adress string) net.Conn {
	com.Log(debug, "connecting to", adress);
	//dest, err1 := net.ResolveUDPAddr("udp", adress)
	cn, err2 := net.Dial("udp", adress)
//	com.Test(err1, "resolving " + adress);
	com.Test(err2, "dialing " + adress);
	cn.SetTimeout(3e9)
	return cn
}

func read(buf []byte, con net.Conn) {
	n, err := con.Read(buf)
	if (err != nil && err.String() == "Timeout() == true") {
		com.Log(6, "Read time-out reached. Read " + fmt.Sprint(n) + " bytes.")
	}
	com.Log(8, "Answer recived!")
	//fmt.Println(buf)
}

func write(buf []byte, con net.Conn) {
	n, err := con.Write(buf)
	if (err != nil && err.String() == "Timeout() == true") {
		com.Log(6, "Write time-out reached. Written " + fmt.Sprint(n) + " bytes.")
	}
	com.Log(8, "Package sent!")
	//fmt.Println(buf)
}

func GetChallengeCode(con net.Conn, id int32) int32 {
	var buffer [32]byte
	com.Log(debug, "Sending Testpackage: Challenge Request")
	b := []byte{Magic1, Magic2, Challenge}
	b = com.AppendN(b, []uint8(com.Int32toByteSlice(id)))
	//fmt.Println(b)
	write(b, con)
	com.Log(7, "Challenge Request package sent")
	read(buffer[:], con)
	res := ParseChall(buffer[5:32])
	return res
}

func ParseChall(bs []byte) int32 {
	i := 0
	for i < len(bs){
		if (bs[i] == 0x00) {
			//bs[i]-uint8(18)
			break
		} else {
			i++
		}
	}
	return com.Str2int32(string(bs))
}

func GetShortPack(con net.Conn, id int32, chall int32) (map[string] string) {
	var buffer [128]byte
	com.Log(7, "Sending Testpackage: Short Status Request")
	b := []byte{Magic1, Magic2, Request}
	b = com.AppendN(b, []uint8(com.Int32toByteSlice(id)))
	b = com.AppendN(b, []uint8(com.Int32toByteSlice(chall)))
	//fmt.Println(b)
	write(b, con)
	com.Log(8, "Short Status request package sent")
	read(buffer[:], con)
	shortmap := ParseShortPackage(buffer[5:])
	return shortmap
}

func ParseShortPackage(bs []byte) (map[string] string) {
	result := make(map[string] string)
	s, rest := com.GetString(bs)
	result["motd"] = s
	s, rest = com.GetString(rest)
	result["gametype"] = s
	s, rest = com.GetString(rest)
	result["map"] = s
	s, rest = com.GetString(rest)
	result["numplayers"] = s
	s, rest = com.GetString(rest)
	result["maxplayers"] = s
	sh, rest := com.GetShort(rest)
	result["hostport"] = fmt.Sprint(sh)
	s, rest = com.GetString(rest)
	result["hostname"] = s
	return result
}

func GetLongPack(con net.Conn, id int32, chall int32) ((map[string] string), []string) {
	var buffer [1024]byte
	com.Log(7, "Sending Testpackage: Long Status Request")
	b := []byte{Magic1, Magic2, Request}
	b = com.AppendN(b, []uint8(com.Int32toByteSlice(id)))
	b = com.AppendN(b, []uint8(com.Int32toByteSlice(chall)))
	b = com.AppendN(b, []byte{0, 0, 0, 0})
	//fmt.Println(b)
	write(b, con)
	com.Log(8, "Long Status request package sent")
	read(buffer[:], con)
	longmap, plist := ParseLongPackage(buffer[16:])
	return longmap, plist
}

func ParseLongPackage(bs []byte) ((map[string] string), []string) {
	result := make(map[string] string)
	as := strings.Split(string(bs), "\x00\x01player_\x00\x00")
	aas := strings.Split(as[0], "\x00")
	//fmt.Println(aas)
	for i := 0; i < (len(aas)-1); i+=2 {
		//fmt.Println([]byte(aas[i]), []byte(aas[i+1]))
		result[aas[i]] = aas[i+1]
	}
	maxp := com.Str2int32(result["maxplayers"])
	var pList []string
	if (maxp > 0 && len(as) > 0) {
		pList = strings.Split(as[1], "\x00")
	}
	return result, pList
}
