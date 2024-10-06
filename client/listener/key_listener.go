package listener

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
)

func StartKeyboardListener(ctx context.Context) (chan string, error) {
	cmd := exec.CommandContext(ctx, "cmd", "/C", "deps\\gkl")

	cmd.Stderr = cmd.Stdout
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(reader)

	ksCh := make(chan string)

	go func() {
		for scanner.Scan() {
			ks := scanner.Text()
			ksCh <- ks
		}

		close(ksCh)
	}()

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("could not start key-listener, %s", err)
	}

	return ksCh, nil
}
