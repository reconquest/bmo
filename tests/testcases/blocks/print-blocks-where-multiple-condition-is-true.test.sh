#!/bin/bash

tests:ensure bmo -b '/beginning/' '/finishing/' -c '/aaa/ && /bbb/' <<EOF
beginning 1
hhaaaxxx
bbb
finishing 1
garbage

garbage1122
asdkjhasld
garbage

beginning 2
baaahhhhh
b
finishing 2
EOF

tests:assert-no-diff stdout <<EOF
beginning 1
hhaaaxxx
bbb
finishing 1
EOF
