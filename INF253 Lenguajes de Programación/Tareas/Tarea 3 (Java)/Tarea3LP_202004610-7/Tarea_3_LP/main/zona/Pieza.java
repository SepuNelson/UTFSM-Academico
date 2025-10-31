package zona;

public class Pieza extends Zona {

    /* 
    ============================================================ 
    .                    A T R I B U T O S  
    ============================================================ 
    */
    
    protected int peso;
    //protected String nombre;

    /* 
    ============================================================ 
    .                   C O N S T R U C T O R  
    ============================================================ 
    */

    public Pieza(boolean usado,int peso){
        set_usado(usado);
        set_nombre("Pieza");
        this.peso = peso;
    }

    /* 
    ============================================================ 
    .                      G E T T E R S  
    ============================================================ 
    */

    public int get_peso(){
        return this.peso;
    }

    /* 
    ============================================================ 
    .                      S E T T E R S  
    ============================================================ 
    */

    public void set_peso(int peso){
        this.peso = peso;
    }
    
}
