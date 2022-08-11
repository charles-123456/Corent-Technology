import mysql.connector as mysql

def DataUpdate(utter_values):
    '''
    Pushes Unmatched intent Data to the Database
    '''
    db = mysql.connect(
                host="localhost",
                database="rasa",
                user="root",
                password="root"
                )

    mycursor = db.cursor()
    
    mysql_insert_query = """INSERT INTO rasainfo(utter_word) VALUES (%s);""".format()
    
    mycursor.execute(mysql_insert_query)
    
    db.commit()

    print("Record inserted successfully into table")