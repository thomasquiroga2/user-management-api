# API REST - Gestión de Usuarios

## Descripción
API REST desarrollada en Golang para gestionar usuarios con operaciones CRUD básicas.

## Requisitos
- Go 1.20+
- Librería `gorilla/mux`

## Ejecución
1. Clonar el repositorio.
2. Instalar dependencias: `go mod tidy`
3. Ejecutar la API: `go run main.go`

## Endpoints
- **POST /users**: Crear un usuario.
- **GET /users**: Listar todos los usuarios.
- **GET /users/{id}**: Obtener un usuario por ID.
- **PUT /users/{id}**: Actualizar un usuario por ID.
- **DELETE /users/{id}**: Eliminar un usuario por ID.
