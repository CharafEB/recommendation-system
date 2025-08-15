from sklearn.cluster import KMeans
import matplotlib.pyplot as plt
import pandas as pd
from update import update 
from update import update

UserList = []
def ClustringUsers():
    #Read the csv file
    UsersData = pd.read_csv('../tempf/Users.csv')

    #build the indetity of the user [x,y] in 2d 
    UsersMatrix = UsersData[['xfeature', 'yfeature']]

    #creating 3 cluster of users (the is no - k - number is good enagh) 
    kmeans = KMeans(n_clusters=3, n_init="auto").fit(UsersMatrix)
    cluster_colors = ['red', 'blue', 'green']

    #
    UsersData['Cluster'] = kmeans.labels_
    UsersGroups = UsersData.groupby('Cluster')['user_id'].apply(list).to_dict()
    for clusterid, users in UsersGroups.items():
        data = [(clusterid,users)]
        UserList.append(data)
    
    update([(clusterID, UserID) for users in UserList for clusterID, usersID in users for UserID in usersID])
