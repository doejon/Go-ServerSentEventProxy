# Golang Server Sent Events proxy
Proxying Server Sent Events (javascript's 'EventSource') I could not figure out
why golang's default proxy would not forward any server sent event flushes.
When looking into golang's proxy implementation realized proxy does not
flush response body by default - but does offer a FlushInterval property.

This tiny library wraps the default http reverse proxy which works with SSE by
setting flush interval to 100ms;

To get a feeling for SSE I'd recommend https://github.com/kljensen/golang-html5-sse-example

License: MIT
