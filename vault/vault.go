package vault

import (
	"io/ioutil"
	"os"

	"github.com/fsnotify/fsnotify"
)

type AuditData struct {
	auditContent []byte
}

var CH chan interface{}
var filePath = "/Users/aly/work/vault/audit/audit.log"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (ad *AuditData) auditFileCheck() []byte {
	_, err := os.Stat(filePath)
	check(err)
	ad.auditContent, err = ioutil.ReadFile(filePath)
	check(err)
	return ad.auditContent
}

func truncateFile(filepath string) {
	if err := os.Truncate(filePath, 0); err != nil {
		check(err)
	}
}

// WatchFile is a function to work on runtime to monitor any change
// on the FilePath and print as STDOUT it then truncate the file content to keep it clean
func WatchFile() {
	// create instance from the AuditData
	audit := &AuditData{}
	watcher, err := fsnotify.NewWatcher()
	check(err)
	defer watcher.Close()
	// starting to listen on events in runtime
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					ad := audit.auditFileCheck()
					CH <- ad
					defer close(CH)
				}
			case err := <-watcher.Errors:
				check(err)
			}
		}
	}()
	err = watcher.Add(filePath)
	check(err)
	<-make(chan struct{})
}
