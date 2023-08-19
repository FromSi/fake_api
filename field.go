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
		case "latitude":
			return gofakeit.Latitude()
		case "longitude":
			return gofakeit.Longitude()
		case "boolean":
			return gofakeit.Bool()
		case "string_name":
			return gofakeit.Name()
		case "string_first_name":
			return gofakeit.FirstName()
		case "string_middle_name":
			return gofakeit.MiddleName()
		case "string_last_name":
			return gofakeit.LastName()
		case "string_gender":
			return gofakeit.Gender()
		case "string_ssn":
			return gofakeit.SSN()
		case "string_hobby":
			return gofakeit.Hobby()
		case "string_email":
			return gofakeit.Email()
		case "string_username":
			return gofakeit.Username()
		case "string_country":
			return gofakeit.Country()
		case "string_country_abr":
			return gofakeit.CountryAbr()
		case "string_city":
			return gofakeit.City()
		case "string_state":
			return gofakeit.State()
		case "string_street":
			return gofakeit.Street()
		case "string_street_name":
			return gofakeit.StreetName()
		case "string_street_number":
			return gofakeit.StreetNumber()
		case "string_street_prefix":
			return gofakeit.StreetPrefix()
		case "string_street_suffix":
			return gofakeit.StreetSuffix()
		case "string_zip":
			return gofakeit.Zip()
		case "string_gametag":
			return gofakeit.Gamertag()
		case "string_beer_alcohol":
			return gofakeit.BeerAlcohol()
		case "string_beer_blg":
			return gofakeit.BeerBlg()
		case "string_beer_hop":
			return gofakeit.BeerHop()
		case "string_beer_ibu":
			return gofakeit.BeerIbu()
		case "string_beer_malt":
			return gofakeit.BeerMalt()
		case "string_beer_name":
			return gofakeit.BeerName()
		case "string_beer_style":
			return gofakeit.BeerStyle()
		case "string_beer_yeast":
			return gofakeit.BeerYeast()
		case "string_noun":
			return gofakeit.Noun()
		case "string_noun_common":
			return gofakeit.NounCommon()
		case "string_noun_concrete":
			return gofakeit.NounConcrete()
		case "string_noun_abstract":
			return gofakeit.NounAbstract()
		case "string_noun_collective_people":
			return gofakeit.NounCollectivePeople()
		case "string_noun_collective_animal":
			return gofakeit.NounCollectiveAnimal()
		case "string_noun_collective_thing":
			return gofakeit.NounCollectiveThing()
		case "string_noun_countable":
			return gofakeit.NounCountable()
		case "string_noun_uncountable":
			return gofakeit.NounUncountable()
		case "string_verb":
			return gofakeit.Verb()
		case "string_verb_action":
			return gofakeit.VerbAction()
		case "string_verb_linking":
			return gofakeit.VerbLinking()
		case "string_verb_helping":
			return gofakeit.VerbHelping()
		case "string_adverb":
			return gofakeit.Adverb()
		case "string_adverb_manner":
			return gofakeit.AdverbManner()
		case "string_adverb_degree":
			return gofakeit.AdverbDegree()
		case "string_adverb_place":
			return gofakeit.AdverbPlace()
		case "string_adverb_time_definite":
			return gofakeit.AdverbTimeDefinite()
		case "string_adverb_time_indefinite":
			return gofakeit.AdverbTimeIndefinite()
		case "string_adverb_frequency_definite":
			return gofakeit.AdverbFrequencyDefinite()
		case "string_adverb_frequency_indefinite":
			return gofakeit.AdverbFrequencyIndefinite()
		case "string_preposition":
			return gofakeit.Preposition()
		case "string_preposition_simple":
			return gofakeit.PrepositionSimple()
		case "string_preposition_double":
			return gofakeit.PrepositionDouble()
		case "string_preposition_compound":
			return gofakeit.PrepositionCompound()
		case "string_adjective":
			return gofakeit.Adjective()
		case "string_adjective_descriptive":
			return gofakeit.AdjectiveDescriptive()
		case "string_adjective_quantitative":
			return gofakeit.AdjectiveQuantitative()
		case "string_adjective_proper":
			return gofakeit.AdjectiveProper()
		case "string_adjective_demonstrative":
			return gofakeit.AdjectiveDemonstrative()
		case "string_adjective_possessive":
			return gofakeit.AdjectivePossessive()
		case "string_adjective_interrogative":
			return gofakeit.AdjectiveInterrogative()
		case "string_adjective_indefinite":
			return gofakeit.AdjectiveIndefinite()
		case "string_pronoun":
			return gofakeit.Pronoun()
		case "string_pronoun_personal":
			return gofakeit.PronounPersonal()
		case "string_pronoun_object":
			return gofakeit.PronounObject()
		case "string_pronoun_possessive":
			return gofakeit.PronounPossessive()
		case "string_pronoun_reflective":
			return gofakeit.PronounReflective()
		case "string_pronoun_demonstrative":
			return gofakeit.PronounDemonstrative()
		case "string_pronoun_interrogative":
			return gofakeit.PronounInterrogative()
		case "string_pronoun_relative":
			return gofakeit.PronounRelative()
		case "string_connective":
			return gofakeit.Connective()
		case "string_connective_time":
			return gofakeit.ConnectiveTime()
		case "string_connective_comparative":
			return gofakeit.ConnectiveComparative()
		case "string_connective_complaint":
			return gofakeit.ConnectiveComplaint()
		case "string_connective_listing":
			return gofakeit.ConnectiveListing()
		case "string_connective_casual":
			return gofakeit.ConnectiveCasual()
		case "string_connective_examplify":
			return gofakeit.ConnectiveExamplify()
		case "string_question":
			return gofakeit.Question()
		case "string_quote":
			return gofakeit.Quote()
		case "string_phrase":
			return gofakeit.Phrase()
		case "string_word":
			return gofakeit.Word()
		case "string_sentence":
			return gofakeit.Sentence(f.GetCorrectMax())
		case "string_url":
			return gofakeit.URL()
		case "string_image_url":
			return gofakeit.ImageURL(f.GetCorrectMax(), f.GetCorrectMax())
		case "string_uuid":
			return gofakeit.UUID()
		case "string_color":
			return gofakeit.Color()
		case "string_hex_color":
			return gofakeit.HexColor()
		case "string_safe_color":
			return gofakeit.SafeColor()
		case "string_phone":
			return gofakeit.Phone()
		case "string_phone_formatted":
			return gofakeit.PhoneFormatted()
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
		case "string_car_maker":
			return gofakeit.CarMaker()
		case "string_car_model":
			return gofakeit.CarModel()
		case "string_car_type":
			return gofakeit.CarType()
		case "string_car_fuel_type":
			return gofakeit.CarFuelType()
		case "string_car_transmission_type":
			return gofakeit.CarTransmissionType()
		case "string_fruit":
			return gofakeit.Fruit()
		case "string_vegetable":
			return gofakeit.Vegetable()
		case "string_breakfast":
			return gofakeit.Breakfast()
		case "string_lunch":
			return gofakeit.Lunch()
		case "string_dinner":
			return gofakeit.Dinner()
		case "string_snack":
			return gofakeit.Snack()
		case "string_dessert":
			return gofakeit.Dessert()
		case "string_flip_a_coin":
			return gofakeit.FlipACoin()
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
		case "latitude":
			return max(-90, value)
		case "longitude":
			return max(-180, value)
		case "boolean":
			return max(0, value)
		case "string_name":
			return max(1, value)
		case "string_first_name":
			return max(1, value)
		case "string_middle_name":
			return max(1, value)
		case "string_last_name":
			return max(1, value)
		case "string_gender":
			return max(1, value)
		case "string_ssn":
			return max(1, value)
		case "string_hobby":
			return max(1, value)
		case "string_email":
			return max(5, value)
		case "string_username":
			return max(1, value)
		case "string_country":
			return max(1, value)
		case "string_country_abr":
			return max(1, value)
		case "string_city":
			return max(1, value)
		case "string_state":
			return max(1, value)
		case "string_street":
			return max(1, value)
		case "string_street_name":
			return max(1, value)
		case "string_street_number":
			return max(1, value)
		case "string_street_prefix":
			return max(1, value)
		case "string_street_suffix":
			return max(1, value)
		case "string_zip":
			return max(1, value)
		case "string_gametag":
			return max(1, value)
		case "string_beer_alcohol":
			return max(1, value)
		case "string_beer_blg":
			return max(1, value)
		case "string_beer_hop":
			return max(1, value)
		case "string_beer_ibu":
			return max(1, value)
		case "string_beer_malt":
			return max(1, value)
		case "string_beer_name":
			return max(1, value)
		case "string_beer_style":
			return max(1, value)
		case "string_beer_yeast":
			return max(1, value)
		case "string_noun":
			return max(1, value)
		case "string_noun_common":
			return max(1, value)
		case "string_noun_concrete":
			return max(1, value)
		case "string_noun_abstract":
			return max(1, value)
		case "string_noun_collective_people":
			return max(1, value)
		case "string_noun_collective_animal":
			return max(1, value)
		case "string_noun_collective_thing":
			return max(1, value)
		case "string_noun_countable":
			return max(1, value)
		case "string_noun_uncountable":
			return max(1, value)
		case "string_verb":
			return max(1, value)
		case "string_verb_action":
			return max(1, value)
		case "string_verb_linking":
			return max(1, value)
		case "string_verb_helping":
			return max(1, value)
		case "string_adverb":
			return max(1, value)
		case "string_adverb_manner":
			return max(1, value)
		case "string_adverb_degree":
			return max(1, value)
		case "string_adverb_place":
			return max(1, value)
		case "string_adverb_time_definite":
			return max(1, value)
		case "string_adverb_time_indefinite":
			return max(1, value)
		case "string_adverb_frequency_definite":
			return max(1, value)
		case "string_adverb_frequency_indefinite":
			return max(1, value)
		case "string_preposition":
			return max(1, value)
		case "string_preposition_simple":
			return max(1, value)
		case "string_preposition_double":
			return max(1, value)
		case "string_preposition_compound":
			return max(1, value)
		case "string_adjective":
			return max(1, value)
		case "string_adjective_descriptive":
			return max(1, value)
		case "string_adjective_quantitative":
			return max(1, value)
		case "string_adjective_proper":
			return max(1, value)
		case "string_adjective_demonstrative":
			return max(1, value)
		case "string_adjective_possessive":
			return max(1, value)
		case "string_adjective_interrogative":
			return max(1, value)
		case "string_adjective_indefinite":
			return max(1, value)
		case "string_pronoun":
			return max(1, value)
		case "string_pronoun_personal":
			return max(1, value)
		case "string_pronoun_object":
			return max(1, value)
		case "string_pronoun_possessive":
			return max(1, value)
		case "string_pronoun_reflective":
			return max(1, value)
		case "string_pronoun_demonstrative":
			return max(1, value)
		case "string_pronoun_interrogative":
			return max(1, value)
		case "string_pronoun_relative":
			return max(1, value)
		case "string_connective":
			return max(1, value)
		case "string_connective_time":
			return max(1, value)
		case "string_connective_comparative":
			return max(1, value)
		case "string_connective_complaint":
			return max(1, value)
		case "string_connective_listing":
			return max(1, value)
		case "string_connective_casual":
			return max(1, value)
		case "string_connective_examplify":
			return max(1, value)
		case "string_question":
			return max(1, value)
		case "string_quote":
			return max(1, value)
		case "string_phrase":
			return max(1, value)
		case "string_word":
			return max(1, value)
		case "string_sentence":
			return max(1, value)
		case "string_url":
			return max(10, value)
		case "string_image_url":
			return max(1, value)
		case "string_uuid":
			return max(36, value)
		case "string_color":
			return max(1, value)
		case "string_hex_color":
			return max(7, value)
		case "string_safe_color":
			return max(1, value)
		case "string_phone":
			return max(1, value)
		case "string_phone_formatted":
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
		case "string_car_maker":
			return max(1, value)
		case "string_car_model":
			return max(1, value)
		case "string_car_type":
			return max(1, value)
		case "string_car_fuel_type":
			return max(1, value)
		case "string_car_transmission_type":
			return max(1, value)
		case "string_fruit":
			return max(1, value)
		case "string_vegetable":
			return max(1, value)
		case "string_breakfast":
			return max(1, value)
		case "string_lunch":
			return max(1, value)
		case "string_dinner":
			return max(1, value)
		case "string_snack":
			return max(1, value)
		case "string_dessert":
			return max(1, value)
		case "string_flip_a_coin":
			return max(1, value)
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
		case "latitude":
			return min(90, value)
		case "longitude":
			return min(180, value)
		case "boolean":
			return min(1, value)
		case "string_name":
			return min(255, value)
		case "string_first_name":
			return min(255, value)
		case "string_middle_name":
			return min(255, value)
		case "string_last_name":
			return min(255, value)
		case "string_gender":
			return min(255, value)
		case "string_ssn":
			return min(255, value)
		case "string_hobby":
			return min(255, value)
		case "string_email":
			return min(255, value)
		case "string_username":
			return min(255, value)
		case "string_country":
			return min(255, value)
		case "string_country_abr":
			return min(255, value)
		case "string_city":
			return min(255, value)
		case "string_state":
			return min(255, value)
		case "string_street":
			return min(255, value)
		case "string_street_name":
			return min(255, value)
		case "string_street_number":
			return min(255, value)
		case "string_street_prefix":
			return min(255, value)
		case "string_street_suffix":
			return min(255, value)
		case "string_zip":
			return min(255, value)
		case "string_gametag":
			return min(255, value)
		case "string_beer_alcohol":
			return min(255, value)
		case "string_beer_blg":
			return min(255, value)
		case "string_beer_hop":
			return min(255, value)
		case "string_beer_ibu":
			return min(255, value)
		case "string_beer_malt":
			return min(255, value)
		case "string_beer_name":
			return min(255, value)
		case "string_beer_style":
			return min(255, value)
		case "string_beer_yeast":
			return min(255, value)
		case "string_noun":
			return min(255, value)
		case "string_noun_common":
			return min(255, value)
		case "string_noun_concrete":
			return min(255, value)
		case "string_noun_abstract":
			return min(255, value)
		case "string_noun_collective_people":
			return min(255, value)
		case "string_noun_collective_animal":
			return min(255, value)
		case "string_noun_collective_thing":
			return min(255, value)
		case "string_noun_countable":
			return min(255, value)
		case "string_noun_uncountable":
			return min(255, value)
		case "string_verb":
			return min(255, value)
		case "string_verb_action":
			return min(255, value)
		case "string_verb_linking":
			return min(255, value)
		case "string_verb_helping":
			return min(255, value)
		case "string_adverb":
			return min(255, value)
		case "string_adverb_manner":
			return min(255, value)
		case "string_adverb_degree":
			return min(255, value)
		case "string_adverb_place":
			return min(255, value)
		case "string_adverb_time_definite":
			return min(255, value)
		case "string_adverb_time_indefinite":
			return min(255, value)
		case "string_adverb_frequency_definite":
			return min(255, value)
		case "string_adverb_frequency_indefinite":
			return min(255, value)
		case "string_preposition":
			return min(255, value)
		case "string_preposition_simple":
			return min(255, value)
		case "string_preposition_double":
			return min(255, value)
		case "string_preposition_compound":
			return min(255, value)
		case "string_adjective":
			return min(255, value)
		case "string_adjective_descriptive":
			return min(255, value)
		case "string_adjective_quantitative":
			return min(255, value)
		case "string_adjective_proper":
			return min(255, value)
		case "string_adjective_demonstrative":
			return min(255, value)
		case "string_adjective_possessive":
			return min(255, value)
		case "string_adjective_interrogative":
			return min(255, value)
		case "string_adjective_indefinite":
			return min(255, value)
		case "string_pronoun":
			return min(255, value)
		case "string_pronoun_personal":
			return min(255, value)
		case "string_pronoun_object":
			return min(255, value)
		case "string_pronoun_possessive":
			return min(255, value)
		case "string_pronoun_reflective":
			return min(255, value)
		case "string_pronoun_demonstrative":
			return min(255, value)
		case "string_pronoun_interrogative":
			return min(255, value)
		case "string_pronoun_relative":
			return min(255, value)
		case "string_connective":
			return min(255, value)
		case "string_connective_time":
			return min(255, value)
		case "string_connective_comparative":
			return min(255, value)
		case "string_connective_complaint":
			return min(255, value)
		case "string_connective_listing":
			return min(255, value)
		case "string_connective_casual":
			return min(255, value)
		case "string_connective_examplify":
			return min(255, value)
		case "string_question":
			return min(255, value)
		case "string_quote":
			return min(255, value)
		case "string_phrase":
			return min(255, value)
		case "string_word":
			return min(255, value)
		case "string_sentence":
			return min(2048, value)
		case "string_url":
			return min(255, value)
		case "string_image_url":
			return min(2000, value)
		case "string_uuid":
			return min(36, value)
		case "string_color":
			return min(255, value)
		case "string_hex_color":
			return min(7, value)
		case "string_safe_color":
			return min(255, value)
		case "string_phone":
			return min(20, value)
		case "string_phone_formatted":
			return min(255, value)
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
		case "string_car_maker":
			return min(255, value)
		case "string_car_model":
			return min(255, value)
		case "string_car_type":
			return min(255, value)
		case "string_car_fuel_type":
			return min(255, value)
		case "string_car_transmission_type":
			return min(255, value)
		case "string_fruit":
			return min(255, value)
		case "string_vegetable":
			return min(255, value)
		case "string_breakfast":
			return min(255, value)
		case "string_lunch":
			return min(255, value)
		case "string_dinner":
			return min(255, value)
		case "string_snack":
			return min(255, value)
		case "string_dessert":
			return min(255, value)
		case "string_flip_a_coin":
			return min(255, value)
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
		case "latitude":
			return "float64"
		case "longitude":
			return "float64"
		case "boolean":
			return "bool"
		case "string_name":
			return "string"
		case "string_first_name":
			return "string"
		case "string_middle_name":
			return "string"
		case "string_last_name":
			return "string"
		case "string_gender":
			return "string"
		case "string_ssn":
			return "string"
		case "string_hobby":
			return "string"
		case "string_email":
			return "string"
		case "string_username":
			return "string"
		case "string_country":
			return "string"
		case "string_country_abr":
			return "string"
		case "string_city":
			return "string"
		case "string_state":
			return "string"
		case "string_street":
			return "string"
		case "string_street_name":
			return "string"
		case "string_street_number":
			return "string"
		case "string_street_prefix":
			return "string"
		case "string_street_suffix":
			return "string"
		case "string_zip":
			return "string"
		case "string_gametag":
			return "string"
		case "string_beer_alcohol":
			return "string"
		case "string_beer_blg":
			return "string"
		case "string_beer_hop":
			return "string"
		case "string_beer_ibu":
			return "string"
		case "string_beer_malt":
			return "string"
		case "string_beer_name":
			return "string"
		case "string_beer_style":
			return "string"
		case "string_beer_yeast":
			return "string"
		case "string_noun":
			return "string"
		case "string_noun_common":
			return "string"
		case "string_noun_concrete":
			return "string"
		case "string_noun_abstract":
			return "string"
		case "string_noun_collective_people":
			return "string"
		case "string_noun_collective_animal":
			return "string"
		case "string_noun_collective_thing":
			return "string"
		case "string_noun_countable":
			return "string"
		case "string_noun_uncountable":
			return "string"
		case "string_verb":
			return "string"
		case "string_verb_action":
			return "string"
		case "string_verb_linking":
			return "string"
		case "string_verb_helping":
			return "string"
		case "string_adverb":
			return "string"
		case "string_adverb_manner":
			return "string"
		case "string_adverb_degree":
			return "string"
		case "string_adverb_place":
			return "string"
		case "string_adverb_time_definite":
			return "string"
		case "string_adverb_time_indefinite":
			return "string"
		case "string_adverb_frequency_definite":
			return "string"
		case "string_adverb_frequency_indefinite":
			return "string"
		case "string_preposition":
			return "string"
		case "string_preposition_simple":
			return "string"
		case "string_preposition_double":
			return "string"
		case "string_preposition_compound":
			return "string"
		case "string_adjective":
			return "string"
		case "string_adjective_descriptive":
			return "string"
		case "string_adjective_quantitative":
			return "string"
		case "string_adjective_proper":
			return "string"
		case "string_adjective_demonstrative":
			return "string"
		case "string_adjective_possessive":
			return "string"
		case "string_adjective_interrogative":
			return "string"
		case "string_adjective_indefinite":
			return "string"
		case "string_pronoun":
			return "string"
		case "string_pronoun_personal":
			return "string"
		case "string_pronoun_object":
			return "string"
		case "string_pronoun_possessive":
			return "string"
		case "string_pronoun_reflective":
			return "string"
		case "string_pronoun_demonstrative":
			return "string"
		case "string_pronoun_interrogative":
			return "string"
		case "string_pronoun_relative":
			return "string"
		case "string_connective":
			return "string"
		case "string_connective_time":
			return "string"
		case "string_connective_comparative":
			return "string"
		case "string_connective_complaint":
			return "string"
		case "string_connective_listing":
			return "string"
		case "string_connective_casual":
			return "string"
		case "string_connective_examplify":
			return "string"
		case "string_question":
			return "string"
		case "string_quote":
			return "string"
		case "string_phrase":
			return "string"
		case "string_word":
			return "string"
		case "string_sentence":
			return "string"
		case "string_url":
			return "string"
		case "string_image_url":
			return "string"
		case "string_uuid":
			return "string"
		case "string_color":
			return "string"
		case "string_hex_color":
			return "string"
		case "string_safe_color":
			return "string"
		case "string_phone":
			return "string"
		case "string_phone_formatted":
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
		case "string_car_maker":
			return "string"
		case "string_car_model":
			return "string"
		case "string_car_type":
			return "string"
		case "string_car_fuel_type":
			return "string"
		case "string_car_transmission_type":
			return "string"
		case "string_fruit":
			return "string"
		case "string_vegetable":
			return "string"
		case "string_breakfast":
			return "string"
		case "string_lunch":
			return "string"
		case "string_dinner":
			return "string"
		case "string_snack":
			return "string"
		case "string_dessert":
			return "string"
		case "string_flip_a_coin":
			return "string"
		default:
			return "error"
	}
}