# Microservicio de Eliminación de Libros

## Desarrollador
Ricardo Florez

## Descripción
Este microservicio es responsable de eliminar libros del sistema. Forma parte de una arquitectura de microservicios para la gestión de una biblioteca digital.

## Características
- Implementado en Go 1.21
- Arquitectura limpia (Clean Architecture)
- Endpoints RESTful
- Integración con MongoDB
- Despliegue en Kubernetes

## Estructura del Proyecto
```
.
├── config/         # Configuración de la base de datos
├── models/         # Modelos de datos
├── repositories/   # Capa de acceso a datos
├── services/      # Lógica de negocio
├── k8s/           # Configuración de Kubernetes
└── tests/         # Pruebas unitarias
```

## API Endpoint
- **DELETE** `/libros/{id}`
  - Puerto: 30085 (NodePort)
  - Elimina un libro específico por su ID

### Ejemplo de Respuesta Exitosa
```json
{
    "mensaje": "Libro eliminado exitosamente"
}
```

### Respuestas de Error
```json
{
    "error": "Libro no encontrado"
}
```

## Configuración Kubernetes
- Deployment con 3 réplicas
- Service tipo NodePort (30085)
- Conexión a MongoDB mediante Service Discovery

## Variables de Entorno
- MONGODB_URI: URI de conexión a MongoDB

## Despliegue
```bash
# Construir la imagen
docker build -t libro-delete:latest .

# Desplegar en Kubernetes
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

## Pruebas
```bash
# Ejecutar pruebas unitarias
go test ./...

# Probar el endpoint
curl -X DELETE http://localhost:30085/libros/5f7b5e1b9d3e2a1b4c7d8e9f
```

## Monitoreo
El servicio puede ser monitoreado mediante:
- Logs de Kubernetes
- Métricas de contenedor
- Estado del Service y Deployment

## Consideraciones de Seguridad
- Validación de ID antes de la eliminación
- Registro de operaciones de eliminación
- Verificación de permisos (si aplica)
- Soft delete vs Hard delete configurables

## Recuperación de Errores
- Manejo de errores de conexión a MongoDB
- Rollback en caso de fallos
- Notificación de errores críticos 