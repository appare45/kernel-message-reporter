package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func main() {
	var output string
	flag.StringVar(&output, "o", "", "Output file")
	flag.Parse()
	var err error
	out := os.Stderr
	if output != "" {
		out, err = os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}
	commands := flag.Args()
	if len(commands) < 1 {
		log.Fatal("Usage: kernel-message-reporter <CMD>")
	}
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	kmsg, err := os.OpenFile("/dev/kmsg", os.O_RDONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer kmsg.Close()
	kmsg.Seek(0, io.SeekEnd)

	cmd.Run()

	msgs, err := readKmsgs(kmsg)

	for _, m := range msgs {
		fmt.Fprint(out, m)
	}
}

func readKmsgs(kmsg *os.File) (messages []string, err error) {
	buf := make([]byte, 1024)
	for {
		kmsg.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		var n int
		n, err = kmsg.Read(buf)
		if errors.Is(err, os.ErrDeadlineExceeded) {
			break
		}
		if err != nil {
			return
		}
		if _, msg, ok := strings.Cut(string(buf[:n]), ";"); ok {
			messages = append(messages, msg)
		}
	}
	return
}
