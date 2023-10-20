# Saludos en Go

Este paquete proporciona una forma simple de retornar saludos

## Instalacion
Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/KronosSolucionesTIC/greetings.git
```

## Uso
Aqui tienes un ejemplo de como usar el paquete:

```go
package main

import (
    "fmt"
    "github.com/KronosSolucionesTIC/greetings"
)

func main() {
    message, err := greetings.Hello("Diego")
    if err != nil {
        fmt.Println("Ocurrio un error:", err)
        return
    }

    fmt.Println(message)
}
```

Este ejemplo importa el paquete github.com/KronosSolucionesTIC/greetings y llama a la funcion Hello saludo personalizado. Si occurre un error, se imprime mensaje de error.