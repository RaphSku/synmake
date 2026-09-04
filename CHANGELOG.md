# Changelog

## v0.3.0
### Features
- New `variables` config field allowing the user to specify variables with different assignment operators.
- Introducing the following `functions`: require environment variable function, require
file function, require directory function, require confirmation function, require tool version function.

### Changed
- `targets` is an array now and not a map and therefore the `name` is part of the target configuration.
- There is no `versionTemplate` anymore, instead there is a better `require_tool` function that can be used in targets.

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
