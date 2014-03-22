ghcontribs
--

Command-line app to count the number of GitHub contributions made by some users
in the last year. Since this data is not easily available in the Github API, it
uses the calendar data from the user's profile page, so expect this to break as
soon as they change that UI.

Functionality is also available as a Go library if that is what you prefer. [Docs are here](https://godoc.org/github.com/aliafshar/ghcontribs)

    $ ghcontribs --help
    Usage: ghcontrib [options] username, [username, ...]
      -after="": Date after which to count. e.g. 2013-Sep-30
    exit status 2

    $ ghcontribs aliafshar
    aliafshar 111

    $ ghcontribs --after 2013-Sep-30 aliafshar
    aliafshar 57

    $ ghcontribs aliafshar rakyll
    aliafshar 111
    rakyll 964
