build:
	#保存bin目录
	mkdir -p bin
	#变量保存bin目录路径
	PROJECT_BIN_DIR :=$(CURDIR)/bin
	#编译
	cd api-gateway/cmd/api-gateway/ && go build . && mv api-gateway $(PROJECT_BIN_DIR)
	cd connector/cmd/connector/ && go build . && mv connector $(PROJECT_BIN_DIR)
	cd job/cmd/job/ && go build . && mv job $(PROJECT_BIN_DIR)
	cd logic/logic/cmd/logic/ && go build . && mv logic $(PROJECT_BIN_DIR)
	cd logic/message/cmd/message/ && go build . && mv message $(PROJECT_BIN_DIR)
	cd logic/relationship/cmd/relationship/ && go build . && mv relationship $(PROJECT_BIN_DIR)
	cd logic/user/cmd/user/ && go build . && mv user $(PROJECT_BIN_DIR)
	cd online/cmd/online/ && go build . && mv online $(PROJECT_BIN_DIR)