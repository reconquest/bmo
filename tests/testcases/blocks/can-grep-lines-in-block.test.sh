#!/bin/bash

tests:ensure bmo \
    -b '/begin/' '/end/' \
    -c '/foo/' \
    -t '/query_time/' <<EOF
begin 1
foo
query_time 100
write_ops 200 read_ops 300
end 1

garbage

begin 2
bar
query_time 200
write_ops 100 read_ops 500
end 2

begin 3
foo
query_time 300
write_ops 100 read_ops 500
end 3

begin 4
foo
garbage
end 4
EOF

tests:assert-no-diff stdout <<EOF
query_time 100
query_time 300
EOF
