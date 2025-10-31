#lang scheme

;================================
;               P R E G U N T A   1 
;================================

(define (checkear cantidad lista)
  (cond
    ((= cantidad 0) (if (null? lista) (display "true") (display "false")))
    ((null? lista) (if (= cantidad 0) (display "true") (display "false")))
    (else (checkear (- cantidad 1) (cdr lista)))))


;EJEMPLOS
(checkear 4 '(1 2 5 4))
(newline)
(checkear 2 '(a a c d e w w q t a v))
(newline) 
(checkear 0 '())