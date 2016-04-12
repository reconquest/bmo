package main

import "github.com/docopt/docopt-go"

var (
	version = "1.0"
	usage   = `bmo - missing tool for parsing structured text.

Usage:
    bmo [-w] (-b <awk_range_begin> <awk_range_end> <awk_condition>)
            [(-v <name> <expression>)]...
              [--debug]
    bmo -h | --help
    bmo --version

Options:
    -b --blocks  Match blocks using <awk_range_begin>, <awk_range_end> and
                 output it only if <awk_condition> is true for at least on line
                 in the block.
    -v --var     Match variable <name> using <expression>.
    -w           Enumerate words when printing a line.
    --debug      Debug mode.
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
