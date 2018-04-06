#!/bin/bash

echo '1..3'

desc='Exits zero for successful run'
bin/myhead main.go &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Defaults to 10 lines'
line_count=$(bin/myhead <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') | wc -l)
[[ ${line_count} -eq '10' ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Prints requested number of lines'
line_count=$(bin/myhead -n=1 <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') | wc -l)
[[ ${line_count} -eq '1' ]] && echo ok "${desc}" || echo not ok "${desc}"
