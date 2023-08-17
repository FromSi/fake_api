package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"reflect"
)

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

func validateFieldsFromRequest(w http.ResponseWriter, r *http.Request, request map[string]interface{}, fields []Field, isRequired bool) (interface{}, int) {
	errs := []RouteErrorList{}

	for _, field := range fields {
		isFound := false
		
		for k, v := range request {
			if k != field.Name {
				continue
			}

			isFound = true
			err := RouteErrorList{}
			err.Index = k
	
			if field.Name == "" {
				err.Errors = append(err.Errors, "Field not found")
			} else if field.Required == true && v == nil {
				err.Errors = append(err.Errors, "Field is required")
			} else if field.GetType() != reflect.TypeOf(v).String() {
				err.Errors = append(err.Errors, "Field has incorrect type")
			} else if field.GetType() == "string" {
				if field.GetCorrectMin() > len(v.(string)) {
					err.Errors = append(err.Errors, "Field has incorrect min length")
				}
	
				if field.GetType() == "string" && field.GetCorrectMax() < len(v.(string)) {
					err.Errors = append(err.Errors, "Field has incorrect max length")
				}
			} else if field.GetType() == "float64" {
				if field.GetCorrectMin() > int(v.(float64)) {
					err.Errors = append(err.Errors, "Field has incorrect min value")
				}
	
				if field.GetCorrectMax() < int(v.(float64)) {
					err.Errors = append(err.Errors, "Field has incorrect max value")
				}
			}
	
			if len(err.Errors) > 0 {
				errs = append(errs, err)
			}
		}

		if isFound == false && field.Required == true && isRequired == true {
			err := RouteErrorList{}
			err.Index = field.Name
			err.Errors = append(err.Errors, "Field not found")

			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs, http.StatusAccepted
	} else {
		return RouteSuccess{ Data: "Success: todo rewrite!!!"}, http.StatusCreated
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