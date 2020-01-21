ROOT :=  $(shell pwd)

.PHONY: format
format:
	find ${ROOT} -type f -name "*.go" \
	| xargs -L1 goimports -w -l
