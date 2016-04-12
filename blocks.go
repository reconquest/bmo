package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/seletskiy/hierr"
)

func parseBlocks(args map[string]interface{}) error {
	var (
		awkRangeBegin   = args["<awk_range_begin>"].(string)
		awkRangeEnd, _  = args["<awk_range_end>"].(string)
		awkCondition, _ = args["<awk_condition>"].(string)

		awkEnumerateWords = awkBool(args["-w"])

		awkVarNames       = args["<name>"].([]string)
		awkVarExpressions = args["<expression>"].([]string)
		awkFormat, _      = args["<format>"].(string)
		awkSort, _        = args["<var>"].(string)

		debug = args["--debug"].(bool)
	)

	awkVars := map[string]string{}
	for i, name := range awkVarNames {
		awkVars[name] = awkVarExpressions[i]
	}

	if awkRangeEnd == "" {
		awkRangeEnd = awkRangeBegin
	}

	awkHandleRangeEndLine := awkBool(awkRangeBegin != awkRangeEnd)

	vars := map[string]interface{}{
		"range_begin":           awkRangeBegin,
		"range_end":             awkRangeEnd,
		"condition":             awkCondition,
		"enumerate_words":       awkEnumerateWords,
		"handle_range_end_line": awkHandleRangeEndLine,
		"vars":                  awkVars,
		"format":                awkFormat,
		"sort":                  awkSort,
	}

	var awkProgram bytes.Buffer
	err := awkParseBlocks.Execute(&awkProgram, vars)
	if err != nil {
		return hierr.Errorf(err, "can't prepare awk program")
	}

	if debug {
		fmt.Fprintln(os.Stderr, awkProgram.String())
	}

	command := exec.Command("awk", awkProgram.String())
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
