package levellog

import (
	"flag"
	"testing"
)

var (
	userLevelString = flag.String("l", "debug", "log level")
)

func TestInit(t *testing.T) {
	initLoggers(TRACE)
	if currentLevel != TRACE {
		t.Fatalf("expected %s, found %s", TRACE, currentLevel)
	}
}

func TestSetLevel(t *testing.T) {
	SetLevel("TRACE")
	if currentLevel != TRACE {
		t.Fatalf("expected %s, found %s", TRACE, currentLevel)
	}
}

func TestParseLevel(t *testing.T) {
	flag.Parse()
	l, err := ParseLevel(*userLevelString)
	if err != nil {
		t.Errorf("error parsing level: %v", err)
	}
	t.Logf("user level is : %s", l)
}

func TestPrintfLevel(t *testing.T) {
	flag.Parse()
	l, err := ParseLevel(*userLevelString)
	if err != nil {
		t.Errorf("error parsing level: %v", err)
	}
	initLoggers(WARN)
	Printf(l, "this is a message at level %s", l)
}
