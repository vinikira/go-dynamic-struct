package API

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Sexo   string `json:"sexo"`
	Idade  int    `json:"idade"`
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

	if err = json.Unmarshal(body, &p); err != nil {
		return []Person{}, err
	}

	if err = json.Unmarshal(body, &others); err != nil {
		return []Person{}, err
	}

	for i := range p {
		p[i].Outros = make(map[string]interface{})
		for k, v := range others[i] {
			if k != "nome" && k != "email" && k != "sexo" && k != "idade" {
				p[i].Outros[k] = v
			}
		}
	}

	return p, nil
}
