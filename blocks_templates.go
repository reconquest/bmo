package main

import (
	"text/template"
)

var awkParseBlocks = template.Must(template.New("blocks").Parse(`
	function enumerate_words(line) {
		gsub(/\w+/, "x", line);
		for (number = 1; number <= NF; number++) {
			word = $number;
			repeat_count = length(word) - length(number);
			replacement = number;
			if (repeat_count > 0  && number != NF) {
				replacement = sprintf("%d%*s", number, repeat_count, "");
			}
			sub(/x/, replacement, line);
		}
		return line;
	}

	function handle_block_inner() {
		line = $0

		if ({{ .enumerate_words }}) {
			line = enumerate_words(line) "\n" line;
		}

		if (block_contents) {
			block_contents = block_contents "\n" line;
		} else {
			block_contents = line;
		}
	}

	function handle_block_end() {
		current_line = $0;

		$0 = block_contents;
		if ({{ $.condition }}) {
			print;
		}

		$0 = current_line;
	}

	in_range && ({{ $.range_end }}) {
		if ({{ .handle_range_end_line }}) {
			handle_block_inner()
		}
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
