package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type RouteLoginRequest struct {
	Fields []Field `json:"fields" xml:"fields>field"`
}

type RouteLoginResponse struct {
	Token string `json:"token" xml:"token"`
}

type RouteError struct {
	Message interface{} `json:"message" xml:"message"`
}

func getFieldsFromAuthorization(r *http.Request) ([]Field, interface{}) {
	token := r.Header.Get("Authorization")

	if token == "" {
		return []Field{}, "Authorization token is empty"
	}

	tokenParts := strings.Split(token, " ")

	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return []Field{}, "Invalid token format"
	}
	
	fields, err := jwtDecodeFields(tokenParts[1])

	if err != nil {
		return []Field{}, err
	}

	return fields, nil
}

func beforeRoute(w http.ResponseWriter, r *http.Request, method string) error {
	contentType(w, r)

	if r.Method != method {
		message := fmt.Sprintf("Only %s requests are allowed", method)

		w.WriteHeader(http.StatusMethodNotAllowed)
		encodeResponse(w, r, RouteError{ Message: message })

		return errors.New(message)
	}

	return nil
}

func afterRoute(w http.ResponseWriter, r *http.Request, response interface{}, statusCode int) error {
	w.WriteHeader(statusCode)

	if response == nil {
		return nil
	}

	if err := encodeResponse(w, r, response); err != nil {
		encodeResponse(w, r, RouteError{ Message: err.Error() })

		return errors.New(err.Error())
	}

	return nil
}

func contentType(w http.ResponseWriter, r *http.Request) string {
	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		contentType = "application/json"
	}

	switch contentType {
		case "application/xml":
			w.Header().Set("Content-Type", "application/xml")
		default:
			w.Header().Set("Content-Type", "application/json")
	}

	return contentType
}

func encodeResponse(w http.ResponseWriter, r *http.Request, response interface{}) error {
	contentType := contentType(w, r)

	switch contentType {
		case "application/xml":			
			if err := xml.NewEncoder(w).Encode(response); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				xml.NewEncoder(w).Encode(RouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
		default:
			if err := json.NewEncoder(w).Encode(response); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(RouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
	}

	return nil
}

func decodeRequest(w http.ResponseWriter, r *http.Request, request interface{}) error {
	contentType := contentType(w, r)

	switch contentType {
		case "application/xml":
			if err := xml.NewDecoder(r.Body).Decode(&request); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				xml.NewEncoder(w).Encode(RouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
		case "application/json":
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(RouteError{ Message: err.Error() })
		
				return errors.New(err.Error())
			}
	}

	return nil
}

func routeLogin(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPost); err != nil {
		return
	}

	var fields []Field

	if err := decodeRequest(w, r, &fields); err != nil {
		return
	}

	if len(fields) == 0 {
		w.WriteHeader(http.StatusBadRequest)	
		encodeResponse(w, r, RouteError{ Message: "Fields are required" })

		return
	}

	if err := verifyFields(fields); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	token, err := jwtEncode(jwt.MapClaims{ "fields": fields })

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeResponse(w, r, RouteError{ Message: err.Error() })

		return
	}

	response := RouteLoginResponse{
		Token: token,
	}

	if err := afterRoute(w, r, response, http.StatusOK); err != nil {
		return
	}
}

func routeShow(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodGet); err != nil {
		return
	}

	fields, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	for _, field := range fields {
		result := map[string]interface{}{}

		// TODO: rewrite code
		// Свитч просто набосок, нужно сделать рандом исходя от типа и возможно перенести в отдельную функцию

		switch field.Type {
			case "uint8":
				result[field.Name] = uint8(0)
			case "uint16":
				result[field.Name] = uint16(0)
			case "uint32":
				result[field.Name] = uint32(0)
			case "int8":
				result[field.Name] = int8(0)
			case "int16":
				result[field.Name] = int16(0)
			case "int32":
				result[field.Name] = int32(0)
			case "float32":
				result[field.Name] = float32(0)
			case "float64":
				result[field.Name] = float64(0)
			case "string":
				result[field.Name] = ""
			case "bool":
				result[field.Name] = false
			default:
				result[field.Name] = nil
		}
	}

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusOK); err != nil {
		return
	}
}

func routeList(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodGet); err != nil {
		return
	}

	fields, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	for _, field := range fields {
		println(field.Name) // TODO: rewrite code
	}

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusOK); err != nil {
		return
	}
}

func routeCreate(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPost); err != nil {
		return
	}

	fields, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	var request map[string]interface{}

	if err := decodeRequest(w, r, &request); err != nil {
		return
	}

	for k, v := range request {
		println(k, v) // TODO: rewrite code
	}

	for _, field := range fields {
		println(field.Name) // TODO: rewrite code
	}

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func routePut(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPut); err != nil {
		return
	}

	fields, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	var request map[string]interface{}

	if err := decodeRequest(w, r, &request); err != nil {
		return
	}

	for k, v := range request {
		println(k, v) // TODO: rewrite code
	}

	for _, field := range fields {
		println(field.Name) // TODO: rewrite code
	}

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func routePatch(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPatch); err != nil {
		return
	}

	fields, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	var request map[string]interface{}

	if err := decodeRequest(w, r, &request); err != nil {
		return
	}

	for k, v := range request {
		println(k, v) // TODO: rewrite code
	}

	for _, field := range fields {
		println(field.Name) // TODO: rewrite code
	}

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func routeDelete(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodDelete); err != nil {
		return
	}

	fields, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	for _, field := range fields {
		println(field.Name) // TODO: rewrite code
	}

	if err := afterRoute(w, r, nil, http.StatusNoContent); err != nil {
		return
	}
}

func runHttpServer() {
	http.HandleFunc("/login", routeLogin)
	http.HandleFunc("/show", routeShow)
	http.HandleFunc("/list", routeList)
	http.HandleFunc("/create", routeCreate)
	http.HandleFunc("/put", routePut)
	http.HandleFunc("/patch", routePatch)
	http.HandleFunc("/delete", routeDelete)

	fmt.Println("Server started on :8081")
	
	http.ListenAndServe(":8081", nil)
}