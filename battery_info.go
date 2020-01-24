package main
import (
    "fmt"
    "strings"
    "io/ioutil"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
    bcap, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
    check(err)
    cap := strings.TrimSpace(string(bcap))
    bstate, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/status")
    check(err)
    state := strings.TrimSpace(string(bstate))
    fmt.Print("\r\n",state,":",cap,"%","\r\n","\r\n")
}
