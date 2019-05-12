package auditlog

import (
	"os"
)

var logFileName = "./run.log"

//TODO: consider concurrent write log, it's not so easy.
func logToFile(msg []byte) (err error) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	msg = append(msg, '\n')
	finfo, err := f.Stat()
	if err != nil {
		return
	}
	_, err = f.WriteAt(msg, finfo.Size())
	return
}
