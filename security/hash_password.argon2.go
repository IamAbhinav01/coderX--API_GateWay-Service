package bcrypt

import "golang.org/x/crypto/argon2"


type ArgonHandlers interface{

}

type Argon2Configs struct{

}

func NewArgonConfig() *Argon2Configs{
	return &Argon2Configs{
		
	}

}

func HashPassword(password string) string {

	argon2.IDKey()
}