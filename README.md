# Go Challenge

Repositorio que contiene el código fuente del *Go Challenge* de agosto 2020.

## Sobre el challenge

Los requerimientos y objetivos del *challenge* se encuentran documentados en el siguiente documento: `docs/Challenge.pdf`.

## Deploy

El proyecto utiliza [Docker](https://www.docker.com/) y [Docker Compose](https://docs.docker.com/compose/) para orquestar su deploy.

### Construir el stack de servicios

Copiar el archivo `deployments/.env-example` a `deployments/.env` y editar la configuración de entorno *(los valores por defecto funcionan con el stack de docker-compose)*.

```shell script
make docker-build
```

### Iniciar el stack de servicios

```shell script
make docker-up
```

### Remover el stack de servicios

```shell script
make docker-down
make docker-clear
```

## Herramientas de desarrollo

### Live Reload

El proyecto viene configurado con [Air](https://github.com/cosmtrek/air) como live reload para el desarrollo de forma local.

Ejemplo de uso:

```shell script
make go-live
```

> Este paquete no forma parte del proyecto por lo que su instalación debe realizarse de forma global.

## Arquitectura

La arquitectura se basa en patrones de `Domain Driven Design` y `Hexagonal Architecture`.

### Estructura del proyecto

La estructura raíz del proyecto se basa en el siguiente estándar: <https://github.com/golang-standards/project-layout>.
La elección del mismo se basa en el hecho de que los patrones allí documentados son ampliamente reconocidos y adoptados por una gran cantidad de proyectos *Go*. 

#### Directorios

##### `/api`

Contiene las definiciones de la API para ser utilizadas con [Swagger](https://swagger.io/) y [Postman](https://www.postman.com/).

##### `/build`

Contiene los binarios generados.

> Esta carpeta no forma parte del proyecto y se genera de forma dinámica.

##### `/cmd`

Contiene los códigos fuente de los diferentes puntos de entrada a binarios. Este proyecto solo maneja un punto de entrada a una API.

##### `/deployments`

Contiene la configuración de los contenedores de [Docker](https://www.docker.com/) y la configuración de las variables de entorno.

##### `/docs`

Contiene los documentos de especificaciones del proyecto.

##### `/internal`

Contiene los códigos fuente del proyecto.

##### `/scripts`

Contiene los scripts de migraciones de base de datos.

##### `/storage`

Contiene los logs de la aplicación y los datos de volúmenes de contenedores Docker.

##### `/tmp`

Contiene archivos temporales de herramientas o pasos de compilación.

> Esta carpeta no forma parte del proyecto y se genera de forma dinámica.

### Algoritmo de búsqueda de sucursal cercana

Si bien existen formas simples de implementar esta característica - por ejemplo utilizar el tipo de dato `point` de MySQL 8 y la función espacial `ST_distance_sphere` - 
se opta por delegar el cálculo al dominio de la aplicación por los siguientes motivos:

* unificar la lógica de dominio en un solo lado dado que:
  * delegar la lógica a la base de datos implica que ante un cambio de infraestructura se debe re escribir lógica de dominio.
  * impide trabajar con diferentes infraestructuras por ambiente *(in memory en tests por ejemplo)*.
* se contempla que la geografía impacta en las distancias, por ejemplo:
  * obstáculos como montañas, rios, etc.
* al no saber si existen caminos o si la geografía lo permite se puede optar por utilizar servicios de GIS o APIs externas de cálculo de carretera.

Para este proyecto se utilizá [GraphHopper Directions API](https://www.graphhopper.com/) como servicio de cálculo de rutas.

> Nota: al momento de evaluar el proyecto contemplar que la *API_KEY* corresponde a un plan free con un límite de 500 solicitudes por día.

### Paquetes de terceros

A continuación se listan los paquetes de terceros utilizados y la motivación de su elección:

* [Chi](https://github.com/go-chi/chi): para el routing de la API.
  * paquete ampliamente reconocido que goza de una gran popularidad.
  * cuenta con muy pocos issues abiertos.
  * es compatible con la interface `net/http` lo que facilita a futuro su reemplazo por otro paquete compatible.
  * soporte para middleware integrado respetando el estándar `func(http.Handler) http.Handler` lo que facilita a futuro reutilizar los middlewares (Ej: con [Gorila Mux](https://github.com/gorilla/mux))
  * buen equilibrio entre performance y uso de recursos respecto de otros routers.
* [MySQL](https://github.com/go-sql-driver/mysql): para trabajar con el motor de base de datos MySQL.
  * paquete ampliamente reconocido que goza de una gran popularidad.
  * cuenta con muy pocos issues abiertos.
  * implementa la API de la interface `database/sql` de la librería estándar de *Go*.
* [Migrate](https://github.com/golang-migrate/migrate): para la ejecución de migraciones de base de datos.
  * paquete ampliamente reconocido que goza de una gran popularidad.
  * cuenta con muy pocos issues abiertos.
  * permite trabajar con múltiples motores de base de datos.
* [Zap](https://github.com/uber-go/zap): para el registro de los logs del webserver de la API.
  * paquete ampliamente reconocido que goza de una gran popularidad.
  * cuenta con muy pocos issues abiertos.
  * es desarrollado y mantenido por Uber.
  * los benchmarks lo posicionan como uno de los paquetes más rápidos y con menor asignación de memoria de su tipo.
  * aporta características de logging superadoras respecto del paquete `log` de la librería estándar de *Go*.

> Nota: con el fin de no desvirtuar el propósito del *challenge* se trató de buscar un balance intermedio en la cantidad de paquetes de terceros utilizados.

## Endpoints de la API

### Buscar una Sucursal

Este servicio RESTful permite ver el detalle de una Sucursal.

#### Request

* *Method*: GET
* *URL*: /api/v1/offices/**{ID}**
  * La URL contiene el identificador de búsqueda

#### Response

* Content-Type: application/json

```json
{"id": 1, "address": "El Salvador 5700 - CABA", "latitude": -34.582375, "longitude": -58.436463}
```

#### Ejemplo

```shell script
curl --request GET 'http://127.0.0.1:8000/api/v1/offices/1'
```

### Crear una Sucursal

Este servicio RESTful permite crear una nueva Sucursal.

#### Request

* *Method*: POST
* *URL*: /api/v1/offices
* *Body* json

```json
{"address": "El Salvador 5700 - CABA", "latitude": -34.582375, "longitude": -58.436463}
```

#### Response

* Content-Type: application/json

```json
{"id": 1, "address": "El Salvador 5700 - CABA", "latitude": -34.582375, "longitude": -58.436463}
```

#### Ejemplo

```shell script
curl --request POST 'http://127.0.0.1:8000/api/v1/offices' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address": "El Salvador 5700 - CABA",
    "latitude": -34.582375,
    "longitude": -58.436463
}'
```

### Buscar Sucursal cercana

Este servicio RPC permite buscar la Sucursal más cercana.

#### Request

* *Method*: GET
* *URL*: rpc/v1/searches/officeByProximity?latitude=**{LATITUD}**&longitude=**{LONGITUD}**
  * La URL requiere el parámetro *latitude* para indicar la búsqueda por latitud
  * La URL requiere el parámetro *longitude* para indicar la búsqueda por longitud

#### Response

* Content-Type: application/json

```json
{"id": 1, "address": "El Salvador 5700 - CABA", "latitude": -34.582375, "longitude": -58.436463}
```

#### Ejemplo

```shell script
curl --request GET 'http://127.0.0.1:8000/rpc/v1/searches/officeByProximity?latitude=-34.582932&longitude=-58.437875'
```

### Check de Base de Datos

Este servicio RPC permite hacer un ping a la base de datos para validar su estado.

#### Request

* *Method*: GET
* *URL*: rpc/v1/systems/health-checks/database/ping

#### Response

* Content-Type: text/plain

```text
PONG
```

#### Ejemplo

```shell script
curl --request GET 'http://127.0.0.1:8000/rpc/v1/systems/health-checks/database/ping'
```
