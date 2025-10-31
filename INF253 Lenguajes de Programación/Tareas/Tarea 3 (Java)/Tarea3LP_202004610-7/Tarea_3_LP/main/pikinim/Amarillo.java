package pikinim;

public class Amarillo extends Pikinim{

    /* 
    ============================================================ 
    .                   C O N S T R U C T O R  
    ============================================================ 
    */

    public Amarillo(double cantidad){
        set_ataque(1);
        set_capacidad(3);
        set_cantidad(cantidad);
    }

    /* 
    ============================================================ 
    .                      M E T O D O S  
    ============================================================ 
    */

    public void multiplicar(double cantidad){
        this.set_cantidad(cantidad * 1.5);
    }
}
