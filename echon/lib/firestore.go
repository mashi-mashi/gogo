package lib

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

var basePath = "gogo/v1"

func AddDoc(path string, data map[string]interface{}) map[string]interface{} {
	ctx := context.Background()
	app := FirebaseApp(ctx)

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	fmt.Println("add", data)
	addData := merge(data, map[string]interface{}{
		"createdAt": firestore.ServerTimestamp,
		"updatedAt": firestore.ServerTimestamp,
		"deleted":   false,
	})
	_, _, err = client.Collection(basePath+path).Add(ctx, addData)

	if err != nil {
		fmt.Println(err)
	}

	return addData
}

func AddDocWithId(path string, id string, data map[string]interface{}) map[string]interface{} {
	ctx := context.Background()
	app := FirebaseApp(ctx)

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	addData := merge(data, map[string]interface{}{
		"createdAt": firestore.ServerTimestamp,
		"updatedAt": firestore.ServerTimestamp,
		"deleted":   false,
	})

	_, err = client.Collection(basePath+path).Doc(id).Set(ctx, addData, firestore.MergeAll)

	if err != nil {
		fmt.Println(err)
	}

	return addData
}

func SetDoc(path string, id string, data map[string]interface{}) map[string]interface{} {
	ctx := context.Background()
	app := FirebaseApp(ctx)

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	setData := merge(data, map[string]interface{}{
		"updatedAt": firestore.ServerTimestamp,
	})

	_, err = client.Collection(basePath+path).Doc(id).Set(ctx, setData, firestore.MergeAll)

	if err != nil {
		fmt.Println(err)
	}

	return setData
}

func merge(xs ...map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, x := range xs {
		for k, v := range x {
			result[k] = v
		}
	}
	return result
}
