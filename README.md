# TSDNS
This is a TSDNS Implementation in GO.

---
[![WTFPL](https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/WTFPL_badge.svg/320px-WTFPL_badge.svg.png)](LICENSE.md)

This work is licensed under the "Do **W**hat **T**he **F**uck You Want To **P**ublic **L**icense". To see the full License Text go to the [LICENSE File](LICENSE.md)
 
---
## Example
Install: `go get github.com/Freaxbot/TSDNS`
```go
package main

import (
    "log"
    "github.com/Freaxbot/TSDNS"
)

func main()  {
    res :=goTSDNS.Lookup("public.teamspeak.com", "127.0.0.1")
    log.Println(("TSDNS :: " + res))
}
```
