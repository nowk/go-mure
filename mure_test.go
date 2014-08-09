package mure_test

import (
	"io/ioutil"
	"testing"
)
import . "github.com/nowk/go-mure"

func TestReaders(t *testing.T) {
	readers := Readers{
		Files: []string{"test/good.txt", "bad.txt"},
	}

	ch, er := readers.Subscribe()

	i := readers.Len()
	for {
		if i == 0 {
			break
		}

		select {
		case r := <-ch:
			if val := r.Name(); "good.txt" != val {
				t.Errorf("Expected Name() to be 'good.txt', got '%s'", val)
			}

			if val := r.Size(); int64(13) != val {
				t.Errorf("Expected Size() to be 13, got %d", val)
			}

			bytes, _ := ioutil.ReadAll(r) // ensure defer gets called
			if val := string(bytes[:12]); "Hello World!" != val {
				t.Errorf("Expected file contents to be 'Hello World!', got '%s'", val)
			}
		case e := <-er:
			if e != nil {
				if val := e.Error(); "open bad.txt: no such file or directory" != val {
					t.Error("Expected error to be 'open bad.txt: no such file or directory', got '%s'", val)
				}
			} else {
				i--
			}
		}
	}
}
