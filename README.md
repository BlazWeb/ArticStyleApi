# ArticStyle REST API

Rest Api hecha en Go consumida por ArticStyle *(Microservices API)*.

## Toma de contacto con la aplicación
Primeramente para hacer que la API empiece a escuchar las peticiones tendremos que consumilar mediante un puerto, por defecto está el :8080, para asignarle uno diferentes deberás crear una variable de entorno con:
```js
PORT // Sirve para que la API empiece a escuchar mediante el puerto seleccionado, por defecto esta el :8080
```
Para hacer una conexión con la base de datos requerimos de variables de entorno para cumplir con su funcionamiento.
```js
DB_ADDR // Esta nos sirve para darle la IP de donde está consumida la base de datos
DB_TABLE // En está le daremos el nombre de la tabla de la base de datos
DB_USER // En esta pondremos el usuario de la base de datos, por defecto queda asignada como root
DB_PASS // En esta pondremos la contraseña para poder acceder a la base de datos
```

## ¿Cómo funciona?
Esto es una prueba de **README**
