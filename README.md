# go-kata

# Getting Started

## Setup

* Set up Go.
* Create directory `$GOPATH/src/github.com/anniehedgpeth`
* Run `git clone https://github.com/anniehedgpeth/go-kata.git`

## Every time you start the kata

* Close slack and email
* Set a timer for 30 minutes
* Create a new branch `20190104-annie`
* Check into that branch
* Never merge to master
* When the timer goes off, quit

# Kata

## Basics

_After every step run the program and see the result_
1. Create a program in the `main` package (and folder) that outputs `Hello, World!`
2. Extract `World` into a local variable `name`
3. Change the name to your name. It should now say, for example `Hello, Annie!`
4. Extract this into a method called `sayHi` that takes a name parameter and passes nothing.
5. Make the variable name a package variable
6. If the name is literally `Annie` add `Your name is like the owner of this repo!` to the end of the message
7. Extract logic from previous step to `isOwner` method returning a `bool`
8. After the greeting, output a list of the names of members of your immediate family, using the classic `for i := 0 ...` loop, where the output contains the index of the family member, for example: `0: Annie` followed by `1: Michael`
9. Refactor this loop to use the `range` keyword
10. Drop the index and refactor this to use the `range` keyword omitting the index
11. Change your list to include the favorite color of each family member by their name. Output it like this: `Annie: green` followed by `Michael: blue`. When looping continue to use the `range` keyword.
12. Refactor this into a method called `sayFavoriteColorsOfFamily`
13. Create a new list of family members with their day of the month number that they were born. Then output these, for example: `John: 22` followed by `Peter: 5`.
14. Refactor this into a method called `sayBirthdaysOfFamily`. At this point, your `main` method should be like this:
```golang
func main() {
  sayHi()
  sayFavoriteColorsOfFamily()
  sayBirthdaysOfFamily()
}
```
15. Create another list that tracks the weekly allowance of each family member and output this: `Michael: $20.50` - make sure you use the cents place for your values.
16. Refactor this into a method `sayAllowanceOfFamily`
17. Create a `person` struct that contains the person's name, favorite color, birthday number, and allowance. Refactor various lists into one global list with all of this information, and ensure your methods still work. 
18. Migrate your methods to a `greeting` package and call that  package from your `main` package
19. Whenever you talk about your family in the above methods, refer to them in alphabetical order (make sure your original list is _not_ in alphabetical order)
20. Whenever you refer to your family by name, refer to them by all-caps, like `MICHAEL: blue`
21. Put the `SayHi` method and associated variable and the family-oriented greetings in different files in the same `greeting` package

# Intermediate

_Listing the topic each step covers to facilitate easy learning/lookup._

1. Remove method calls from `main()` method
2. (REST Server) Create a `webserver` package, that when calling `http://localhost/sayHi` it returns the results of the `SayHi` method in the `greeting` package
2. (API Client) Change your `main` package to call the api endpoint instead of the method directly. Output the results
3. Create an endpoint for `sayFavoriteColorsOfFamily` and call it from `main`
4. Create an endpoint for `sayBirthdaysOfFamily` and call it from `main`
5. Create an endpoint for `sayAllowanceOfFamily` and call it from `main`
6. (JSON Serialization) Create a `http://localhost/familyMembers` endpoint that returns back your family members, deserialized as `person` structs, and output them
7. (PUT API calls) Create an endpoint at `http://localhost/you` where a `PUT` allows you to change the name in memory every time you call `sayHi`. Note that the name is different.
8. (Errors) If someone sets the name to be blank, `setName` method should return an error and you should return an error HTTP response code
9. (Templates) Create an html endpoint at `http://localhost/report` that returns back an HTML formatted list of all family members
10. (Methods) Create `isRich()` method on `person` that will return true if their allowance is greater than `10.50`.
11. (Anonymous functions) create `familyMembers(where func(person member) bool)` method to filter family members
12. (Anonymous functions) create `http://localhost/richFamilyMembers` endpoint to output family members that are rich to HTML
13. (Interfaces) Create `pet` struct with the name of the pet and its favorite toy.
14. (Interfaces) Create `http://localhost/animals` endpoint which outputs all family members and pets in the family by their name, using an `animal` interface.
15. (Channels) Modify the `richFamilyMembers` filter to process whether each family member is rich concurrently.