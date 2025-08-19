package helper

import (
	"encoding/json"
	"fmt"
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
		return fmt.Errorf("Error getting X_MAX env %v", err)
	}
	ymax, err = strconv.Atoi(os.Getenv("Y_MAX"))
	if err != nil {
		return fmt.Errorf("Error getting Y_MAX env %v", err)
	}
	return nil
}

func GetMaxXY() (int, int) {
	return xmax, ymax
}

var requiredEnv = map[string]string{
	"APP_ENV": "DEV,PRODUCTION", "X_MAX": "number, above 0", "Y_MAX": "number, above 0", "IO_THREADS": "number, above 0", "VALKEY_HOST": "e.g. 127.0.0.1:6379", "VALKEY_PASSWORD": "e.g. password123", "PORT": "port num,e.g. 8081, must be set in ui/.env ( VITE_APP_URL=localhost:8081 ) too", "VALKEY_MAX_MEMORY": "e.g. 256mb, 1gb, 2gb"}

func RequiredEnv() error {
	for k, v := range requiredEnv {
		if os.Getenv(k) == "" {
			return fmt.Errorf("missing env variable: %s. Helper: %s", k, v)
		}
	}
	return nil
}
