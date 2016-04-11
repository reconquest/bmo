#!/bin/bash

tests:ensure bmo -b '/^# Time: \w/' '/^# Time: \w/' '/bbb/ || /ccc/' <<EOF
# Time: 1
aaa
# Time: 2
bbb
# Time: 3
ccc
# Time: 4
ddd
EOF

tests:assert-no-diff stdout <<EOF
# Time: 2
bbb
# Time: 3
ccc
EOF
