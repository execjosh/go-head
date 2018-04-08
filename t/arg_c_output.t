#!/bin/bash

echo '1..3'


desc='Prints requested number of bytes (1) for single file'
line_count=$(bin/myhead -c=1 <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') 2> /dev/null | wc -c)
[[ ${line_count} -eq '1' ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Prints requested number of bytes (1) for multiple files, has header'
line_count=$(bin/myhead -c=1 <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') 2> /dev/null  | wc -l)
[[ ${line_count} -eq '3' ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Prints requested number of bytes (2) for multiple files, has header'
line_count=$(bin/myhead -c=2 <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') <(printf '1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n') 2> /dev/null  | wc -l)
[[ ${line_count} -eq '5' ]] && echo ok "${desc}" || echo not ok "${desc}"
