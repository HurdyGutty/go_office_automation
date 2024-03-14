package getTemplate

import "math/rand"

// GetTemplate returns a random template

func GetTemplate(URI_list []string) string {
	return URI_list[rand.Intn(len(URI_list))]
}
