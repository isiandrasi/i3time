// i3time outputs an endless json stream of curent date and time to stdout.
package main

import (
    "os"
    "time"
)

const prefix = `{"version":1}[`
const layout = `
[{"full_text":"📅 2006-01-02 🕒 15:04:05"}],`

func main() {
    writer := os.Stdout
    writer.Write([]byte(prefix))
    buffer := make([]byte, len(layout))
    for {
        buffer = buffer[:0]
        buffer = time.Now().AppendFormat(buffer, layout)
        writer.Write(buffer)
        time.Sleep(time.Second)
    }
}
