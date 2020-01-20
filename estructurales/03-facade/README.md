# Facade (Fachada)

Permite ocultar la complejidad de un sistema proveyendo una interfaz sencilla para el cliente.

Unifica las diferentes interfaces del sistema en una sola, con esto la perspectiva del cliente es el uso de un sistema de fácil acceso.

## Ejemplo

Haremos un proceso para registrar los comentarios de una persona en un post de un blog.

Para realizar esto, el sistema debe ejecutar varios procesos:

- Validar que el usuario esté logueado y activo.
- Validar que el usuario tiene permiso de comentar.
- Registrar el comentario.
- Notificar al creador del post que llegó un comentario.

Son demasiados procesos para que el cliente los haga uno a uno, e incluso puede ser que se salte algunos. Por esta razón la fachada se encargará de enmascarar todo ese proceso en uno solo:

- Registrar comentario.

## Video

