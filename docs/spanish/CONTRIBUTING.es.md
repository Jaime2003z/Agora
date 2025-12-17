# Contribuir a Agora

Antes que nada, gracias por tu interés en contribuir a **Agora**.

Agora es un motor de coordinación abierto y sin permisos diseñado para mejorar la forma en que los humanos colaboran a escala local, regional y global.  
Esto es **software de infraestructura**, no un producto, no una startup y no una campaña política.

Se dan la bienvenida a contribuciones de desarrolladores, investigadores, auditores, diseñadores y pensadores de sistemas.

---

## Filosofía del proyecto

Antes de contribuir, por favor entiende los principios fundamentales de Agora:

- **Mérito sobre autoridad**  
  La influencia se gana con contribución y ejecución, no con estatus.

- **Abierto por defecto**  
  El código, las decisiones y los fallos son públicos y auditables.

- **Mínima confianza, máxima verificación**  
  Los sistemas deben ser verificables, deterministas y resistentes a la captura.

- **Pensamiento protocolo-primero**  
  Agora es un protocolo, no una plataforma. Evitar la centralización por diseño.

- **Resiliencia a largo plazo sobre crecimiento a corto plazo**  
  Priorizamos corrección, seguridad y sostenibilidad.

Si estos principios chocan con tus objetivos, este puede no ser el proyecto adecuado para ti —y está bien.

---

## Formas de contribuir

Puedes contribuir de muchas maneras:

### 1. Código
- Protocolo principal (Go)
- Primitivas criptográficas (Rust)
- Redes (libp2p)
- Almacenamiento y gestión de estado
- Lógica de gobernanza y ejecución
- Herramientas, CLI, SDKs

### 2. Investigación y diseño
- Mecanismos de gobernanza
- Sistemas de reputación
- Modelos económicos
- Vectores de ataque y modelado de amenazas
- Análisis de escalabilidad y rendimiento

### 3. Auditoría y revisión
- Revisión de código
- Análisis de seguridad
- Verificación formal (cuando corresponda)
- Revisión de documentación

### 4. Documentación
- Mejorar claridad y precisión
- Escribir explicaciones técnicas
- Diagramas y especificaciones del protocolo

### 5. Pruebas
- Tests unitarios
- Tests de integración
- Escenarios adversarios
- Participación en testnet

---

## Estructura de los repositorios

Estructura general:

core/ → Implementación del protocolo (AGPLv3)  
crypto/ → Primitivas criptográficas (Rust)  
clients/ → Aplicaciones para usuarios finales  
sdk/ → SDKs para terceros  
docs/ → Especificaciones y documentos de diseño  
testnet/ → Pruebas de red y escenarios

Si no estás seguro de dónde encaja tu contribución, abre una discusión o un issue.

---

## Flujo de contribución

1. Haz fork del repositorio  
2. Crea una rama desde `main`

feature/<descripción-corta>  
fix/<descripción-corta>  
research/<descripción-corta>

3. Haz commits pequeños y enfocados  
4. Escribe mensajes de commit claros  
5. Abre un Pull Request  
6. Participa en el proceso de revisión

---

## Guías de código

### General
- Prefiere claridad antes que ingenio  
- El comportamiento determinista es obligatorio  
- Evita efectos secundarios ocultos  
- Explícito > implícito

### Go
- Sigue las convenciones estándar de Go  
- Mantén los paquetes pequeños y componibles  
- Evita estado global cuando sea posible  
- Documenta las interfaces públicas

### Rust
- Prioriza seguridad y explicitud  
- Nada de código unsafe sin justificación  
- El código criptográfico debe ser revisado con mucho cuidado

### Pruebas
- Los tests no son opcionales  
- El código sensible a seguridad **debe** incluir tests  
- Las regresiones deben documentarse

---

## Gobernanza de las contribuciones

Agora sigue un **modelo de gobernanza basado en mérito**:

- No hay mantenedores permanentes solo por estatus  
- La influencia crece con contribuciones consistentes y de alta calidad  
- Las contribuciones de mala calidad o maliciosas reducen la confianza

Los cambios importantes en el protocolo requieren:  
- Motivación clara  
- Análisis de compatibilidad hacia atrás  
- Consideraciones de seguridad  
- Revisión de la comunidad

---

## Seguridad

Si descubres una vulnerabilidad de seguridad:

- **No abras un issue público**  
- Sigue el proceso descrito en `SECURITY.md`

La divulgación responsable es crítica.

---

## Licencias

- Protocolo principal: **AGPLv3**  
- Clientes y SDKs: licencias permisivas (MIT / Apache)

Al contribuir, aceptas que tu trabajo se licencie bajo la licencia del proyecto.

---

## Código de conducta

Esperamos:  
- Discusión técnica respetuosa  
- Criticar ideas, no personas  
- Nada de acoso ni batallas ideológicas

Lee `/docs/spanish/CODE_OF_CONDUCT.es.md` para más detalles.

---

## Notas finales

Agora es un esfuerzo a largo plazo.

Este proyecto no está optimizado para:  
- victorias rápidas  
- ciclos de hype  
- valor especulativo

Está optimizado para:  
- corrección  
- legitimidad  
- coordinación a escala

Si esto te emociona —bienvenido.