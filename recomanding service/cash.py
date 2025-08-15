import redis
import os


def CashSave(recomendation):
    r = redis.Redis(
    host= os.getenv("REDISHOST") ,
    port=10528,
    decode_responses=True,
    username=os.getenv("REDISUSERNAME"),
    password=os.getenv("REDISPASSWORD"),
    )

    for user , movies in recomendation:
        #that's is an 24h - one day -
        r.set(user, movies , ex= 60 * 1440)


