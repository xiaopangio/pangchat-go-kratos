BIN_DIR:=$(subst /,\,$(CURDIR)/bin)
ifeq ($(OS),Windows_NT)
	RM = rmdir /S /Q
	MKDIR = mkdir -p
	PKILL = taskkill /F /IM
	MV = move
	EXE = .exe
else
	RM = rm -f
	MKDIR = mkdir -p
	PKILL = pkill -9 -f
	MV = mv
	EXE =
endif
build:
	mkdir $(BIN_DIR) >nul 2>&1 || exit 0
	cd api-gateway/cmd/api-gateway/ && go build . && $(MV) api-gateway$(EXE) $(BIN_DIR)
	cd connector/cmd/connector/ && go build . && $(MV) connector$(EXE)  $(BIN_DIR)
	cd job/cmd/job/ && go build . && $(MV) job$(EXE)  $(BIN_DIR)
	cd logic/logic/cmd/logic/ && go build . && $(MV) logic$(EXE)  $(BIN_DIR)
	cd logic/message/cmd/message/ && go build . && $(MV) message$(EXE)  $(BIN_DIR)
	cd logic/relationship/cmd/relationship/ && go build . && $(MV) relationship$(EXE)  $(BIN_DIR)
	cd logic/user/cmd/user/ && go build . && $(MV) user$(EXE)  $(BIN_DIR)
	cd online/cmd/online/ && go build . && $(MV) online$(EXE)  $(BIN_DIR)
run:
	make build
ifeq ($(OS),Windows_NT)
	cd $(BIN_DIR) && start /B .\api-gateway$(EXE) -conf ../api-gateway/configs/config.yaml
	cd $(BIN_DIR) && start /B .\connector$(EXE) -conf ../connector/configs/config.yaml
	cd $(BIN_DIR) && start /B .\job$(EXE) -conf ../job/configs/config.yaml
	cd $(BIN_DIR) && start /B .\logic$(EXE) -conf ../logic/logic/configs/config.yaml
	cd $(BIN_DIR) && start /B .\message$(EXE) -conf ../logic/message/configs/config.yaml
	cd $(BIN_DIR) && start /B .\relationship$(EXE) -conf ../logic/relationship/configs/config.yaml
	cd $(BIN_DIR) && start /B .\user$(EXE) -conf ../logic/user/configs/config.yaml
	cd $(BIN_DIR) && start /B .\online$(EXE) -conf ../online/configs/config.yaml
else
	$(BIN_DIR)/api-gateway -conf ./api-gateway/configs/config.yaml&
	$(BIN_DIR)/connector -conf ./connector/configs/config.yaml&
	$(BIN_DIR)/job -conf ./job/configs/config.yaml&
	$(BIN_DIR)/logic -conf ./logic/logic/configs/config.yaml&
	$(BIN_DIR)/message -conf ./logic/message/configs/config.yaml&
	$(BIN_DIR)/relationship -conf ./logic/relationship/configs/config.yaml&
	$(BIN_DIR)/user -conf ./logic/user/configs/config.yaml&
	$(BIN_DIR)/online -conf ./online/configs/config.yaml&
endif
restart:
	make stop
	make run
stop:
ifeq ($(OS),Windows_NT)
	$(PKILL) api-gateway$(EXE) >nul 2>&1 || exit 0
	$(PKILL) connector$(EXE) >nul 2>&1 || exit 0
	$(PKILL) job$(EXE) >nul 2>&1 || exit 0
	$(PKILL) logic$(EXE) >nul 2>&1 || exit 0
	$(PKILL) message$(EXE) >nul 2>&1 || exit 0
	$(PKILL) relationship$(EXE) >nul 2>&1 || exit 0
	$(PKILL) user$(EXE) >nul 2>&1 || exit 0
	$(PKILL) online$(EXE) >nul 2>&1 || exit 0
else
	$(PKILL) api-gateway
	$(PKILL) connector
	$(PKILL) job
	$(PKILL) logic
	$(PKILL) message
	$(PKILL) relationship
	$(PKILL) user
	$(PKILL) online
endif
	make clean
clean:
	$(RM) $(BIN_DIR)\*.exe
	$(RM) $(BIN_DIR)\cache
	$(RM) $(BIN_DIR)\logs
