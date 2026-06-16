package main

import (
	"errors"
	"fmt"
)

// ItemCarrito representa un producto y la cantidad seleccionada.
type ItemCarrito struct {
	producto *Producto
	cantidad int
}

// Carrito almacena temporalmente los productos seleccionados por el cliente.
type Carrito struct {
	items []ItemCarrito
}

// NewCarrito crea un carrito vacío.
func NewCarrito() *Carrito {
	return &Carrito{
		items: []ItemCarrito{},
	}
}

// AgregarProducto valida stock, reduce inventario y agrega el producto al carrito.
func (c *Carrito) AgregarProducto(producto *Producto, cantidad int) error {
	if producto == nil {
		return errors.New("el producto no existe")
	}

	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	if !producto.TieneStock(cantidad) {
		return errors.New("no existe stock suficiente para ese producto")
	}

	// Se reduce el stock una vez que el producto es aceptado en el carrito.
	err := producto.ReducirStock(cantidad)
	if err != nil {
		return err
	}

	// Si el producto ya existe en el carrito, se suma la cantidad.
	for i := range c.items {
		if c.items[i].producto.ID() == producto.ID() {
			c.items[i].cantidad += cantidad
			return nil
		}
	}

	// Si el producto no existe en el carrito, se agrega como nuevo ítem.
	c.items = append(c.items, ItemCarrito{
		producto: producto,
		cantidad: cantidad,
	})

	return nil
}

// CalcularTotal suma los subtotales de todos los productos agregados.
func (c *Carrito) CalcularTotal() float64 {
	total := 0.0

	for _, item := range c.items {
		total += item.producto.Precio() * float64(item.cantidad)
	}

	return total
}

// MostrarCarrito imprime los productos seleccionados y el total de compra.
func (c *Carrito) MostrarCarrito() {
	if c.EstaVacio() {
		fmt.Println("El carrito está vacío.")
		return
	}

	fmt.Println("\n--- CARRITO DE COMPRAS ---")

	for _, item := range c.items {
		subtotal := item.producto.Precio() * float64(item.cantidad)
		fmt.Printf("Producto: %s | Cantidad: %d | Precio: $%.2f | Subtotal: $%.2f\n",
			item.producto.Nombre(), item.cantidad, item.producto.Precio(), subtotal)
	}

	fmt.Printf("Total de compra: $%.2f\n", c.CalcularTotal())
}

// EstaVacio verifica si el carrito no contiene productos.
func (c *Carrito) EstaVacio() bool {
	return len(c.items) == 0
}
