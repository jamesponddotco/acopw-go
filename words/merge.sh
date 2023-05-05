#!/usr/bin/env bash

TMPFILE="$(mktemp)"
# Make sure we clean up the tmpfile
trap "rm -f ${TMPFILE}" EXIT

# Merge all files starting with "list-" in this directory.
sort -u -S 30% list-* > "${TMPFILE}" && mv "${TMPFILE}" word-list.txt
