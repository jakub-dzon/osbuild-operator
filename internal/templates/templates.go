package templates

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"

	osbuilderprojectflottaiov1alpha1 "github.com/project-flotta/osbuild-operator/api/v1alpha1"
)

func ProcessOSBuildConfigTemplate(textTemplate string, expectedParameters []osbuilderprojectflottaiov1alpha1.Parameter,
	values []osbuilderprojectflottaiov1alpha1.ParameterValue) (string, error) {

	keyTypes := make(map[string]string)
	keyValues := make(map[string]string)

	for _, p := range expectedParameters {
		keyTypes[p.Name] = p.Type
		keyValues[p.Name] = p.DefaultValue
	}

	for _, v := range values {
		if kType, ok := keyTypes[v.Name]; ok {
			if !validateParameter(v.Value, kType) {
				return "", fmt.Errorf("parameter %s of type %s was given %s value, which can't be represented as %[2]s", v.Name, kType, v.Value)
			}
			keyValues[v.Name] = v.Value
		}
	}

	t, err := template.New("template").Parse(textTemplate)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, keyValues)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func validateParameter(value, pType string) bool {
	switch pType {
	case "bool":
		_, err := strconv.ParseBool(value)
		return err == nil
	case "int":
		_, err := strconv.Atoi(value)
		return err == nil
	default:
		// That's a string
		return true
	}
}
