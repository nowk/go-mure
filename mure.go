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

// Subscribe returns a .Reader and error channel then begins the readpiping
func (self *Readers) Subscribe() (<-chan Reader, <-chan error) {
	ch := make(chan Reader)
	er := make(chan error)

	for _, file := range self.Files {
		go self.read(file, ch, er)
	}

	return ch, er
}

// read opens a file and pipes back to the .Reader channel
// This can call the error channel more than once
func (self *Readers) read(file string, ch chan<- Reader, er chan<- error) {
	fi, err := os.Open(file)
	if err != nil {
		er <- err
		return
	}
	defer fi.Close()
	// ch <- bufio.NewReader(fi) // need to close file

	r, w := io.Pipe()
	defer w.Close()

	stat, err := fi.Stat()
	if err != nil {
		er <- err
		return
	}

	ch <- Reader{r, stat.Name(), stat.Size()}
	if _, err := io.Copy(w, fi); err != nil {
		er <- err
	}
}
