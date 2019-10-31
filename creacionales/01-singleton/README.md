# Singleton

Este patrón de diseño nos permite usar una única instancia para toda la aplicación.

Esto es útil para poder compartir información única en toda nuestra app.

En otros lenguajes de programación utilizamos la palabra reservada `static`. Pero en Go no tenemos esa palabra ya que es una arquitectura diferente. Puedes llamar una función de un paquete sin instanciar alguna estructura. Ejemplo:

```go
package saludo

func Saludar() string {
    return "Hola mundo"
}
```
```go
package main

import (
    "fmt"
    "saludo"
)

func main() {
    fmt.Println(saludo.Saludar())
}
```

En el ejemplo anterior vemos que el paquete saludo no tiene estructuras y que tiene una función pública llamada `Saludar()` la cual se puede llamar directamente del paquete.

Se puede pensar que con eso podemos crear una estructura que se llame desde cualquier lado. El problema viene cuando se hacen procesos concurrentes.

En los siguientes video explico cómo se puede solucionar el problema:

Primera parte:

[![Singleton primera parte](http://img.youtube.com/vi/vBdHnWtoAhg/0.jpg)](http://www.youtube.com/watch?v=vBdHnWtoAhg "Singleton - primera parte")

Segunda parte:

[![Singleton segunda parte](http://img.youtube.com/vi/uBoxSrMtOKA/0.jpg)](http://www.youtube.com/watch?v=uBoxSrMtOKA "Singleton - segunda parte")