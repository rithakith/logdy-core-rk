package tests

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogdyE2E_Socket(t *testing.T) {
	// Channels for synchronization
	msgChan := make(chan string, 100)
	socketsReady := make(chan string, 3)

	// Setup wait group for message verification
	var wg sync.WaitGroup
	wg.Add(3) // Expect 3 messages

	// We'll use ports that are likely to be free
	basePort := 9475
	port1 := strconv.Itoa(basePort)
	port2 := strconv.Itoa(basePort + 1)
	port3 := strconv.Itoa(basePort + 2)

	// Start logdy process with -t flag
	cmd := exec.Command("go", "run", "../.", "socket", "-t", port1, port2, port3, "-p", "8082")
	stdout, err := cmd.StdoutPipe()
	assert.NoError(t, err)
	stderr, err := cmd.StderrPipe()
	assert.NoError(t, err)

	// Start reading stdout in background
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "TCP Server is listening") {
				socketsReady <- line
			}
			if strings.Contains(line, "test message on port") {
				select {
				case msgChan <- line:
				default:
				}
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			t.Logf("Logdy Stderr: %s", scanner.Text())
		}
	}()

	// Start the process
	err = cmd.Start()
	assert.NoError(t, err)

	// Wait for all 3 socket servers to report they are listening
	for i := 0; i < 3; i++ {
		select {
		case <-socketsReady:
			// One server ready
		case <-time.After(10 * time.Second):
			t.Fatalf("Timeout waiting for socket servers to start (only %d started)", i)
		}
	}
	time.Sleep(500 * time.Millisecond)

	// Send test messages to each port
	ports := []string{port1, port2, port3}
	for _, port := range ports {
		var conn net.Conn
		var dialErr error
		for retries := 0; retries < 20; retries++ {
			conn, dialErr = net.DialTimeout("tcp", "127.0.0.1:"+port, 1*time.Second)
			if dialErr == nil {
				break
			}
			time.Sleep(200 * time.Millisecond)
		}
		if dialErr != nil {
			t.Fatalf("Failed to connect to port %s: %v", port, dialErr)
		}

		message := fmt.Sprintf("test message on port %s", port)
		fmt.Fprintln(conn, message)
		conn.Close()
	}

	// Collect received messages
	var msgReceived []string
	timeout := time.After(10 * time.Second)

	for len(msgReceived) < 3 {
		select {
		case msg := <-msgChan:
			msgReceived = append(msgReceived, msg)
		case <-timeout:
			t.Fatalf("Timeout waiting for messages to be received. Got %d", len(msgReceived))
		}
	}

	// Kill the process
	cmd.Process.Kill()
	cmd.Wait()

	// Verify we received messages from all ports
	assert.Equal(t, 3, len(msgReceived))
}
