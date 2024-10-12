package sge

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var sgeJobId = regexp.MustCompile(`^Your job (\d+) \("\S+"\) has been submitted\n$`)

// WrapSubmit
func WrapSubmit(submit, script, hjid string, submitArgs []string) (jid string) {
	if hjid != "" {
		submitArgs = append(submitArgs, "-hold_jid", hjid)
	}
	var cmds = append(submitArgs, script)
	var c = exec.Command(submit, cmds...)
	log.Printf("%s [%s]", submit, strings.Join(cmds, "]["))
	var submitLogBytes, err = c.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %v:[%v]", err, submitLogBytes)
	}
	// Your job (\d+) \("script"\) has been submitted
	log.Print(submitLogBytes)
	var submitLogs = sgeJobId.FindStringSubmatch(string(submitLogBytes))
	if len(submitLogs) == 2 {
		jid = submitLogs[1]
	} else {
		log.Fatalf("Error: jid parse error:%v->%v", submitLogBytes, submitLogs)
	}
	return
}

// Run wrap exec.Command(name,args...).Run() with stdout stderr redirect to os.Stdout os.Stderr
func Run(name string, args ...string) error {
	var cmd = exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// add os.Envrion
	// cmd.Env = os.Environ()
	return cmd.Run()
}
