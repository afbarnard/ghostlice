# Copyright (c) 2014 Aubrey Barnard.  This is free software.  See
# LICENSE for details.

.PHONY: test

# Generates the code for the package from the templates
ghostlice.go: genCode.sh allTypes.gotemplate
	bash $^ > $@

# Runs unit tests
test:
	go test
