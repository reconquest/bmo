package main

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"

	"github.com/seletskiy/hierr"
)

var parseBlocksAWKProgram = template.Must(template.New("blocks").Parse(`
	{{ $.RangeBegin }} {
		in_range = 1;
		use_block = 0;
		block_contents = ""
	}

	in_range {
		if (block_contents) {
			block_contents = block_contents "\n" $0
		} else {
			block_contents = $0
		}

		$0 = block_contents
		if (!use_block && ({{ $.Condition }})) {
			use_block = 1;
		}
	}

	{{ $.RangeEnd }} {
		if (use_block) {
			print block_contents
		}

		in_range = 0;
	}
`))

func parseBlocks(args map[string]interface{}) error {
	var AWKProgram bytes.Buffer

	err := parseBlocksAWKProgram.Execute(&AWKProgram, map[string]string{
		"RangeBegin": args["<awk_range_begin>"].(string),
		"RangeEnd":   args["<awk_range_end>"].(string),
		"Condition":  args["<awk_condition>"].(string),
	})

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
