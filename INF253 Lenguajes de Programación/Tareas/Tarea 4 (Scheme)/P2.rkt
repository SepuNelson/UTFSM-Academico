#lang scheme

;====================================
;           P R E G U N T A   2   C O L A
;====================================

(define (cantidades_cola base lista)
  (define (calculate acc remaining)
    ; ; Función auxiliar que procesa la lista de funciones con un número base
    ; ; y devuelve una lista de resultados.
    ; ;
    ; ; acc: Lista acumuladora que almacena los resultados intermedios de
    ; ; aplicar las funciones.
    ; ; remaining: Lista de funciones que se aplicarán al número base.

    (if (null? remaining)
        (reverse acc)
        (calculate (cons ((car remaining) base) acc) (cdr remaining))
    )
  )
  (calculate '() lista)
)

;====================================
;        P R E G U N T A   2   S I M P L E
;====================================

(define (cantidades_simple base lista)
  (map (lambda (f) (f base)) lista)
)

;EJEMPLOS
(cantidades_cola 2 (list (lambda (x) (/ x 2)) (lambda (x) (* x 3))
(lambda (x) (- 2 2))))

(cantidades_cola 2 (list (lambda (x) (/ (+ x 2) 2)) (lambda (x) (* x
4)) (lambda (x) (* (/ x 3) 2))))