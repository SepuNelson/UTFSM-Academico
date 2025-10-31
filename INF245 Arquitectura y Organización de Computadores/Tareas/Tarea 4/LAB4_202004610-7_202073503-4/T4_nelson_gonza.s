.data
	input:		.word 	1		@ Valor de la Función 
	palabra1:	.asciz	"ol"		@ Palabra 1
	palabra2:	.asciz	"ko"		@ Palabra 2
	largo_p1:	.word	2		@ Tamaño palabra 1
	largo_p2:	.word	2		@ Tamaño palabra 2
	abcd1:		.space	26
	abcd2:		.space	26
	n:		.word	5		@ Valor de n
	k:		.word	0		@ Valor de k
	arr:		.word	2,4,5 		@ Arreglo
	t_arr:		.word	3		@ Tamaño del arreglo
	@arr_par:   	.space 	20

.text

main:
    	ldr r0, =input  @ Vemos que opción es 1, 2, 3
	ldrb r0, [r0]   @ Carga 1 byte del valor de la dm en el reg 0
	cmp r0, #1	@ Compara r0 con 1
	beq funcion_1   @ Manda a fun 1
	cmp r0, #2      @ Compara r0 con 2
	beq funcion_2   @ Manda a fun 2
	cmp r0, #3      @ Compara r0 con 3
	beq funcion_3 	@ Manda a fun 3
	bl exit         @ Termina

@==========================================================
@                   F U N C I Ó N   1
@==========================================================

funcion_1:		@ Función de los Anagramas

	ldr r3, =palabra1	@guardamos los inputs en los registros para luego comparar
	ldr r7, =palabra2
	ldr r4, =largo_p1
	ldr r6, =largo_p2
	ldrb r4, [r4]
	mov r1, #0

leerp1:	cmp r1, r4		@mediante un ciclo se almacena la cantidad de caracteres en una especie de abecedario
	beq fin_p1		@que luego servirá para comparar la cantidad de letras en cada una de las palabras
	ldrb r5, [r3,r1]	
	ldr r0, =abcd1		@en la linea 36 se comprueba si se ha llegado al final de la primera palabra, y si
	sub r5, r5, #97		@es así, se procede a hacer el mismo proceso pero con la segunda palabra
	add r0, r0, r5		
	ldrb r5, [r0]
	add r5, r5, #1
	add r1, r1 , #1
	strb r5, [r0]
	b leerp1

fin_p1: ldrb r6, [r6]		@cuando se termina de almacenar cada letra de palabra 1, se reinicia el contador del ciclo
	mov r1, #0		@y se empieza a trabajar con la palabra 2

leerp2:	cmp r1, r6		@se realiza la misma operacion realizada con la palabra 1, pero con la palabra 2
	beq fin_p2		@y en su respectivo abecedario
	ldrb r5, [r7,r1]
	ldr r0, =abcd2
	sub r5, r5, #97
	add r0, r0, r5
	ldrb r5, [r0]
	add r5, r5, #1
	add r1, r1 , #1
	strb r5, [r0]
	b leerp2

fin_p2:	ldr r3, =abcd1		@cuando se termina de almacenar cada letra de la palabra 2, se guardan estos abecedarios
	ldr r7, =abcd2		@dentro de los registros r3 y r7, respectivamente.
	mov r1, #0		@también se activan unas flags que luego serán cambiadas en la subrutina "fin" una vez se comparen
	mov r2, #1		@por completo cada uno de los abecedarios.

cmp_abcd: 			@se realiza un ciclo parecido al de las subrutinas leerp1 y leerp2
	cmp r1, #26
	beq output
	ldrb r4, [r3,r1]
	ldrb r5, [r7,r1]
	cmp r4, r5
	bne fin
	add r1, r1, #1
	b cmp_abcd

fin:	mov r2, #0		
	add r1, r1, #1
	b cmp_abcd

output:	mov r0, #20
	mov r1, #2
	bl printInt
	b exit

@==========================================================
@                   F U N C I Ó N   2
@==========================================================
	
funcion_2:		@ Función Recursiva

	mov r0, #0	@ Limpiar r0 con 0
	ldr r3,	=n  @ Valor de n
	ldrb r3, [r3]
	ldr r4, =k  @ Valor de k
	ldrb r4, [r4]
	bl rec

main_rec:

	mov r2, r0
	mov r0, #20
	mov r1, #2
	bl printInt
	b exit

rec:			@ Main de la función recursiva

	push {r3-r4, lr}
	cmp r4, r3	
	bgt if		@ k > n
	beq elif_1      @ n == k
	cmp r4, #0
	beq elif_2	@ k == 0
	bl else

if:			@ If de la función // Retorna 0

	mov r5, #0
	add r0, r0, r5
	bl f2_fin

elif_1:			@ Elif de la función // Retorna 1
	
	mov r5, #1
	add r0, r0, r5
	bl f2_fin

elif_2:			@ Elif de la función // Retorna 1

	mov r5, #1
	add r0, r0, r5
	bl f2_fin
	

else:			@ Else de la función // Recursivo

	sub r3, r3, #1
	bl rec

else_2:

	sub r4, r4, #1
	bl rec
		

f2_fin:
	pop {r3-r4,pc}	
	 

@==========================================================
@                   F U N C I Ó N   3
@==========================================================

funcion_3:

	mov r0 , #0
	ldr r0, =t_arr
	ldrb r0, [r0]
	@ldr r5, =arr_par

	mov r1, #0
	mov r2, #0
	mov r3, #0
	mov r4, #1
	@mov r6, #0

loop:

	ldr r3, =arr
	ldrb r3, [r3, r1]
	and r3, r3, r4
	cmp r3, #0
	beq es_par
	b es_impar

es_par:

	add r2, r2, #1 	@Aumentamos en 1 el contador de números pares
	add r1, r1, #4 	@Aumentamos en 4 bytes las dirección del elemento del arreglo a leer
	sub r0, r0, #1 	@Aumentamos en 1 a la cantidad de números que hemos leido
	cmp r0, #0
	bne loop 	@Si quedan números por leer se hace el salto.
	b next 		@Si no quedan números imprimimos el resultado

es_impar: 		@Análogo a cuando es par, pero sin aumentar el contador de pares.

	add r1, r1, #4
	sub r0, r0, #1
	cmp r0, #0
	bne loop
	b next

next:

	mov r0, #0
	mov r1, #0
	bl printInt
	b exit

exit:

	wfi
