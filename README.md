# market-backend

## Descripcion

Aplicacion back-end para la [tienda online](https://github.com/asloth/market-frontend)

## Live version
El backend puede ser utilizado en este [link](https://market-backend-production.up.railway.app/)

## Tecnologias
- Lenguaje: Go
- Framework: Fiber
- ORM: Gorm

## Descripcion del proyecto
- database/main.go: Se declara un objeto que va a contener la conexion a la base de datos
- src: Se encuentra el codigo de negocio
  - category: modeulo que se encarga de las categorias
    - handler: Se declaran los manejadores que reciben las solicitudes y los parametros
    - model: Se declara el modelo que representa a la tabla en DB y sus metodos respectivos (listar, etc.)
  - product: modeulo que se encarga de los productos
    - handler: Se declaran los manejadores que reciben las solicitudes y los parametros
    - model: Se declara el modelo que representa a la tabla  en DB y sus metodos respectivos (listar todos, listar por nombre, etc.)
    
 ## Utilizar el API Rest
  ~~~ 
 - /products/category? : GET, endpoint que lista los productos con el parametro OPCIONAL category para listar por categoria
 
fetch(uri + "/products")
        .then(response => response.json())
        .then(data => console.log(data)) 
     
        
 - /products/search : POST, endpoint que recibe el objeto {name, idcategory} y busca por nombre y/o categoria.
 
  `fetch(uri + "/products/search", {
    method: 'POST',
    body: {
      name: "ron",
      idcategory: 2,
    }
    
  })
    .then(response => response.json())
    .then(data => console.log(data))`
        
        
 - /categories : GET, endpoint que lista las categorias y sus ids.
 
  `fetch(uri + "categories")
        .then(response => response.json())
        .then(data => console.log(data))`
