import schedule
import time 
from ClastringUsers import ClustringUsers

def RunService():
    schedule.every(23).hour.do(ClustringUsers)
    while True :
        schedule.run_pending()
        time.sleep(1)

RunService()
