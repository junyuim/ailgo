package core_utils

import gonanoid "github.com/matoous/go-nanoid/v2"

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const alphabet_apparent = "23456789ABCDEFGHJKLMNOPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

func NanoId() (string, error) {
	return gonanoid.Generate(alphabet, 28)
}
