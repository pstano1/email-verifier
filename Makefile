config:
	@rm -r config
	@pkl-gen-go pkl/config.pkl --base-path github.com/pstano1/emailVerfier/

.PHONY: config