package sentinelError

import "github.com/pkg/errors"

var notFound = errors.New("not found")

func test() error {
	return notFound
}
