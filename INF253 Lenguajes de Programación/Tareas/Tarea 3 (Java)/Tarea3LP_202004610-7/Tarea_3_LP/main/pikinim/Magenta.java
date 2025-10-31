package pikinim;

public class Magenta extends Pikinim{

    /* 
    ============================================================ 
    .                   C O N S T R U C T O R  
    ============================================================ 
    */

    public Magenta(double cantidad){
        set_ataque(2);
        set_capacidad(1);
        set_cantidad(cantidad);
    }

    /* 
    ============================================================ 
    .                      M E T O D O S  
    ============================================================ 
    */

    public void multiplicar(double cantidad){
        this.set_cantidad(cantidad * this.get_ataque());
    }
}
