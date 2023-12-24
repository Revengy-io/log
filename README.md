# Get the Revengy.io logger module  
Note that you need to include the v in the version tag.

```
go get github.com/revengy.io/logger@v0.1.0
```


```bash
package main

import (
  "fmt"
  "github.com/revengy.io/logger"
)

func main() {
    logger.Debug("I'm debugging this line")
    logger.Warn("I'm warning this line")
    logger.Error("I'll error this line")
    logger.Info("I'll info this line")
}
```