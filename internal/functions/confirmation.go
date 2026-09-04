package functions

func GetRequireConfirmFunction() string {
	return `define confirm
	@printf "$(1) [y/N] "; \
	read answer; \
	case "$$answer" in \
		y|Y|yes|YES) ;; \
		*) echo "Aborted."; exit 1 ;; \
	esac
endef
`
}
