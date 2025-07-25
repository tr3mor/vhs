//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package main

import (
	"fmt"
	"syscall"
)

func dropUserPrivileges(gid int, uid int) error {
	if err := syscall.Setgid(gid); err != nil {
		return fmt.Errorf("setgid error: %s", err)
	}
	if err := syscall.Setuid(uid); err != nil {
		return fmt.Errorf("setuid error: %s", err)
	}
	return nil
}
