#!/bin/bash

tests:ensure bmo -b '/beginning/' <<EOF
beginning 1
aaa

beginning 2
bbb
EOF

tests:assert-no-diff stdout <<EOF
beginning 1
aaa

beginning 2
bbb
EOF
