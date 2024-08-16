package core_utils

import gonanoid "github.com/matoous/go-nanoid/v2"

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func NanoId() (string, error) {
	return gonanoid.Generate(alphabet, 28)
}
