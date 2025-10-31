package pikinim;

public abstract class Pikinim {

    /* 
    ============================================================ 
    .                    A T R I B U T O S  
    ============================================================ 
    */

    protected int ataque;
    protected int capacidad;
    protected double cantidad;

    /* 
    ============================================================ 
    .                      G E T T E R S  
    ============================================================ 
    */
    
    public int get_ataque(){
        return this.ataque;
    }

    public int get_capacidad(){
        return this.capacidad;
    }
    public double get_cantidad(){
        return this.cantidad;
    }

    /* 
    ============================================================ 
    .                      S E T T E R S  
    ============================================================ 
    */

    public void set_ataque(int ataque){
        this.ataque = ataque;
    }

    public void set_capacidad(int capacidad){
        this.capacidad = capacidad;
    }
    public void set_cantidad(double cantidad){
        this.cantidad = cantidad;
    }

    /* 
    ============================================================ 
    .                      M E T O D O S  
    ============================================================ 
    */

    public void disminuir(int cantidad){
        this.cantidad -= cantidad;
    }

    public abstract void multiplicar(double cantidad);
}