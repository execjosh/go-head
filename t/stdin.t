#!/bin/bash

echo '1..2'


#
# Defaults to stdin
#

desc='Assumes stdin when no file specified'
bin/myhead < <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Assumes stdin when no file specified'
printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n' | bin/myhead &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"
