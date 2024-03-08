package config

import (
	"fmt"
)

func (cnt *content) assemblePhonyToContent() {
	cnt.logger.Info("preparing PHONY string")

	strs := []string{".PHONY:"}
	strs = append(strs, cnt.config.Phony...)
	content := cnt.concatStringsWithWhiteSpaces(strs...)
	content += "\n"
	cnt.Phony = content

	cnt.logger.Info("PHONY string created successfully")
}

func (cnt *content) assembleDefaultToContent() {
	cnt.logger.Info("preparing Default string")

	var content SuggarString
	content.appendString("default:").appendString(" ").appendString("help")

	cnt.Default = content.getString()

	cnt.logger.Info("Default string created successfully")
}

func (cnt *content) assembleTargetsToContent() {
	cnt.logger.Info("preparing Target string")

	var content SuggarString
	for targetName, targetConfig := range cnt.config.Targets {
		if cnt.config.HelpTemplate.Enabled {
			content.appendString(cnt.config.HelpTemplate.Delimiter).appendString(" ").appendString(targetConfig.HelpDescription).lineBreak()
		}
		content.appendString(targetName + ":").appendString(" ")
		preTargets := cnt.concatStringsWithWhiteSpaces(targetConfig.PreTargets...)
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

	cnt.Targets = content.getString()

	cnt.logger.Info("Target string created successfully")
}

func (cnt *content) assembleAll() string {
	var content SuggarString

	content.appendString(cnt.Phony).lineBreak()
	for _, constant := range cnt.Constants {
		content.appendString(constant)
	}
	content.lineBreak().lineBreak()
	content.appendString(cnt.Default).lineBreak().lineBreak()
	content.appendString("preflight:").lineBreak().tab()
	for _, preflight := range cnt.Preflight {
		content.appendString(preflight).lineBreak()
	}
	content.appendString(cnt.Targets)
	content.appendString(cnt.Help)

	return content.getString()
}
