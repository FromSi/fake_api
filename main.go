package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
)

type Field struct {
	Type string `json:"type" xml:"type,attr"`
	Name string `json:"name" xml:"name,attr"`
	Required bool `json:"required" xml:"required,attr"`
	Max int `json:"max" xml:"max,attr"`
	Min int `json:"min" xml:"min,attr"`
}

type HttpRouteLoginRequest struct {
	Fields []Field `json:"fields" xml:"fields>field"`
}

type HttpRouteLoginResponse struct {
	Token string `json:"token" xml:"token"`
}

type HttpRouteError struct {
	Message string `json:"message" xml:"message"`
}

func httpBeforeRoute(w http.ResponseWriter, r *http.Request, method string) error {
	if r.Method != method {
		message := fmt.Sprintf("Only %s requests are allowed", method)

		w.WriteHeader(http.StatusMethodNotAllowed)
		httpEncodeResponse(w, r, HttpRouteError{ Message: message })

		return errors.New(message)
	}

	return nil
}

func httpAfterRoute(w http.ResponseWriter, r *http.Request, response interface{}, statusCode int) error {
	w.WriteHeader(statusCode)

	if response == nil {
		return nil
	}

	if err := httpEncodeResponse(w, r, response); err != nil {
		httpEncodeResponse(w, r, HttpRouteError{ Message: err.Error() })

		return errors.New(err.Error())
	}

	return nil
}

func httpEncodeResponse(w http.ResponseWriter, r *http.Request, response interface{}) error {
	contentType := r.Header.Get("Content-Type")

	switch contentType {
		case "application/xml":
			w.Header().Set("Content-Type", "application/xml")
			
			if err := xml.NewEncoder(w).Encode(response); err != nil {
				println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				xml.NewEncoder(w).Encode(HttpRouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
		default:
			w.Header().Set("Content-Type", "application/json")

			if err := json.NewEncoder(w).Encode(response); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(HttpRouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
	}

	return nil
}

func httpDecodeRequest(w http.ResponseWriter, r *http.Request, request interface{}) error {
	contentType := r.Header.Get("Content-Type")

	switch contentType {
		case "application/xml":
			w.Header().Set("Content-Type", "application/xml")

			if err := xml.NewDecoder(r.Body).Decode(&request); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				xml.NewEncoder(w).Encode(HttpRouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
		case "application/json":
			w.Header().Set("Content-Type", "application/json")

			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(HttpRouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
	}

	return nil
}

func httpRouteLogin(w http.ResponseWriter, r *http.Request) {
	if err := httpBeforeRoute(w, r, http.MethodPost); err != nil {
		return
	}

	var request HttpRouteLoginRequest

	if err := httpDecodeRequest(w, r, &request); err != nil {
		httpEncodeResponse(w, r, HttpRouteError{ Message: err.Error() })

		return
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"fields": request.Fields,
	})

	token, err := jwt.SignedString([]byte("secret"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		httpEncodeResponse(w, r, HttpRouteError{ Message: err.Error() })

		return
	}

	response := HttpRouteLoginResponse{
		Token: token,
	}

	if err := httpAfterRoute(w, r, response, http.StatusOK); err != nil {
		return
	}
}

func httpRouteShow(w http.ResponseWriter, r *http.Request) {
	if err := httpBeforeRoute(w, r, http.MethodGet); err != nil {
		return
	}

	// TODO: write code here

	if err := httpAfterRoute(w, r, HttpRouteError{ Message: "Not implemented" }, http.StatusOK); err != nil {
		return
	}
}

func httpRouteList(w http.ResponseWriter, r *http.Request) {
	if err := httpBeforeRoute(w, r, http.MethodGet); err != nil {
		return
	}

	// TODO: write code here

	if err := httpAfterRoute(w, r, HttpRouteError{ Message: "Not implemented" }, http.StatusOK); err != nil {
		return
	}
}

func httpRouteCreate(w http.ResponseWriter, r *http.Request) {
	if err := httpBeforeRoute(w, r, http.MethodPost); err != nil {
		return
	}

	// TODO: write code here

	if err := httpAfterRoute(w, r, HttpRouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func httpRoutePut(w http.ResponseWriter, r *http.Request) {
	if err := httpBeforeRoute(w, r, http.MethodPut); err != nil {
		return
	}

	// TODO: write code here

	if err := httpAfterRoute(w, r, HttpRouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func httpRoutePatch(w http.ResponseWriter, r *http.Request) {
	if err := httpBeforeRoute(w, r, http.MethodPatch); err != nil {
		return
	}

	// TODO: write code here

	if err := httpAfterRoute(w, r, HttpRouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func httpRouteDelete(w http.ResponseWriter, r *http.Request) {
	if err := httpBeforeRoute(w, r, http.MethodDelete); err != nil {
		return
	}

	if err := httpAfterRoute(w, r, nil, http.StatusNoContent); err != nil {
		return
	}
}

func initHttpServer() {
	http.HandleFunc("/login", httpRouteLogin)
	http.HandleFunc("/show", httpRouteShow)
	http.HandleFunc("/list", httpRouteList)
	http.HandleFunc("/create", httpRouteCreate)
	http.HandleFunc("/put", httpRoutePut)
	http.HandleFunc("/patch", httpRoutePatch)
	http.HandleFunc("/delete", httpRouteDelete)

	fmt.Println("Server started on :8081")
	
	http.ListenAndServe(":8081", nil)
}

func main() {
	initHttpServer()
}
