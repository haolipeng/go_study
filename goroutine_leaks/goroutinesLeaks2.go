package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

type result struct {
	record string
	err    error
}

//Forgotten Sender
//leak is a buggy function.It launches a goroutine that
//blocks receiving from a channel.Nothing will ever be
//sent on that channel and the channel is never closed so
//that goroutine will be blocked forever
func leak() {
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println("We received a value:", val)
	}()
}

//For this leak example you will see a Goroutine that is blocked indefinitely,
//waiting to send a value on a channel

//search simulates a function that finds a record based
//on a search term.It takes 200ms to perform this work.
func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}

//process is the work for the program.It finds a record
//then prints it
func process_one(term string) error {

	record, err := search(term)
	if err != nil {
		return err
	}

	fmt.Println("Received:", record)
	return nil
}

//Think about what the goroutine is doing;it send result struct to channel
//Sending on this channel blocks execution until another goroutine is ready to receive the value.
//In the timeout case,the receiver stops waiting to receive from the goroutine and moves on.
//this will cause the goroutine to block forever waiting for a receiver to appear which can't happen
//This is when the Goroutine leaks,How to fix this leak?
//Fix:change the channel from an unbuffered channel to a buffered channel with a capacity of 1,so sending does't block
func process(term string) error {
	//create a context that will be canceled in 100ms
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	//make a channel for the goroutine to report its result
	//
	ch := make(chan result)

	//launch a goroutine to find the record.Create a result
	//from the returned values to send through the channel
	go func() {
		record, err := search(term)
		ch <- result{record, err}
	}()

	//Block waiting to either receive from the goroutine's
	//channel or for the context to be canceled
	select {
	//receive from the ctx.Done() channel.This case will be executed if the Context gets canceled(100ms duration passes)
	case <-ctx.Done():
		return errors.New("search canceled")
	case result := <-ch:
		if result.err != nil {
			return result.err
		}
		fmt.Println("Received:", result.record)
		return nil
	}
}

func main() {
	startingGs := runtime.NumGoroutine()

	err := process("gophers")
	if err != nil {
		fmt.Println(err)
	}

	//Hold the program from terminating for 1 second to see
	//if any gouroutines created by processRecords terminate
	time.Sleep(time.Second * 2)

	//Capture ending number of goroutines
	endingGs := runtime.NumGoroutine()

	//Report the results
	fmt.Println("========================================")
	fmt.Println("Number of goroutines before:", startingGs)
	fmt.Println("Number of goroutines after :", endingGs)
	fmt.Println("Number of goroutines leaked:", endingGs-startingGs)
}
