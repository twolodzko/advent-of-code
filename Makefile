
.PHONY: test
test:
	@ for f in day-*.jl; do\
		printf "\n$$f:\n";\
		julia $$f;\
	done

.PHONY: format
format:
	@ julia -e 'using JuliaFormatter; format(".")'
