MAKEFLAGS += --no-print-directory --silent
SHELL = bash

last = $(word $(words $(1)),$(1))
tail = $(wordlist 2, $(words $(1)), $(1))
isdir = $(shell test -d $(1); echo $$?)
lt = $(shell test $(1) -lt $(2) 2>/dev/null; echo $$?)
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

.PHONY: main
main: setup
	$(MAKE) solve ARGS="$(foreach v, \
					$(shell $(MAKE) dirs ARGS="$(shell ls $(MAIN))"), \
					$(shell $(MAKE) size ARG=$(v)))"

.PHONY: solve
solve:
	echo "Problem 1:"
	echo $(shell $(MAKE) problem1 \
				ACC=0 \
				ARGS="$(ARGS)")
	echo
	echo "Problem 2:"
	echo $(shell $(MAKE) problem2 \
				ACC=999999999 \
				NEEDED=$(shell expr 30000000 - \( 70000000 - $(shell $(MAKE) size ARG=$(MAIN)) \)) \
				ARGS="$(ARGS)")

.PHONY: setup
setup:
	mkdir $(MAIN)
	$(MAKE) loop

.PHONY: loop
loop:
ifeq ($(call le,$(LINENUMBER),$(INPUTSIZE)),0)
	$(MAKE) eval
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
	$(MAKE) loop LINENUMBER=$(NEXTLINE) DIR=$(MAIN)
else
	$(MAKE) loop LINENUMBER=$(NEXTLINE) DIR=$(shell cd $(DIR)/$(word 3, $(LINE)) && pwd)
endif
else ifeq ($(word 2, $(LINE)),ls)
	$(MAKE) loop LINENUMBER=$(NEXTLINE) DIR=$(DIR)
endif
else ifeq ($(firstword $(LINE)),dir)
	mkdir -p $(DIR)/$(word 2, $(LINE))
	$(MAKE) loop LINENUMBER=$(NEXTLINE) DIR=$(DIR)
else
	echo $(firstword $(LINE)) > $(DIR)/$(word 2, $(LINE))
	$(MAKE) loop LINENUMBER=$(NEXTLINE) DIR=$(DIR)
endif

.PHONY: dirs
dirs:
ifeq ($(words $(ARGS)),0)
	@:
else
ifeq ($(call isdir,$(LOCALPATH)),0)
	echo $(LOCALPATH)
	$(MAKE) dirs \
		DIR="$(LOCALPATH)" \
		ARGS="$(shell ls $(LOCALPATH))"
endif
	$(MAKE) dirs \
		DIR="$(DIR)" \
		ARGS="$(call tail,$(ARGS))"
endif

.PHONY: size
size:
ifeq ($(call isdir,$(ARG)),0)
	$(MAKE) sizesum \
		ACC=0 \
		DIR=$(ARG) \
		ARGS="$(shell ls $(ARG))"
else
	cat $(firstword $(ARG))
endif

.PHONY: sizesum
sizesum:
ifeq ($(words $(ARGS)),0)
	echo $(ACC)
else
	$(MAKE) sizesum \
		ACC=$(shell expr $(ACC) + \
			$(shell $(MAKE) size \
				ACC=0 \
				ARG=$(LOCALPATH))) \
		DIR=$(DIR) \
		ARGS="$(call tail,$(ARGS))"
endif

.PHONY: problem1
problem1:
ifeq ($(words $(ARGS)),0)
	echo $(ACC)
else ifeq ($(call lt, $(firstword $(ARGS)), 100001),0)
	$(MAKE) problem1 \
		ACC=$(shell expr $(ACC) + $(firstword $(ARGS))) \
		ARGS="$(call tail,$(ARGS))"
else
	$(MAKE) problem1 \
		ACC=$(ACC) \
		ARGS="$(call tail,$(ARGS))"
endif

.PHONY: problem2
problem2:
ifeq ($(words $(ARGS)),0)
	echo $(ACC)
else ifeq ($(shell test $(NEEDED) -le $(firstword $(ARGS)) 2>/dev/null && \
					test $(firstword $(ARGS)) -lt $(ACC) 2>/dev/null; \
					 echo $$?),0)
	$(MAKE) problem2 \
		ACC=$(firstword $(ARGS)) \
		ARGS="$(call tail,$(ARGS))" \
		NEEDED=$(NEEDED)
else
	$(MAKE) problem2 \
		ACC=$(ACC) \
		ARGS="$(call tail,$(ARGS))" \
		NEEDED=$(NEEDED)
endif
