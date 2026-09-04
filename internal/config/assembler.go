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

func assembleVariables(variables []Variable) (SuggarString, error) {
	var content SuggarString
	for _, variable := range variables {
		if variable.Export {
			content.appendString(fmt.Sprintf("%s ", "export"))
		}
		operator := variable.Operator.Symbol()
		if operator == "" {
			return content, fmt.Errorf("The specified operator '%s' does not exist, list of available operators: %s", variable.Operator, JoinAssignmentOperators())
		}
		content.appendString(fmt.Sprintf("%s %s %s", variable.Key, operator, variable.Value)).lineBreak()
	}

	return content, nil
}

func assembleDefaultTarget() SuggarString {
	var content SuggarString
	content.appendString(".PHONY: default").lineBreak()
	content.appendString("default:").appendString(" ").appendString("help")

	return content
}

func assembleRequireToolFunction(functionContent string) SuggarString {
	var content SuggarString
	content.appendString(functionContent)

	return content
}

func assembleRequireEnvFunction(functionContent string) SuggarString {
	var content SuggarString
	content.appendString(functionContent)

	return content
}

func assembleRequireFileFunction(functionContent string) SuggarString {
	var content SuggarString
	content.appendString(functionContent)

	return content
}

func assembleRequireDirFunction(functionContent string) SuggarString {
	var content SuggarString
	content.appendString(functionContent)

	return content
}

func assembleRequireConfirmationFunction(functionContent string) SuggarString {
	var content SuggarString
	content.appendString(functionContent)

	return content
}

func assemblePreflightTarget(commands []string) SuggarString {
	var content SuggarString
	if len(commands) == 0 {
		return content
	}

	content.appendString(".PHONY: preflight").lineBreak()
	content.appendString("preflight:").lineBreak().tab()
	for _, command := range commands {
		content.appendString(command).lineBreak()
	}

	return content
}

func assembleTargets(targets []Target, delimiter string) SuggarString {
	var content SuggarString
	for _, targetConfig := range targets {
		targetName := targetConfig.Name
		content.appendString(".PHONY:").appendString(" ").appendString(targetName).lineBreak()
		content.appendString(delimiter).appendString(" ").appendString(targetConfig.HelpDescription).lineBreak()
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
