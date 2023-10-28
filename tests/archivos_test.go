package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"tienda/listausuarios"
	"tienda/usuario"
)

func TestCrearArchivo(t *testing.T) {
	archivo, err := os.Create("ejemplo.txt")
	
	if err != nil {
		t.Fatal(err)
	}

	defer archivo.Close()

	n, err := archivo.WriteString("Hola desde un archivo")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Se escribieron %d bytes", n)
}

func TestUsuarioGuardarCSV(t *testing.T) {
	archivo, err := os.Create("usuario.csv")
	if err !=  nil {
		t.Fatal(err)
	}
	defer archivo.Close()

	u := usuario.NewUsuario(1, "cucei", "cucei@udg.mx", "abcdefg")

	archivo.WriteString(fmt.Sprintf("%d", u.GetID()) + ",")
	archivo.WriteString(u.GetNombre() + ",")
	archivo.WriteString(u.GetEmail() + ",")
	archivo.WriteString(u.GetPassword() + "\n")	

}

func TestListaUsuariosGuardarCSV(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	archivo, err := os.Create("usuarios.csv")
	if err !=  nil {
		t.Fatal(err)
	}
	defer archivo.Close()

	for i := 0; i < lista.Cantidad(); i++ {
		u, err := lista.Obtener(i)
		if err != nil {
			t.Fatal(err)
		}
		archivo.WriteString(fmt.Sprintf("%d", u.GetID()) + ",")
		archivo.WriteString(u.GetNombre() + ",")
		archivo.WriteString(u.GetEmail() + ",")
		archivo.WriteString(u.GetPassword() + "\n")	
	}
}

func TestListaUsuariosRecuperarCSV(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	archivo, err := os.Open("usuarios.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer archivo.Close()


	var linea string
	for {
		n, err := fmt.Fscanf(archivo, "%s\n", &linea)
		if err != nil {
			if err == io.EOF { // End of File
				break
			}
			t.Fatal(err)
		}
		t.Logf("Se leyerón %d bytes", n)
		t.Log(linea)

		campos := strings.Split(linea, ",")
		id, err := strconv.Atoi(campos[0])
		if err != nil {
			t.Fatal(err)
		}
		nombre := campos[1]
		email := campos[2]
		password := campos[3]

		u := usuario.NewUsuario(id, nombre, email, password)
		lista.InsertarFinal(u)
	}

	lista.Imprimir()
}

func TestUsuarioJSON(t *testing.T) {
	archivo, err := os.Create("usuario.json")
	if err != nil {
		t.Fatal(err)
	}
	defer archivo.Close()

	u := usuario.NewUsuario(3, "cucei", "cucei@udg.mx", "1234567")

	b, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))
	archivo.Write(b)
}

func TestListaUsuariosGuardarJSON(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	u1 := usuario.NewUsuario(1, "michel", "michel@udg.mx", "12345")
	u2 := usuario.NewUsuario(2, "davalos", "davalos@udg.mx", "abcdef")
	u3 := usuario.NewUsuario(3, "boites", "boites@udg.mx", "a1b2c3")

	lista.InsertarFinal(u1)
	lista.InsertarFinal(u2)
	lista.InsertarPosicion(u3, 1)

	err := lista.GuardarJSON()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsuarioRecuperarJSON(t *testing.T) {
	archivo, err := os.Open("usuario.json")
	if err != nil {
		t.Fatal(err)
	}
	defer archivo.Close()

	var usuario usuario.Usuario
	t.Logf("%v", usuario)

	err = json.NewDecoder(archivo).Decode(&usuario)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", usuario)
}

func TestListaUsuariosRecuperarJSON(t *testing.T) {
	lista := listausuarios.NewListaUsuarios()

	err := lista.RecuperJSON()
	if err != nil {
		t.Fatal(err)
	}

	lista.Imprimir()
}