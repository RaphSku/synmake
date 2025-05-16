# Changelog

## v0.2.0
### Changed
- Restructured/Refactored code base and removed unnecessary intermediary components
- PHONYs will be added for every given target for now.
  This will change in the future for targets that are actually files.
- New flag `debug` added for `synmake --config=<path_to_config.yaml> --debug` that 
  will run synmake in `debug` mode

## v0.1.1
### Features
- Print the version of synmake with `synmake version`
- Generate example config with `synmake generate config`
- Parse config and generate the Makefile with `synmake --config=<path_to_config.yaml>`
- 2 provided templates: a help template and a min version preflight check template
- currently available: PHONY, variables, targets, help
