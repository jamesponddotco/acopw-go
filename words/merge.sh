#!/usr/bin/env bash

# Merge all files starting with "list-" in this directory.
for file in list-*; do
    cat "${file}" >> "${TMP:-/tmp}/word-list.txt"
done

cat "${TMP:-/tmp}/word-list.txt" \
  | sort -u -S 30% > 'word-list.txt'

exit 0
