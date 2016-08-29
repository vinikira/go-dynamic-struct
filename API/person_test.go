package API

import (
	"regexp"
	"testing"
)

func TestPersonAPI(t *testing.T) {
	var (
		per        []Person
		err        error
		validEmail = regexp.MustCompile(`^[a-z0-9]+@+[a-z0-9]+\.+[a-z0-9]+.[a-z0-9]$`)
	)

	if per, err = GetPeople("http://localhost:8080/"); err != nil {
		t.Fatalf("Ocorreu um erro ao processar a função GetPeople: " + err.Error())
	}

	if len(per) != 5 {
		t.Fatalf("Tamanho do array de Person diferente do esperado")
	}

	for _, v := range per {
		if v.Email == "" {
			t.Fatalf("O campo Email está vazio")
		} else if v.Idade == 0 {
			t.Fatalf("O campo Idade está vazio")
		} else if v.Nome == "" {
			t.Fatalf("O campo Nome está vazio")
		} else if v.Sexo == "" {
			t.Fatalf("O campo Sexo está vazio")
		} else if v.Idade > 120 {
			t.Fatalf("O campo Idade contém um valor muito grande")
		} else if validEmail.MatchString(v.Email) {
			t.Fatalf("O campo Email é inválido")
		}
	}
}
