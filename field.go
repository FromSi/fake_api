package main

import (
	"strconv"
	"math"
	"math/rand"
	"time"
	"github.com/brianvoe/gofakeit/v6"
	"regexp"
)

// Поля для мок ресурса.
type Field struct {
	Type string `json:"type" xml:"type"`
	Name string `json:"name" xml:"name"`
	Required bool `json:"required" xml:"required"`
	Max int `json:"max" xml:"max"`
	Min int `json:"min" xml:"min"`
}

// Валидировать для Field структуры.
func validateFields(fields []Field) []RouteErrorList {
	errs := []RouteErrorList{}

	fieldNames := map[string]bool{}

	for i, f := range fields {
		err := RouteErrorList{}

		if f.Type == "" {
			err.Errors = append(err.Errors, "Field has empty type")
		}

		if f.Name == "" {
			err.Errors = append(err.Errors, "Field has empty name")
		} else if matched, _ := regexp.MatchString("^[a-z_0-9]+$", f.Name); matched == false {
			err.Errors = append(err.Errors, "Field has incorrect format (only lowercase letters, underscores are allowed and numbers)")
		} else if _, ok := fieldNames[f.Name]; ok {
			err.Errors = append(err.Errors, "Field has duplicate name")
		} else {
			fieldNames[f.Name] = true
		}

		if f.GetCorrectMin() > f.GetCorrectMax() {
			err.Errors = append(err.Errors, "Field has incorrect min and max values")
		}

		if len(err.Errors) > 0 {
			err.Index = strconv.Itoa(i)
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Плучить случайное значение для поля.
func (f *Field) GetRandomValue() interface{} {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	switch f.Type {
		case "uint8":
			return rand.Intn(f.GetCorrectMax() - f.GetCorrectMin() + 1) + f.GetCorrectMin()
		case "uint16":
			return rand.Intn(f.GetCorrectMax() - f.GetCorrectMin() + 1) + f.GetCorrectMin()
		case "uint32":
			return rand.Intn(f.GetCorrectMax() - f.GetCorrectMin() + 1) + f.GetCorrectMin()
		case "int8":
			return rand.Intn(f.GetCorrectMax() - f.GetCorrectMin() + 1) + f.GetCorrectMin()
		case "int16":
			return rand.Intn(f.GetCorrectMax() - f.GetCorrectMin() + 1) + f.GetCorrectMin()
		case "int32":
			return rand.Intn(f.GetCorrectMax() - f.GetCorrectMin() + 1) + f.GetCorrectMin()
		case "float":
			return rand.Float32() * float32(f.GetCorrectMax() - f.GetCorrectMin()) + float32(f.GetCorrectMin())
		case "boolean":
			return gofakeit.Bool()
		case "string_name":
			return gofakeit.Name()
		case "string_email":
			return gofakeit.Email()
		case "string_username":
			return gofakeit.Username()
		case "string_country":
			return gofakeit.Country()
		case "string_word":
			return gofakeit.Word()
		case "string_sentence":
			return gofakeit.Sentence(f.GetCorrectMax())
		case "string_url":
			return gofakeit.URL()
		case "string_uuid":
			return gofakeit.UUID()
		case "string_hex_color":
			return gofakeit.HexColor()
		case "string_phone":
			return gofakeit.Phone()
		case "string_credit_card":
			return gofakeit.CreditCard()
		case "string_currency":
			return gofakeit.CurrencyShort()
		case "string_bitcoin_address":
			return gofakeit.BitcoinAddress()
		case "string_emoji":
			return gofakeit.Emoji()
		case "string_ipv4":
			return gofakeit.IPv4Address()
		case "string_ipv6":
			return gofakeit.IPv6Address()
		case "string_date":
			return gofakeit.Date().Format("2006-01-02")
		case "string_date_time":
			return gofakeit.FutureDate().Format("2006-01-02 15:04:05")
		case "string_time":
			return gofakeit.FutureDate().Format("15:04:05")
		default:
			return nil
	}
}

// Получить корректное, минимальное значение для поля.
func (f *Field) GetCorrectMin() int {
	return getCorrectMinByType(f.Type, f.Min)
}

// Получить корректное, минимальное значение для поля по типу. 
// value - значение, которое должно быть корректным.
func getCorrectMinByType(fieldType string, value int) int {
	switch fieldType {
		case "uint8":
			return max(0, value)
		case "uint16":
			return max(0, value)
		case "uint32":
			return max(0, value)
		case "int8":
			return max(math.MinInt8, value)
		case "int16":
			return max(math.MinInt16, value)
		case "int32":
			return max(math.MinInt32, value)
		case "float":
			return max(-999999, value)
		case "boolean":
			return max(0, value)
		case "string_name":
			return max(1, value)
		case "string_email":
			return max(5, value)
		case "string_username":
			return max(1, value)
		case "string_country":
			return max(1, value)
		case "string_word":
			return max(1, value)
		case "string_sentence":
			return max(1, value)
		case "string_url":
			return max(10, value)
		case "string_uuid":
			return max(36, value)
		case "string_hex_color":
			return max(7, value)
		case "string_phone":
			return max(1, value)
		case "string_credit_card":
			return max(19, value)
		case "string_currency":
			return max(1, value)
		case "string_bitcoin_address":
			return max(34, value)
		case "string_emoji":
			return max(1, value)
		case "string_ipv4":
			return max(7, value)
		case "string_ipv6":
			return max(15, value)
		case "string_date":
			return max(10, value)
		case "string_date_time":
			return max(19, value)
		case "string_time":
			return max(8, value)
		default:
			return max(0, value)
	}
}

// Получить корректное, максимальное значение для поля.
func (f *Field) GetCorrectMax() int {
	return getCorrectMaxByType(f.Type, f.Max)
}

// Получить корректное, максимальное значение для поля по типу. 
// value - значение, которое должно быть корректным.
func getCorrectMaxByType(fieldType string, value int) int {
	switch fieldType {
		case "uint8":
			return min(math.MaxUint8, value)
		case "uint16":
			return min(math.MaxUint16, value)
		case "uint32":
			return min(math.MaxUint32, value)
		case "int8":
			return min(math.MaxInt8, value)
		case "int16":
			return min(math.MaxInt16, value)
		case "int32":
			return min(math.MaxInt32, value)
		case "float":
			return min(999999, value)
		case "boolean":
			return min(1, value)
		case "string_name":
			return min(255, value)
		case "string_email":
			return min(255, value)
		case "string_username":
			return min(255, value)
		case "string_country":
			return min(255, value)
		case "string_word":
			return min(255, value)
		case "string_sentence":
			return min(2048, value)
		case "string_url":
			return min(255, value)
		case "string_uuid":
			return min(36, value)
		case "string_hex_color":
			return min(7, value)
		case "string_phone":
			return min(20, value)
		case "string_credit_card":
			return min(19, value)
		case "string_currency":
			return min(20, value)
		case "string_bitcoin_address":
			return min(62, value)
		case "string_emoji":
			return min(4, value)
		case "string_ipv4":
			return min(15, value)
		case "string_ipv6":
			return min(39, value)
		case "string_date":
			return min(10, value)
		case "string_date_time":
			return min(19, value)
		case "string_time":
			return min(8, value)
		default:
			return min(0, value)
	}
}

// Получить тип поля для значения поля
func (f *Field) GetType() string {
	return getTypeByType(f.Type)
}

// Получить тип поля для значения поля по типу
func getTypeByType(fieldType string) string {
	switch fieldType {
		case "uint8":
			return "float64"
		case "uint16":
			return "float64"
		case "uint32":
			return "float64"
		case "int8":
			return "float64"
		case "int16":
			return "float64"
		case "int32":
			return "float64"
		case "float":
			return "float64"
		case "boolean":
			return "bool"
		case "string_name":
			return "string"
		case "string_email":
			return "string"
		case "string_username":
			return "string"
		case "string_country":
			return "string"
		case "string_word":
			return "string"
		case "string_sentence":
			return "string"
		case "string_url":
			return "string"
		case "string_uuid":
			return "string"
		case "string_hex_color":
			return "string"
		case "string_phone":
			return "string"
		case "string_credit_card":
			return "string"
		case "string_currency":
			return "string"
		case "string_bitcoin_address":
			return "string"
		case "string_emoji":
			return "string"
		case "string_ipv4":
			return "string"
		case "string_ipv6":
			return "string"
		case "string_date":
			return "string"
		case "string_date_time":
			return "string"
		case "string_time":
			return "string"
		default:
			return "error"
	}
}