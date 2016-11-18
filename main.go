package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func init() {
	f, _ := os.OpenFile("/var/log/sendmail-logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)
}

func main() {
	conf := LoadConfig()

	body := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		body += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(conf.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	header := "Sendmail Date: " + time.Now().String()
	file.Write(([]byte)(header + "\n" + body))

	sendmail := exec.Command("sendmail", "-t")
	stdin, _ := sendmail.StdinPipe()

	err = sendmail.Start()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(stdin, body+".\n")

	err = sendmail.Wait()
	if err != nil {
		log.Fatal(err)
	}
}