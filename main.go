package main

import (
	"fmt"
	"net/http"

	authroutes "github.com/rbaylon/trixie_mods/auth/routes"
	authtypes "github.com/rbaylon/trixie_mods/auth/types"
)

func main() {
	u := &authtypes.GroupType{Name: "test"}
	fmt.Println(u.Name)

	auth_router := authroutes.GetRouter()

	router := http.NewServeMux() // 1 - root router
	router.Handle("/auth/", http.StripPrefix("/auth", auth_router))

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", router)
}
