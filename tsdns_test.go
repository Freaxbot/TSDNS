package goTSDNS_test

import (
    "log"
    "testing"
    "github.com/Freaxbot/TSDNS"
)

func TestLookup(t *testing.T) {
    res :=goTSDNS.Lookup("public.teamspeak.com", "127.0.0.1")
    log.Println(("TSDNS :: " + res))

}
