package config

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var placeholderRegexp = regexp.MustCompile(`\$\{(.*?)}`)

// resolvePlaceholders preserves the type declared by a complete placeholder's
// default value. Kratos' default resolver always returns strings, which makes
// protojson reject values for bool and numeric fields during Config.Scan.
func resolvePlaceholders(input map[string]any) error {
	return resolvePlaceholderMap(input, input)
}

func resolvePlaceholderMap(values map[string]any, root map[string]any) error {
	for key, value := range values {
		switch typed := value.(type) {
		case string:
			resolved, err := resolvePlaceholderString(typed, root)
			if err != nil {
				return err
			}
			values[key] = resolved
		case map[string]any:
			if err := resolvePlaceholderMap(typed, root); err != nil {
				return err
			}
		case []any:
			for index, item := range typed {
				switch item := item.(type) {
				case string:
					resolved, err := resolvePlaceholderString(item, root)
					if err != nil {
						return err
					}
					typed[index] = resolved
				case map[string]any:
					if err := resolvePlaceholderMap(item, root); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func resolvePlaceholderString(value string, root map[string]any) (any, error) {
	matches := placeholderRegexp.FindAllStringSubmatchIndex(value, -1)
	if len(matches) == 0 {
		return value, nil
	}

	// A placeholder embedded in text is always a string, for example host:port
	// and connection URLs. Only a whole scalar can safely retain a YAML type.
	if len(matches) == 1 && matches[0][0] == 0 && matches[0][1] == len(value) {
		expression := value[matches[0][2]:matches[0][3]]
		resolved, fallback, hasFallback := resolvePlaceholderText(expression, root)
		if !hasFallback || isQuoted(fallback) {
			return resolved, nil
		}

		if fallback == "true" || fallback == "false" {
			parsed, err := strconv.ParseBool(resolved)
			if err != nil {
				return nil, fmt.Errorf("placeholder %q requires a boolean value: %w", expression, err)
			}
			return parsed, nil
		}
		if _, err := strconv.ParseInt(fallback, 10, 64); err == nil {
			parsed, err := strconv.ParseInt(resolved, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("placeholder %q requires an integer value: %w", expression, err)
			}
			return parsed, nil
		}
		if _, err := strconv.ParseFloat(fallback, 64); err == nil {
			parsed, err := strconv.ParseFloat(resolved, 64)
			if err != nil {
				return nil, fmt.Errorf("placeholder %q requires a number value: %w", expression, err)
			}
			return parsed, nil
		}
		return resolved, nil
	}

	return placeholderRegexp.ReplaceAllStringFunc(value, func(match string) string {
		name := match[2 : len(match)-1]
		resolved, _, _ := resolvePlaceholderText(name, root)
		return resolved
	}), nil
}

func resolvePlaceholderText(expression string, root map[string]any) (value string, fallback string, hasFallback bool) {
	name, fallback, hasFallback := strings.Cut(strings.TrimSpace(expression), ":")
	current := root
	parts := strings.Split(name, ".")
	for index, part := range parts {
		configured, ok := current[part]
		if !ok {
			break
		}
		if index == len(parts)-1 {
			return fmt.Sprint(configured), fallback, hasFallback
		}
		next, ok := configured.(map[string]any)
		if !ok {
			break
		}
		current = next
	}
	if configured, ok := os.LookupEnv(name); ok {
		return configured, fallback, hasFallback
	}
	if isQuoted(fallback) {
		return fallback[1 : len(fallback)-1], fallback, hasFallback
	}
	return fallback, fallback, hasFallback
}

func isQuoted(value string) bool {
	return len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"'
}
