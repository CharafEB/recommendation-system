import psycopg2 
import os
from dotenv import load_dotenv
import io
load_dotenv()

#update : do update the cluster that user on 
def update(users):
    try:
        conn = psycopg2.connect(os.getenv("POSTGERS_API_LINE"))  
        print("Successfully connected to the database.")
        cursor = conn.cursor()
        
       
        query = "UPDATE users SET cluster_id = %s WHERE user_id = %s"
        
        cursor.executemany(query, users)
        
        conn.commit()  
        print(f"Updated {cursor.rowcount} rows successfully.")
        
    except psycopg2.Error as e:
        print(f"Error connecting to the database: {e}")
        conn.rollback()  
    finally:
        if cursor:
            cursor.close()
        if conn:
            conn.close()