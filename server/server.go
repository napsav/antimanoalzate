package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

var down string
var up string

const (
	SIZE_W = 350
	SIZE_H = 200
)

type MyMainWindow struct {
	*walk.MainWindow
}

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

var upTextEdit, downTextEdit, ipTextEdit *walk.TextEdit
var tasto *walk.PushButton

func main() {
	messaggio := "AntiManoAlzate - v1.0.2 by Saverio Napolitano\n\n"

	http.HandleFunc("/up", upHandler)
	http.HandleFunc("/down", downHandler)

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
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
			if strings.Contains(ip.String(), "192.168.1.") || strings.Contains(ip.String(), "192.168.0.") {
				fmt.Printf("Possibile indirizzo ip: " + ip.String() + "\n")
				ipTextEdit.SetText(ip.String())
			}
		}
	}

	mw := new(MyMainWindow)
	MainWindow{
		Visible:  false,
		AssignTo: &mw.MainWindow,
		Title:    "AntiManoAlzate - Server 1.1",
		Layout:   VBox{},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					Label{Text: messaggio},
					Label{Text: "Nell'app inserisci l'indirizzo che trovi sotto che inizia per '192.168.1.xxx' o '192.168.0.xxx'\nSe non funziona il primo, prova altri indirizzi."},
					Label{Text: "Tasto da premere per il volume giù SENZA VIRGOLETTE (consigliato: pagedown)\n (consulta la lista su github per i tasti disponibili)"},
					TextEdit{AssignTo: &downTextEdit},
					Label{
						Text: "Tasto da premere per il volume su SENZA VIRGOLETTE (consigliato: pageup)\n (consulta la lista su github per i tasti disponibili)",
					},
					TextEdit{AssignTo: &upTextEdit},
					Label{
						Text: "Indirizzo IP",
					},
					TextEdit{AssignTo: &ipTextEdit, ReadOnly: true},
				},
			},
			PushButton{
				AssignTo:  &tasto,
				Text:      "Inizia",
				OnClicked: mw.test,
			},
		},
	}.Create()

	defaultStyle := win.GetWindowLong(mw.Handle(), win.GWL_STYLE) // Gets current style
	newStyle := defaultStyle &^ win.WS_THICKFRAME                 // Remove WS_THICKFRAME
	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, newStyle)

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)

	mw.Run()

}

func server() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var running bool

func (mw *MyMainWindow) test() {
	if running {
		walk.MsgBox(mw, "Errore", "Server già in esecuzione", walk.MsgBoxIconError)
	} else if upTextEdit.Text() == "" || downTextEdit.Text() == "" {
		walk.MsgBox(mw, "Errore", "Non hai impostato i tasti!", walk.MsgBoxIconError)
	} else {
		go server()
		up = upTextEdit.Text()
		down = downTextEdit.Text()
		running = true
		tasto.SetText("Avviato con successo!")
	}
}
