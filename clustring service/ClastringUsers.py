from sklearn.cluster import KMeans
import matplotlib.pyplot as plt
import pandas as pd
from update import update

def ClustringUsers():
    # Read the csv file
    UsersData = pd.read_csv('../tempf/Users.csv')

    # Build the identity of the user [x, y] in 2d 
    UsersMatrix = UsersData[['xfeature', 'yfeature']]

    # Creating 3 clusters of users
    kmeans = KMeans(n_clusters=3, n_init="auto").fit(UsersMatrix)
    cluster_colors = ['red', 'blue', 'green']

    # Assign clusters
    UsersData['Cluster'] = kmeans.labels_
    UsersGroups = UsersData.groupby('Cluster')['user_id'].apply(list).to_dict()

    

    # Plot clusters
    plt.figure(figsize=(10, 6))
    for i in range(3):
        cluster_points = UsersMatrix[kmeans.labels_ == i]
        plt.scatter(cluster_points['xfeature'], cluster_points['yfeature'],
                    c=cluster_colors[i],
                    marker='o',
                    label=f'Cluster {i+1}')

    # Plot cluster centers
    plt.scatter(kmeans.cluster_centers_[:, 0], kmeans.cluster_centers_[:, 1], c='black', marker='X', label='Centers')
    # Example point
    plt.scatter(x=0.69195914, y=-0.17124348, c='slategray', marker="*")

    plt.xlabel('xfeature')
    plt.ylabel('yfeature')
    plt.title('User Clustering')
    plt.legend()
    plt.savefig('kmeansWUser2.png')
    plt.close()

ClustringUsers()