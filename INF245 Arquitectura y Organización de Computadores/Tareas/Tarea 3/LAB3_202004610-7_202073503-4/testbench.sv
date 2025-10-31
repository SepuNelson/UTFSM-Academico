module top_tb;
  reg clk;
  reg reset;
  
  
  wire [5:0] dir;
  wire [18:0] ins;
  wire [2:0] op;
  wire [7:0] A;
  wire [7:0] B;
  wire [7:0] resultado;
  
  top dut (
    .dir(dir),
    .ins(ins),
    .op(op),
    .A(A),
    .B(B),
    .resultado(resultado),
    .clk(clk),
    .reset(reset)
  );
  
  
  initial begin
    clk = 0;
    reset = 1;
    #10 reset = 0;
    
    
    repeat(20) begin
      #5 clk = ~clk;
    end
    
    $display("<%b> <%b> <%b> <%b> <%b>", ins, op, A, B, resultado);
        
    #5 $finish;
  end

endmodule
