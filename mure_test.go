package mure_test

import "testing"
import . "github.com/nowk/go-mure"

func TestReaders(t *testing.T) {
	readers := Readers{
		Files: []string{"test/good.txt", "bad.txt"},
	}

	ch, er := readers.Subscribe()

	for i := 0; i < 2; i++ {
		select {
		case r := <-ch:
			if val := r.Name(); "good.txt" != val {
				t.Errorf("Expected Name() to be 'good.txt', got '%s'", val)
			}

			if val := r.Size(); int64(13) != val {
				t.Errorf("Expected Size() to be 13, got %d", val)
			}

			bytes := make([]byte, 12)
			r.Read(bytes)
			if val := string(bytes); "Hello World!" != val {
				t.Errorf("Expected file contents to be 'Hello World!', got '%s'", val)
			}
		case e := <-er:
			if val := e.Error(); "open bad.txt: no such file or directory" != val {
				t.Error("Expected error to be 'open bad.txt: no such file or directory', got '%s'", val)
			}
		}
	}
}
