package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lector = bufio.NewReader(os.Stdin)

func main() {
	productos := cargarProductosIniciales()
	carrito := NewCarrito()

	fmt.Println("===================================")
	fmt.Println("   SISTEMA COMERCIO AL PASO")
	fmt.Println("===================================")

	usuario := registrarUsuario()

	opcion := 0

	for opcion != 6 {
		mostrarMenu()
		opcion = leerEntero("Seleccione una opción: ")

		switch opcion {
		case 1:
			mostrarProductos(productos)

		case 2:
			agregarProductoAlCarrito(productos, carrito)

		case 3:
			carrito.MostrarCarrito()

		case 4:
			generarYProcesarPedido(usuario, carrito)

		case 5:
			mostrarProductos(productos)

		case 6:
			fmt.Println("Gracias por utilizar Comercio al Paso.")

		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}

// registrarUsuario solicita los datos básicos del cliente.
func registrarUsuario() *Usuario {
	for {
		fmt.Println("\n--- REGISTRO DE CLIENTE ---")

		nombre := leerTexto("Ingrese su nombre: ")
		correo := leerTexto("Ingrese su correo: ")
		telefono := leerTexto("Ingrese su teléfono: ")

		usuario, err := NewUsuario(1, nombre, correo, telefono)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Por favor, ingrese nuevamente los datos.")
			continue
		}

		fmt.Println("Cliente registrado correctamente.")
		return usuario
	}
}

// mostrarMenu presenta las opciones principales del sistema.
func mostrarMenu() {
	fmt.Println("\n========== MENÚ PRINCIPAL ==========")
	fmt.Println("1. Mostrar productos disponibles")
	fmt.Println("2. Agregar producto al carrito")
	fmt.Println("3. Ver carrito de compras")
	fmt.Println("4. Generar pedido y procesar pago")
	fmt.Println("5. Consultar stock actualizado")
	fmt.Println("6. Salir")
	fmt.Println("====================================")
}

// agregarProductoAlCarrito permite seleccionar producto y cantidad.
func agregarProductoAlCarrito(productos []*Producto, carrito *Carrito) {
	mostrarProductos(productos)

	idProducto := leerEntero("\nIngrese el ID del producto que desea comprar: ")
	producto := buscarProductoPorID(productos, idProducto)

	if producto == nil {
		fmt.Println("Error: no existe un producto con ese ID.")
		return
	}

	cantidad := leerEntero("Ingrese la cantidad: ")

	err := carrito.AgregarProducto(producto, cantidad)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Producto agregado correctamente al carrito.")
}

// generarYProcesarPedido crea el pedido y permite seleccionar forma de pago.
func generarYProcesarPedido(usuario *Usuario, carrito *Carrito) {
	pedido, err := NewPedido(1, usuario, carrito)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pedido.MostrarResumen()

	fmt.Println("\n--- MÉTODOS DE PAGO ---")
	fmt.Println("1. Tarjeta")
	fmt.Println("2. Transferencia bancaria")

	opcionPago := leerEntero("Seleccione el método de pago: ")

	var metodo MetodoPago

	switch opcionPago {
	case 1:
		numeroTarjeta := leerTexto("Ingrese el número de tarjeta: ")
		titular := leerTexto("Ingrese el nombre del titular: ")
		metodo = NewPagoTarjeta(numeroTarjeta, titular)

	case 2:
		banco := leerTexto("Ingrese el nombre del banco: ")
		cuenta := leerTexto("Ingrese el número de cuenta: ")
		metodo = NewPagoTransferencia(banco, cuenta)

	default:
		fmt.Println("Método de pago no válido.")
		return
	}

	err = metodo.ProcesarPago(pedido.Total())
	if err != nil {
		fmt.Println("Error al procesar el pago:", err)
		return
	}

	pedido.MarcarPagado()

	fmt.Println("\nCompra realizada correctamente.")
	fmt.Println("Método utilizado:", metodo.NombreMetodo())
	pedido.MostrarResumen()
}

// leerTexto permite leer cadenas de texto completas desde consola.
func leerTexto(mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

// leerEntero permite leer y validar números enteros.
func leerEntero(mensaje string) int {
	for {
		fmt.Print(mensaje)
		texto, _ := lector.ReadString('\n')
		texto = strings.TrimSpace(texto)

		numero, err := strconv.Atoi(texto)
		if err != nil {
			fmt.Println("Debe ingresar un número válido.")
			continue
		}

		return numero
	}
}
