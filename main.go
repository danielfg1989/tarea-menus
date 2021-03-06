package main

import "fmt"

func main() {

	var opc1 int
	var nav string
	var contr string
	nav = "si"

	for nav == "si" {

		fmt.Println("Menú principal\n1-Cambiar su contraseña\n2-Editar perfil\n3-Foro\n4/Finalizar programa")
		fmt.Scanf("%d", &opc1)

		switch opc1 {
		case 1:
			fmt.Println("Ha ingresado al menu para cambiar su contraseña\n ingrese una nueva contraseña")
			fmt.Scanf("%s", &contr)
			fmt.Println("Su nueva contraseña es", contr)
			fmt.Println("Desea seguir navegando, si o no?")
			fmt.Scanf("%s", &nav)
			if nav == "no" {
				break
			}

		case 2:
			var opc2 int
			fmt.Println("Menú de Perfil\n1-Cambiar foto de perfil\n2-Cambiar nombre de usuario\n3-Editar biografía\n4-Volver al menú principal")
			fmt.Println("Ingrese una opcion")
			fmt.Scanf("%d", &opc2)
			switch opc2 {
			case 1:
				fmt.Println("Foto de perfil cambiada")
				fmt.Println("Desea seguir navegando, si o no?")
				fmt.Scanf("%s", &nav)
				if nav == "no" {
					break
				}
			case 2:
				fmt.Println("Nombre de Usuario cambiado")
				fmt.Println("Desea seguir navegando, si o no?")
				fmt.Scanf("%s", &nav)
				if nav == "no" {
					break
				}
			case 3:
				fmt.Println("Biografia editada")
				fmt.Println("Desea seguir navegando, si o no?")
				fmt.Scanf("%s", &nav)
				if nav == "no" {
					break
				}
			case 4:
				fmt.Println("Volver al menu principal")
				break
			default:
				fmt.Println("No se pudo detectar la opcion ingresada")
			}
		case 3:
			var opc3 int
			fmt.Println("Menú de foro\n1-Publicar nueva entrada en el foro\n2-Editar una entrada del foro\n3-Eliminar una entrada del foro\n4-Volver al menú principal")
			fmt.Println("Ingrese una opcion")
			fmt.Scanf("%d", &opc3)
			switch opc3 {
			case 1:
				fmt.Println("Nueva entrada publicada")
				fmt.Println("Desea seguir navegando, si o no?")
				fmt.Scanf("%s", &nav)
				if nav == "no" {
					break
				}
			case 2:
				fmt.Println("Entrada editada con exito")
				fmt.Println("Desea seguir navegando, si o no?")
				fmt.Scanf("%s", &nav)
				if nav == "no" {
					break
				}
			case 3:
				fmt.Println("Entrada eliminada")
				fmt.Println("Desea seguir navegando, si o no?")
				fmt.Scanf("%s", &nav)
				if nav == "no" {
					break
				}
			case 4:
				fmt.Println("Volver al menu principal")
				break
			default:
				fmt.Println("No se pudo detectar la opcion ingresada")
			}
		case 4:
			nav = "no"
		default:
			fmt.Println("No se pudo detectar la opcion ingresada")
		}
	}
}
