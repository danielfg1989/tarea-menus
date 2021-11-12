package main

import "fmt"

func main() {
	fmt.Println("Bienvenidos a la tienda de Duque")
	fmt.Println("Ingrese su nombre:")

	var nombre string
	fmt.Scanf("%s", &nombre)

	fmt.Println("Bienvenido", nombre)

	fmt.Println("Ingrese la opcion que desee comprar: \n 1-Para Bebidas \n 2-Para Alimentos \n 3-Para Medicamentos \n 4-Para otros elementos")

	var opc int
	fmt.Scanf("%d", &opc)

	switch opc {
	case 1:
		fmt.Println("Bienvenidos a las bebidas")
		fmt.Println("1-Cocacola\n2-Mazamorra\n3-Masato\n4-Aguapanela\n5-Finalizar")
		var beb int
		fmt.Scanf("%d", &beb)
		switch beb {
		case 1:
			fmt.Println("Usted ha comprado una cocacola")
		case 2:
			fmt.Println("Usted ha comprado una Mazamorra")
		case 3:
			fmt.Println("Usted ha comprado un Masato")
		case 4:
			fmt.Println("Usted ha comprado una Aguapanela")
		case 5:
			fmt.Println("El programa finalizara")
			break
		}
		break
	case 2:
		fmt.Println("Bievendios a los alimentos")
		break
	case 3:
		fmt.Println("Bienvenidos a la seccion de medicamentos")
		break

	}

}
