package main
import "fmt"

func main() {

    var w1, w2, w3, w4 string
    var percobaan int
    berhasil := true

    percobaan = 1

    for percobaan <= 5 {

    fmt.Print("Percobaan ", percobaan, ": ")
    fmt.Scan(&w1, &w2, &w3, &w4)

    if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
            berhasil = false
        }

    percobaan = percobaan + 1
    }

    fmt.Println("berhasil:", berhasil)
}
