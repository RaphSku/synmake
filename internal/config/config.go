package config

import (
	"strings"

	"go.uber.org/zap"
)

type VariableOperator string

func (o VariableOperator) Symbol() string {
	switch o {
	case Assignment:
		return "="
	case ImmediateAssignment:
		return ":="
	case ConditionalAssignment:
		return "?="
	case AppendingAssignment:
		return "+="
	case ShellAssignment:
		return "!="
	default:
		return ""
	}
}

const (
	Assignment            VariableOperator = "Assignment"
	ImmediateAssignment   VariableOperator = "ImmediateAssignment"
	ConditionalAssignment VariableOperator = "ConditionalAssignment"
	AppendingAssignment   VariableOperator = "AppendingAssignment"
	ShellAssignment       VariableOperator = "ShellAssignment"
)

var assignmentOperators = []VariableOperator{Assignment, ImmediateAssignment, ConditionalAssignment, AppendingAssignment, ShellAssignment}

func JoinAssignmentOperators() string {
	parts := make([]string, len(assignmentOperators))
	for i, v := range assignmentOperators {
		parts[i] = string(v)
	}

	return strings.Join(parts, ", ")
}

type Variable struct {
	Key      string           `yaml:"key"`
	Value    string           `yaml:"value"`
	Operator VariableOperator `yaml:"operator"`
	Export   bool             `yaml:"export,omitempty"`
}

type Target struct {
	Name            string   `yaml:"name"`
	HelpDescription string   `yaml:"helpDescription"`
	PreTargets      []string `yaml:"preTargets,omitempty"`
	Commands        []string `yaml:"commands"`
	Display         bool     `yaml:"display"`
}

type HelpTemplate struct {
	Enabled   bool   `yaml:"enabled"`
	Delimiter string `yaml:"delimiter"`
}

type OptionalTemplates struct {
	HelpTargetTemplate HelpTemplate `yaml:"helpTemplate"`
}

type RequireToolFunction struct {
	Enabled bool `yaml:"enabled"`
}

type RequireEnvFunction struct {
	Enabled bool `yaml:"enabled"`
}

type RequireFileFunction struct {
	Enabled bool `yaml:"enabled"`
}

type RequireDirFunction struct {
	Enabled bool `yaml:"enabled"`
}

type RequireConfirmationFunction struct {
	Enabled bool `yaml:"enabled"`
}

type OptionalFunctions struct {
	RequireToolFunction         RequireToolFunction         `yaml:"requireTool"`
	RequireEnvFunction          RequireEnvFunction          `yaml:"requireEnv"`
	RequireFileFunction         RequireFileFunction         `yaml:"requireFile"`
	RequireDirFunction          RequireDirFunction          `yaml:"requireDir"`
	RequireConfirmationFunction RequireConfirmationFunction `yaml:"requireConfirm"`
}

type Config struct {
	Variables []Variable        `yaml:"variables"`
	Targets   []Target          `yaml:"targets"`
	Templates OptionalTemplates `yaml:"templates"`
	Functions OptionalFunctions `yaml:"functions"`
}

type ConfigManager struct {
	logger *zap.Logger
	Config Config
	File   FileInterface
}

func NewConfigManager(logger *zap.Logger, file FileInterface) (*ConfigManager, error) {
	return &ConfigManager{
		logger: logger,
		Config: Config{},
		File:   file,
	}, nil
}

func (cm *ConfigManager) Close() error {
	return cm.File.Close()
}
