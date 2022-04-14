package pipe

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func Test1(t *testing.T) {
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
	cmd := exec.Command("/bin/sh")

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.ExtraFiles = append(cmd.ExtraFiles, inR, outW)

	cmd.Stdin = inR
	cmd.Stdout = outW
	cmd.Stderr = outW

	err = cmd.Start()
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

		err := cmd.Wait()
		if err != nil {
			fmt.Println("cmd.Wait failed")
			return
		}
		done <- struct{}{}

		fmt.Println("Finished")
	}()

	<-done
}
