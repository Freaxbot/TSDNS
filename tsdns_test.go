package goTSDNS

import (
    "github.com/Freaxbot/TSDNS"
    "log"
    "testing"
)

func TestLookup(t *testing.T) {
    res :=goTSDNS.Lookup("public.teamspeak.com", "127.0.0.1")
    log.Println(("TSDNS :: " + res))

}
