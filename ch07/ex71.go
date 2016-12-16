package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*c += 1
	}

	return int(*c), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		*c += 1
	}

	return int(*c), nil
}

func main() {
	var counter1 WordCounter
	fmt.Fprintf(&counter1, "The words are %s", "hello, world")
	fmt.Println(counter1)

	var counter2 LineCounter
	fmt.Fprintf(&counter2, "The words \n are %s \n %s", "hello", "world")
	fmt.Println(counter2)
}
