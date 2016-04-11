#!/bin/bash

tests:put input <<EOF
beginning 1
aaa
finishing 1
garbage

garbage1122
asdkjhasld
garbage

beginning 2
b
b
finishing 2
EOF

tests:eval cat input \| bmo -b '/beginning/' '/finishing/' '1'
tests:assert-success

tests:assert-no-diff stdout <<EOF
beginning 1
aaa
finishing 1
beginning 2
b
b
finishing 2
EOF
