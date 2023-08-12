package main

import (
	"strconv"
)

type Field struct {
	Type string `json:"type" xml:"type,attr"`
	Name string `json:"name" xml:"name,attr"`
	Required bool `json:"required" xml:"required,attr"`
	Max int `json:"max" xml:"max,attr"`
	Min int `json:"min" xml:"min,attr"`
}

func verifyFields(fields []Field) map[string][]string {
	errs := map[string][]string{}

	for i, f := range fields {
		err := []string{}

		if f.Type == "" {
			err = append(err, "Field has empty type")
		}

		if f.Name == "" {
			err = append(err, "Field has empty name")
		}

		if f.GetCorrectMin() > f.GetCorrectMax() {
			err = append(err, "Field has incorrect min and max values")
		}

		if len(err) > 0 {
			errs[strconv.Itoa(i)] = err
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (f *Field) GetCorrectMin() int {
	return getCorrectMinByType(f.Type, f.Min)
}

func getCorrectMinByType(fieldType string, value int) int {
	switch fieldType {
		case "uint8":
			return max(0, value)
		case "uint16":
			return max(0, value)
		case "uint32":
			return max(0, value)
		case "int8":
			return max(-128, value)
		case "int16":
			return max(-32768, value)
		case "int32":
			return max(-2147483648, value)
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

func (f *Field) GetCorrectMax() int {
	return getCorrectMaxByType(f.Type, f.Max)
}

func getCorrectMaxByType(fieldType string, value int) int {
	switch fieldType {
		case "uint8":
			return min(255, value)
		case "uint16":
			return min(65535, value)
		case "uint32":
			return min(4294967295, value)
		case "int8":
			return min(127, value)
		case "int16":
			return min(32767, value)
		case "int32":
			return min(2147483647, value)
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
