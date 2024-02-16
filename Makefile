config:
	@rm -r config
	@pkl-gen-go pkl/config.pkl --base-path github.com/pstano1/emailVerfier/

build:
	@chmod +x start.sh
	@go mod tidy && go mod download
	@cd frontend && npm install
	@./start.sh

.PHONY: config build