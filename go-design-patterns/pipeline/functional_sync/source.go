package functional_sync

import (
	"bufio"
	"context"
	"io"
	"log"
	"strings"
	"time"
)

// 产生一次数据
func dataSource(ctx context.Context) (chan any, error) {
	out := make(chan any)

	go func() {
		defer close(out)

		r := bufio.NewReader(strings.NewReader(LINES))
		input, err := io.ReadAll(r)
		if err != nil {
			log.Printf("Read error %v\n", err.Error())
			return
		}

		out <- string(input)
	}()

	return out, nil
}

func dataTimerSource(ctx context.Context) (chan any, error) {
	out := make(chan any)

	go func() {
		defer close(out)

		for {
			select {
			case <-time.After(time.Second):
				r := bufio.NewReader(strings.NewReader(LINES))
				input, err := io.ReadAll(r)
				if err != nil {
					log.Printf("Read error %v\n", err.Error())
					return
				}

				out <- string(input)

			case <-ctx.Done():
				log.Println("Receive exit msg!")
				return
			}
		}
	}()

	return out, nil
}
