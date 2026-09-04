package functions

func GetRequireFileFunction() string {
	return `define require_file
	@if [ ! -f "$(1)" ]; then \
		echo "Required file $(1) does not exist."; \
		exit 1; \
	fi
endef
`
}

func GetRequireDirFunction() string {
	return `define require_dir
	@if [ ! -d "$(1)" ]; then \
		echo "Required directory $(1) does not exist."; \
		exit 1; \
	fi
endef
`
}
