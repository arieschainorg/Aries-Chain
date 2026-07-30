all: evmd-env

evmd-env: evmd-rmi
	docker build --tag ariesd -f evmd-env/Dockerfile \
		$(shell git rev-parse --show-toplevel)

evmd-rmi:
	docker rmi ariesd 2>/dev/null; true

.PHONY: all evmd-env evmd-rmi
