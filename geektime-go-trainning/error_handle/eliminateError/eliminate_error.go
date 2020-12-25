package eliminateError

import (
	"fmt"
	"io"
)

type Header struct {
	Key   string
	Value string
}

type Status struct {
	code   int
	reason string
}

type errWriter struct {
	writer io.Writer
	err    error
}

func (ew *errWriter) Write(p []byte) (int, error) {
	//在写入操作之前，先判断上一次操作是否成功
	if ew.err != nil {
		return 0, ew.err
	}

	var written int
	//将错误保存起来
	written, ew.err = ew.writer.Write(p)
	return written, nil
}

//能有效减少错误的写法
func WriteResponseGood(w io.Writer, s Status, headers []Header, body io.Reader) error {
	ew := &errWriter{writer: w}
	fmt.Fprintf(ew, "Http /1.1 %d %s\r\n", s.code, s.reason)

	for _, header := range headers {
		fmt.Fprintf(w, "%s: %s\r\n", header.Key, header.Value)
	}

	fmt.Fprintf(w, "\r\n")

	io.Copy(w, body)

	return ew.err
}

//一般写法
func WriteResponse(w io.Writer, s Status, headers []Header, body io.Reader) error {
	var err error
	_, err = fmt.Fprintf(w, "Http /1.1 %d %s\r\n", s.code, s.reason)
	if err != nil {
		return err
	}

	//write headers
	for _, header := range headers {
		_, err = fmt.Fprintf(w, "%s: %s\r\n", header.Key, header.Value)
		if err != nil {
			return err
		}
	}
	//write \r\n
	_, err = fmt.Fprintf(w, "\r\n")
	if err != nil {
		return err
	}

	//write body
	_, err = io.Copy(w, body)
	return err
}
