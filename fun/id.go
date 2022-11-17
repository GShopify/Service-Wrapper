package fun

import (
	"fmt"
	"regexp"
	"strings"
)

func NewId(namespace string, id string) string {
	return fmt.Sprintf("gid://gshopify/%s/%s", strings.ToLower(namespace), id)
}

func ParseId(id string, namespace string) (string, error) {
	s := regexp.MustCompile(`^gid://gshopify/(\w+)/(\w+)$`).FindAllStringSubmatch(id, -1)
	if len(s) != 1 || len(s[0]) != 3 {
		return "", fmt.Errorf("illegal ID format: %s", id)
	}

	if s[0][1] != namespace {
		return "", fmt.Errorf("illegal ID namespace: %s", namespace)
	}

	return s[0][2], nil
}
