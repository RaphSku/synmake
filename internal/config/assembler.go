package config

import (
	"fmt"
)

func concatStringsWithWhiteSpaces(strs ...string) string {
	result := ""
	for i := range strs {
		result += strs[i]
		if i != len(strs)-1 {
			result += " "
		}
	}

	return result
}

func assembleVariables(variables []string) SuggarString {
	var content SuggarString
	for i := range variables {
		content.appendString(variables[i]).lineBreak()
	}

	return content
}

func assembleDefaultTarget() SuggarString {
	var content SuggarString
	content.appendString(".PHONY: default").lineBreak()
	content.appendString("default:").appendString(" ").appendString("help")

	return content
}

func assemblePreflightTarget(commands []string) SuggarString {
	var content SuggarString
	content.appendString(".PHONY: preflight").lineBreak()
	content.appendString("preflight:").lineBreak().tab()
	for _, command := range commands {
		content.appendString(command).lineBreak()
	}

	return content
}

func assembleTargets(targetMap map[string]Target, delimiter string) SuggarString {
	helpDelimiter := ""
	if delimiter != "" {
		helpDelimiter = delimiter
	}

	var content SuggarString
	for targetName, targetConfig := range targetMap {
		content.appendString(".PHONY:").appendString(" ").appendString(targetName).lineBreak()
		content.appendString(helpDelimiter).appendString(" ").appendString(targetConfig.HelpDescription).lineBreak()
		content.appendString(targetName + ":").appendString(" ")
		preTargets := concatStringsWithWhiteSpaces(targetConfig.PreTargets...)
		content.appendString(preTargets).lineBreak().tab()
		for index, command := range targetConfig.Commands {
			if targetConfig.Display {
				content.appendString(fmt.Sprintf("@%s", command)).lineBreak()
				if index != len(targetConfig.Commands)-1 {
					content.tab()
				}
				continue
			}
			content.appendString(command).lineBreak()
			if index != len(targetConfig.Commands)-1 {
				content.tab()
			}
		}
		content.lineBreak()
	}

	return content
}
