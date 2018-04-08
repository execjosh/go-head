#!/bin/bash

echo '1..3'


#
# -n=LINES FILE
#

desc='Exits non-zero for equal-to-zero LINES'
bin/myhead -n=0 main.go &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"

desc='Exits non-zero for less-than-zero LINES'
bin/myhead -n=-1 main.go &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"


#
# -n=LINES
#

desc='Exits zero for no input, assumes stdin'
bin/myhead -n=1 < /dev/null &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"
