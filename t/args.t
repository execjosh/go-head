#!/bin/bash

echo '1..7'


#
# No args
#

desc='Exits zero for no args, assumes stdin'
bin/myhead < /dev/null &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"


#
# -n=<COUNT> validation
#

desc='Exits non-zero for equal-to-zero n'
bin/myhead -n=0 main.go &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"

desc='Exits non-zero for less-than-zero n'
bin/myhead -n=-1 main.go &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"


#
# FILE validation
#

desc='Exits zero for no file, assumes stdin'
bin/myhead -n=1 < /dev/null &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Exits non-zero for non-existent file'
bin/myhead non-existent-file &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"

desc='Exits non-zero for non-existent directory'
bin/myhead non-existent-dir/ &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"

desc='Exits non-zero for directory'
bin/myhead t &> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"
