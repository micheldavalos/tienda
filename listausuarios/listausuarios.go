package listausuarios

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"tienda/usuario"
)

var (
	ErrorPosicionInvalida = errors.New("posición invalida")
	ErrorListaVacia = errors.New("la lista está vacía")
	ErrorUsuarioNoEncontrado = errors.New("no existe el usuario")
)

type ListaUsuarios struct {
	data []*usuario.Usuario
}

func NewListaUsuarios() *ListaUsuarios {
	return &ListaUsuarios{}
}

func (l *ListaUsuarios) InsertarFinal(dato *usuario.Usuario) {
	l.data = append(l.data, dato)
}

func (l *ListaUsuarios) InsertarInicio(dato *usuario.Usuario) {
	nuevo := []*usuario.Usuario{dato}
	l.data = append(nuevo, l.data...)
}

func (l *ListaUsuarios) InsertarPosicion(dato *usuario.Usuario, pos int) error {
	if pos < 0 || pos >= len(l.data) {
		return ErrorPosicionInvalida
	}

	primerParte := l.data[:pos]
	segundaParte := l.data[pos:]

	var nuevo []*usuario.Usuario
	nuevo = append(nuevo, primerParte...)
	nuevo = append(nuevo, dato)
	nuevo = append(nuevo, segundaParte...)

	l.data = nuevo

	return nil
}

func (l *ListaUsuarios) Imprimir() {
	for i := 0; i < len(l.data); i++ {
		fmt.Println(l.data[i])
	}
}

func (l *ListaUsuarios) EliminarFinal() error {
	if len(l.data) == 0 {
		return ErrorListaVacia
	}

	l.data = l.data[:len(l.data)-1]
	return nil
}

func (l *ListaUsuarios) EliminarInicio() error {
	if len(l.data) == 0 {
		return ErrorListaVacia
	}

	l.data = l.data[1:]

	return nil
}

func (l *ListaUsuarios) EliminarPosicion(pos int) error {
	if pos < 0 || pos >= len(l.data) {
		return ErrorPosicionInvalida
	}
	// l.data[pos+1:] = [3, 4, 5]
	l.data = append(l.data[:pos], l.data[pos+1:]...)

	return nil
}

func (l *ListaUsuarios) Obtener(posicion int) (*usuario.Usuario, error) {
	switch {
	case len(l.data) == 0:
		return nil, ErrorListaVacia
	case posicion < 0 || posicion >= len(l.data):
		return nil, ErrorPosicionInvalida
	}

	return l.data[posicion], nil
}

func (l *ListaUsuarios) BuscarByID(id int) (*usuario.Usuario, error) {
	for i := 0; i < len(l.data); i++ {
		if l.data[i].GetID() == id {
			return l.data[i], nil
		}
	}

	return nil, ErrorUsuarioNoEncontrado
}

func (l *ListaUsuarios) BuscarByNombre(nombre string) (*usuario.Usuario, error) {
	for i := 0; i < len(l.data); i++ {
		if l.data[i].GetNombre() == nombre {
			return l.data[i], nil
		}
	}

	return nil, ErrorUsuarioNoEncontrado
}

func (l *ListaUsuarios) BuscarByEmail(email string) (*usuario.Usuario, error) {
	for i := 0; i < len(l.data); i++ {
		if l.data[i].GetEmail() == email {
			return l.data[i], nil
		}
	}

	return nil, ErrorUsuarioNoEncontrado
}

func (l *ListaUsuarios) OrdenarByID() {
	sort.Slice(l.data, func(i, j int) bool {
		return l.data[i].GetID() < l.data[j].GetID()
	})
}

func (l *ListaUsuarios) OrdenarByNombre() {
	sort.Slice(l.data, func(i, j int) bool {
		return l.data[i].GetNombre() < l.data[j].GetNombre()
	})
}

func (l *ListaUsuarios) OrdenarByEmail() {
	sort.Slice(l.data, func(i, j int) bool {
		return l.data[i].GetEmail() < l.data[j].GetEmail()
	})
}

func (l *ListaUsuarios) Cantidad() int {
	return len(l.data)
}

func (l *ListaUsuarios) GuardarCSV() error {
	archivo, err := os.Create("usuarios.csv")
	if err !=  nil {
		return err
	}
	defer archivo.Close()

	for i := 0; i < l.Cantidad(); i++ {
		u, err := l.Obtener(i)
		if err != nil {
			return err
		}
		archivo.WriteString(fmt.Sprintf("%d", u.GetID()) + ",")
		archivo.WriteString(u.GetNombre() + ",")
		archivo.WriteString(u.GetEmail() + ",")
		archivo.WriteString(u.GetPassword() + "\n")	
	}

	return nil
}

func (l *ListaUsuarios) RecuperarCSV() error {
	archivo, err := os.Open("usuarios.csv")
	if err != nil {
		return err
	}
	defer archivo.Close()

	var linea string
	for {
		_, err := fmt.Fscanf(archivo, "%s\n", &linea)
		if err != nil {
			if err == io.EOF { // End of File
				break
			}
			return err
		}

		campos := strings.Split(linea, ",")
		id, err := strconv.Atoi(campos[0])
		if err != nil {
			return err
		}
		nombre := campos[1]
		email := campos[2]
		password := campos[3]

		u := usuario.NewUsuario(id, nombre, email, password)
		l.InsertarFinal(u)
	}

	return nil
}

func (l *ListaUsuarios) GuardarJSON() error {
	archivo, err := os.Create("usuarios.json")
	if err != nil {
		return err
	}
	defer archivo.Close()

	b, err := json.MarshalIndent(l.data, "", " ")
	if err != nil {
		return err
	}

	archivo.Write(b)

	return nil
}

func (l *ListaUsuarios) RecuperJSON() error {
	archivo, err := os.Open("usuarios.json")
	if err != nil {
		return err
	}
	defer archivo.Close()

	err = json.NewDecoder(archivo).Decode(&l.data)
	if err != nil {
		return err
	}

	return nil
}