package main

import (
	"text/template"
)

var awkParseBlocks = template.Must(template.New("blocks").Parse(`
{{ range $name, $expression := $.vars }}
function extract_var_{{ $name }}() {
	{{ $expression }}

	return 0;
}
{{ end }}

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
	_line = $0

	{{ range $name, $_ := $.vars }}
	if (!{{ $name }}) {
		{{ $name }} = extract_var_{{ $name }}()
	}{{ end }}

	if ({{ $.enumerate_words }}) {
		_line = enumerate_words(_line) "\n" _line;
	}

	if (_block_contents) {
		_block_contents = _block_contents "\n" _line;
	} else {
		_block_contents = _line;
	}
}

function handle_block_end() {
	current_line = $0;

	$0 = _block_contents;
	if ({{ $.condition }}) {
		print {{ $.format }};
	}

	$0 = current_line;
}

_in_range && ({{ $.range_end }}) {
	if ({{ $.handle_range_end_line }}) {
		handle_block_inner()
	}
	handle_block_end();
	_in_range = 0;
}

{{ $.range_begin }} {
	_in_range = 1;
	_use_block = 0;
	_block_contents = ""

	{{ range $name, $_ := $.vars }}
	{{ $name }} = 0;{{ end }}
}

_in_range {
	handle_block_inner()
}
`))
