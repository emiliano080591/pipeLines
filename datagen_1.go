package main

import (
	"bufio"
	"github.com/google/uuid"
	"io"
	"os"
	"strings"
)

const(
	filePath="guids.txt"
)

func generateData() <-chan uuid.UUID {
	c:=make(chan uuid.UUID)

	go func() {
		file,_:=os.Open(filePath)
		defer file.Close()

		reader:=bufio.NewReader(file)
		for  {
			line,err:=reader.ReadString('\n')
			if err==io.EOF {
				break
			}

			line=strings.TrimSuffix(line,"\n")
			guid,err:=uuid.Parse(line)

			if err != nil {
				continue
			}
			c<-guid
		}
		close(c)
	}()
	return c
}
