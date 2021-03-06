package plugin

const Usage = ShortUsage + "\n" + LongUsage

const ShortUsage = `
   cf local stage   <name> [ (-b <name> | -b <URL> | -b <zip>)... ]
                           [ (-p <dir> | -p <zip>) (-d <dir> [-r]) -e ]
                           [ (-s <app> | -f <app>) ]
   cf local run     <name> [ (-i <ip>) (-p <port>) (-s <app>) (-f <app>) ]
                           [ (-d <dir> [-r -w] | (-d <dir> [-r]) [-t]) ]
   cf local export  <name> [ (-r <ref>) ]
   cf local pull    <name>
   cf local push    <name> [-e -k]
   cf local help
   cf local version`

const LongUsage = `
STAGE OPTIONS:
   stage <name>   Build a droplet using the app in the current directory and
                     the environment variables and service bindings specified
                     in local.yml.
                     Droplet filename: <name>.droplet

   -b <name>      Use one or more official CF buildpacks (specified by name).
                     Default: (uses detection)
   -b <url>       Use one or more buildpacks specified by git repository URL
                     or zip file URL (HTTP or HTTPS).
                     Default: (uses detection)
   -b <zip>       Use one or more buildpacks specified by local zip file path.
                     Default: (uses detection)
   -p <dir>       Use the specified directory as the app directory.
                     Default: current working directory
   -p <zip>       Use the specified ZIP file contents as the app directory.
                     Note that JAR and WAR files use ZIP file format.
                     Default: current working directory
   -d <dir>       Mount the provided directory into the app root before
                     staging. Should be used with a VCS to track changes.
                     Ideal for use with: cf local run <name> -d <dir>
                     Use of -r is STRONGLY RECOMMENDED to avoid staging with
                     files that should be ignored.
                     Default: (not mounted)
   -r             When used with -d, rsync any files that were created,
                     modified, or moved during staging into the specified
                     directory. The directory is mounted elsewhere in the
                     container. No files are deleted.
                     Default: false
   -e             If buildpacks are explicitly specified then select one of
                     them using the buildpack detection process instead of
                     applying all of them using the multi-buildpack process.
                     Default: false
   -s <app>       Use the service bindings from the specified remote CF app
                     instead of the service bindings in local.yml.
                     Default: (uses local.yml)
   -f <app>       Same as -s, but re-writes the service bindings to match
                     what they would be if they were tunneled through the app
                     with: cf local run <name> -f <app>
                     Default: (uses local.yml)

RUN OPTIONS:
   run <name>     Run a droplet with the configuration specified in local.yml.
                     Droplet filename: <name>.droplet

   -i <ip>        Listen on the specified interface IP
                     Default: localhost
   -p <port>      Listen on the specified port
                     Default: (arbitrary free port)
   -d <dir>       Mount the specified directory into the app root.
                     If empty, the app root is copied into the directory.
                     If not empty, the app root is replaced by the directory.
                     Default: (not mounted)
   -r             When used with -d, rsync the contents of the specified
                     directory into the container app root. The directory is
                     mounted elsewhere in the container. No files are deleted.
                     Default: false, Invalid: without -d
   -w             When used with -d, restart the app when the contents of the
                     specified directory are changed.
                     Default: false, Invalid: with -t, without -d
   -t             Start a shell (Bash) with the same environment as the app.
                     Default: false, Invalid: with -w
   -s <app>       Use the service bindings from the specified remote CF app
                     instead of the service bindings in local.yml.
                     Default: (uses local.yml or app provided by -f)
   -f <app>       Tunnel service connections through the specified remote CF
                     app. This re-writes the service bindings in the container
                     environment in order to use the tunnel. The service
                     bindings from the specified app will be used if -s is not
                     also passed.
                     Default: (uses local.yml)

EXPORT OPTIONS:
   export <name>  Export a standalone Docker image using the specified droplet
                     and configuration from local.yml.
                     Droplet filename: <name>.droplet

   -r <ref>       Tag the exported image with the provided reference.
                     Default: none

PULL OPTIONS:
   pull <name>    Download the droplet, environment variables, environment
                     variable groups, and start command of the named remote
                     CF app. The local.yml file is updated with the downloaded
                     configuration.
                     Droplet filename: <name>.droplet

PUSH OPTIONS:
   push <name>    Push a droplet to a remote CF app and restart the app.
                     Droplet filename: <name>.droplet

   -e             Additionally replace the remote app environment variables
                     with the environment variables from local.yml. This does
                     not read or replace environment variable groups.
                     Default: false
   -k             Do not restart the application after pushing the droplet.
                     The current droplet will continue to run until the next
                     restart.
                     Default: false

ENVIRONMENT:
   DOCKER_HOST    Docker daemon address
                     Default: /var/run/docker.sock

SAMPLE: local.yml

applications:
- name: first-app
  buildpacks:
  - some_buildpack
  - some_other_buildpack
  command: "some start command"
  memory: 2G
  disk_quota: 4G
  staging_env:
    SOME_STAGING_VAR: "some staging value"
  running_env:
    SOME_RUNNING_VAR: "some running value"
  env:
    SOME_VAR: "some value"
  services:
    (( VCAP_SERVICES object in YAML ))
`
