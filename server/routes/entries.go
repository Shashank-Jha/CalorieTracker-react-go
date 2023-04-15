package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var entryCollection *mongo.Collection = openCollection(Client,"calories")


func AddEntry(c *gin.Context){

}

func GetEntries(c *gin.Context){
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx,bson.M{})
	if(err != nil){
		c.JSON(http.StatusInternalServer, gin.H{"error":err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != {
		c.JSON(http.StatusInternalServer,gin.H{"error":err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntriesByIngredient(c *gin.Context){

}

func GetEntryById(c *gin.Context){

}

func UpdateIngredient(c *gin.Context){

}

func UpdateEntry(c *gin.Context){

}
