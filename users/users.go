package users

import (
	"fmt"
	"time"

	"github.com/vcuadrosv/go/modelos"
)

func AltaUsuarios() {
	u := new(modelos.User)
	u.AddUser(10, "Val", time.Now(), true)

	fmt.Println(u)

}
