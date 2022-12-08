.PHONY: run
run:
	docker run --rm -it $$(docker build -q .) \
		bash solution.sh example.txt
	docker run --rm -it $$(docker build -q .) \
		bash solution.sh problem.txt
