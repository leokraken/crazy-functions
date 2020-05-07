package function

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	query.Get("type")

	birds := []string{"bird 1", "birds2"}

	adjetives := []string{"crazy", "loco"}

	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	birdNum := random.Intn(len(birds))
	bird := strings.Replace(birds[birdNum], " ", "_", -1)
	adjNum := random.Intn(len(adjetives))
	adjetive := strings.Replace(adjetives[adjNum], " ", "_", -1)

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s-%s", adjetive, bird)))
}
