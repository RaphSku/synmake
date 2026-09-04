# synmake
![Version](https://img.shields.io/badge/version-v0.3.0-orange)
![CI](https://github.com/RaphSku/synmake/workflows/Go%20CI/badge.svg)
## What is synmake?
You will not need synmake, when you have to write a simple Makefile but synmake should rather help you in times where you cannot remember, e.g. how to setup a preflight check for checking the minimum required version of a tool or how to print a help window that describes the available targets.

Nowadays, the majority of people are used to writing YAML. synmake allows you to write the specification in YAML and is doing the generation of the Makefile for you.

## How to install synmake
You can install synmake with the following command:
```bash
go install github.com/RaphSku/synmake@latest
```

## How to use synmake
To print synmake's version, use
```bash
synmake version
```
If you want to generate an example configuration YAML, simply run:
```bash
synmake generate config
```
Adjust it to your needs and then you can simply generate your Makefile via
```bash
synmake --config=<path/to/your/config.yaml>
```
The Makefile will be created in the directory in which you ran this command.
If something goes wrong, you can enable debugging with
```bash
synmake --config=<path/to/your/config.yaml> --debug
```

## Understanding the config schema
An example configuration file might look like this:
```yaml
---
variables:
  - key: SHELL
    value: /bin/bash
    operator: ImmediateAssignment
  - key: .SHELLFLAGS
    value: -eu -o pipefail -c
    operator: ImmediateAssignment
targets:
  - name: targetA
    helpDescription: targetA just prints an output
    commands:
      - echo "Hello World"
      - echo "This is how you specify commands!"
    display: false
  - name: targetB
    helpDescription: targetB just prints an output
    preTargets:
      - targetA
    commands:
      - echo "This is targetB!"
      - echo "How are you doing?"
    display: true
templates:
  helpTemplate:
    enabled: true
    delimiter: '##'
functions:
  requireTool:
    enabled: true
```

### Targets
Each item in targets specifies a new target where you can specify 4 attributes:
- `name` -> The name of the target.
- `helpDescription` -> This field describes the target and will be added above the target as documentation. If you have set
  `helpTemplate.enabled` to `true`, the `helpTemplate.delimiter` will override the default comment delimiter which is `#`.
- `commands` -> Here you can specify the commands that should be appear under the given target.
- `display` -> If `display` is `true`, the commands will be printed when you run the given target, otherwise they are
  supressed by the command modifier `@`.

### Templates
Templates are additional helper targets that can be optionally added to the Makefile. Currently, 1 templates is available:
- `helpTemplate` -> It adds a help target to the Makefile, such that when you run `make`, a help description is shown
  with all the available targets and their help description.

### Functions
- `require_tool` verifies that a required CLI tool is installed and optionally checks that its version falls within a specified minimum and/or maximum range. By default, the version is determined using `<tool> version`, but a custom version command can be provided for tools that use a different version syntax.
