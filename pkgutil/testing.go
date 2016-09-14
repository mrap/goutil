package pkgutil

import "fmt"

func TestResult(actual, expected interface{}, labelFormat string, formatArgs ...interface{}) string {
	return fmt.Sprintf("%s => %v. Expected: %v", fmt.Sprintf(labelFormat, formatArgs...), actual, expected)
}
