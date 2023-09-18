//go:build windows

package launcher

import (
	"log"
	"os/exec"
	"syscall"

	"github.com/kolesnikovae/go-winjob"
	"golang.org/x/sys/windows"
)

func killGroup(pid int) {
	terminateProcess(pid)
}

func (l *Launcher) osSetupCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: windows.CREATE_SUSPENDED | windows.CREATE_NEW_PROCESS_GROUP,
	}
}

func (l *Launcher) afterStartOsSetupCmd(cmd *exec.Cmd) {
	job, _ := winjob.Create("",
		winjob.LimitKillOnJobClose,
		winjob.LimitBreakawayOK)
	if err := job.Assign(cmd.Process); err != nil {
		log.Fatal(err)
	}
	if err := winjob.ResumeProcess(cmd.Process.Pid); err != nil {
		log.Fatal(err)
	}
}

func terminateProcess(pid int) {
	handle, err := syscall.OpenProcess(syscall.PROCESS_TERMINATE, true, uint32(pid))
	if err != nil {
		return
	}

	_ = syscall.TerminateProcess(handle, 0)
	_ = syscall.CloseHandle(handle)
}
