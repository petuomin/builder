# Server bind address
bind_addr = "0.0.0.0:8080"

# Server URL - passed to builds t
url = http://builder.juhovuori.net

# List of projects to manage
projects = [
  "https://raw.githubusercontent.com/juhovuori/builder/master/project.hcl"
]

# State store configuration
state_store {
  type = "file"
  directory = "/tmp/builder"
}
