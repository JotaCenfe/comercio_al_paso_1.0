# comercio_al_paso_1.0

## Descripción del proyecto

**Comercio al Paso** es un sistema básico de gestión de e-commerce desarrollado en el lenguaje de programación **Go**.

El sistema simula el funcionamiento principal de una tienda virtual mediante consola. Permite registrar un cliente, visualizar productos disponibles, agregar productos al carrito de compras, generar un pedido y simular un proceso de pago.

Este proyecto corresponde al **Aprendizaje Autónomo 2** de la materia **Programación Orientada a Objetos**. Se desarrolla como continuación del Aprendizaje Autónomo 1, en el cual se realizó la planificación del sistema, incluyendo su alcance, módulos, funcionalidades y diagrama UML.

En esta segunda etapa se implementa el sistema en código, aplicando estructuras, métodos, encapsulación, interfaces, manejo de errores y organización modular.

---

## Objetivo

Implementar un sistema básico de comercio electrónico en consola utilizando el lenguaje Go, aplicando los principios de programación orientada a objetos revisados en la asignatura.

---

## Funcionalidades principales

El sistema permite realizar las siguientes acciones:

- Registrar los datos básicos de un cliente.
- Mostrar una lista de productos disponibles.
- Buscar productos mediante su identificador.
- Agregar productos al carrito de compras.
- Validar disponibilidad de stock.
- Calcular subtotales y total de compra.
- Generar un pedido con estado inicial pendiente de pago.
- Seleccionar un método de pago.
- Simular pago con tarjeta o transferencia bancaria.
- Actualizar el estado del pedido a pagado.
- Consultar el stock actualizado de los productos.

---

## Estructura del proyecto

```text
comercio_al_paso/
│
├── go.mod
├── main.go
├── producto.go
├── usuario.go
├── carrito.go
├── pedido.go
├── pago.go
└── auxiliares.go
