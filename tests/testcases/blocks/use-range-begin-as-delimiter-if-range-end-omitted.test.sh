#!/bin/bash

tests:ensure bmo -b '/time/' <<EOF
time 1
aaa  1
time 2
aaa  2
time 3
bbb  3
EOF

tests:assert-no-diff stdout <<EOF
time 1
aaa  1
time 2
aaa  2
EOF
