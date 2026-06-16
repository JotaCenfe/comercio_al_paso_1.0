package main

import "fmt"

// mostrarProductos imprime todos los productos disponibles.
func mostrarProductos(productos []*Producto) {
	fmt.Println("\n--- PRODUCTOS DISPONIBLES ---")

	for _, producto := range productos {
		producto.MostrarInformacion()
	}
}

// buscarProductoPorID localiza un producto según el ID ingresado.
func buscarProductoPorID(productos []*Producto, id int) *Producto {
	for _, producto := range productos {
		if producto.ID() == id {
			return producto
		}
	}

	return nil
}

// cargarProductosIniciales crea una lista base de productos para la tienda.
func cargarProductosIniciales() []*Producto {
	productos := []*Producto{}

	p1, _ := NewProducto(1, "Audífonos Bluetooth", "Audífonos inalámbricos recargables", 18.50, 10)
	p2, _ := NewProducto(2, "Mouse Gamer", "Mouse óptico con luces RGB", 15.75, 8)
	p3, _ := NewProducto(3, "Teclado Mecánico", "Teclado mecánico compacto", 35.00, 5)
	p4, _ := NewProducto(4, "Memoria USB 64GB", "Unidad flash USB 3.0", 9.99, 15)
	p5, _ := NewProducto(5, "Cargador Tipo C", "Cargador rápido para celular", 12.25, 12)

	productos = append(productos, p1, p2, p3, p4, p5)

	return productos
}
