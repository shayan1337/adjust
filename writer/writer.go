package writer

import "fmt"

type StdOutWriter struct {
}

func NewStdOutWriter() *StdOutWriter {
	return &StdOutWriter{}
}

func (this *StdOutWriter)Write(url string, value string) {
	fmt.Println(url, value)
}
