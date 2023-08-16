package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Return a greeting
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Nombre vacío")
	}
	// Devuelve saludo en forma de mensaje
	message := fmt.Sprint(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}
