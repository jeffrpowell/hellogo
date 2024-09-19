package main

import (
	"fmt"
	"net/http"

	"github.com/jeffrpowell/hellogo/internal/constants"
	//_ "github.com/jeffrpowell/hellogo/internal/database" //blank import to run init()
	_ "github.com/jeffrpowell/hellogo/internal/handlers" //blank import to run init()
)

func main() {
	fmt.Println("####################################")
	fmt.Println("#             HELLOGO              #")
	fmt.Println("####################################")
	fmt.Println()
	fmt.Println("Server is running at http://localhost:" + constants.PORT)
	http.ListenAndServe(":"+constants.PORT, constants.ROUTER)
}
