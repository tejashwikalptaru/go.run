//go:build !android && !ios

package utils

func IsMobile() bool {
	return false
}
