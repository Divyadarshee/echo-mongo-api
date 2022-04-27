### Source: https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-echo-version-2gdg


Folders and their uses:
0. In general folders help in modularizing the project
1. configs: is for modularizing project configuration files
2. contorllers: is for modularizing application logics
3. models: is for modularizing data and database logics
4. responses: is for modularizing files describing the response we wnat our API to give
5. routes: is for modularizing URL pattern and handler information


# Setup Environment Variable
Next, we must modify the copied connection string with the user's password we created earlier and change the database name. To do this, first, we need to create a .env file in the root directory, and in this file, add the snippet below:

```
 mongodb+srv://dvd:dvd6900@cluster0.huu8r.mongodb.net/golangdb?retryWrites=true&w=majority
``` 

# Load Environment Variable
 With that done, we need to create a helper function to load the environment variable using the github.com/joho/godotenv library we installed earlier. To do this, we need to navigate to the configs folder and in this folder, create an env.go file

# env.go does the following:
Import the required dependencies.
Create an EnvMongoURI function that checks if the environment variable is correctly loaded and returns the environment variable.

# Connecting to MongoDB
To connect to the MongoDB database from our application, first we need to navigate to the configs folder and in this folder, create a setup.go file

# setup.go does the following:
Import the required dependencies.
Create a ConnectDB function that first configures the client to use the correct URI and check for errors. Secondly, we defined a timeout of 10 seconds we wanted to use when trying to connect. Thirdly, check if there is an error while connecting to the database and cancel the connection if the connecting period exceeds 10 seconds. Finally, we pinged the database to test our connection and returned the client instance.
Create a DB variable instance of the ConnectDB. This will come in handy when creating collections.
Create a GetCollection function to retrieve and create collections on the database.

# Setup API Route Handler and Response Type
Route Handler
With that done, we need to create a user_route.go file inside the routes folder to manage all the user-related routes in our application

# Response Type
Next, we need to create a reusable struct to describe our API’s response. To do this, navigate to the responses folder and in this folder, create a user_response.go file

# user_response.go does the following:
creates a UserResponse struct with Status, Message, and Data property to represent the API response type.

** PS: json:"status", json:"message", and json:"data" are known as struct tags. Struct tags allow us to attach meta-information to corresponding struct properties. In other words, we use them to reformat the JSON response returned by the API.

# Finally, Creating REST API’s
Next, we need a model to represent our application data. To do this, we need to navigate to the models folder, and in this folder, create a user_model.go file 

# user_model.go does the following:
Import the required dependencies.
Create a User struct with required properties. We added omitempty and validate:"required" to the struct tag to tell Fiber to ignore empty fields and make the field required, respectively.

# Create a User Endpoint
With the model setup, we can now create a function to create a user. To do this, we need to navigate to the controllers folder, and in this folder, create a user_controller.go file

# createuser in user_controller.go does the following
Import the required dependencies.
Create userCollection and validate variables to create a collection and validate models using the github.com/go-playground/validator/v10 library we installed earlier on, respectively.
Create a CreateUser function that returns an error. Inside the function, we first defined a timeout of 10 seconds when inserting user into the document, validating both the request body and required field using the validator library. We returned the appropriate message and status code using the UserResponse struct we created earlier. Secondly, we created a newUser variable, inserted it using the userCollection.InsertOne function and check for errors if there are any. Finally, we returned the correct response if the insert was successful.

# getauser in user_controller.go does the following:
Import the required dependencies.
Create a GetAUser function that returns an error. Inside the function, we first defined a timeout of 10 seconds when finding a user in the document, a userId variable to get the user’s id from the URL parameter and a user variable. We converted the userId from a string to a primitive.ObjectID type, a BSON type MongoDB uses. Secondly, we searched for the user using the userCollection.FindOne, pass the objId as a filter and use the Decode attribute method to get the corresponding object. Finally, we returned the decoded response.

# editauser in user_controller.go does the following:
The EditAUser function above does the same thing as the CreateUser function. However, we included an update variable to get updated fields and updated the collection using the userCollection.UpdateOne. Lastly, we searched for the updated user’s details and returned the decoded response.

# deleteauser in user_controller.go does the following:
The DeleteAUser function follows the previous steps by deleting the matched record using the userCollection.DeleteOne. We also checked if an item was successfully deleted and returned the appropriate response.

