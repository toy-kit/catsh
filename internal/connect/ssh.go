package connect

import (
	"bufio"
	"bytes"
	"catsh/internal/asciicast"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/olivere/ndjson"
	"golang.org/x/crypto/ssh"
)

type SshConnect struct {
	Session *ssh.Session
	Reader  chan []byte
	Writer  chan []byte
	Close   func()
}

func SshOpen(connect *SshConnect) error {
	defer connect.Close()
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("123456"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", "192.168.1.1:22", config)
	if err != nil {
		return err
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}
	go func() {
		for data := range connect.Writer {
			stdin.Write(data)
		}
	}()
	// asciicast
	var buf bytes.Buffer
	result := ndjson.NewWriter(&buf)
	enc := json.NewEncoder(&buf)
	castStdout := asciicast.NewStream(1.0)
	defer castStdout.Close()
	header := &asciicast.Header{
		Version:   2,
		Command:   "Command",
		Title:     "Title",
		Width:     80,
		Height:    40,
		Timestamp: time.Now().Unix(),
		Duration:  asciicast.Duration(castStdout.Duration().Seconds()),
	}
	if err := enc.Encode(&header); err != nil {
		return err
	}
	go func() {
		s := bufio.NewScanner(stdout)
		s.Split(bufio.ScanBytes)
		for s.Scan() {
			castStdout.Write(s.Bytes())
			connect.Reader <- s.Bytes()
		}
		if s.Err() != nil {
			log.Println("scan:", s.Err())
		}
	}()
	go func() {
		s := bufio.NewScanner(stderr)
		s.Split(bufio.ScanBytes)
		for s.Scan() {
			connect.Reader <- s.Bytes()
		}
		if s.Err() != nil {
			log.Println("scan:", s.Err())
		}
	}()
	// Request pseudo terminal
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		return err
	}
	// Start remote shell
	if err := session.Shell(); err != nil {
		return err
	}
	if err = session.Wait(); err != nil {
		return err
	}
	for _, f := range castStdout.Frames {
		if err := result.Encode([]interface{}{f.Time, "o", string(f.EventData)}); err != nil {
			log.Println(err)
		}
	}
	err = os.WriteFile("a.cast", buf.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
