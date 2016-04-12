#!/bin/bash

tests:ensure bmo -b '/beginning/' '/finishing/' <<EOF
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

tests:assert-no-diff stdout <<EOF
beginning 1
aaa
finishing 1
beginning 2
b
b
finishing 2
EOF
