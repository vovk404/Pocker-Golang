package authentication

type Autentication interface {
	Login(entry LoginRequest) error
}