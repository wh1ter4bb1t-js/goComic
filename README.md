# goComic

goComic is a cli tool written in go that scrapes your favorite childhood favorite comic from gocomics.com.
It will give you a single days comic relative to the date that you give it in its config file

## Usage
  ### How it works
  
  For goComic to work properly a config file is needed with the following attributes:
   - start_date: a date that you would like to be relative to the daily comics that you get. (Ie. if you put in todays date you will get the very first comic for each, if you put in a date 15 days prior you will be getting the 15th days comic);
   - folder: the folder that you would like the comics to be saved to. the folder will be made if it is does not exist
   - comics: an array of objects with the following:
     - url: the url for the actual comic. (Ie. https://www.gocomics.com/garfield)
     - title: the tile that you would like the file to have 
     - date: the start day of the comic, this can be found easily by going to the first day release of the comic

  ### Example config.json
    {
      "start_date": "2021/12/18",
      "folder": "{YOUR HOME DIRECTORY}/comics/",
      "comics": [
        {
          "url": "https://www.gocomics.com/foxtrotclassics", 
          "title": "Foxtrot Classics", 
          "date": "2007/01/01"
        },
        {
          "url": "https://www.gocomics.com/foxtrot", 
          "title": "Foxtrot", 
          "date": "1988/04/11"
        },
        {
          "url": "https://www.gocomics.com/getfuzzy",
          "title": "Get Fuzzy",
          "date": "1999/09/06"
        },
        {
          "url": "https://www.gocomics.com/calvinandhobbes", 
          "title": "Calvin and Hobbes", 
          "date": "1985/11/18"
        },
        { 
          "url": "https://www.gocomics.com/garfield", 
          "title": "Garfield", 
          "date": "1978/06/19"
        },
        {
          "url": "https://www.gocomics.com/pickles", 
          "title": "Pickles", 
          "date": "2003/01/01"
        },
        {
          "url": "https://www.gocomics.com/little-moments-of-love",
          "title": "Catana Comics", 
          "date": "2018/09/24"
        }
      ]
    }

  configs should be saved in your config directory
  ```
  $HOME/.config/goComic/config.json
  ```
