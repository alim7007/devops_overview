#!/usr/bin/env bash

name="${1:-world}"

if [ "$1" = "-u" ]; then
shift
name="${1:-world}"
name=$(echo "$name" | tr '[:lower:]' '[:upper:]')
shift
else
  shift  # remove the name only
fi

message="${*:-No messages yet}"

echo "Hello, $name. Your message: $message"


#remove the first argument if exist
# if [ -n "$1" ]; then
#   shift
# fi


# ./greet.sh -u
# ./greet.sh -u David
# ./greet.sh -u David welcome to the team!
# ./greet.sh David welcome to the team!
