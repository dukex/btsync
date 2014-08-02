## BTSync API

## USAGE

```
package main

import (
	"log"

	"github.com/dukex/btsync"
)

func main() {
	b, _ := btsync.New("http://admin:admin@0.0.0.0:8888")
	log.Println(b.GetFolders(""))
	log.Println(b.GetFolders("key1234"))

	log.Println(b.AddFolder("/tmp/fromGO", "key1234", ""))
	log.Println(b.RemoveFolder("key1234"))

	log.Println(b.GetFiles("key1234", ""))
	log.Println(b.GetFiles("key1234", "dev_folder"))
}
```
