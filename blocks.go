package main

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/seletskiy/hierr"
)

func parseBlocks(args map[string]interface{}) error {
	var AWKProgram bytes.Buffer

	vars := map[string]string{
		"RangeBegin": args["<awk_range_begin>"].(string),
		"RangeEnd":   args["<awk_range_end>"].(string),
		"Condition":  args["<awk_condition>"].(string),
	}

	err := awkParseBlocks.Execute(&AWKProgram, vars)

	if err != nil {
		return hierr.Errorf(err, "can't prepare AWK AWKProgram")
	}

	command := exec.Command("awk", AWKProgram.String())
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err = command.Run()
	if err != nil {
		return hierr.Errorf(err, "can't run AWK AWKProgram")
	}

	return nil
}
