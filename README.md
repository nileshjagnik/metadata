Metadata Handler
==========
We are creating a metadata library to be used in a server that helps us retrieve metadata information and artwork about media files. Here is a high level picture of how the library works: 

![Metadata Library Architecture](docs/img/metadata_server.jpg)

Details
=======
* In our implementation of this library L, the origin server acts as a client to the library
* The main entry point to L will be a method/function call with a media name(string) whose metadata we return in a structured way
* Format of the metadata returned will be a json structure with details such as cast, director, and url of cover art images that are related to it. There may be more than one call for this information separately
* The library keeps a transparent cache as part of the implementation
* Whenever L gets a request from the client, it queries a local database (the cache) to find if the requested information is present in cache already
* If it is present, L reads the data from cache and returns it
* Otherwise S will request metadata from an online API, return the results as soon as possible and cache it
* The cache should use a caching policy, like LRU

TV Metadata
============
```go
data, err := metadata.GetMetadata("Modern Family - 1x03 - Come Fly with Me.avi","tv")
```

The above code will return a json string in the following format:-

```json
{
   "SeriesName":"Modern Family",
   "Banner_Url":"http://thetvdb.com/banners/",
   "Actors":"|Julie Bowen|Ty Burrell|Jesse Tyler Ferguson|Eric Stonestreet|Sofia Vergara|Ed O'Neill|Rico Rodriguez|Nolan Gould|Sarah Hyland|Ariel Winter|Aubrey Anderson-Emmons|",
   "Overview":"This mockumentary explores the many different types of a modern family through the stories of a gay couple, comprised of Mitchell and Cameron, and their daughter Lily, a straight couple, comprised of Phil and Claire, and their three kids, Haley, Alex, and Luke, and a multi-ethnic couple, which is comprised of Jay and Gloria, and her son Manny.",
   "Banner":"graphical/95011-g11.jpg",
   "FanArt":"fanart/original/95011-2.jpg",
   "Poster":"posters/95011-3.jpg",
   "Rating":"8.8",
   "FirstAired":"2009-09-23"
}
```

Movie Metadata
============

```go
data, err := metadata.GetMetadata("Alvin and the Chipmunks Chipwrecked (2011).avi","movie")
```

The above code will return a json string in the following format:-

```json
{
   "Id":55301,
   "Backdrop_path":"/qc0NReaia4MNsr8Ens054x6udrX.jpg",
   "Poster_path":"/31IPWvH2l4ycyK8EsMOmmwiJC7n.jpg",
   "Credits":{
      "Id":55301,
      "Cast":[
         {
            "Character":"Eleanor (voice)",
            "Name":"Amy Poehler",
            "Profile_path":"/l24EoxgHjv5RuPym4Uew0kmxngQ.jpg"
         },
         {
            "Character":"Zoe",
            "Name":"Jenny Slate",
            "Profile_path":"/bKyUz4N566bXBEhj1TkFfMiNFyT.jpg"
         },
         {
            "Character":"Captain Correlli",
            "Name":"Andy Buckley",
            "Profile_path":"/6HGgc8tWDiiiaTlxDXGB2TWydyL.jpg"
         },
         {
            "Character":"Theodore (voice)",
            "Name":"Jesse McCartney",
            "Profile_path":"/puqSuyl7MMZvx2nnRVegk9uVlOS.jpg"
         },
         {
            "Character":"David Seville",
            "Name":"Jason Lee",
            "Profile_path":"/67wSoVHxlOqtRhh3KJFBYf2qrDJ.jpg"
         },
         {
            "Character":"Jeanette (voice)",
            "Name":"Anna Faris",
            "Profile_path":"/2meONfWoGleESr78e9oxWIFhfMn.jpg"
         },
         {
            "Character":"Brittany (voice)",
            "Name":"Christina Applegate",
            "Profile_path":"/irUHUE17vDaSSYsOOJAxnHyEiSo.jpg"
         },
         {
            "Character":"Alvin (voice)",
            "Name":"Justin Long",
            "Profile_path":"/4mG8e1MAtkHAFewm2yJI3M1f59s.jpg"
         }
      ],
      "Crew":[
         {
            "Department":"Directing",
            "Name":"Mike Mitchell",
            "Job":"Director",
            "Profile_path":"/1nYMBY0XeYUeWsZlokoccYOU7h3.jpg"
         },
         {
            "Department":"Writing",
            "Name":"Jonathan Aibel",
            "Job":"Writer",
            "Profile_path":""
         },
         {
            "Department":"Writing",
            "Name":"Glenn Berger",
            "Job":"Writer",
            "Profile_path":""
         }
      ]
   },
   "Media_type":"movie",
   "Config":{
      "Images":{
         "Base_url":"http://image.tmdb.org/t/p/",
         "Secure_base_url":"https://image.tmdb.org/t/p/",
         "Backdrop_sizes":[
            "w300",
            "w780",
            "w1280",
            "original"
         ],
         "Logo_sizes":[
            "w45",
            "w92",
            "w154",
            "w185",
            "w300",
            "w500",
            "original"
         ],
         "Poster_sizes":[
            "w92",
            "w154",
            "w185",
            "w342",
            "w500",
            "w780",
            "original"
         ],
         "Profile_sizes":[
            "w45",
            "w185",
            "h632",
            "original"
         ],
         "Still_sizes":[
            "w92",
            "w185",
            "w300",
            "original"
         ]
      },
      "Change_keys":[
         "adult",
         "also_known_as",
         "alternative_titles",
         "biography",
         "birthday",
         "budget",
         "cast",
         "character_names",
         "crew",
         "deathday",
         "general",
         "genres",
         "homepage",
         "images",
         "imdb_id",
         "name",
         "original_title",
         "overview",
         "plot_keywords",
         "production_companies",
         "production_countries",
         "releases",
         "revenue",
         "runtime",
         "spoken_languages",
         "status",
         "tagline",
         "title",
         "trailers",
         "translations"
      ]
   },
   "Imdb_id":"tt1615918",
   "Overview":"Playing around while aboard a cruise ship, the Chipmunks and Chipettes accidentally go overboard and end up marooned in a tropical paradise. They discover their new turf is not as deserted as it seems.",
   "Title":"Alvin and the Chipmunks: Chipwrecked",
   "Release_date":"2011-12-14"
}
```


