package main

import (
	"errors"
	"fmt"
)

// Producto representa un artículo disponible para la venta.
// Sus atributos están en minúscula para aplicar encapsulación.
type Producto struct {
	id          int
	nombre      string
	descripcion string
	precio      float64
	stock       int
}

// NewProducto funciona como constructor para crear productos válidos.
func NewProducto(id int, nombre string, descripcion string, precio float64, stock int) (*Producto, error) {
	if id <= 0 {
		return nil, errors.New("el ID del producto debe ser mayor a cero")
	}

	if nombre == "" {
		return nil, errors.New("el nombre del producto no puede estar vacío")
	}

	if precio <= 0 {
		return nil, errors.New("el precio del producto debe ser mayor a cero")
	}

	if stock < 0 {
		return nil, errors.New("el stock no puede ser negativo")
	}

	return &Producto{
		id:          id,
		nombre:      nombre,
		descripcion: descripcion,
		precio:      precio,
		stock:       stock,
	}, nil
}

// Métodos getter para acceder a los datos protegidos.
func (p *Producto) ID() int {
	return p.id
}

func (p *Producto) Nombre() string {
	return p.nombre
}

func (p *Producto) Precio() float64 {
	return p.precio
}

func (p *Producto) Stock() int {
	return p.stock
}

// MostrarInformacion imprime los datos básicos del producto.
func (p *Producto) MostrarInformacion() {
	fmt.Printf("ID: %d | Producto: %s | Descripción: %s | Precio: $%.2f | Stock: %d\n",
		p.id, p.nombre, p.descripcion, p.precio, p.stock)
}

// TieneStock verifica si existe stock suficiente para la cantidad solicitada.
func (p *Producto) TieneStock(cantidad int) bool {
	return cantidad > 0 && cantidad <= p.stock
}

// ReducirStock descuenta la cantidad comprada del stock disponible.
func (p *Producto) ReducirStock(cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	if !p.TieneStock(cantidad) {
		return errors.New("stock insuficiente para realizar la compra")
	}

	p.stock -= cantidad
	return nil
}
