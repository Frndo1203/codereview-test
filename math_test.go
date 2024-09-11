package main

func TestSoma(t, *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalido: Resultado %d. Espado: %d", total, 30)
	}
}