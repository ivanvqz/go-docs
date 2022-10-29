# Punteros
## Definición
Un puntero es una variable que almacena la dirección de memoria de otra variable. Para definir un puntero se utiliza el operador `*` seguido del nombre de la variable. Por ejemplo:
    
    
    int *puntero;

Operador de dirección de memoria: ampersand (&)
Accede a la dirección de memoria de una variable. Sintaxis: &                   
    <nombre_de_la_variable>.

Declaración y asignación de un puntero

// declarando el puntero p, que almacenará
// una dirección de memoria de tipo string
var p *string

// asignando a p la dirección de memoria de fruit,
// p almacena, apunta o referencia la dirección de memoria de fruit
p = &fruit
Operador de Derreferenciación o Indirección
Accede al valor almacenado en la dirección de memoria donde apunta un puntero. Sintaxis: *<nombre_del_puntero>.