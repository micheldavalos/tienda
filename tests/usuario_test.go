package tests

import (
	"testing"
	"tienda/usuario"
)

func TestUsuarioNew(t *testing.T) {
	u := usuario.NewUsuario(1, "michel", "michel@udg.mx", "1234567")

	t.Log(u)
}

func TestUsuarioGetters(t *testing.T) {
	u := usuario.NewUsuario(1, "michel", "michel@udg.mx", "1234567")

	t.Log(u.GetID())
	t.Log(u.GetNombre())
	t.Log(u.GetEmail())
	t.Log(u.GetPassword())
}

func TestUsuarioSetters(t *testing.T) {
	u := usuario.NewUsuario(1, "michel", "michel@udg.mx", "1234567")

	t.Log(u)

	// modificar con setters
	u.SetID(10)
	u.SetNombre("Davalos")
	u.SetEmail("davalos@udg.mx")
	u.SetPassword("abcdefg")

	t.Log(u)
}