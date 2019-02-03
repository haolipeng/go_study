package mockretriever

import "fmt"

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	fmt.Println("mock retriever url is ", url)

	return url
}
