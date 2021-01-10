package main

import (
    "fmt"
	"log"
	"net"
	"net/http"
	"github.com/go-vgo/robotgo"
)

var down string;
var up string;

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func upHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
    if r.URL.Path != "/up" {
        http.Error(w, "404 non trovato.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Solo richieste GET.", http.StatusNotFound)
        return
	}
	robotgo.KeyTap(up)
    fmt.Fprintf(w, "up!")
}

func downHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
    if r.URL.Path != "/down" {
        http.Error(w, "404 non trovato.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Solo richieste GET.", http.StatusNotFound)
        return
    }
	robotgo.KeyTap(down)
    fmt.Fprintf(w, "down!")
}

func main() {
	http.HandleFunc("/up", upHandler)
	http.HandleFunc("/down", downHandler)
	fmt.Printf("Server in ascolto alla porta 8080\n")
	fmt.Printf("Nell'app inserisci l'indirizzo qui sotto che inizia per '192.168.1.xxx' o '192.168.0.xxx'\n")
    fmt.Printf("------------------------------------\n")
	ifaces, err := net.Interfaces()
if err != nil{
	log.Fatal(err)
}
for _, i := range ifaces {
    addrs,err := i.Addrs()
    if err != nil{
		log.Fatal(err)
	}
    for _, addr := range addrs {
        var ip net.IP
        switch v := addr.(type) {
        case *net.IPNet:
                ip = v.IP
        case *net.IPAddr:
                ip = v.IP
        }
        fmt.Printf("Possibile indirizzo ip: "+ip.String()+"\n")
    }
}
fmt.Printf("------------------------------------\n")
fmt.Printf("Tasto da premere per il volume giù SENZA VIRGOLETTE (consulta la lista su github per i tasti disponibili)\n")
fmt.Scanln(&down)
fmt.Printf("Tasto da premere per il volume su SENZA VIRGOLETTE (consulta la lista su github per i tasti disponibili)\n")
fmt.Scanln(&up)
fmt.Printf("----------PRONTO---------\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

