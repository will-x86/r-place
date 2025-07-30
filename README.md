This is the code for an r-place clone that can only be used via code:) 


Elements:
cmd/ -- entrypoint for golang 
pkg/ -- packages for golang 
ui/ -- React userinterface for requests
curl/ -- Basic curl commands for interfacing
.air.toml -- hot reloading



# Cache:
- K:V are stored in appendonly valkey db
- On startup all key:pair's are fetched up to xmax and ymax ( see .env.example ) 
- Two caches are stored, both the fully marshalled 
