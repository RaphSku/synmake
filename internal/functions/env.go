package functions

func GetRequireEnvFunction() string {
	return `define require_env
	@if [ -z "$$($(1))" ]; then \
		echo "Required environment variable $(1) is not set."; \
		exit 1; \
	fi
endef
`
}
