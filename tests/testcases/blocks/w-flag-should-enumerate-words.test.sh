#!/bin/bash

tests:ensure bmo -w -b '/begin/' '/end/' <<EOF
begin
bbbbbb 		blah dddd ff h
	tttttt ttt  ttt ttt ttt ttt ttt ttt ttt ttt ttt ttt ttt
	yyyyyy yyy  yyy yyy yyy yyy yyy yyy yyy yyy yy yy yy
end
EOF

tests:assert-no-diff stdout <<EOF
1
begin
1      		2    3    4  5
bbbbbb 		blah dddd ff h
	1      2    3   4   5   6   7   8   9   10  11  12  13
	tttttt ttt  ttt ttt ttt ttt ttt ttt ttt ttt ttt ttt ttt
	1      2    3   4   5   6   7   8   9   10  11 12 13
	yyyyyy yyy  yyy yyy yyy yyy yyy yyy yyy yyy yy yy yy
1
end
EOF
