package config

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io"

	"go.uber.org/zap"
)

func GenerateExampleYamlConfig(logger *zap.Logger, w io.StringWriter) error {
	var variables []Variable

	variableA := Variable{
		Key:      "SHELL",
		Value:    "/bin/bash",
		Operator: ImmediateAssignment,
	}
	variables = append(variables, variableA)

	variableB := Variable{
		Key:      ".SHELLFLAGS",
		Value:    "-eu -o pipefail -c",
		Operator: ImmediateAssignment,
	}
	variables = append(variables, variableB)

	tA := Target{
		Name:            "targetA",
		HelpDescription: "targetA just prints an output",
		Commands:        []string{"echo \"Hello World\"", "echo \"This is how you specify commands!\""},
		Display:         false,
	}

	tB := Target{
		Name:            "targetB",
		HelpDescription: "targetB just prints an output",
		PreTargets:      []string{"targetA"},
		Commands:        []string{"echo \"This is targetB!\"", "echo \"How are you doing?\""},
		Display:         true,
	}

	config := Config{
		Variables: variables,
		Targets:   []Target{tA, tB},
		Templates: OptionalTemplates{
			HelpTemplate{
				Enabled:   true,
				Delimiter: "##",
			},
		},
		Functions: OptionalFunctions{
			RequireToolFunction{
				Enabled: true,
			},
			RequireEnvFunction{
				Enabled: true,
			},
			RequireFileFunction{
				Enabled: true,
			},
			RequireDirFunction{
				Enabled: true,
			},
			RequireConfirmationFunction{
				Enabled: true,
			},
		},
	}

	buffer := &bytes.Buffer{}
	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(2)
	err := encoder.Encode(config)
	if err != nil {
		logger.Error("Config could not be marshalled", zap.String("func", "GenerateExampleYamlConfig"), zap.Error(err))
		return err
	}

	yamlTemplate := YAML_DOCUMENT_SEPARATOR + buffer.String()
	_, err = w.WriteString(yamlTemplate)
	if err != nil {
		logger.Error("Failed to write to config.yaml", zap.String("func", "GenerateExampleYamlConfig"), zap.Error(err))
		return err
	}

	return nil
}
