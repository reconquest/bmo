#!/bin/bash

tests:ensure bmo \
    -b '/begin/' '/end/' \
    -v 'time'='if (/query_time/) { return int(\$2); }' \
    -s 'time' <<EOF
begin 1
query_time 100
end 1
begin 2
query_time 300
end 2
begin 3
query_time 200
end 3
EOF

tests:assert-no-diff stdout <<EOF
begin 1
query_time 100
end 1
begin 3
query_time 200
end 3
begin 2
query_time 300
end 2
EOF
