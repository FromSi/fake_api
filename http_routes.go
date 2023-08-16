package main

import (
	"net/http"
	"math/rand"
	"time"
	"github.com/golang-jwt/jwt/v5"
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

type RouteSuccess struct {
	Data interface{} `json:"data" xml:"data"`
}

func routeLogin(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPost); err != nil {
		return
	}

	var reuqest RouteLoginRequest

	if err := decodeRequest(w, r, &reuqest); err != nil {
		return
	}

	if len(reuqest.Fields) == 0 {
		w.WriteHeader(http.StatusBadRequest)	
		encodeResponse(w, r, RouteError{ Message: "Fields are required" })

		return
	}
	
	if err := verifyFields(reuqest.Fields); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	token, err := jwtEncode(jwt.MapClaims{ "fields": reuqest.Fields })

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

	result := map[string]interface{}{}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, field := range fields {
		result[field.Name] = field.GetRandomValue()
	}

	if err := afterRoute(w, r, RouteSuccess{ Data: result }, http.StatusOK); err != nil {
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

	result := make([]map[string]interface{}, 20)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range result {
		result[i] = map[string]interface{}{}

		for _, field := range fields {
			result[i][field.Name] = field.GetRandomValue()
		}
	}

	if err := afterRoute(w, r, RouteSuccess{ Data: result }, http.StatusOK); err != nil {
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