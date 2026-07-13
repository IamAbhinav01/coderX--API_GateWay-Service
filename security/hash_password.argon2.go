package security

import (
	"Coderx/config/env"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)


type ArgonHandlers interface{
	HashPassword(password string) string
}

type Argon2Configs struct{

	HashRaw []byte
	Salt []byte
	TimeCost uint32 
	MemoryCost uint32 
	Threads uint8 
	KeyLength uint32
	
}

func NewArgonConfig() (*Argon2Configs, error) {

	timeCost := env.GetInt("TimeCost")
	
	memoryCost := env.GetInt("MemoryCost")

	threads := env.GetInt("Threads")

	keyLength := env.GetInt("KeyLength")

	return &Argon2Configs{

		TimeCost: uint32(timeCost),
		MemoryCost: uint32(memoryCost),
		Threads: uint8(threads),
		KeyLength: uint32(keyLength),

	}, nil
}


func generateCryptographicallySalt(saltsize uint32) ([]byte,error){

	salt := make([]byte,saltsize)
	_,err:=rand.Read(salt) 
	
	if err != nil{
		fmt.Printf("Salt Generation failed");
		return nil,err
	}

	return salt,nil;

}

func (config *Argon2Configs) HashPassword(password string) (string,error) {

	saltSize := env.GetInt("salt_size")

	salt, err := generateCryptographicallySalt(uint32(saltSize))

	if err!= nil{
		fmt.Println("Error while generating the cryptographic salt")
		return "",err
	}

	config.Salt = salt



	config.HashRaw = argon2.IDKey([]byte(password),config.Salt,config.TimeCost,config.MemoryCost,config.Threads,config.KeyLength)

	encodeHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
        argon2.Version,
        config.MemoryCost,
        config.TimeCost,
        config.Threads,
        base64.RawStdEncoding.EncodeToString(config.Salt),
        base64.RawStdEncoding.EncodeToString(config.HashRaw),
	)

	return encodeHash,nil
}


// func main(){
// 	config, err := NewArgonConfig()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	hash, err := config.HashPassword("iamAbhinavSunil@2005")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Secure hash generated: ", hash)
// }