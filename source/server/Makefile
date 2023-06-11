build:
	#保存bin目录
	mkdir -p bin
	#变量保存bin目录路径
	BIN_DIR = $(shell pwd)/bin/
	cd api-gateway/cmd/api-gateway/ && go build . && mv api-gateway ${BIN_DIR}
	cd connector/cmd/connector/ && go build . && mv connector ${BIN_DIR}
	cd job/cmd/job/ && go build . && mv job ${BIN_DIR}
	cd logic/logic/cmd/logic/ && go build . && mv logic ${BIN_DIR}
	cd logic/message/cmd/message/ && go build . && mv message ${BIN_DIR}
	cd logic/relationship/cmd/relationship/ && go build . && mv relationship ${BIN_DIR}
	cd logic/user/cmd/user/ && go build . && mv user ${BIN_DIR}
	cd online/cmd/online/ && go build . && mv online ${BIN_DIR}