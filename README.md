# GolangSeminar
Este proyecto fué realizado como entrega para un seminario.

Para ver la consigna original, ir a [INSTRUCTIONS](/INSTRUCTIONS.md)

Utilicé 'fmt' para obtener el input del usuario en main.go, continúo solicitándolo hasta que inserte uno válido. Si lo es, imprimo el resultado y finalizo el programa.

La estructura Result consiste, como dice la consigna, en un tipo, un largo y un valor. Está declarada en result/result.go junto a una función NewResult() y otra ParseStringToResult().

Toda la lógica está en ParseStringToResult() y es la siguiente:
- Se asegura que el texto cumpla con un largo mínimo de 5 (2 dígitos para el tipo, 2 para el largo y por lo menos 1 para el valor)
- Obtiene el tipo
- Obtiene el largo
- Parsea el largo de string a int y verifica que no haya errores
- Verifica que el largo declarado por el usuario no sea mayor que el largo verdadero que tiene el valor que insertó
- Obtiene el valor
- De haber llegado hasta aquí sin devolver errores, devuelve los 3 valores y error 'nil'

Por último, en result_test.go están los tests implementados para el paquete result, con un coverage del 100%