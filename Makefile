
.PHONY: all
all: format test

.PHONY: test
test:
	@ for f in day-*.jl; do\
		printf "\n$$f:\n";\
		julia $$f;\
	done

.PHONY: format
format: fmt.so
	@ julia --sysimage fmt.so fmt.jl

fmt.so:
	@ julia -e "using PackageCompiler; \
		create_sysimage(:JuliaFormatter, sysimage_path = \"fmt.so\", \
			precompile_execution_file = \"fmt.jl\", replace_default = false)"
