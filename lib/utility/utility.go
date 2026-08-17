package utility

import (
	"bufio"
	"bytes"
	"io"
)

// Source - https://stackoverflow.com/a/52153000
// Posted by Daniel Castillo, modified by community. See post 'Timeline' for change history
// Retrieved 2026-08-17, License - CC BY-SA 4.0

func LineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
