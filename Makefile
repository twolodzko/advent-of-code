
JULIA = julia --project=. 

.PHONY: all
all: format test

.PHONY: test
test:
	@ for f in day-*.jl; do\
		printf "\n$$f:\n";\
		$(JULIA) $$f;\
	done

.PHONY: format
format: fmt.so
	@ $(JULIA) --sysimage fmt.so fmt.jl

fmt.so:
	@ $(JULIA) -e "using PackageCompiler; \
		create_sysimage(:JuliaFormatter, sysimage_path = \"fmt.so\", \
			precompile_execution_file = \"fmt.jl\", replace_default = false)"

.PHONY: repl
repl:
	$(JULIA)
