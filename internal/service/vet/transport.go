// This allows the service to be acceded
// via HTTP requests (REST)

package vet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ..
type HTTPService interface{
	Register(*gin.Engine)
}

type endpoint struct{
	method string
	path string
	function gin.HandlerFunc
}

type httpService struct{
	endpoints []*endpoint
}

// NewHTTPTransport ..
func NewHTTPTransport(s Service) HTTPService{
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

// Register ..
func (s httpService) Register(r *gin.Engine){
	for _, e := range s.endpoints{
		r.Handle(e.method, e.path, e.function)
	}
} 


func makeEndpoints (s Service) []*endpoint{
	list := []*endpoint{}
	list = append(list, &endpoint{
		method: "GET",
		path: "/animals",
		function: getAll(s),
	})
	list = append(list, &endpoint{
		method: "GET",
		path: "/animal/:id",
		function: getOne(s),
	})
	list = append(list, &endpoint{
		method: "POST",
		path: "/animal",
		function: addOne(s),
	})
	list = append(list, &endpoint{
		method: "DELETE",
		path: "/animal/:id",
		function: deleteOne(s),
	})
	list = append(list, &endpoint{
		method: "PUT",
		path: "/animal/:id",
		function: replaceOne(s),
	})
	return list
}



func getAll(s Service) gin.HandlerFunc{
	return func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"Animals": s.FindAll(),
		})
	}
}
func getOne(s Service) gin.HandlerFunc{
	return func (c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		c.JSON(http.StatusOK, gin.H{
			"Animal": s.FindByID(id),
		})
	}
}
func addOne(s Service) gin.HandlerFunc{
	return func (c *gin.Context){
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		var animal Animal
		if err = json.Unmarshal(body, &animal); err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		s.AddAnimal(animal)
		c.JSON(http.StatusOK, gin.H{
			"Animal": "Successfuly added",
		})
	}
}
func deleteOne(s Service) gin.HandlerFunc{
	return func (c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		s.DeleteAnimal(id)
		c.JSON(http.StatusOK, gin.H{
			"Animal": "Successfuly deleted (ID: "+c.Param("id")+")",
		})
	}
}
func replaceOne(s Service) gin.HandlerFunc{
	return func (c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		var animal Animal
		if err = json.Unmarshal(body, &animal); err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		s.ReplaceAnimal(id, animal)
		c.JSON(http.StatusOK, gin.H{
			"Animal": "Successfuly updated (ID: "+c.Param("id")+")",
		})
	}
}