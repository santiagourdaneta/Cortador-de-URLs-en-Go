# 🔗 Cortador de URLs Mágico

### Descripción del Proyecto
Un simple y rápido servicio para acortar URLs, desarrollado con **Go** y el framework **Gin-Gonic**. Este proyecto demuestra cómo construir una aplicación web sin base de datos, manejando la lógica de forma eficiente en la memoria del servidor. Es ideal para aprender los fundamentos del desarrollo backend, manejo de concurrencia y despliegue de aplicaciones Go ligeras.

### 🚀 Características Clave
- **Sin Recarga de Página**: Interfaz de usuario dinámica con JavaScript y `fetch`.
- **Backend de Alto Rendimiento**: Construido con Go y Gin para una respuesta ultrarrápida.
- **Sin Base de Datos**: Todas las URLs se almacenan en la memoria del servidor, lo que lo hace muy rápido.
- **Fácil Despliegue**: Un único binario ejecutable para una implementación sencilla.
- **Diseño Responsivo**: Utiliza **Bulma CSS** para una UI/UX moderna y amigable.

### 🛡️ Seguridad y Robustez
El proyecto garantiza un funcionamiento seguro y confiable.
- **Rate Limiting**: Se implementó una protección contra ataques de fuerza bruta y DDoS para limitar el número de peticiones por usuario.
- **Pruebas Completas**: Se desarrollaron y ejecutaron pruebas unitarias, de integración y de estrés para validar el comportamiento del servidor bajo diferentes cargas y escenarios.
- **Código Depurado**: Se corrigieron errores de compilación y de lógica, asegurando que la base del proyecto es sólida y confiable para futuras expansiones.

### 🛠️ Tecnologías Usadas
- **Go**: El lenguaje principal del backend.
- **Gin-Gonic**: Framework web para gestionar las rutas y el servidor.
- **gin-rate-limit**: Middleware para implementar la limitación de peticiones.
- **Bulma CSS**: Framework CSS ligero para el diseño de la interfaz.

### ⚙️ Cómo Ejecutar el Proyecto
1. Clona el repositorio:
   ```bash
   git clone https://github.com/santiagourdaneta/Cortador-de-URLs-en-Go/
   cd Cortador-de-URLs-en-Go

2. Instala las dependencias:

go mod tidy

3. Ejecuta el servidor:

go run main.go 

4. Abre tu navegador y navega a http://localhost:8080.

🧪 Cómo Ejecutar las Pruebas

Pruebas unitarias y de integración:

go test

Pruebas de estrés (con k6):

Asegúrate de que el servidor esté corriendo en una terminal y luego ejecuta:

k6 run --vus 100 --duration 30s stress_test.js

📝 Licencia
Este proyecto está bajo la licencia MIT.
