#!/bin/bash

echo "1..3"


#
# Url
#

desc='Handles HTTP URL'
bin/myhead http://golang.org &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Handles HTTPS URL'
bin/myhead https://golang.org &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"

desc='Handles mix of URLs and files'
bin/myhead https://golang.org main.go https://golang.org Makefile &> /dev/null
[[ $? -eq 0 ]] && echo ok "${desc}" || echo not ok "${desc}"
