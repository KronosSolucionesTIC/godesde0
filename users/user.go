package users

import (
	"fmt"
	"time"

	"github.com/KronosSolucionesTIC/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
