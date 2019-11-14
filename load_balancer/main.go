u, _ := url.Parse("http://localhost:8080")
rp := httputil.NewSingleHostReverseProxy(u)

http.HandlerFunc(rp.ServerHTTP)


func lb(w http.ResponseWriter, r *http.Request) {
	peer := serverPool.GetNextPeer()
	if  nil != peer {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not Avilable", http.StatusServiceUnavailable)
}

server := http.Server {
	Addr:		fmt.Sprintf(":%d", port),
	Handler:	http.handlerFunc(lb),
}

proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
	
}