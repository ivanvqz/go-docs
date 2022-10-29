# Apuntes de Go
<!-- Tabla de contenidos -->


## Declaración de variables
Las variables se declaran de la siguiente manera:
```go
var nombre_variable tipo_variable = valor_variable
```
Se especifica que se declara una variable, el nombre de la variable. También el tipo de la variable, al final se le declara el valor de la variable.
Puedes declarar múltiples variables de la siguiente manera:
```go
var (
    nombre_variable tipo_variable = valor_variable
    nombre_variable tipo_variable = valor_variable
    nombre_variable tipo_variable = valor_variable
)
```
La manera más corta y sencila de declarar variables es la siguiente:
```go
nombre_variable := valor_variable
```
El valor es declarada por el mismo go.

## Constantes
```go
const nombre_constante = valor_constante
```
No puede acortarse más, ya que el compilador infiere que es una variable.

## Comentarios
// - para comentarios de una línea
/* */ - para comentarios de varias líneas

# Reglas de formato para documentación:
- Comentarios de una línea deben comenzar con el carácter `//` seguido de un espacio en blanco.
- Los comentarios multilínea se usan para código que no se va a utilizar
Algo práctico de go es que puedes visualiar la documentación con el comando 
```go
go doc --All
```
```

# Slices
Los slices son apuntadores de arrays, no poseen datos, solo apuntan a un array.
```go
    set := [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	animales := set[0:3]
	fmt.Println(animales)   // [a b c]
```
a la hora de seleccionar un rango en un slice, este exluye el último elemento.
Características:
- No tienen un tamaño fijo
- Pueden tener un tamaño máximo y capacidad máxima

## map
son estructuras de clave valor, como los diccionarios de python.
```go
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3
    fmt.Println(m)  // map[a:1 b:2 c:3]
```
## Estrucutras
Permiten almacenar colecciones de campos, simmilares a las clases de otros lenguajes.
```go
type Persona struct {
    nombre string
    edad int
}
```
# Manejo de errores
La biblioteca estándar de Go ofrece dos métodos para crear errores: errors.New del paquete errors (se muestra en el vídeo) y fmt.Errorf. Ambos métodos funcionan igual bajo el capó, devuelven un valor de tipo error y al msimo tiempo especifican un mensaje de error personalizado.

Sin embargo, es preferible "la función fmt.Errorf, ya que formatea el mensaje de error usando fmt.Sprintf, permitiendo generar errores descriptivos al agregar sucesivamente información de contexto adicional al mensaje de error original" (Donovan and Kernighan, 2015).

El mensaje de error ("no name provided") se escribe en minúscula, porque el usuario que posteriormente maneje de forma correcta el error lo debe presentar en un contexto explicativo (línea 16-19).

# paquetes
Es una carpeta que tiene una colección de archivos que provee una funcionalidad.
Cada archivo del paquete debe contener nombre del paquete donde pertenece.
Identificadores a nivel paquete
Son identificadores (variables, constantes, estructuras, funciones) que se pueden compartir entre los archivos que pertenecen al mismo paquete.
No se puede declarar y asignar con el operador de variable corta.
Identificadores exportables
Un identificador puede ser exportado, es decir, puede usarse en otros paquetes. Si la primera letra del identificador es:
minúscula, el identificador es no exportable.
mayúscula, el identificador es exportable.
Para importar un identificador, se debe usar como prefijo el nombre del paquete, es decir: <nombre_paquete>.identificador.