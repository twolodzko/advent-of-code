#!/bin/bash
set -e

WORKDIR=$(mktemp -d)
INPUTFILE=$(realpath "$1")

cd "${WORKDIR}"

cat "${INPUTFILE}" \
| grep -v '$ ls' \
| sed -E 's/^dir ([A-Za-z0-9\._-]+)$/mkdir -p \1/' \
| sed -E 's/^([0-9]+) ([A-Za-z0-9\._-]+)$/fallocate -l \1 \2/' \
| sed -E 's/^\$ //' \
| sed "s#cd /#cd ${WORKDIR}#" \
| bash

SIZES=$(mktemp)
du -b "${WORKDIR}" > "${SIZES}"

# Problem 1

paste \
    <( cut -f1 "${SIZES}" ) \
    <( cut -f2 "${SIZES}" | xargs -i sh -c 'find {} -mindepth 1 -type d | wc -l' ) \
| awk '{ print ($1 - (4096 * ($2 + 1))) }' \
| awk '{ if ($1 <= 100000) { s += $1 } } END { print s }'

# Problem 2

paste \
    <( cut -f1 "${SIZES}" ) \
    <( cut -f2 "${SIZES}" | xargs -i sh -c 'find {} -mindepth 1 -type d | wc -l' ) \
    <( paste \
            <( du -bs "${WORKDIR}" ) \
            <( find "${WORKDIR}" -mindepth 1 -type d | wc -l ) \
        | awk '{ print $1 - (($3 + 1) * 4096) }' \
        | xargs yes 2> /dev/null \
        | head -n "$( wc -l "${SIZES}" | cut -d' ' -f1 )" \
    ) \
| awk 'OFS="\t" { print ($1 - (4096 * ($2 + 1))) , $3 }' \
| awk '{ if ($1 >= (30000000 - (70000000 - $2))) { print } }' \
| sort -k1 \
| head -n1 \
| cut -f1

rm -rf "${WORKDIR}" "$SIZES"
