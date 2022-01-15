.PHONY: test
test:
	cd "$(PWD)/extract-pdf-page-num" && make test-all
	cd "$(PWD)/invert-image-color" && make test-all
