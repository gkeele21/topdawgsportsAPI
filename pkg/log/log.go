package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/MordFustang21/nova"
)

var l = log.New(os.Stderr, "", 0)

// Info defines the json structure for the log
type Info struct {
	Host    string    `json:"host"`
	Message string    `json:"message"`
	Errors  []string  `json:"errors"`
	Time    time.Time `json:"time"`
	LogType string    `json:"type"`
}

// Println creates json formatted log for use in logstash
func Println(logType, message string, args ...interface{}) {
	info := Info{
		Message: message,
	}

	for _, arg := range args {
		info.Errors = append(info.Errors, fmt.Sprintf("%v", arg))
	}

	host, _ := os.Hostname()
	info.Host = host

	info.Time = time.Now()
	info.LogType = logType

	data, err := json.Marshal(info)
	if err != nil {
		fmt.Println("couldn't json info:", err)
	}

	l.Println(string(data))
}

func LogRequest(req *nova.Request) {
	fmt.Printf("%s : %#v\n", time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006"), req)
}

func LogRequestData(object interface{}) {
	fmt.Printf("%s : Data Passed : %#v\n", time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006"), object)
}
