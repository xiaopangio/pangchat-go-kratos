BIN_DIR:=$(CURDIR)/bin
build:
	#保存bin目录
	mkdir -p bin
	#编译
	cd api-gateway/cmd/api-gateway/ && go build . && mv api-gateway $(BIN_DIR)
	cd connector/cmd/connector/ && go build . && mv connector $(BIN_DIR)
	cd job/cmd/job/ && go build . && mv job $(BIN_DIR)
	cd logic/logic/cmd/logic/ && go build . && mv logic $(BIN_DIR)
	cd logic/message/cmd/message/ && go build . && mv message $(BIN_DIR)
	cd logic/relationship/cmd/relationship/ && go build . && mv relationship $(BIN_DIR)
	cd logic/user/cmd/user/ && go build . && mv user $(BIN_DIR)
	cd online/cmd/online/ && go build . && mv online $(BIN_DIR)
run:
	#运行
	$(BIN_DIR)/api-gateway -conf ./api-gateway/conf/config.yaml&
	$(BIN_DIR)/connector -conf ./connector/conf/config.yaml&
	$(BIN_DIR)/job -conf ./job/conf/config.yaml&
	$(BIN_DIR)/logic -conf ./logic/logic/conf/config.yaml&
	$(BIN_DIR)/message -conf ./logic/message/conf/config.yaml&
	$(BIN_DIR)/relationship -conf ./logic/relationship/conf/config.yaml&
	$(BIN_DIR)/user -conf ./logic/user/conf/config.yaml&
	$(BIN_DIR)/online -conf ./online/conf/config.yaml&