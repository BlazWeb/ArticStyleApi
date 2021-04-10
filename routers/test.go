package routers

import ( 
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Rest Api ArticStyles. Att: ArticDev")
}