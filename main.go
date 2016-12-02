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
	body := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		body += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer SendMail(body)

	conf, err := LoadConfig()
	if err != nil {
		log.Print(err)
		return
	}
	file, err := os.OpenFile(conf.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()

	header := "Sendmail Date: " + time.Now().String()
	file.Write(([]byte)(header + "\n" + body))
}

func SendMail(body string) {
	sendmail := exec.Command("sendmail", "-t", "-i")
	stdin, _ := sendmail.StdinPipe()

	io.WriteString(stdin, body)
	stdin.Close()

	err := sendmail.Run()
	if err != nil {
		log.Fatal(err)
	}
}
