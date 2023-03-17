# Simple GO HTTP Server

This server implements:
 * graceful exit that does the following:
    * closes HTTP server after all connections are handled
    * closes all dependency connections
    * closes the process
 * a way to pass modules to all handlers 

The server currenty implements the following modules:
 * Redis


The required environment variables are the following:
```ini
PORT=8081

REDIS_CLIENT_ADDRESS="host.docker.internal:6379"
REDIS_CLIENT_USERNAME=""
REDIS_CLIENT_PASSWORD=""
REDIS_CLIENT_DB_INDEX=0
REDIS_CLIENT_CONN_NAME="Redis-Client"
```