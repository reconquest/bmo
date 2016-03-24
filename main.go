package main

import "github.com/docopt/docopt-go"

const usage = `bmo - missing tool for parsing structured text.

Usage:
    bmo -h | --help
    bmo (-b | --blocks) <awk_range_begin> <awk_range_end> <awk_condition>

Options:
    -h --help    Show this help.
    -b --blocks  Match blocks using <awk_range_begin>, <awk_range_end> and
                 output it only if <awk_condition> is true for at least on line
                 in the block.
`

func main() {
	args, err := docopt.Parse(usage, nil, true, "1.0", false)
	if err != nil {
		panic(err)
	}

	switch {
	case args["--blocks"].(bool):
		parseBlocks(args)
	}
}
