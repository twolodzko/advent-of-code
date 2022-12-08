MAKEFLAGS += --no-print-directory --silent
SHELL = bash
INPUTFILE := example.txt
MAIN = tmp

define make
	$(MAKE) -f files.Makefile
endef

.PHONY: run
run:
	docker run --rm -it $$(docker build -q .) \
		make -f files.Makefile solve INPUTFILE=example.txt
	docker run --rm -it $$(docker build -q .) \
		make -f files.Makefile solve INPUTFILE=problem.txt

.PHONY: solve
solve: setup problem1 problem2

.PHONY: setup
setup:
	set -e && \
	mkdir -p $(MAIN) && \
	cat $(INPUTFILE) \
	| grep -v '$$ ls' \
	| sed -E 's/^dir ([A-Za-z0-9\._-]+)$$/mkdir -p \1/' \
	| sed -E 's/^([0-9]+) ([A-Za-z0-9\._-]+)$$/fallocate -l \1 \2/' \
	| sed -E 's/^\$$ //' \
	| sed "s#cd /#cd $(PWD)/$(MAIN)#" \
	| bash

.PHONY: problem1
problem1:
	paste \
		<( du -b $(MAIN) | sort -k2 | awk '{ print $$2 "\t" $$1 }' ) \
		<( du -b $(MAIN) | sort -k2 | cut -f2 | xargs -I % sh -c 'find % -mindepth 1 -type d | wc -l' ) \
	| awk '{ print $$1 "\t" ($$2 - (4096 * ($$3 + 1))) }' \
	| awk '{ if ( $$2 <= 100000 ) { print } }' \
	| awk '{ s += $$2 } END { print s }'

.PHONY: problem2
problem2:
	paste \
		<( du -b $(MAIN) | sort -k2 | awk '{ print $$2 "\t" $$1 }' ) \
		<( du -b $(MAIN) | sort -k2 | cut -f2 | xargs -I % sh -c 'find % -mindepth 1 -type d | wc -l' ) \
		<( paste \
				<( du -bs $(MAIN) ) \
				<( find $(MAIN) -mindepth 1 -type d | wc -l ) \
			| awk '{ print $$1 - (($$3 + 1) * 4096) }' \
			| xargs yes 2> /dev/null \
			| head -n $$( du -b $(MAIN) | wc -l ) \
		   ) \
		| awk '{ print ($$2 - (4096 * ($$3 + 1))) "\t" $$4 }' \
		| awk '{ if ( $$1 >= (30000000 - (70000000 - $$2))) { print } }' \
		| sort -k1 \
		| head -n1 \
		| cut -f1

.PHONY: sizes
sizes:
	paste \
		<( du -b $(MAIN) | sort -k2 | awk '{ print $$2 "\t" $$1 }' ) \
		<( du -b $(MAIN) | sort -k2 | cut -f2 | xargs -I % sh -c 'find % -mindepth 1 -type d | wc -l' ) \
	| awk '{ print $$1 "\t" ($$2 - (4096 * ($$3 + 1))) }'
