package functions

func GetRequireToolDefineFunction() string {
	return `define require_tool
	@command -v $(1) >/dev/null 2>&1 || { \
		echo "$(1) is not installed."; \
		exit 1; \
	}; \
	VERSION="$$( \
		$(if $(strip $(4)),$(4),$(1) version) 2>/dev/null \
		| grep -Eo '[0-9]+(\.[0-9]+){0,2}' \
		| head -n1 \
	)"; \
	if [ -z "$$VERSION" ]; then \
		echo "Could not determine $(1) version."; \
		exit 1; \
	fi; \
	version_cmp() { \
		awk -v a="$$1" -v b="$$2" 'BEGIN { \
			split(a, A, "."); \
			split(b, B, "."); \
			for (i = 1; i <= 3; i++) { \
				x = A[i] + 0; \
				y = B[i] + 0; \
				if (x < y) { print -1; exit; } \
				if (x > y) { print 1; exit; } \
			} \
			print 0; \
		}'; \
	}; \
	if [ -n "$(2)" ] && [ "$$(version_cmp "$$VERSION" "$(2)")" -lt 0 ]; then \
		echo "$(1) version $(2) or greater is required. Installed: $$VERSION"; \
		exit 1; \
	fi; \
	if [ -n "$(3)" ] && [ "$$(version_cmp "$$VERSION" "$(3)")" -gt 0 ]; then \
		echo "$(1) version $(3) or lower is required. Installed: $$VERSION"; \
		exit 1; \
	fi
endef
`
}
