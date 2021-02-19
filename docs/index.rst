================
numerologySolver
================
---------------------------------
Simple words-to-numbers converter 
---------------------------------

Adds the values of letters in a word (case-insensitive, A=1, B=2, ..., Z=26) to find its value.

Can calculate the value of a given word, or find all known words with a given value.

DISCLAIMER!!
------------

I don't believe in numerology in any way.  This was actually written in response to an Internet meme that made me slightly irritated ;-) but nonetheless if you *do* like numerology for whatever reason you're welcome to use it and go in peace. Or use it for logic puzzles, school projects, whatever.

Started as an exercise in testing performance of various go data handling primitives, so don't expect much development beyond the initial feature set, but feel free to clone, update, whatever.  I'll accept PRs if they're nice and clear.

Getting Started
---------------

!! WARNING !! The code does not ship with a word list.  Before the ``/number`` API will work you'll need one in the form of a \n delimited list of words.  `words_alpha` from https://github.com/dwyl/english-words works well for this.  By default the code searches for 'words_aplha.txt' in the project root directory but you can configure this if you like.
  
Once you have a word list (or if you only want to use the ``/word`` API) run ``run.sh`` to bring up the server and make a test call.

Usage
-----

Runs as a web service - see code of ``run.sh`` for connection details and ENV overrides.  Examples:

Get all words with a particular value (and optional length: set Length to 0 for all known words):
=================================================================================================

::

  $ http PUT 127.0.0.1:8080/number Number=100 Length=14  
  HTTP/1.1 200 OK
  Content-Length: 101
  Content-Type: application/json
  Date: Wed, 23 Jan 2019 15:52:05 GMT
  {
     "SUCCESS": "There are 2 known words with numerological value 100: [batrachoididae biddulphiaceae]\n"
  }

Get the value of a word (note that words don't need to be known to the dictionary for this operation):  
======================================================================================================

::

   $ http PUT 127.0.0.1:8080/word Word="GitHub"
   HTTP/1.1 200 OK
   Content-Length: 49
   Content-Type: application/json
      Date: Wed, 23 Jan 2019 15:53:50 GMT
   {
      "SUCCESS": "Numerological value of GitHub: 67\n"
   }

