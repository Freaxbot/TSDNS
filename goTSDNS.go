package goTSDNS

import (
    "strings"
    "strconv"
    "net"
)

const DEFAUL_TSDNS_PORT = 41144;

/*/
    This is a TSDNS client. To check the DNS records from a TSDNS server.

    status code
    error | error (can't find TSDNS or IP is wrong) the error message is set to the addr string.
    404 | 404 TSDNS can't find the right domain.
    204 | no response TSDNS. Ignored our request..
 */



func Lookup(domain string, addr string) string  {
    strEcho := domain
    servAddr := addr;
    if (!(strings.Contains(addr, ":"))){
        servAddr = addr + ":" + strconv.Itoa(DEFAUL_TSDNS_PORT);
    }

    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        //println("ResolveTCPAddr failed:", err.Error())
        return "ResolveTCPAddr failed:" +  err.Error()
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        //println("Dial failed:" +  err.Error());
        return "error"
    }

    _, err = conn.Write([]byte(strEcho))
    if err != nil {
        //println("Socket closed:" +  err.Error());
        return "error"
    }

    reply := make([]byte, 512)
    _, err = conn.Read(reply)
    if err != nil {
        return "204"
    }
    conn.Close()

    return string(reply)
}
