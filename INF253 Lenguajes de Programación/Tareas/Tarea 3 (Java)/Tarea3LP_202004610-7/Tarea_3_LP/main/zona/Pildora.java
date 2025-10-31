package zona;

public class Pildora extends Zona {
    
    /* 
    ============================================================ 
    .                    A T R I B U T O S  
    ============================================================ 
    */

    protected int cantidad;
    //protected String nombre;

    /* 
    ============================================================ 
    .                   C O N S T R U C T O R  
    ============================================================ 
    */

    public Pildora(boolean usado,int cantidad){
        set_usado(usado);
        set_nombre("Píldora");
        this.cantidad = cantidad;
    }

    /* 
    ============================================================ 
    .                      G E T T E R S  
    ============================================================ 
    */

    public int get_cantidad(){
        return this.cantidad;
    }

    /* 
    ============================================================ 
    .                      S E T T E R S  
    ============================================================ 
    */

    public void set_cantidad(int cantidad){
        this.cantidad = cantidad;
    }
}
