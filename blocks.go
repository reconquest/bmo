package main

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/seletskiy/hierr"
)

func parseBlocks(args map[string]interface{}) error {
	var (
		awkRangeBegin     = args["<awk_range_begin>"].(string)
		awkRangeEnd       = args["<awk_range_end>"].(string)
		awkCondition      = args["<awk_condition>"].(string)
		awkEnumerateWords = awkBool(args["-w"])

		awkHandleRangeEndLine = awkBool(awkRangeBegin != awkRangeEnd)
	)

	vars := map[string]string{
		"range_begin":           awkRangeBegin,
		"range_end":             awkRangeEnd,
		"condition":             awkCondition,
		"enumerate_words":       awkEnumerateWords,
		"handle_range_end_line": awkHandleRangeEndLine,
	}

	var AWKProgram bytes.Buffer
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

func awkBool(raw interface{}) string {
	booled, _ := raw.(bool)
	if booled {
		return "1"
	}

	return "0"
}
