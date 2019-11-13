# Builder (Constuctor)

Permite separar la construcción de un objeto (complejo o no) de su representación con el fin de que el mismo proceso de creación pueda crear diferentes representaciones.

- Product: Objeto principal construido. Representa el objeto bajo construcción.
- Builder: define la interface con los métodos que deberán cumplir los constructores.
- Concrete Builder: Implementa la interface builder para entregar el objeto concreto.
- Director: Construye el objeto utilizando la interface.

Fuente código: [http://blog.ralch.com/tutorial/design-patterns/golang-builder/](http://blog.ralch.com/tutorial/design-patterns/golang-builder/)

Video de ejemplo:

[![Builder (constructor)](http://img.youtube.com/vi/j6Uvm3kzEjU/0.jpg)](https://www.youtube.com/watch?v=j6Uvm3kzEjU)