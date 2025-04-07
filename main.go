package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Run sg_ses with simulated input
	cmd := exec.Command("sg_ses", "--status", "--inhex=ses_areca_all.hex", "--index=vs,1", "--get=Voltage")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	buf := bytes.NewBuffer(out)
	voltsString := buf.String()
	// Sanitze and Math here
	mVolts := strings.TrimSpace(voltsString)
	volts, _ := strconv.ParseFloat(mVolts, 64)
	volts = volts / 100

	// Output in Influx line protocol
	fmt.Printf("voltage_sensor,sensor_id=1 voltage=%.2f\n", volts)

}
