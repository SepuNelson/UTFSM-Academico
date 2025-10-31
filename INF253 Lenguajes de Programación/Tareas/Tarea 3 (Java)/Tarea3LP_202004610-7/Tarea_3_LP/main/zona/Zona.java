package zona;

public class Zona {

    /* 
    ============================================================ 
    .                    A T R I B U T O S  
    ============================================================ 
    */

    private boolean usado;
    protected String nombre;
    

    /* 
    ============================================================ 
    .                      G E T T E R S  
    ============================================================ 
    */

    public boolean get_usado(){
        return this.usado;
    }

    public String get_nombre(){
        return this.nombre;
    }

    /* 
    ============================================================ 
    .                      S E T T E R S  
    ============================================================ 
    */

    public void set_usado(boolean usado){
        this.usado = usado;
    }

    public void set_nombre(String nombre){
        this.nombre = nombre;
    }

    /* 
    ============================================================ 
    .                      M E T O D O S  
    ============================================================ 
    */

    public void Interactuar(){
        /* NO SE QUE SE HACE CON LOS 3 COLORES */
    }
    
}
