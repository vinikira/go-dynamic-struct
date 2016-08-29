package API

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Nome   string
	Email  string
	Sexo   string
	Idade  int
	Outros map[string]interface{}
}

func GetPeople(url string) ([]Person, error) {
	var (
		resp   *http.Response
		err    error
		body   []byte
		p      []Person
		others []map[string]interface{}
	)

	if resp, err = http.Get(url); err != nil {
		return []Person{}, err
	}

	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return []Person{}, err
	}

	if err = json.Unmarshal(body, &others); err != nil {
		return []Person{}, err
	}

	for _, v := range others {
		var per Person
		per.Outros = make(map[string]interface{})
		for k, val := range v {
			switch k {
			case "nome":
				per.Nome = val.(string)
			case "email":
				per.Email = val.(string)
			case "sexo":
				per.Sexo = val.(string)
			case "idade":
				per.Idade = int(val.(float64))
			default:
				per.Outros[k] = val
			}
		}
		p = append(p, per)
	}

	return p, nil
}
