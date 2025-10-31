//==================================================================
//                  M O D U L O   C O N T A D O R
//==================================================================

module Contador (
  output wire [5:0] dir1,
  input wire clk1,
  input wire reset1
);
  
  reg [5:0] dir0 = 0;
  
  always @(posedge clk1, posedge reset1) begin
    if (reset1) 
      dir0 <= 0;
    else 
      dir0 <= dir0 + 1;
  end
  
  assign dir1 = dir0;
  
endmodule

//==================================================================
//                      M O D U L O   R O M
//==================================================================

module ROM (
  input wire [5:0] dir2,
  output wire [18:0] ins1
);

  reg [18:0] mem [0:63];

  initial begin
    
    // Asignar los valores del Ejemplo
    
    mem[0] = 19'b000_00010111_00010011;
    mem[1] = 19'b001_00000111_01001100;
    mem[2] = 19'b010_00011111_00000101;
    mem[3] = 19'b011_00011111_00000010;
    mem[4] = 19'b100_01011101_01010010;
    mem[5] = 19'b101_01011101_01010010;
    mem[6] = 19'b110_01011101_01010010;
    mem[7] = 19'b111_01011101_01010010;
    
  end
  
  integer addr;
  
  always @(*) begin
    addr = $bin2dec(dir2);
  end

  assign ins1 = mem[addr];

endmodule

//==================================================================
//                   M O D U L O   S P L I T T E R
//==================================================================

module Splitter(
  input wire [18:0] ins2,
  output reg [2:0] op1,
  output reg [7:0] A1,
  output reg [7:0] B1
);
  
  always @* begin
    op1 = ins2[2:0];
    A1 = ins2[15:8];
    B1 = ins2[7:0];
  end
  
endmodule

//==================================================================
//                      M O D U L O   A L U
//==================================================================

module ALU (
  input wire [2:0] op2,
  input wire [7:0] A2,
  input wire [7:0] B2,
  output reg [7:0] resultado1
);
  reg [7:0] carry [0:8];

  always_comb begin
    case (op2)
      3'b000: begin // Suma
        carry[0] = 1'b0; 
        for (int i = 0; i < 8; i = i + 1) begin
          carry[i+1] = (A2[i] & B2[i]) | (A2[i] & carry[i]) | (B2[i] & carry[i]);
          resultado1[i] = A2[i] ^ B2[i] ^ carry[i];
        end
      end
      3'b001: begin // Resta
        carry[0] = 1'b1; 
        for (int i = 0; i < 8; i = i + 1) begin
          carry[i+1] = (A2[i] & ~B2[i]) | (A2[i] & carry[i]) | (~B2[i] & carry[i]);
          resultado1[i] = A2[i] ^ ~B2[i] ^ carry[i];
        end
      end
      3'b010: resultado1 = A2 << B2; // Bitshift izquierda
      3'b011: resultado1 = A2 >> B2; // Bitshift derecha
      3'b100: resultado1 = A2 & B2; // Composición
      3'b101: resultado1 = A2 | B2; // Disyunción
      3'b110: resultado1 = A2 ^ B2; // Exclusión
      3'b111: resultado1 = ~A2; // Negación
      default: resultado1 = 8'b0;
    endcase
  end
endmodule


//==================================================================
//                M O D U L O   R E S U L T A D O
//==================================================================

module Resultado (
  input wire [7:0] resultado2,
  input wire clk,
  input wire reset
);
  
endmodule

//==================================================================
//                   M O D U L O   T O P
//==================================================================

module top (
  output reg [5:0] dir,
  output reg [18:0] ins,
  output reg [2:0] op,
  output reg [7:0] A,
  output reg [7:0] B,
  output reg [7:0] resultado,
  input wire clk,
  input wire reset
);
  
  //=============================
  //  D E C L A R A C I O N E S
  //=============================
  
  wire [5:0] dir1;
  wire [18:0] ins1;
  wire [2:0] op1;
  wire [7:0] A1;
  wire [7:0] B1;
  wire [7:0] resultado1;
  
  //=======================
  //  I N S T A N C I A S
  //=======================
  
  Contador contador (
    .dir1(dir1),
    .clk1(clk),
    .reset1(reset)
  );
  
  ROM rom (
    .dir2(dir1),
    .ins1(ins1)
  );
  
  Splitter splitter (
    .ins2(ins1),
    .op1(op1),
    .A1(A1),
    .B1(B1)
  );
  
  ALU alu (
    .op2(op1),
    .A2(A1),
    .B2(B1),
    .resultado1(resultado1)
  );
  
  Resultado resultadoModulo (
    .resultado2(resultado1),
    .clk(clk),
    .reset(reset)
  );
  
  //=======================
  //  C O N E X I O N E S  
  //=======================
  
  // Conectar dir del contador a dir del ROM
  assign dir1 = contador.dir1; 
  
  // Conectar ins del ROM a entrada del Splitter
  assign ins1 = rom.ins1;
  
  // Conectar op del Splitter a op del ALU
  assign op1 = splitter.op1; 
  // Conectar A del Splitter a A del ALU
  assign A1 = splitter.A1; 
  // Conectar B del Splitter a B del ALU
  assign B1 = splitter.B1; 
  
  // Conectar resultado del ALU a resultado del Resultado
  assign resultado1 = resultadoModulo.resultado2; 
  
  // Incrementar la dirección en cada flanco de subida del reloj
  always @(posedge clk, posedge reset) begin
    if (reset)
      dir <= 0;
    else
      dir <= dir + 1;
  end
  
  // Asignar las salidas
  always @(posedge clk) begin
    ins <= ins1;
    op <= op1;
    A <= A1;
    B <= B1;
    resultado <= resultado1;
  end
  
endmodule