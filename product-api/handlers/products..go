package handlers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/nicholasjackson/building-microservices-youtube/product-api/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p := r.URL.Path
		r: = regexp.MustCompile(`/([0-9]+)`)
		r.FindAllStringSubmatch()
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Product: %#v", prod)
	data.AddProduct(prod)
}
