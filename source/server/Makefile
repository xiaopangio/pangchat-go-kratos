build:
	#保存bin目录
	mkdir -p bin
	#变量保存bin目录路径
	BIN_DIR :=$(CURDIR)/bin/
	#编译
	cd api-gateway/cmd/api-gateway/ && go build . && cp api-gateway $(BIN_DIR)
	cd connector/cmd/connector/ && go build . && cp connector $(BIN_DIR)
	cd job/cmd/job/ && go build . && cp job $(BIN_DIR)
	cd logic/logic/cmd/logic/ && go build . && cp logic $(BIN_DIR)
	cd logic/message/cmd/message/ && go build . && cp message $(BIN_DIR)
	cd logic/relationship/cmd/relationship/ && go build . && cp relationship $(BIN_DIR)
	cd logic/user/cmd/user/ && go build . && cp user $(BIN_DIR)
	cd online/cmd/online/ && go build . && cp online $(BIN_DIR)