"""
=================================================
.             L I B R E R Í A S                 .
=================================================

"""
import pandas as pd
import pyodbc
import zipfile
import csv

import warnings
warnings.filterwarnings("ignore", category=UserWarning, message="pandas only support SQLAlchemy connectable")


"""
=================================================
.              C O N E X I Ó N                  .
=================================================

"""
try:
    connection = pyodbc.connect('DRIVER={SQL Server};'
                                'SERVER=LAPTOP-9S0IBRT4\SQLEXPRESS;'
                                'DATABASE=FutUSM;'
                                'UID=Nelson;'
                                'PWD=6357661605;')
    print("\nConexion exitosa! Espere un momento...\n")
except:
    print("\nERROR . . . ERROR . . . ERROR\n")

cursor = connection.cursor()

"""
=================================================
.           C R E A R   T A B L A S             .
=================================================

"""

cursor.execute("DROP TABLE IF EXISTS mundiales")

cursor.execute("""CREATE TABLE mundiales (
                id INTEGER IDENTITY(1,1) PRIMARY KEY  NOT NULL,
                Year INT NOT NULL,
                Position INT NOT NULL,
                Team VARCHAR(100) NOT NULL,
                Games_Played INT NOT NULL,
                Win INT NOT NULL,
                Draw INT NOT NULL,
                Loss INT NOT NULL,
                Goals_For INT NOT NULL,
                Goals_Against INT NOT NULL,
                Goal_Difference INT NOT NULL,
                Points INT NOT NULL)""")

cursor.execute("DROP TABLE IF EXISTS world_cup_summary")

cursor.execute("""CREATE TABLE world_cup_summary (
                id INTEGER IDENTITY(1,1) PRIMARY KEY  NOT NULL,
                YEAR INT NOT NULL,
                HOST VARCHAR(100) NOT NULL,
                CHAMPION VARCHAR(100) NOT NULL,
                RUNNER_UP VARCHAR(100) NOT NULL,
                THIRD_PLACE VARCHAR(100) NOT NULL,
                TEAMS INT NOT NULL,
                MATCHES_PLAYED INT NOT NULL,
                GOALS_SCORED INT NOT NULL,
                AVG_GOALS_PER_GAME FLOAT NOT NULL)""")

connection.commit()

"""
=================================================
.                 T A B L A S                   .
=================================================

"""

with zipfile.ZipFile("FIFA.zip", 'r') as zip_ref:
    lista_archivos = zip_ref.namelist()
    c = 0
    for nombre_archivo in lista_archivos:
        c += 1
        cond = 0
        if c == 23:
            with zip_ref.open(nombre_archivo, 'r') as archivo_interno:
                contenido_csv = archivo_interno.read().decode('utf-8')
                lineas = contenido_csv.split('\n')
                lector_csv = csv.reader(lineas)
                for row in lector_csv:
                    if cond != 0 and row != []:
                        year,host,champion,runner_up,third_place,teams,matches_played,goals_scored,avg_goals_per_game = row
                        consulta = 'INSERT INTO world_cup_summary VALUES(?,?,?,?,?,?,?,?,?);'
                        cursor.execute(consulta,(int(year),host,champion,runner_up,third_place,int(teams),int(matches_played),int(goals_scored),float(avg_goals_per_game)))
                    cond += 1
        else:
            anio = int(nombre_archivo[7:11])
            with zip_ref.open(nombre_archivo, 'r') as archivo_interno:
                contenido_csv = archivo_interno.read().decode('utf-8')
                lineas = contenido_csv.split('\n')
                lector_csv = csv.reader(lineas)
                for row in lector_csv:
                    if cond != 0 and row != []:
                        position,team,games_played,win,draw,loss,goals_for,goals_against,goal_difference,points = row
                        if '−'in goal_difference:
                            goal_difference = goal_difference.replace("−", "-")
                        consulta = 'INSERT INTO mundiales VALUES(?,?,?,?,?,?,?,?,?,?,?);'
                        cursor.execute(consulta,(anio,position,team,games_played,win,draw,loss,goals_for,goals_against,goal_difference,points))
                    cond += 1

connection.commit()

"""
=================================================
.             F U N C I O N E S                 .
=================================================

"""

def texto_respuesta():
    """
    Parametro: No posee
    Función: Printea un texto antes de cada respuesta
    Return: No retorna "algo"
    """
    print("\n=================================================")
    print(".                  RESPUESTA                    .")
    print("=================================================\n")
    return

"""
=================================================
.                   M A I N                     .
=================================================

"""

print("=================================================")
print(".                Interacciones                  .")
print("=================================================")
print(". 0   Para Terminar el Programa                 .")
print(". 1   Mostrar Campeones                         .")
print(". 2   Mostrar goleadores                        .")
print(". 3   Mostrar Tercer Lugar más veces            .")
print(". 4   Mostrar País más goles recibidos          .")
print(". 5   Buscar un país                            .")
print(". 6   Top 3 países en el mundial                .")
print(". 7   Mayor cantidad ganados                    .")
print(". 8   Países ganando en casa                    .")
print(". 9   Más veces en el podio                     .")
print(". 10  Mayores rivales                           .")
print("=================================================\n")

