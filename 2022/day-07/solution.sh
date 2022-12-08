#!/bin/bash
set -e

export MAIN=tmp
export INPUTFILE="$1"

mkdir -p "${MAIN}"

cat "${INPUTFILE}" \
| grep -v '$ ls' \
| sed -E 's/^dir ([A-Za-z0-9\._-]+)$/mkdir -p \1/' \
| sed -E 's/^([0-9]+) ([A-Za-z0-9\._-]+)$/fallocate -l \1 \2/' \
| sed -E 's/^\$ //' \
| sed "s#cd /#cd $(pwd)/${MAIN}#" \
| bash

echo "Problem 1"

paste \
    <( du -b "${MAIN}" | sort -k2 | awk '{ print $2 "\t" $1 }' ) \
    <( du -b "${MAIN}" | sort -k2 | cut -f2 | xargs -I % sh -c 'find % -mindepth 1 -type d | wc -l' ) \
| awk '{ print $1 "\t" ($2 - (4096 * ($3 + 1))) }' \
| awk '{ if ( $2 <= 100000 ) { print } }' \
| awk '{ s += $2 } END { print s }'

echo "Problem 2"

paste \
<( du -b "${MAIN}" | sort -k2 | awk '{ print $2 "\t" $1 }' ) \
<( du -b "${MAIN}" | sort -k2 | cut -f2 | xargs -I % sh -c 'find % -mindepth 1 -type d | wc -l' ) \
<( paste \
        <( du -bs "${MAIN}" ) \
        <( find "${MAIN}" -mindepth 1 -type d | wc -l ) \
    | awk '{ print $1 - (($3 + 1) * 4096) }' \
    | xargs yes 2> /dev/null \
    | head -n "$( du -b "${MAIN}" | wc -l )" \
    ) \
| awk '{ print ($2 - (4096 * ($3 + 1))) "\t" $4 }' \
| awk '{ if ( $1 >= (30000000 - (70000000 - $2))) { print } }' \
| sort -k1 \
| head -n1 \
| cut -f1
