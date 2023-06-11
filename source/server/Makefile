build:
	BIN_DIR := $(shell pwd)/bin
	cd api-gateway && make build && cd cmd/api-gateway && cp api-gateway $(BIN_DIR)
	cd connector && make build && cd cmd/connector && cp connector $(BIN_DIR)
	cd job && make build && cd cmd/job && cp job $(BIN_DIR)
	cd logic/logic && make build && cd cmd/logic && cp logic $(BIN_DIR)
	cd logic/message && make build && cd cmd/message && cp message $(BIN_DIR)
	cd logic/relationship && make build && cd cmd/relationship && cp relationship $(BIN_DIR)
	cd logic/user && make build && cd cmd/user && cp user $(BIN_DIR)
	cd online && make build && cd cmd/online && cp online $(BIN_DIR)