# go-mure

[![Build Status](https://drone.io/github.com/nowk/go-mure/status.png)](https://drone.io/github.com/nowk/go-mure/latest)
[![Build Status](https://travis-ci.org/nowk/go-mure.svg?branch=master)](https://travis-ci.org/nowk/go-mure)

Multiple file readers to channel

## Example

    readers := mure.NewReaders("file-1.txt", "file-2.txt", "file-3.txt")
    ch, er := readers.Subscribe()

    for {
      select {
        case r, ok := <-ch:
          if !ok {
            return // done
          }

          out, err := os.Create(fmt.Sprintf("out_%s", r.Name()))
          if err != nil {
            panic(err)
          }

          io.Copy(out, r)
          out.Close()
        case err := <-er:
          log.Print("error:", err)
      }
    }

### License

MIT
