package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Run sg_ses with simulated input
	cmd := exec.Command("sg_ses", "--status", "--inhex=ses_areca_all.hex", "--page=ec")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating pipe:", err)
		os.Exit(1)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(stdout)
	var volts string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Voltage") && strings.Contains(line, "volts") {
			fields := strings.Fields(line)
			volts = fields[1]
		}
	}

	// Output in Influx line protocol
	fmt.Printf("voltage_sensor,sensor_id=1 voltage=%s\n", volts)

}
