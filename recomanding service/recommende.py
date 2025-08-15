import pandas as pd
from sklearn.metrics.pairwise import cosine_similarity
from GetMovies import GetMovies
from cash import CashSave

UserData = pd.read_csv('../tempf/UsersRating.csv')


user_item_matrix = UserData.pivot_table(index='UserName', columns='FilmName', values='Rating', fill_value=0)


user_similarity = cosine_similarity(user_item_matrix)
user_similarity_df = pd.DataFrame(user_similarity, index=user_item_matrix.index, columns=user_item_matrix.index)

#get all the similar users foe a spicifice one 
def recommend_similar_users(user, similarity_df, top_n=10):

    similar_users = similarity_df[user].sort_values(ascending=False)

    return similar_users[1:top_n+1]


def GetRecommendation():
    #Start whith getting all users
    RecommendList = []
    all_users = UserData['UserName'].unique()
    for user in all_users:
        SimilarUsers = recommend_similar_users(user , user_similarity_df).keys()
        recommendMovies = GetMovies(UserData , SimilarUsers).tolist()
        RecommendList.append((user , recommendMovies))
    
    #Save all the users with there movie in the cash for 24h to make is easi to get 
    CashSave(RecommendList)
        



#print(GetMovies(UserData , recommend_similar_users('User1', user_similarity_df).keys()))