#!/bin/bash

echo '1..2'


#
# -c=BYTES
#

desc='Exits non-zero for equal-to-zero BYTES'
bin/myhead -c=0 main.go &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"

desc='Exits non-zero for less-than-zero BYTES'
bin/myhead -c=-1 main.go &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"
