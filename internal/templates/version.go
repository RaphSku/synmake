package templates

import "fmt"

func GetVersionTemplate(library string, minVersion string) (string, []string) {
	constants := []string{fmt.Sprintf("MIN_VERSION=%s", minVersion)}
	content := fmt.Sprintf(`@if ! test -x "$(shell which %[1]s)"; then \
		echo "%[1]s is not installed. Please install it. Minimum Version required: %[2]s"; \
		exit 1; \
	elif [ "$$(($(shell %[1]s version | cut -d ' ' -f2 | cut -d '.' -f1 | grep -E -o "[0-9]+") >= $(shell echo $(MIN_VERSION) | cut -d '.' -f1)))" -eq 0 ] || \
		[ "$$(($(shell %[1]s version | cut -d ' ' -f2 | cut -d '.' -f2) >= $(shell echo $(MIN_VERSION) | cut -d '.' -f2)))" -eq 0 ] || \
		[ "$$(($(shell %[1]s version | cut -d ' ' -f2 | cut -d '.' -f3 | grep -E -o "^[0-9]+") >= $(shell echo $(MIN_VERSION) | cut -d '.' -f3)))" -eq 0 ]; then \
		echo "%[1]s version $(MIN_VERSION) or greater is required!"; \
		exit 1; \
	fi
`, library, minVersion)

	return content, constants
}
