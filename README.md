# synmake
![Version](https://img.shields.io/badge/version-v0.1.1-orange)
![CI](https://github.com/RaphSku/synmake/workflows/Go%20CI/badge.svg)
## What is synmake?
You will not need synmake, when you have to write a simple Makefiles but synmake should rather help you in times where you cannot remember, e.g. how to setup a preflight check for checking the minimum required version of a tool or how to print a help window that describes the available targets. 

Nowadays, the majority of people are used to writing YAML. synmake allows you to write the specification in YAML and is doing the generation of the Makefile for you.

## How to install synmake
You can either install synmake via Go1.21+ with the following command:
```bash
go install github.com/RaphSku/synmake@latest
```

Alternatively, check the release page and install the binary. There are binaries provided for Windows, MacOS and Linux.

## How to use synmake
If you want to check which version you are using, just check with
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

## Understanding the config schema
The config file might look like this:
```yaml
phony:
  - default
  - preflight
  - targetA
  - targetB
  - help
targets:
  targetA:
    helpDescription: targetA just prints an output
    commands:
      - echo "Hello World"
      - echo "This is how you specify commands!"
    display: false
  targetB:
    helpDescription: targetB just prints an output
    preTargets:
      - targetA
    commands:
      - echo "This is targetB!"
      - echo "How are you doing?"
    display: true
helpTemplate:
  enabled: true
  delimiter: '##'
versionTemplate:
  enabled: true
  library: example
  minVersion: 0.1.0
```
1. Phony Targets:

The phony section lists the names of phony targets (targets that are not actual files or commands) such as default, preflight, targetA, targetB, and help.

2. Targets:

Each target (targetA, targetB) has a helpDescription field providing a brief description of what the target does.
The commands field lists the commands to be executed when the target is invoked.
The display field specifies whether the commands should be displayed when running a target. If you specify `Display` as false, the commands will be shown, as well as the resulting output.
For targetB, there is a preTargets field that specifies dependencies on other targets (targetA in this case).

3. Help and Version Templates:

The helpTemplate section configures the help template:
enabled specifies whether the help template is enabled or not (true/false).
delimiter specifies the delimiter used in the help template.
The versionTemplate section configures the version template:
enabled specifies whether the version template is enabled or not (true/false).
library specifies the library or module used for checking the version.
minVersion specifies the minimum version required for the library/module.
