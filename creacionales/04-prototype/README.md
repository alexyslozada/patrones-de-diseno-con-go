# Prototype

Este patrón permite crear objetos a partir de instancias ya creadas.

La idea es evitar el costo de creación de los objetos, utilizando un objeto creado previamente y clonándolo.

Cuando hablamos del costo de creación de un objeto, hablamos de que para crearlo se pudo haber consultado varios elementos de una base de datos, consultando otros servicios, etc. Por lo que clonarlo es una idea genial para duplicar la información ya consultada.

## Actores

* Prototipo: Interface que deben implementar los objetos clonables.
* Prototipo concreto: Estructura o tipo que implementa el prototipo.
* Cliente: Crea el nuevo objeto solicitando al prototipo que se clone.

## Video
