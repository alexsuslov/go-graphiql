package graphiql

import (
	"fmt"
	"net/http"
)


var (
	Endpoint = "/graphql"
	Content = `<!DOCTYPE html>
<head><style>
    body {
        height: 100vh;
        margin: 0;
        width: 100%;
        overflow: hidden;
    }
    </style><link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/sweetalert/1.1.3/sweetalert.min.css><script src=https://cdnjs.cloudflare.com/ajax/libs/sweetalert/1.1.3/sweetalert.min.js></script><script src=https://cdnjs.cloudflare.com/ajax/libs/fetch/3.0.0/fetch.min.js></script><script src=https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js></script><script src=https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js></script><link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.12.0/graphiql.min.css><script src=https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.12.0/graphiql.min.js></script><script>
    (function() {
		const inputPlaceholder = window.location.origin + '%s';
        var PROMPT_OPTIONS = {
            title: "GraphQL Endpoint",
            text: "Please give the GraphQL HTTP Endpoint",
            type: "input",
            showCancelButton: false,
            inputPlaceholder,
        };
        document.addEventListener('DOMContentLoaded', function() {
            swal(PROMPT_OPTIONS, function(endpoint) {
                if (!endpoint) {
                    endpoint = inputPlaceholder;
                }

                function fetcher(params) {
                    var options = {
                        method: 'post',
                        headers: { 'Accept': 'application/json', 'Content-Type': 'application/json' },
                        body: JSON.stringify(params),
                        credentials: 'include',
                    };
                    return fetch(endpoint, options)
                        .then(function(res) { return res.json() });
                }
                var body = React.createElement(GraphiQL, { fetcher: fetcher, query: '', variables: '' });
                ReactDOM.render(body, document.body);
            });
        });
    }());
    </script></head> <body> </body> </html>`)


// ServeGraphiQL is a handler function for HTTP servers
func ServeGraphiQL(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(fmt.Sprintf(Content, Endpoint)))
}
