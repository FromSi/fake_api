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

	jwtToken, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return []Field{}, err.Error()
	}

	if !jwtToken.Valid {
		return []Field{}, "Invalid token"
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		return []Field{}, "Invalid token claims"
	}

	fieldsClaim, exists := claims["fields"]

	if !exists {
		return []Field{}, "Token doesn't contain fields"
	}

	fieldsClaimArray, ok := fieldsClaim.([]interface{})

	if !ok {
		return []Field{}, "Invalid token fields"
	}

	var fields []Field

	for _, fieldData := range fieldsClaimArray {
		fieldMap, ok := fieldData.(map[string]interface{})

		if !ok {
			return []Field{}, "Invalid token field"
		}

		var field Field

		if f, ok := fieldMap["type"].(string); ok {
			field.Type = f
		} else {
			field.Type = ""
		}

		if f, ok := fieldMap["name"].(string); ok {
			field.Name = f
		} else {
			field.Name = ""
		}

		if f, ok := fieldMap["required"].(bool); ok {
			field.Required = f
		} else {
			field.Required = false
		}

		if f, ok := fieldMap["max"].(float64); ok {
			field.Max = int(f)
		} else {
			field.Max = 0
		}

		if f, ok := fieldMap["min"].(float64); ok {
			field.Min = int(f)
		} else {
			field.Min = 0
		}

		fields = append(fields, field)
	}

	if err := verifyFields(fields); err != nil {
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

	var request RouteLoginRequest

	if err := decodeRequest(w, r, &request); err != nil {
		encodeResponse(w, r, RouteError{ Message: err.Error() })

		return
	}

	if len(request.Fields) == 0 {
		w.WriteHeader(http.StatusBadRequest)	
		encodeResponse(w, r, RouteError{ Message: "Fields are required" })

		return
	}

	if err := verifyFields(request.Fields); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"fields": request.Fields,
	})

	token, err := jwt.SignedString([]byte("secret"))

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

	_, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	// TODO: write code here

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusOK); err != nil {
		return
	}
}

func routeList(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodGet); err != nil {
		return
	}

	_, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	// TODO: write code here

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusOK); err != nil {
		return
	}
}

func routeCreate(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPost); err != nil {
		return
	}

	_, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	// TODO: write code here

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func routePut(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPut); err != nil {
		return
	}

	_, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	// TODO: write code here

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func routePatch(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodPatch); err != nil {
		return
	}

	_, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
	}

	// TODO: write code here

	if err := afterRoute(w, r, RouteError{ Message: "Not implemented" }, http.StatusAccepted); err != nil {
		return
	}
}

func routeDelete(w http.ResponseWriter, r *http.Request) {
	if err := beforeRoute(w, r, http.MethodDelete); err != nil {
		return
	}

	_, err := getFieldsFromAuthorization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: err })

		return
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