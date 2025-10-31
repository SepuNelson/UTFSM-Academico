package zona;

public class Enemigo extends Zona {
    
    /* 
    ============================================================ 
    .                    A T R I B U T O S  
    ============================================================ 
    */

    protected int vida;
    protected int peso;
    protected int ataque;
    

    /* 
    ============================================================ 
    .                   C O N S T R U C T O R  
    ============================================================ 
    */

    public Enemigo(boolean usado, int vida, int peso, int ataque){
        set_usado(usado);
        set_nombre("Enemigo");
        this.vida = vida;
        this.peso = peso;
        this.ataque = ataque;
    }

    /* 
    ============================================================ 
    .                      G E T T E R S  
    ============================================================ 
    */

    public int get_vida(){
        return this.vida;
    }

    public int get_peso(){
        return this.peso;
    }

    public int get_ataque(){
        return this.ataque;
    }

    /* 
    ============================================================ 
    .                      S E T T E R S  
    ============================================================ 
    */

    public void set_peso(int peso){
        this.peso = peso;
    }

    public void set_vida(int vida){
        this.vida = vida;
    }

    public void set_ataque(int ataque){
        this.ataque = ataque;
    }
    
}
