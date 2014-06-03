package event

import (
	"github.com/bit4bit/gami"
	"testing"
)

func TestAgentLogoff(t *testing.T) {
	fixture := map[string]string{
		"Agent": "Agent",
		"Uniqueid": "UniqueID",
		"Logintime": "LoginTime",
	}

	ev := gami.AMIEvent{
		Id: "AgentLogoff",
		Privilege: []string{"all"},
		Params: fixture,
	}
	
	evtype := New(&ev)
	if _, ok := evtype.(AgentLogoff); !ok {
		t.Fatal("AgentLogoff type assertion")
	}
	
	testEvent(t, fixture, evtype)
}
