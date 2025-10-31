package pikinim;

public class Cyan extends Pikinim{

    /* 
    ============================================================ 
    .                   C O N S T R U C T O R  
    ============================================================ 
    */

    public Cyan(double cantidad){
        set_ataque(1);
        set_capacidad(1);
        set_cantidad(cantidad);
    }

    /* 
    ============================================================ 
    .                      M E T O D O S  
    ============================================================ 
    */

    public void multiplicar(double cantidad){
        this.set_cantidad(cantidad * 3);
    }
}