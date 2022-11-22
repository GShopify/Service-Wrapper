package model

import (
	"fmt"
	"regexp"
)

func NewId(namespace Gid, id string) string {
	return fmt.Sprintf("gid://gshopify/%s/%s", namespace, id)
}

func ParseId(namespace Gid, id string) (string, error) {
	s := regexp.MustCompile(`^gid://gshopify/(\w+)/([\w-]+)$`).FindAllStringSubmatch(id, -1)
	if len(s) != 1 || len(s[0]) != 3 {
		return "", fmt.Errorf("illegal ID format: %s", id)
	}

	if Gid(s[0][1]) != namespace {
		return "", fmt.Errorf("illegal ID namespace: %s", namespace)
	}

	return s[0][2], nil
}
