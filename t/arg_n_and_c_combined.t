#!/bin/bash

echo '1..1'


#
# Both -n and -c specified
#

desc='Exits non-zero when both -n and -c are specified'
bin/myhead -n=1 -c=1 main.go #&> /dev/null
[[ $? -eq 0 ]] && echo not ok "${desc}" || echo ok "${desc}"
