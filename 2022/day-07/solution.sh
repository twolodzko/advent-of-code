#!/bin/bash
set -e

export MAIN=tmp
export INPUTFILE="$1"

rm -r "${MAIN}"
mkdir -p "${MAIN}"

cat "${INPUTFILE}" \
| grep -v '$ ls' \
| sed -E 's/^dir ([A-Za-z0-9\._-]+)$/mkdir -p \1/' \
| sed -E 's/^([0-9]+) ([A-Za-z0-9\._-]+)$/fallocate -l \1 \2/' \
| sed -E 's/^\$ //' \
| sed "s#cd /#cd $(pwd)/${MAIN}#" \
| bash

# Problem 1

paste \
    <( du -b "${MAIN}" | cut -f1 ) \
    <( du -b "${MAIN}" | cut -f2 | xargs -I % sh -c 'find % -mindepth 1 -type d | wc -l' ) \
| awk '{ print ($1 - (4096 * ($2 + 1))) }' \
| awk '{ if ($1 <= 100000) { s += $1 } } END { print s }'

# Problem 2

paste \
    <( du -b "${MAIN}" | cut -f1 ) \
    <( du -b "${MAIN}" | cut -f2 | xargs -I % sh -c 'find % -mindepth 1 -type d | wc -l' ) \
    <( paste \
            <( du -bs "${MAIN}" ) \
            <( find "${MAIN}" -mindepth 1 -type d | wc -l ) \
        | awk '{ print $1 - (($3 + 1) * 4096) }' \
        | xargs yes 2> /dev/null \
        | head -n "$( du -b "${MAIN}" | wc -l )" \
        ) \
| awk '{ print ($1 - (4096 * ($2 + 1))) "\t" $3 }' \
| awk '{ if ($1 >= (30000000 - (70000000 - $2))) { print } }' \
| sort -k1 \
| head -n1 \
| cut -f1
