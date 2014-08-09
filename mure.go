package mure

import (
	"io"
	"os"
)

// Reader impelements the io.Reader interface and can be set basic meta data
type Reader struct {
	io.Reader
	name string
	size int64
}

func (self Reader) Name() string {
	return self.name
}

func (self Reader) Size() int64 {
	return self.size
}

// Readers

type Readers struct {
	Files []string
}

func NewReaders(files ...string) (r *Readers) {
	r = &Readers{
		Files: files,
	}

	return
}

// Len returns the number of files to read
func (self Readers) Len() int {
	return len(self.Files)
}

// Subscribe returns a .Reader and error channel then begins the readpiping
func (self *Readers) Subscribe() (<-chan Reader, chan error) {
	n := len(self.Files)
	ch := make(chan Reader, n)
	done := make(chan error, n)

	for _, file := range self.Files {
		go self.read(file, ch, done)
	}

	return ch, done
}

// read opens a file and pipes back to the .Reader channel
func (self *Readers) read(file string, ch chan<- Reader, done chan<- error) {
	f, err := os.Open(file)

	defer func() {
		if f != nil {
			f.Close()
		}
		done <- nil
	}()

	if err != nil {
		done <- err
		return
	}

	r, w := io.Pipe()
	defer w.Close()

	stat, err := f.Stat()
	if err != nil {
		done <- err
		return
	}

	ch <- Reader{
		r,
		stat.Name(),
		stat.Size(),
	}

	if _, err := io.Copy(w, f); err != nil {
		done <- err
	}
}
