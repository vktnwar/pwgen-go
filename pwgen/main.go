package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)


func generatePassword(length int, charset string, r *rand.Rand) string {
     password := ""
    
     for i := 0; i < length; i++ {
        // Passo 2a: gerar índice aleatório
        randomIndex := r.Intn(len(charset))
        
        // Passo 2b: pegar caractere nessa posição
        randomChar := charset[randomIndex]
        
        // Passo 2c: adicionar à senha
        password += string(randomChar)
    }
    
    // Passo 3: retornar senha
    return password
}
	 
func main() {
	lenght := flag.Int("l", 12, "Password length")
	includeUpper := flag.Bool("upper", true, "Upper include")
	includeLower := flag.Bool("lower", true, "Lower include")
	includeNumbers := flag.Bool("n", true, "Numbers include")
	includeSymbols := flag.Bool("sy", false, "Symbols include")
	count := flag.Int("Count", 1, "pwns")

	flag.Parse()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const (
    lowercase = "abcdefghijklmnopqrstuvwxyz"
    uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"  
    numbers   = "0123456789"
    symbols   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	)

	var charset string

	if *includeUpper {
		charset += uppercase
	}
	if *includeLower {
		charset += lowercase	
	}
	if *includeNumbers {
		charset += numbers
	}
	if *includeSymbols {
		charset  += symbols
	}

	if charset == "" {
		fmt.Println("Error: String vazia")
		os.Exit(1)
	}

	fmt.Printf("Gerando %d senha(s) com %d caracteres:\n\n", *count, *lenght)

	for i := 0; i < *count; i++ {
    	password := generatePassword(*lenght, charset, r)
    	fmt.Printf("Senha %d: %s\n", i+1, password)
	}

}
