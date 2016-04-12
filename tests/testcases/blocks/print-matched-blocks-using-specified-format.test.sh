#!/bin/bash

tests:ensure bmo \
    -b '/begin/' '/end/' \
    -v 'time' 'if (/query_time/) { return int(\$2); }' \
    -v 'read_ops' 'if (/read_ops/) { return int(\$4); }' \
    -f '\"time:\" time \"\n\" \"read_ops:\" read_ops' <<EOF
begin 1
bar
query_time 100
write_ops 200 read_ops 300
end 1

garbage

begin 2
foo
query_time 300
write_ops 100 read_ops 500
end 2
EOF

tests:assert-no-diff stdout <<EOF
time:100
read_ops:300
time:300
read_ops:500
EOF
