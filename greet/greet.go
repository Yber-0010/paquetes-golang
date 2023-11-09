package greet

/* variables a nivel de paquete solo se declaran con var y no operador corto := */
var emoji = "🍔"

// GreetEnglish retorna saludo en ingles.
func English() string {
	return "Hi" + emoji
}

// Italian retorna saludo en Italiano.
func Italian() string {
	return "Ciao" + emoji
}

/* se puede compartir todo entre paquetes variables constantes funciones */
/* que exportar y que no exportar */
/* si empieza con mayuscula se exporta */
/* si empieza con minuscula no se exporta */
