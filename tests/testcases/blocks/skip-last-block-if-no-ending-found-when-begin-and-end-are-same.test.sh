#!/bin/bash

tests:put input <<EOF
time 1
aaa  1
time 2
aaa  2
time 3
aaa  3
EOF

tests:eval cat input \| bmo -b '/time/' '/time/' '/aaa/'
tests:assert-success

tests:assert-no-diff stdout <<EOF
time 1
aaa  1
time 2
time 2
aaa  2
time 3
EOF
