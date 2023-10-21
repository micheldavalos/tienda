package usuario

type Usuario struct {
	ID       int
	Nombre   string
	Email    string
	Password string
}

func NewUsuario(id int, nombre, email, password string) *Usuario {
	return &Usuario{
		ID: id,
		Nombre: nombre,
		Email: email,
		Password: password,
	}
}

func (u *Usuario) GetID() int {
	return u.ID
}

func (u *Usuario) GetNombre() string {
	return u.Nombre
}

func (u *Usuario) GetEmail() string {
	return u.Email
}

func (u *Usuario) GetPassword() string {
	return u.Password
}

func (u *Usuario) SetID(id int) {
	u.ID = id
}

func (u *Usuario) SetNombre(nombre string) {
	u.Nombre = nombre
}

func (u *Usuario) SetEmail(email string) {
	u.Email = email
}

func (u *Usuario) SetPassword(password string) {
	u.Password = password
} 
