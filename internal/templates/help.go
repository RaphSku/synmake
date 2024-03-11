package templates

func GetHelpTemplate() string {
	content := `help:
	@echo "----------------------------------"
	@echo "Welcome to make! Enjoy the flight."
	@echo "Makefile - make [\033[38;5;154mtarget\033[0m]"
	@echo "----------------------------------"
	@echo
	@echo "Targets:"
	@awk '/^[a-zA-z\-_0-9%:\\]+/ { \
		description = match(descriptionLine, /^## (.*)/); \
		if (description) { \
			target = $$1; \
			description = substr(descriptionLine, RSTART + 3, RLENGTH); \
			gsub("\\\\", "", target); \
			gsub(":+$$", "", target); \
			printf "    \033[38;5;154m%-15s\033[0m %s\n", target, description; \
		} \
	} \
	{ descriptionLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"
`

	return content
}
