import schedule
import time 
from recommende import GetRecommendation


def RunService():
    schedule.every(24).hour.do(GetRecommendation)
    while True :
        schedule.run_pending()
        time.sleep(1)

RunService()