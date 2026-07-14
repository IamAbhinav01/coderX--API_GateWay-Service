package security

import (
	"Coderx/config/env"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)


type ArgonHandlers interface{
	HashPassword(password string) (string,error)
	VerifyPassword(Hashpassword string , providedPassword string) (bool,error)
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

func parseArgon2Hash(Hashpassword string) (*Argon2Configs,error){
	
	componenets := strings.Split( Hashpassword, "$")

	if len(componenets) != 6{
		fmt.Printf("Invalid Password Hash format")
		return nil,fmt.Errorf("invalid hash format structure")
	}

	if !strings.HasPrefix(componenets[1],"argon2id"){
		fmt.Printf("unsupported algorithm variant")
		return nil,fmt.Errorf("unsupported algorithm variant")
	}

	var version int
	fmt.Sscanf(componenets[2], "v=%d", &version)

	if version != argon2.Version{
		fmt.Printf("unsupported argon2 version")
		return nil,fmt.Errorf("unsupported argon2 version")
	}

	config:= &Argon2Configs{}

	_,cerr:=fmt.Sscanf(componenets[3], "m=%d,t=%d,p=%d", 
        &config.MemoryCost, &config.TimeCost, &config.Threads)

	if cerr!= nil{
		return nil,cerr
	}

	salt,err := base64.RawStdEncoding.DecodeString(componenets[4])

	if err != nil{
		fmt.Printf("salt decoding failed")
		return nil,err
	}

	config.Salt = salt

	hash,Hasherr := base64.RawStdEncoding.DecodeString(componenets[5])

	if Hasherr!= nil{
		fmt.Println("hash decoding failed")
		return nil,Hasherr
	}

	config.HashRaw = hash
	config.KeyLength = uint32(len(hash))

	return config,nil
	
}

func (config *Argon2Configs) VerifyPassword(Hashpassword string , providedPassword string) (bool,error){

	parsedConfig , err := parseArgon2Hash(Hashpassword)

	if err!= nil{
		fmt.Println("Error happend while parsing the password")
		return false,err
	}

	computedHash := argon2.IDKey([] byte(providedPassword),parsedConfig.Salt,parsedConfig.TimeCost,parsedConfig.MemoryCost,parsedConfig.Threads,parsedConfig.KeyLength)

	match := subtle.ConstantTimeCompare(computedHash,parsedConfig.HashRaw) == 1

	return match,nil
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