package main

// Ver 0.1

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/XKBlib.h>
#include <stdbool.h>

bool isNumLockOn() {
    Display* display = XOpenDisplay(NULL);
    unsigned int n;
    XkbGetIndicatorState(display, XkbUseCoreKbd, &n);
    XCloseDisplay(display);
    return (n & 0x02) == 2;
}
*/
import "C"
import (
	"github.com/getlantern/systray"
	"io/ioutil"
	"log"
	"time"
)

func checkNumLock() bool {
	return bool(C.isNumLockOn())
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	var numLockActive = checkNumLock()
	var numLockIcon []byte
	var err error

	if numLockActive {
		numLockIcon, err = ioutil.ReadFile("/usr/share/numstatus/num_on.png")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		numLockIcon, err = ioutil.ReadFile("/usr/share/numstatus/num_off.png")
		if err != nil {
			log.Fatal(err)
		}
	}

	systray.SetIcon(numLockIcon)
	// systray.SetTitle("NumLock Status")

	mQuitOrig := systray.AddMenuItem("Выход", "Выход из программы")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()

	go func() {
		for {
			time.Sleep(400 * time.Millisecond)
			numLockActive := checkNumLock()
			if numLockActive {
				numLockIcon, err = ioutil.ReadFile("/usr/share/numstatus/num_on.png")
				if err != nil {
					log.Fatal(err)
				}
			} else {
				numLockIcon, err = ioutil.ReadFile("/usr/share/numstatus/num_off.png")
				if err != nil {
					log.Fatal(err)
				}
			}
			systray.SetIcon(numLockIcon)
		}
	}()
}

func onExit() {
	// Очистка ресурсов перед выходом
}
