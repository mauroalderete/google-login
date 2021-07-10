# login

Recupera un token de acceso a partir del consentimiento aceptado de un usuario de google utilizando rayquen-google/golang/auth

# Objetivo

Construir una aplicación de consola que permite recuperar un token valido para acceder a los servicios api de google, como google spreadsheet

## Detalles

Se utiliza la libreria gitlab.com/rayquen-google/golang/auth

Se requiere previamente una credencial para aplicaciones clientes generada por google console. Estas credenciales deben generarse para cada aplicación de destino

# Instalar

Para instalar las librerias faltantes ejecutar:

```bash
go mod vendor
```

# Compilar

Para compilar correr:

```
go build
```

# Uso

> Para usarlo es necesario contar con un archivo de credencial de cliente otorgado por google. Se puede obtener gestionando las credenciales con Google Console.

Generar el token de acceso es sencillo. Sólo debe ejecutar el script por consola, por ejemplo:

```bash
./login -credential credential.json -token token.json -workdir ./
```

Donde

- *credential.json* es el nombre del archivo de la credencial de cliente otorgada por google
- *token* es el nombre del archivo con el que se guardara el token solicitado
- *workdir* es la carpeta donde se encuentra el archivo de credencial y donde se guardara el archivo con el token

Se mostrará en la salida del promtp un enlace que debera abrir en un navegador para otorgar los permisos necesarios con una cuenta de Google. Al final del proceso se le entregara un hash que debera copiar y pegar en el promtp. Tras presionar enter, se realizaran las comprobaciones necesarias y se generara un archivo token con el nombre dado en el argumento y dentro de la carpeta de trabajo.

Para mas información del uso de login ejecute

```bash
./login -help
```
