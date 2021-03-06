# SearchDoc

Search documentation for a specific langauge.

Supports the same documentation formats that [Dash](http://kapeli.com/dash) and [Zeal](http://zealdocs.org) do.

## Build
To build the executable, simply run `make`.

## Usage

### List Downloaded Docsets
```
searchdoc --list
```

### Show Downloadable Docsets
```
searchdoc --download_list
```

### Download a new docset
```
searchdoc --download DOCSET_NAME
```

### Perform a query
```
searchdoc --language LANG --query QUERY
```

## Tips

It is recommended that you create an alias for a specific language rather than
specifying it everytime in the command line arguments.

# Inserting documentation
The application uses the same format as zeal and Dash. Look at https://kapeli.com/docsets#dashDocset for more info.
