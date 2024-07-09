package parser

import (
	"regexp"
	"strings"
)

// msg define
var msgDefine = regexp.MustCompile(`^\s*@msg\s+[a-zA-Z_]+\s*$`)

// common define
var comDefine = regexp.MustCompile(`^\s*@common\s+[a-zA-Z_ ]+$`)

// nobind define
var noBindDefine = regexp.MustCompile(`^\s*@nobind\s*$`)

// handle define, if handle is define , not have nobind
var handleDefine = regexp.MustCompile(`^\s*@handle\s*$`)

func IsMsg(ss string) bool {
	return MsgParser(ss) != nil
}

func MsgParser(ss string) *Msg {
	lines := strings.Split(ss, "\n")
	var svm string
	var nbd, handle bool
	for _, line := range lines {
		if msgDefine.MatchString(line) {
			svm = msgDefine.FindString(line)
		}
		if !nbd {
			nbd = noBindDefine.MatchString(line)
		}
		if !handle {
			handle = handleDefine.MatchString(line)
		}
	}
	if svm == "" {
		return nil
	}
	return &Msg{
		ID:     strings.TrimSpace(strings.ReplaceAll(svm, "@msg", "")),
		NoBind: nbd,
		Handle: handle,
	}
}

type Msg struct {
	ID     string
	Name   string
	NoBind bool
	Handle bool
}

func CommonsParser(ss string) []string {
	lines := strings.Split(ss, "\n")
	var svm string
	for _, line := range lines {
		if comDefine.MatchString(line) {
			svm = comDefine.FindString(line)
			break
		}
	}

	cmids := strings.Split(strings.ReplaceAll(svm, "@common", ""), " ")
	var rest []string
	for _, id := range cmids {
		if isEmpty(id) {
			continue
		}
		rest = append(rest, strings.TrimSpace(id))
	}
	return rest
}

func isEmpty(s string) bool {
	if s == "" {
		return true
	}
	if strings.TrimSpace(s) == "" {
		return true
	}
	return false
}
