MAKEFLAGS += --no-print-directory --silent
SHELL = bash

le = $(shell test $(1) -le $(2) 2>/dev/null; echo $$?)

INPUTFILE := $(PWD)/example
INPUTSIZE = $(firstword $(shell wc -l $(INPUTFILE)))
LINENUMBER := 1
NEXTLINE := $(shell expr $(LINENUMBER) + 1)
LINE := $(shell head -n $(LINENUMBER) $(INPUTFILE) | tail -n 1)
MAIN = tmp
DIR := $(MAIN)
ACC := 0
ARGS := ""
LOCALPATH := $(DIR)/$(firstword $(ARGS))

define make
	$(MAKE) -f files.Makefile
endef

.PHONY: main
main: setup problem1 problem2

.PHONY: setup
setup:
	mkdir $(MAIN)
	$(make) loop

.PHONY: loop
loop:
ifeq ($(call le,$(LINENUMBER),$(INPUTSIZE)),0)
	$(make) eval
else
	@:
endif

.PHONY: eval
eval:
ifeq ($(LINE),)
	@:
else ifeq ($(firstword $(LINE)),$)
ifeq ($(word 2, $(LINE)),cd)
ifeq ($(word 3, $(LINE)),/)
	$(make) loop LINENUMBER=$(NEXTLINE) DIR=$(MAIN)
else
	$(make) loop LINENUMBER=$(NEXTLINE) DIR=$(shell cd $(DIR)/$(word 3, $(LINE)) && pwd)
endif
else ifeq ($(word 2, $(LINE)),ls)
	$(make) loop LINENUMBER=$(NEXTLINE) DIR=$(DIR)
endif
else ifeq ($(firstword $(LINE)),dir)
	mkdir -p $(DIR)/$(word 2, $(LINE))
	$(make) loop LINENUMBER=$(NEXTLINE) DIR=$(DIR)
else
	fallocate -l $(firstword $(LINE)) $(DIR)/$(word 2, $(LINE))
	$(make) loop LINENUMBER=$(NEXTLINE) DIR=$(DIR)
endif

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

size:
	paste \
		<( du -bs $(MAIN) ) \
		<( find $(MAIN) -mindepth 1 -type d | wc -l ) \
	| awk '{ print $$1 - ( ( $$3 + 1 ) * 4096 ) }'
