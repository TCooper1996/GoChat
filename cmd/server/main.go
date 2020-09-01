package GoChat

import (
	C "GoChat/internal/pkg/controllers"
	M "GoChat/internal/server/models"
	"fmt"
	"os"
)

var useGui = true

func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] == "-noGui" {
			useGui = false
		}
	}

	M.InitRoom()
	M.InitUser()
	server := M.NewUser("Server", nil)
	M.NewRoom("Main", server.ID)

	if useGui {
		C.InitGui()
	} else {
		fmt.Println("")
	}

}
