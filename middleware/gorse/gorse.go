package gorse

import "github.com/zhenghaoz/gorse/client"

var GorseInstance *client.GorseClient

func InitGorse() {
	GorseInstance = client.NewGorseClient("http://127.0.0.1:8087", "")
}
