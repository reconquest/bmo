package main

import (
	"text/template"
)

var awkParseBlocks = template.Must(template.New("blocks").Parse(`
	number_words = {{ .number_words }}

	function handle_block_inner() {
		if (block_contents) {
			block_contents = block_contents "\n" $0
		} else {
			block_contents = $0
		}
	}

	function handle_block_end() {
		current_line = $0

		$0 = block_contents
		if ({{ $.condition }}) {
			print block_contents
		}

		$0 = current_line
	}

	in_range && ({{ $.range_end }}) {
		handle_block_inner()
		handle_block_end();
		in_range = 0;
	}

	{{ $.range_begin }} {
		in_range = 1;
		use_block = 0;
		block_contents = ""
	}

	in_range {
		handle_block_inner()
	}
`))
