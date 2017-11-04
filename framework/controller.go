package framework

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	// "github.com/betacraft/binding"
	// "github.com/betacraft/ressy/logger"
)

// a function to read thhe JSON body and return a map of string and interface
func ReadBody(r *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	bodyMap := make(map[string]interface{})

	err := decoder.Decode(&bodyMap)
	if err != nil {
		return bodyMap, err
	}
	/*	if os.Getenv("ENV") != "prod" {
		logger.Debug(bodyMap)
	}*/
	return bodyMap, nil
}

// a function to read json body and return an interface
func ReadJSONBody(r *http.Request) (interface{}, error) {
	var response interface{}
	err := json.NewDecoder(r.Body).Decode(response)
	return response, err
}

// func Bind(r *http.Request, fm binding.FieldMapper) error {
// 	err := binding.Bind(r, fm)
// 	if err != nil {
// 		return errors.New(err.Error())
// 	}
// 	return nil
// }

func BindWithInterface(c *http.Request, m interface{}) error {
	err := json.NewDecoder(c.Body).Decode(m)
	if err != nil {
		log.Printf("json error %s", err)
	}
	return err
}

func WriteError(w http.ResponseWriter, r *http.Request, c int, err error) {
	// logger.Err(err.Error() + "\n" + r.UserAgent())
	log.Printf(err.Error() + "\n" + r.UserAgent())
	requstDump, _ := httputil.DumpRequest(r, true)
	// logger.Warn(err, "\n", strings.Trim(string(requstDump), "\n\r"))
	log.Printf(err, "\n", strings.Trim(string(requstDump), "\n\r"))
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := JSONResponse{"message": err.Error(), "success": false}
	w.Write(res.ByteArray())
}

func WriteShortError(w http.ResponseWriter, err error) {
	// logger.Err(err.Error())
	log.Printf(err.Error())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := JSONResponse{"message": err.Error(), "success": false}
	// logger.Err(err)
	log.Printf(err.Error())
	w.Write(res.ByteArray())
}
func WriteResponse(w http.ResponseWriter, c int, r JSONResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(c)
	w.Write(r.ByteArray())
}
