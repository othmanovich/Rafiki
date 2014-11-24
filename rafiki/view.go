

package rafiki


import (
        "os"
        "os/exec"
        "fmt"
        )


func ClearScreen() {

    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()

}



func PrintOrange(msg string) {

        fmt.Printf("\x1b[31;1m%s\x1b[0m\n", msg)

}


