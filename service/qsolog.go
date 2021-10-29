package service

import (
	"encoding/gob"
	"github.com/s51ds/n1mmweb/udp"
	"log"
	"os"
	"sync"
)

func init() {
	loadQsoLog()

}

var (
	QsoLogFileName = "NOV-2021.log"
	qsoLog         = make(map[string]udp.QsoInfo)
	qsoLogMu       sync.RWMutex
)

func saveQsoLog() {
	if logFile, err := os.Create(QsoLogFileName); err != nil {
		wd, _ := os.Getwd()
		log.Fatalln("saveQsoLog", err.Error(), wd+QsoLogFileName)
	} else {
		defer logFile.Close()
		encoder := gob.NewEncoder(logFile)
		qsoLogMu.RLock()
		defer qsoLogMu.RUnlock()
		if err := encoder.Encode(&qsoLog); err != nil {
			log.Fatalln("saveQsoLog", err.Error())
		}

	}

}

func loadQsoLog() {
	if logFile, err := os.Open(QsoLogFileName); err != nil {
		wd, _ := os.Getwd()
		log.Println("loadQsoLog", err.Error(), wd+QsoLogFileName)
	} else {
		defer logFile.Close()
		decoder := gob.NewDecoder(logFile)
		qsoLogMu.Lock()
		defer qsoLogMu.Unlock()
		if err := decoder.Decode(&qsoLog); err != nil {
			log.Fatalln("loadQsoLog", err.Error())
		}
	}
}