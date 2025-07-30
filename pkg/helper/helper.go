package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

func ReturnJsonError(w http.ResponseWriter, e error, code int) {
	w.WriteHeader(code)
	body, err := json.Marshal(map[string]any{"error": e.Error()})
	if err != nil {
		log.Println("Error marshalling.. string ? ")
	}
	w.Write(body)

}

var xmax int
var ymax int

func SetMaxXY() error {
	var err error
	xmax, err = strconv.Atoi(os.Getenv("X_MAX"))
	if err != nil {
		return err
	}
	ymax, err = strconv.Atoi(os.Getenv("Y_MAX"))
	if err != nil {
		return err
	}
	return nil
}

func GetMaxXY() (int, int) {
	return xmax, ymax
}
