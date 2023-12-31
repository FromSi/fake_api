package main

import (
	"github.com/golang-jwt/jwt/v5"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Функция для кодирования данных в JWT
func jwtEncode(data jwt.MapClaims) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	return jwt.SignedString([]byte("secret"))
}

// Функция для декодирования данных из JWT
func jwtDecode(token string) (jwt.MapClaims, error) {
	data, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if !data.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := data.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	return claims, nil
}

// Функция для кодирования полей в JWT
func jwtDecodeFields(token string) ([]Field, interface{}) {
	claims, err := jwtDecode(token)

	if err != nil {
		return []Field{}, err.Error()
	}

	fieldsClaim, exists := claims["fields"]

	if !exists {
		return []Field{}, "Token doesn't contain fields"
	}

	fieldsClaimArray, ok := fieldsClaim.([]interface{})

	if !ok {
		return []Field{}, "Invalid token fields"
	}

	fakeApiMaxFieldsInObject, err := strconv.Atoi(os.Getenv("FAKE_API_MAX_FIELDS_IN_OBJECT"))

	if err != nil {
		return []Field{}, err.Error()
	}

	if len(fieldsClaimArray) > fakeApiMaxFieldsInObject {
		return []Field{}, fmt.Sprintf("Maximum fields in object is %d", fakeApiMaxFieldsInObject)
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

	if err := validateFields(fields); err != nil {
		return []Field{}, err
	}

	return fields, nil
}