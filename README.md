# Fleet Go SDK

Go SDK for interacting with the [Fleet](https://fleetdm.com/) API.

_Note: this is incomplete, as I only implemented what I needed (+ low-hanging fruit). PRs welcome!._

## Usage

You can retrieve your API key via: https://fleetdm.com/docs/rest-api/rest-api#retrieve-your-api-token

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wolveix/fleet-go"
)

func main() {
	key := "your_api_key_here"
	service := fleet.New("https://your.fleet.domain", key, 15*time.Second, false)

	user, err := service.FindMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Email)
}
```

## Retrieve Hosts

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wolveix/fleet-go"
)

func main() {
	key := "your_api_key_here"
	service := fleet.New("https://your.fleet.domain", key, 15*time.Second, false)

	hosts, err := service.FindHosts()
	if err != nil {
		log.Fatal(err)
	}
	
	for _, host := range hosts {
		fmt.Printf("Host details:\n- Hostname: %s\n- OS Version: %s\n\n", host.Hostname, host.OSVersion)
    }
}
```

## License

BSD licensed. See the [LICENSE](LICENSE) file for details.