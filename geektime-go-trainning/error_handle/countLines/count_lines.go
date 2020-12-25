package countLines

import (
	"bufio"
	"io"
)

func CountLinesGood(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}

	return lines, sc.Err()
}

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}
