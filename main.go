package main

import "github.com/docopt/docopt-go"

var (
	version = "3.0"
	usage   = `bmo - missing tool for parsing structured text.

Usage:
    bmo [-w] (-b <awk_range_begin> [<awk_range_end>])
            [(-v <name> <expression>)]...
            [(-c <awk_condition>)]
            [(-f <format>)]
            [(-s <var>)]
            [--debug]
    bmo -h | --help
    bmo --version

Options:
    -b --blocks     Match blocks using <awk_range_begin> and <awk_range_end>.
    -v --var        Match variable <name> using <expression>.
    -s --sort       Sort blocks by <var> variable.
    -c --condition  Use blocks only if <awk_condition> is true.
    -w              Enumerate words when printing a line.
    --debug         Debug mode.
    -h --help       Show this help.
    --version       Show version.
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