flag = True
while flag == True:

    while True:
        try:
            cond = int(input("\nIngrese una de las opciones: "))
            break
        
        except ValueError:
            print("\nERROR . . . ERROR . . . ERROR")
            print("Ingrese una de las opciones señaladas\n")

    # Para Terminar el Programa
    if cond == 0:

        flag = False
        connection.close()

    # Mostrar Campeones
    elif cond == 1:

        texto_respuesta()
        query = "SELECT CHAMPION, YEAR FROM world_cup_summary ORDER BY YEAR DESC"
        df = pd.read_sql_query(query, connection)
        for index, row in df.iterrows():
            print("◦ " + str(row['YEAR']) + " -> " + row['CHAMPION'] )

    # Mostrar goleadores 
    elif cond == 2:

        texto_respuesta()
        query = "SELECT TOP 5 Team, SUM(Goals_For) AS Total_Goals FROM mundiales GROUP BY Team ORDER BY Total_Goals DESC"
        df = pd.read_sql_query(query, connection)
        for index, row in df.iterrows():
            print("◦ " + str(row['Total_Goals']) + " -> " + row['Team'])

    # Mostrar Tercer Lugar más veces
    elif cond == 3:

        texto_respuesta()
        query = "SELECT TOP 5 Team, COUNT(*) as Third_Place_Count FROM mundiales WHERE Position = 3 GROUP BY Team ORDER BY Third_Place_Count DESC"
        df = pd.read_sql_query(query, connection)
        for index, row in df.iterrows():
            print("◦ " + str(row['Third_Place_Count']) + " -> " + row['Team'] )
        

    # Mostrar País más goles recibidos
    elif cond == 4:

        texto_respuesta()
        query = "SELECT Team, SUM(Goals_Against) AS Total_Goals_Against FROM mundiales GROUP BY Team ORDER BY Total_Goals_Against DESC"
        df = pd.read_sql_query(query, connection)
        print("◦ " + str(df.iloc[0]['Total_Goals_Against']) + " -> " + df.iloc[0]['Team'] )

    # Buscar un país
    elif cond == 5:

        flag_cond_5 = True
        while(flag_cond_5 == True):
            pais_buscar = str(input("Ingrese el país que desea buscar: "))
            query = "SELECT * FROM mundiales WHERE Team = '" + pais_buscar + "' ORDER BY Year DESC"
            df = pd.read_sql_query(query, connection)
            if not df.empty:
                flag_cond_5 =  False
                texto_respuesta()
                df = df.drop(columns=['id'])
                df = df.drop(columns=['Team'])
                print(df.to_string(index=False))
            else:
                print("\nERROR . . . ERROR . . . ERROR")
                print("No se encuentra el nombre en la Base de Datos")
                print("Por favor ingresar un nombre válido\n")

    # Top 3 países en el mundial
    elif cond == 6:

        texto_respuesta()
        cursor.execute("DROP VIEW IF EXISTS mundiales_aux")
        cursor.execute('CREATE VIEW mundiales_aux AS SELECT Year, Team FROM mundiales')

        query = "SELECT TOP 3 Team, STRING_AGG(Year, ', ') as Years_Participated FROM mundiales_aux GROUP BY Team ORDER BY COUNT(*) DESC"
        df = pd.read_sql_query(query, connection)
        for index, row in df.iterrows():
            print("◦ " + row["Team"] + " -> " + str(row["Years_Participated"]))

    # Mayor cantidad ganados
    elif cond == 7:

        texto_respuesta()
        query = "SELECT TOP 1 Team, CAST(SUM(Win) AS FLOAT) / CAST(SUM(Games_Played) AS FLOAT) AS Win_Rate FROM mundiales GROUP BY Team ORDER BY Win_Rate DESC"
        df = pd.read_sql_query(query, connection)
        print("◦ " + df.iloc[0]["Team"] + " -> " + str(round(df.iloc[0]["Win_Rate"] * 100,2)) + "%")

    # Países ganando en casa
    elif cond == 8:

        texto_respuesta()
        query = "SELECT HOST, CHAMPION, YEAR FROM world_cup_summary ORDER BY YEAR DESC"
        df = pd.read_sql_query(query, connection)
        df = df[df["HOST"] == df["CHAMPION"]]
        for index, row in df.iterrows():
            print("◦ " + str(row["YEAR"]) + " -> " + row["HOST"])
        
    # Más veces en el podio
    elif cond == 9:

        texto_respuesta()
        query = "SELECT TOP 1 Team, COUNT(*) AS Podium_Count FROM mundiales WHERE Position IN (1, 2, 3) GROUP BY Team ORDER BY Podium_Count DESC"
        df = pd.read_sql_query(query, connection)
        print("◦ " + df.iloc[0]["Team"] + " -> " + str(df.iloc[0]["Podium_Count"]))

    # Mayores rivales
    elif cond == 10:

        texto_respuesta()
        query = "SELECT TOP 1 First.Team AS First_Country, Second.Team AS Second_Country, COUNT(*) AS Match_Count FROM mundiales AS First JOIN mundiales AS Second ON First.Year = Second.Year AND First.Position = 1 AND Second.Position = 2 WHERE First.Team < Second.Team GROUP BY First.Team, Second.Team ORDER BY Match_Count DESC"
        df = pd.read_sql_query(query, connection)
        for index, row in df.iterrows():
            print("◦ " + row["First_Country"] + " v/s " + row["Second_Country"] + " -> " + str(row["Match_Count"]) + " veces")

    # Error al ingresar una opción
    else:
        print("\nERROR . . . ERROR . . . ERROR")
        print("Ingrese una de las opciones señaladas\n")