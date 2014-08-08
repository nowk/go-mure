# go-mure

[![Build Status](https://travis-ci.org/nowk/go-mure.svg?branch=master)](https://travis-ci.org/nowk/go-mure)

Multiple file readers to channel

## Example

    readers := mure.NewReaders("file-1.txt", "file-2.txt", "file-3.txt")
    ch, er := readers.Subscribe()

    for {
      select {
        case r := <-ch:
          // r impelements io.Reader
          //
          // provides these additional methods
          // r.Name() => the original file name 
          // r.Size() => the original file size
        case e := <-er:
          // error is returned
      }
    }

### License

MIT