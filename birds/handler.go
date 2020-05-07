package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	query.Get("type")

	birds := GetBirds()
	adjectives := GetAdjectives()

	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	birdNum := random.Intn(len(birds))
	bird := strings.Replace(birds[birdNum], " ", "_", -1)
	adjNum := random.Intn(len(adjectives))
	adjective := strings.Replace(adjectives[adjNum], " ", "_", -1)

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s-%s\n", adjective, bird)))
}

func ReadFileToArray(filename string) []string {
	var data []string
	file, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func GetBirds() []string {
	return ReadFileToArray("./birds.json")
}

func GetAdjectives() []string {
	return ReadFileToArray("./adjectives.json")
}
