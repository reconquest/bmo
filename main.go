package main

import "github.com/docopt/docopt-go"

var (
	version = "1.0"
	usage   = `bmo - missing tool for parsing structured text.

Usage:
    bmo (-b | --blocks) <awk_range_begin> <awk_range_end> <awk_condition>
    bmo -h | --help
    bmo --version

Options:
    -b --blocks  Match blocks using <awk_range_begin>, <awk_range_end> and
                 output it only if <awk_condition> is true for at least on line
                 in the block.
    -h --help    Show this help.
    --version    Show version.
`
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		panic(err)
	}

	switch {
	case args["--blocks"].(bool):
		parseBlocks(args)
	}
}
