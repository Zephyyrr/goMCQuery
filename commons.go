package commons

import (
  "log"
  "time"
  "os"
  "io"
  "fmt"
)
var (
  debug = 0
)

/**
* Logging function
*/
func Log(level int, v ... string) {
    if debug >= level {
    	ret := fmt.Sprint(v);
        log.Printf("CLIENT: " + ret);
    }
}

/**
* Testar om ett fel instäffat och utför lämplig åtgärd.
*/
func Test(err os.Error, s string) {
	if (err != nil) {
		Log(1, "Error:", s)
		time.Sleep(1e9)
		os.Exit(0)
	}
}

func Exit(code int) {
	fmt.Println("CLOSED!")
	os.Exit(code)
}

/**
* Sätter debug variabeln för loggning.
*/
func SetDebugLevel(level int) {
  debug = level
}

func Str2int32(s string) int32 {
	tio := int32(1)
	res := int32(0)
	for i := len(s)-1; i >= 0; i-- {
		if s[i] > 47 && s[i] < 58 {
			res += int32((s[i]-48))*tio
			tio = tio*10
		}
	}
	return res
}

/**
 * Appends b to a  and returns the new slice.
 */
func AppendN(a []byte, b []byte) []byte {
	for _, val := range b {
		a = append(a, val)
	}
	return a
}

func Int32toByteSlice(a int32) []byte {
	res := make([]byte, 4)
	for i := 0; i < len(res); i++ {
		res[i] = byte(a >> uint(24-i*8))// && int32(0xFF))
	}
	return res
}

func GetString(bs []byte) (string, []byte) {
	i := 0
	for i = 0; bs[i] != 0; i++ {
	}
	return string(bs[:i]), bs[i+1:]
}

func GetShort(bs []byte) (int16, []byte) {
	var res int16
	res = res | int16(bs[1])
	res = res | (int16(bs[0]) << 8)
	return res, bs[2:]
}

func PrintMap(w io.Writer, m map[string] string) {
	//TODO
}

