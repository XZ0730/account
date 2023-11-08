MODULE = github.com/XZ0730/runFzu

.PHONY: target
target:
	sh build.sh

.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print