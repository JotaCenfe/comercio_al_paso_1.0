package main

import (
	"errors"
	"fmt"
)

const (
	EstadoPendiente = "Pendiente de pago"
	EstadoPagado    = "Pagado"
)

// Pedido relaciona un usuario con su carrito de compras.
type Pedido struct {
	id      int
	usuario *Usuario
	carrito *Carrito
	estado  string
}

// NewPedido crea un pedido validando que exista usuario y carrito.
func NewPedido(id int, usuario *Usuario, carrito *Carrito) (*Pedido, error) {
	if id <= 0 {
		return nil, errors.New("el ID del pedido debe ser mayor a cero")
	}

	if usuario == nil {
		return nil, errors.New("no existe un usuario registrado")
	}

	if carrito == nil || carrito.EstaVacio() {
		return nil, errors.New("no se puede generar un pedido con el carrito vacío")
	}

	return &Pedido{
		id:      id,
		usuario: usuario,
		carrito: carrito,
		estado:  EstadoPendiente,
	}, nil
}

// Total devuelve el valor total del pedido.
func (p *Pedido) Total() float64 {
	return p.carrito.CalcularTotal()
}

// MarcarPagado actualiza el estado del pedido luego del pago.
func (p *Pedido) MarcarPagado() {
	p.estado = EstadoPagado
}

// MostrarResumen imprime el resumen final del pedido.
func (p *Pedido) MostrarResumen() {
	fmt.Println("\n--- RESUMEN DEL PEDIDO ---")
	fmt.Println("Pedido Nro.:", p.id)
	p.usuario.MostrarDatos()
	p.carrito.MostrarCarrito()
	fmt.Println("Estado del pedido:", p.estado)
}
