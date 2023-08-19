package main

import (
	"net/http"
	"math/rand"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"encoding/xml"
	"encoding/json"
)

// Структура для request body в роуте /login.
type RouteLoginRequest struct {
	Fields []Field `json:"fields" xml:"fields>field"`
}

// Структура для response body в роуте /login.
type RouteLoginResponse struct {
	Token string `json:"token" xml:"token"`
}

// Структура для response body для динамичеких полей
type ObjectResponse struct {
	XMLName xml.Name `xml:"object" json:"-"`
	Data map[string]interface{} `json:"object" xml:"-"`
}

// Структура для response body для массива динамичеких полей
type ObjectsResponse struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Data []ObjectResponse `json:"data" xml:"object"`
}

// Структура для response body, если будет ошибка. Если будут две и более ошибки, то используется RouteErrorList.
type RouteError struct {
	Message interface{} `json:"message" xml:"message"`
}

// Структура для response body, если будет успех.
type RouteSuccess struct {
	Data interface{} `json:"data" xml:"data"`
}

// Структура для ошибок валидации. Обычно используется там, где есть массив ошибок и RouteError не подходит.
type RouteErrorList struct {
	Index string `json:"index" xml:"index,attr"`
	Errors []string `json:"errors" xml:"errors>error"`
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

	if len(reuqest.Fields) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, r, RouteError{ Message: "Fields count must be less than 50" })

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

// Сериализация объекта в XML.
func (d ObjectResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

// Сериализация объекта в JSON.
func (d ObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Data)
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

	result := map[string]interface{}{}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, field := range fields {
		result[field.Name] = field.GetRandomValue()
	}

	if err := afterRoute(w, r, ObjectResponse{ Data: result }, http.StatusOK); err != nil {
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

	result := ObjectsResponse{}
	result.Data = make([]ObjectResponse, 20)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range result.Data {
		result.Data[i] = ObjectResponse{}
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

	var request map[string]interface{}

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

	var request map[string]interface{}

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

	var request map[string]interface{}

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