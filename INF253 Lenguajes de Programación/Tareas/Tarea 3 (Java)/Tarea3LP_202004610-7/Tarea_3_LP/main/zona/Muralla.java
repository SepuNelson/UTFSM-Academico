package zona;
import pikinim.Pikinim;

public class Muralla extends Zona {

    /* 
    ============================================================ 
    .                    A T R I B U T O S  
    ============================================================ 
    */
    
    protected double vida;
    //protected String nombre;

    /* 
    ============================================================ 
    .                   C O N S T R U C T O R  
    ============================================================ 
    */

    public Muralla(boolean usado,double vida){
        set_usado(usado);
        set_nombre("Muralla");
        this.vida = vida;
    }

    /* 
    ============================================================ 
    .                      G E T T E R S  
    ============================================================ 
    */

    public double get_vida(){
        return this.vida;
    }

    /* 
    ============================================================ 
    .                      S E T T E R S  
    ============================================================ 
    */
    
    public void set_vida(double vida){
        this.vida = vida;
    }

    /* 
    ============================================================ 
    .                      M E T O D O S  
    ============================================================ 
    */

    public boolean TryRomper(Pikinim Cyan, Pikinim Magenta, Pikinim Amarillo){
        double ataque_total = Cyan.get_ataque() * Cyan.get_cantidad() + Magenta.get_ataque() * Magenta.get_cantidad() + Amarillo.get_ataque() * Amarillo.get_cantidad();
        this.set_vida((this.vida - ataque_total));

        if(this.get_vida() < 0){this.set_vida(0);}

        if(this.vida == 0){return true;}
        else{return false;}
    }
}
