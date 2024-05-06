package main

import (
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/net"
	"github.com/sirupsen/logrus"
)

type InterfaceData struct {
	Tx uint64 `json:"tx"`
	Rx uint64 `json:"rx"`

	TxSpeed uint64 `json:"tx_speed"`
	RxSpeed uint64 `json:"rx_speed"`
}

var measureInterval uint64 = ParseEnvUint("MEASURE_INTERVAL", 5)

var lock sync.RWMutex

var interfacesMap = map[string]InterfaceData{}

func StartUpdateTask() {
	ticker := time.NewTicker(time.Second * time.Duration(measureInterval))

	for {
		<-ticker.C
		data, err := net.IOCounters(true)
		if err != nil {
			logrus.Infof("update io counters failed: %v", err)
			continue
		}

		lock.Lock()
		for _, v := range data {
			if v.BytesSent == 0 && v.BytesRecv == 0 {
				continue
			}

			data, ok := interfacesMap[v.Name]
			if !ok {
				interfacesMap[v.Name] = InterfaceData{
					Tx: v.BytesSent,
					Rx: v.BytesRecv,

					TxSpeed: 0,
					RxSpeed: 0,
				}

				continue
			}

			interfacesMap[v.Name] = InterfaceData{
				Tx:      v.BytesSent,
				Rx:      v.BytesRecv,
				TxSpeed: (v.BytesSent - data.Tx) / measureInterval,
				RxSpeed: (v.BytesRecv - data.Rx) / measureInterval,
			}
		}

		lock.Unlock()
		// UpdateInterfaces()
	}
}
