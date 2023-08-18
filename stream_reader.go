package echoopenai

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type streamable interface {
	ChatCompletionStreamResponse
}

type streamReader[T streamable] struct {
	emptyMessagesLimit uint
	isFinished         bool

	scanner        *bufio.Scanner
	response       *http.Response
	errAccumulator ErrorAccumulator
	marshaler      Marshaller
}

func (stream *streamReader[T]) Recv() (response T, err error) {
	if stream.isFinished {
		err = io.EOF
		return
	}

	response, err = stream.scanLines()
	return
}

func (stream *streamReader[T]) scanLines() (response T, err error) {
	var emptyMessagesCount uint

	for stream.scanner.Scan() {
		rawLine, scanErr := stream.scanner.Bytes(), stream.scanner.Err()
		if scanErr != nil {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return *new(T), fmt.Errorf("error, %w", respErr.Error)
			}
			return *new(T), scanErr
		}

		if len(rawLine) == 0 {
			stream.emptyMessagesLimit++
		}

		if emptyMessagesCount > stream.emptyMessagesLimit {
			return *new(T), ErrTooManyEmptyStreamMessages
		}

		headerBytes := []byte("data: ")
		trimmedSpaceLine := bytes.TrimSpace(rawLine)

		if !bytes.HasPrefix(trimmedSpaceLine, headerBytes) {
			writeErr := stream.errAccumulator.Write(trimmedSpaceLine)
			if writeErr != nil {
				return *new(T), writeErr
			}

			continue
		}

		noPrefixLine := bytes.TrimPrefix(trimmedSpaceLine, headerBytes)
		if string(noPrefixLine) == "[DONE]" {
			stream.isFinished = true
			return *new(T), io.EOF
		}

		unmarshalErr := stream.marshaler.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return *new(T), unmarshalErr
		}

		return response, nil
	}
	return
}

func (stream *streamReader[T]) unmarshalError() (errResp *ErrorResponse) {
	errBytes := stream.errAccumulator.Bytes()
	if len(errBytes) == 0 {
		return
	}

	err := stream.marshaler.Unmarshal(errBytes, &errResp)
	if err != nil {
		errResp = nil
	}

	return
}

func (stream *streamReader[T]) Close() {
	stream.response.Body.Close()
}
