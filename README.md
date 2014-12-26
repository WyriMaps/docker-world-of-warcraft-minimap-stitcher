# Minimap sticher docker container ##

Ready to go docker container to stitcher the World of Warcraft minimaps.

## Usage ##

Before using this container be sure to have the minimap tile extracted and converted from `BLP` to `PNG`. 

```sh
docker run -v /path/to/source/pngs:/opt/map_tiles -v /path/to/output/directory:/opt/maps golang-test-stitcher
```
