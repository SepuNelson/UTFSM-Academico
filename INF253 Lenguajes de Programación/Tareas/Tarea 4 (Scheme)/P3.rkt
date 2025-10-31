#lang scheme

;================================
;               P R E G U N T A   3 
;================================

(define (armar_lista stock)
  (define (armar_lista_aux stock acumulada)
    ; ; Función auxiliar que verifica el stock para determinar la
    ; ; cantidad de ingredientes que faltan comprar.
    ; ;
    ; ; stock : Lista que contiene pares de cantidad necesaria
    ; ; e ingredientes disponibles.
    ; ; acumulada : Lista acumuladora que almacena los ingredientes
    ; ; que faltan y su cantidad a comprar.

    (cond
      ((null? stock) (reverse acumulada))
      (else
       (let* ((cantidad (car (car stock)))
              (ingredientes (cadr (car stock)))
              (disponibles (length ingredientes)))
         (if (< cantidad disponibles)
             (armar_lista_aux (cdr stock) acumulada)
             (let ((cant_comprar (- cantidad disponibles)))
               (if (> cant_comprar 0)
                   (armar_lista_aux (cdr stock) (cons (list cant_comprar (car ingredientes)) acumulada))
                   (armar_lista_aux (cdr stock) acumulada)))))))
  )
  (armar_lista_aux stock '())
)

; EJEMPLO
(armar_lista '((5 (cebolla cebolla)) (3 (tomate tomate tomate)) (2 (ajo))))
