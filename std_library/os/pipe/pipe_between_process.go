package pipe

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPipeBetweenProcess(t *testing.T) {
	var (
		inR  *os.File
		inW  *os.File
		outR *os.File
		outW *os.File
		done chan struct{}

		err error
	)
	done = make(chan struct{})

	//send data to process
	inR, inW, err = os.Pipe()
	if err != nil {
		return
	}

	//process send data to out
	outR, outW, err = os.Pipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	//first three entries correspond to standard input, standard output and standard error
	process, err := os.StartProcess("/bin/sh", nil, &os.ProcAttr{
		Files: []*os.File{inR, outW, outW}, //so the last two param is outW
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	//send data or command to new process
	go func() {
		w := bufio.NewWriter(inW)
		for i := 0; i < 3; i++ {
			time.Sleep(2 * time.Second)
			_, err := w.WriteString("date\n")
			if err != nil {
				return
			}

			err = w.Flush() //flush write buffer
			if err != nil {
				fmt.Println("writer Flush failed")
				return
			}
		}
		inW.Close()
		outW.Close()
	}()

	//read data from new process
	go func() {
		s := bufio.NewScanner(outR)
		for s.Scan() {
			fmt.Println(s.Text())
		}

		//send signal to kill process
		process.Signal(os.Kill)
		done <- struct{}{}

		fmt.Println("Finished")
	}()

	//wait process finish
	_, err = process.Wait()
	if err != nil {
		return
	}

	<-done
}
