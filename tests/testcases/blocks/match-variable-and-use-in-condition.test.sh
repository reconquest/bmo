#!/bin/bash

tests:ensure bmo \
    -b '/begin/' '/end/' "time == 300" \
    -v 'time' 'if (/query_time/) { return int(\$2); }' <<EOF
begin 1
bar
query_time 100
end 1

garbage

begin 2
foo
query_time 300
end 2
EOF

tests:assert-no-diff stdout <<EOF
begin 2
foo
query_time 300
end 2
EOF
