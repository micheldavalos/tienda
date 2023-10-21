package tests

import (
	"testing"
	"tienda/listausuarios"
	"tienda/usuario"
)

func TestListaUsuariosInsertarFinal(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)

	t.Log(lista)
	
	lista.Imprimir()
}

func TestListaUsuariosInsertarInicio(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")

	lista.InsertarInicio(u1)
	lista.InsertarInicio(u2)

	t.Log(lista)
	
	lista.Imprimir()
}

func TestListaUsuariosInsertarPosicion(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	t.Log(lista)
	
	lista.Imprimir()
}

func TestListaUsuariosEliminarFinal(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	t.Log(lista)
	
	lista.Imprimir()

	lista.EliminarFinal()

	lista.Imprimir()
}

func TestListaUsuariosEliminarInicio(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	t.Log(lista)
	
	lista.Imprimir()

	lista.EliminarInicio()

	lista.Imprimir()
}

func TestListaUsuariosEliminarPosicion(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	t.Log(lista)
	
	lista.Imprimir()

	err := lista.EliminarPosicion(1)
	if err != nil {
		t.Fatal(err)
	}

	lista.Imprimir()
}

func TestListaUsuariosObtener(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	t.Log(lista)
	
	lista.Imprimir()

	u, err := lista.Obtener(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(u)
}

func TestListaUsuariosBuscarByID(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	u, err := lista.BuscarByID(3)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(u)
}

func TestListaUsuariosBuscarByNombre(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	u, err := lista.BuscarByNombre("michel")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(u)
}

func TestListaUsuariosBuscarByEmail(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	u, err := lista.BuscarByEmail("boites@udg.mx")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(u)
}

func TestListaUsuariosOrdenarByID(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	lista.Imprimir()

	lista.OrdenarByID()

	lista.Imprimir()
}

func TestListaUsuariosOrdenarByNombre(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	lista.Imprimir()

	lista.OrdenarByNombre()

	lista.Imprimir()
}

func TestListaUsuariosOrdenarByEmail(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	lista.Imprimir()

	lista.OrdenarByEmail()

	lista.Imprimir()
}