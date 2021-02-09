package main

import "github.com/docopt/docopt-go"

var (
	version = "3.1"
	usage   = `bmo - missing tool for parsing structured text.

Usage:
    bmo [-w] [-n] (-b <awk_range_begin> [<awk_range_end>])
            [(-v <name> <expression>)]...
            [(-c <awk_condition>)]
            [(-f <format>)]
            [(-s <var> [<how>])]
            [--debug]
    bmo -h | --help
    bmo --version

Options:
  -b --blocks     Match blocks using <awk_range_begin> and <awk_range_end>.
                   If <awk_range_end> is omitted, than <awk_range_begin> will
                   be used as blocks delimiter.
  -v --var        Match variable <name> using <expression>.
  -s --sort       Sort blocks by <var> variable using <how> func.
                   The <how> will be passed as third param to asorti().
                   Tip: man gawk and find PROCINFO["sorted_in"]
  -c --condition  Use blocks only if <awk_condition> is true.
  -f --format     Output blocks using specified format.
  -w              Enumerate words when printing a line.
  -n              Add a new line after each block.
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
