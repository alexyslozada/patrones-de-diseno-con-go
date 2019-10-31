# Builder (Constuctor)

Este patrón de diseño nos permite encapsular la lógica de creación de los objetos. Se usa cuando el proceso de construcción de un objeto es complejo y se requiere entregar diferentes representaciones.

- Builder: define la interface con los métodos que deberán cumplir los constructores.
- Concrete Builder: Implementa la interface builder para entregar el objeto concreto.
- Director: Construye el objeto utilizando la interface.
- Product: Objeto principal construido. Representa el objeto bajo construcción.

Fuente código: [http://blog.ralch.com/tutorial/design-patterns/golang-builder/](http://blog.ralch.com/tutorial/design-patterns/golang-builder/)

Video de ejemplo:

