
def GetMovies(Connaction , selected_users):
  
    filtered_data = Connaction[Connaction['UserName'].isin(selected_users)]

    unique_movies = filtered_data['FilmName'].unique()

    return unique_movies

