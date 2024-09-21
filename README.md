# Aplicación de Lista de Tareas

Esta es una aplicación de Lista de Tareas con una API backend en Go y un frontend en Vue.js. Permite a los usuarios crear, leer, actualizar y eliminar tareas, con una interfaz de usuario intuitiva y un modo oscuro.

## Características

- API RESTful en Go
- Frontend en Vue.js con Tailwind CSS
- Modo oscuro
- Operaciones CRUD completas para las tareas
- Retroalimentación visual para las acciones del usuario
- Dockerizado para fácil despliegue

## Estructura del Proyecto

- `backend/`: Contiene la API de Go
    - `cmd/api/`: Punto de entrada de la aplicación
    - `internal/`: Lógica interna de la aplicación
        - `handlers/`: Manejadores de las rutas HTTP
        - `models/`: Definiciones de los modelos de datos
        - `database/`: Configuración y operaciones de la base de datos
        - `utils/`: Utilidades generales
- `frontend/`: Contiene el frontend de Vue.js
    - `src/`: Código fuente de Vue.js
        - `components/`: Componentes de Vue
        - `assets/`: Recursos estáticos
- `docker-compose.yml`: Configuración de Docker Compose

## Requisitos Previos

- Docker y Docker Compose (para ejecutar con Docker)
- Go 1.16+ (para desarrollo local del backend)
- Node.js 14+ (para desarrollo local del frontend)
- SQLite (incluido en Go, no se requiere instalación adicional)

## Ejecutar la Aplicación con Docker

1. Clona este repositorio:
   ```
   git clone https://github.com/FernandoJCa/fullstack-project.git
   cd todo-app
   ```

2. Construye y ejecuta los contenedores:
   ```
   docker-compose up --build
   ```

3. Accede a la aplicación en `http://localhost:4243`.

## Desarrollo Local

### Backend (Go)

1. Navega al directorio `backend`:
   ```
   cd backend
   ```

2. Instala las dependencias:
   ```
   go mod download
   ```

3. Ejecuta el servidor:
   ```
   go run cmd/api/main.go
   ```

El backend estará disponible en `http://localhost:8080`.

### Frontend (Vue.js)

1. Navega al directorio `frontend`:
   ```
   cd frontend
   ```

2. Instala las dependencias:
   ```
   npm install
   ```

3. Ejecuta el servidor de desarrollo:
   ```
   npm run dev
   ```

El frontend estará disponible en `http://localhost:4243`.

## API Endpoints

- `GET /todos`: Obtiene todas las tareas
- `GET /todos/{id}`: Obtiene una tarea específica
- `POST /todos`: Crea una nueva tarea
- `PUT /todos/{id}`: Actualiza una tarea existente
- `DELETE /todos/{id}`: Elimina una tarea

Para más detalles, consulta la documentación de Swagger en `http://localhost:8080/swagger/index.html` cuando el backend esté en ejecución.