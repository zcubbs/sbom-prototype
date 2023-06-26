package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomInt64(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomArtifactName() string {
	return RandomString(10)
}

func RandomArtifactVersion() string {
	a := strconv.Itoa(RandomInt(1, 20))
	b := strconv.Itoa(RandomInt(1, 20))
	c := strconv.Itoa(RandomInt(1, 20))
	return "v" + a + "." + b + "." + c
}

func RandomArtifactType() string {
	var artifactTypes = []string{"image", "file", "directory", "repository"}
	return artifactTypes[rand.Intn(len(artifactTypes))]
}
