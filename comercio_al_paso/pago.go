package main

import (
	"errors"
	"fmt"
)

// MetodoPago es una interfaz que permite aplicar polimorfismo.
// Cualquier estructura que implemente ProcesarPago y NombreMetodo
// puede ser utilizada como método de pago.
type MetodoPago interface {
	ProcesarPago(monto float64) error
	NombreMetodo() string
}

// PagoTarjeta representa un pago realizado con tarjeta.
type PagoTarjeta struct {
	numeroTarjeta string
	titular       string
}

// NewPagoTarjeta crea un método de pago con tarjeta.
func NewPagoTarjeta(numeroTarjeta string, titular string) *PagoTarjeta {
	return &PagoTarjeta{
		numeroTarjeta: numeroTarjeta,
		titular:       titular,
	}
}

// ProcesarPago simula el pago con tarjeta.
func (p *PagoTarjeta) ProcesarPago(monto float64) error {
	if monto <= 0 {
		return errors.New("el monto del pago debe ser mayor a cero")
	}

	if len(p.numeroTarjeta) < 4 {
		return errors.New("el número de tarjeta ingresado no es válido")
	}

	if p.titular == "" {
		return errors.New("el titular de la tarjeta no puede estar vacío")
	}

	fmt.Printf("Procesando pago con tarjeta por $%.2f...\n", monto)
	fmt.Println("Pago con tarjeta aprobado.")
	return nil
}

func (p *PagoTarjeta) NombreMetodo() string {
	return "Tarjeta"
}

// PagoTransferencia representa un pago realizado por transferencia bancaria.
type PagoTransferencia struct {
	banco  string
	cuenta string
}

// NewPagoTransferencia crea un método de pago por transferencia.
func NewPagoTransferencia(banco string, cuenta string) *PagoTransferencia {
	return &PagoTransferencia{
		banco:  banco,
		cuenta: cuenta,
	}
}

// ProcesarPago simula el pago mediante transferencia bancaria.
func (p *PagoTransferencia) ProcesarPago(monto float64) error {
	if monto <= 0 {
		return errors.New("el monto del pago debe ser mayor a cero")
	}

	if p.banco == "" {
		return errors.New("el banco no puede estar vacío")
	}

	if p.cuenta == "" {
		return errors.New("la cuenta no puede estar vacía")
	}

	fmt.Printf("Procesando transferencia bancaria por $%.2f...\n", monto)
	fmt.Println("Transferencia confirmada.")
	return nil
}

func (p *PagoTransferencia) NombreMetodo() string {
	return "Transferencia bancaria"
}
