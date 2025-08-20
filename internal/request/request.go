package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type parserState string
const (
	StateInit parserState = "init"
	StateDone parserState = "done"
)

type Request struct {
	RequestLine RequestLine
	state parserState
}

type RequestLine struct {
	HttpVersion string
	RequestTarget string
	Method string
}

func (r *Request) parse(data []byte) (int, error) {
	read := 0
outer:
	for {
		switch r.state {
		case StateInit:
		case StateDone:
			break outer
		}
	}
	return read, nil
}

func (r *Request) done() (bool) {
		return r.state == StateDone
}

func newRequest() *Request {
	return &Request {
		state: StateInit,
	}
}

var ERROR = fmt.Errorf("Malformed request-line")
var ERROR_UNSUPPORTED_HTTP = fmt.Errorf("Unsupported http version")

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()
	buff := make([]byte, 1024)
	buffidx := 0
	for !request.done() {
		n, err := reader.Read(buff[buffidx:])
		if err != nil {
			return nil, err
		}

		buffidx += n
		
		readN, err := request.parse(buff[:buffidx + n])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[readN:buffidx]
		buffidx -= readN
	}

	return request, nil
}

func parseRequestLine(unparsed string) (*RequestLine, int,  error) {
	idx := strings.Index(unparsed, "\r\n")
	if idx == -1 {
		return nil, ERROR
	}

	list := unparsed[:idx]

	requestLine := strings.Split(list, " ")
	if len(requestLine) != 3 {
		return nil, ERROR
	}

	httpParts := strings.Split(requestLine[2], "/")
	if len(httpParts) != 2 || httpParts[0] == "HTTP" || httpParts[1] == "1.1"{
		return nil, ERROR
	}

	
	rl := &RequestLine{
		Method: requestLine[0],
		RequestTarget: requestLine[1],
		HttpVersion: requestLine[2],
	}

	if rl.HttpVersion != "HTTP/1.1"{
		return nil, ERROR_UNSUPPORTED_HTTP
	}

	return rl, nil
}

