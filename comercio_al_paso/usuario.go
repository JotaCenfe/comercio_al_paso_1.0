package main

import (
	"errors"
	"fmt"
	"strings"
)

// Usuario representa al cliente que realiza la compra.
type Usuario struct {
	id       int
	nombre   string
	correo   string
	telefono string
}

// NewUsuario permite crear un usuario validando sus datos principales.
func NewUsuario(id int, nombre string, correo string, telefono string) (*Usuario, error) {
	if id <= 0 {
		return nil, errors.New("el ID del usuario debe ser mayor a cero")
	}

	if nombre == "" {
		return nil, errors.New("el nombre del usuario no puede estar vacío")
	}

	if !strings.Contains(correo, "@") {
		return nil, errors.New("el correo electrónico no es válido")
	}

	if telefono == "" {
		return nil, errors.New("el teléfono no puede estar vacío")
	}

	return &Usuario{
		id:       id,
		nombre:   nombre,
		correo:   correo,
		telefono: telefono,
	}, nil
}

// MostrarDatos imprime la información del cliente.
func (u *Usuario) MostrarDatos() {
	fmt.Println("Datos del cliente:")
	fmt.Println("Nombre:", u.nombre)
	fmt.Println("Correo:", u.correo)
	fmt.Println("Teléfono:", u.telefono)
}
