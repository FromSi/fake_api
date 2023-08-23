package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// Структура для request body в роуте /login.
type RouteLoginRequest struct {
	XMLName xml.Name `xml:"fields" json:"-"`
	Fields []Field `json:"fields" xml:"field"`
}

// Структура для response body в роуте /login.
type RouteLoginResponse struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Token string `json:"token" xml:"token"`
}

// Структура для response body для динамичеких полей
type Item struct {
	XMLName xml.Name `xml:"-" json:"-"`
	Data map[string]interface{} `json:"object" xml:"-"`
}

// Структура для response body для массива динамичеких полей
type Data struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Data []Item `json:"data" xml:"item"`
}

// Структура для response body, если будет ошибка. Если будут две и более ошибки, то используется RouteErrorList.
type RouteError struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Message interface{} `json:"message" xml:"message"`
}

// Структура для response body, если будет успех.
type RouteSuccess struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Data interface{} `json:"data" xml:"data"`
}

// Структура для ошибок валидации. Обычно используется там, где есть массив ошибок и RouteError не подходит.
type RouteErrorList struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Index string `json:"index" xml:"index,attr"`
	Errors []string `json:"errors" xml:"errors>error"`
}

// Костыль для сериализации RouteErrorList в XML.
type RouteErrorListXML struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Data []RouteErrorList `json:"data" xml:"item"`
}

// Сериализация объекта в XML.
func (d Item) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "Item" {
		start.Name.Local = "data"
	}

    if err := e.EncodeToken(start); err != nil {
        return err
    }

    for key, value := range d.Data {
        if err := e.EncodeElement(value, xml.StartElement{Name: xml.Name{Local: key}}); err != nil {
            return err
        }
    }

    return e.EncodeToken(start.End())
}

// Десериализация объекта из XML. Пример XML: <data> <id>123</id><id2>123</id2> </data>
func (o *Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	o.Data = make(map[string]interface{})

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch token := token.(type) {
		case xml.StartElement:
			var value string

			if err := d.DecodeElement(&value, &token); err != nil {
				return err
			}

			o.Data[token.Name.Local] = value
		case xml.EndElement:
			if token == start.End() {
				return nil
			}
		}
	}
}

// Сериализация объекта в JSON.
func (d Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Data)
}

// Десериализация объекта из JSON.
func (o *Item) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	o.Data = v

	return nil
}

// Роут для авторизации. Возвращает JWT токен. Поля для токена передаются в request body.
// Примеры запросов по request body можно посмотреть в файлах example_payload.json и example_payload.xml.
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

	fakeApiMaxFieldsInObject, err := strconv.Atoi(os.Getenv("FAKE_API_MAX_FIELDS_IN_OBJECT"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeResponse(w, r, RouteError{ Message: err.Error() })
	}

	if len(reuqest.Fields) > fakeApiMaxFieldsInObject {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: fmt.Sprintf("Maximum fields in object is %d", fakeApiMaxFieldsInObject) })

		return 
	}
	
	if err := validateFields(reuqest.Fields); err != nil {
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

// Роут для получения объекта со случайными данными. 
// Данные для полей берутся в зависимости от информации JWT токена.
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

	result := Item{}
	result.Data = map[string]interface{}{}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, field := range fields {
		result.Data[field.Name] = field.GetRandomValue()
	}

	if err := afterRoute(w, r, result, http.StatusOK); err != nil {
		return
	}
}

// Роут для получения массива объектов со случайными данными.
// Данные для полей берутся в зависимости от информации JWT токена.
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

	result := Data{}
	result.Data = make([]Item, 20)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range result.Data {
		result.Data[i] = Item{}
		result.Data[i].Data = map[string]interface{}{}

		for _, field := range fields {
			result.Data[i].Data[field.Name] = field.GetRandomValue()
		}
	}

	if err := afterRoute(w, r, result, http.StatusOK); err != nil {
		return
	}
}

// Роут для создания объекта со случайными данными.
// Данные для полей берутся в зависимости от информации JWT токена.
// В случае успеха возвращается статус 201 Created.
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

	var request Item

	if err := decodeRequest(w, r, &request); err != nil {
		return
	}

	validate, statusCode := validateFieldsFromRequest(w, r, request, fields, true, http.StatusCreated)

	if err := afterRoute(w, r, validate, statusCode); err != nil {
		return
	}
}

// Роут для обновления объекта в целом, со случайными данными.
// Данные для полей берутся в зависимости от информации JWT токена.
// В случае успеха возвращается статус 200 OK.
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

	var request Item

	if err := decodeRequest(w, r, &request); err != nil {
		return
	}

	validate, statusCode := validateFieldsFromRequest(w, r, request, fields, true, http.StatusOK)

	if err := afterRoute(w, r, validate, statusCode); err != nil {
		return
	}
}

// Роут для обновления объекта частично, со случайными данными.
// Данные для полей берутся в зависимости от информации JWT токена.
// В случае успеха возвращается статус 200 OK.
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

	var request Item

	if err := decodeRequest(w, r, &request); err != nil {
		return
	}

	validate, statusCode := validateFieldsFromRequest(w, r, request, fields, false, http.StatusOK)

	if err := afterRoute(w, r, validate, statusCode); err != nil {
		return
	}
}

// Роут для удаления объекта.
// Данные для полей берутся в зависимости от информации JWT токена.
// В случае успеха возвращается статус 204 No Content.
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