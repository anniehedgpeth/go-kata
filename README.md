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

Start off using [Visual Studio Code](https://code.visualstudio.com/)

Graduate to [Go Land](https://www.jetbrains.com/go/) and use the [refactorings](https://www.jetbrains.com/help/go/discover-goland.html#RefactoringBasics) to speed up. Only use the keyboard.

# Background
Use this background information for the kata.

Simpson family:
---------------
| name | birthday | favorite color | allowance |
|---|---|---|---|
| Homer| 22-September | 10.50 |
| Marge | 14-February | red | 20.34 |
| Lisa | 4-July | purple | 15.50 |
| Bart | 1-June | black | 5.50 |
| Maggie | 1-April | white | 6.50 |

# Kata

## Basics

This gets you through most of the basic elements of golang.

_After every step run the program and see the result_
1. Create a program in the `main` package (and folder) that outputs `Hello, World!`
2. (Variables) Extract `World` into a local variable `name`
3. (Variables) Change the name to your name. It should now say, for example `Hello, Annie!`
4. (Functions) Extract the greeting into a function called `sayHi` that takes a name parameter and returns nothing.
5. (Packages) Make the variable name a package variable
6. (If statement) In the `sayHi` function, if the name is literally `Annie` add `Your name is like the owner of this repo!` to the end of the greeting
7. (Functions returning value) Extract whether name is literally `Annie` expression to an `isOwner` method returning a `bool`
8. (Loops) After the greeting (`sayHi`), in the `main` function output a list of the names of members of the Simpson family (above), using the classic `for i := 0 ...` loop, where the output contains the index of the family member, for example: `0: Marge` followed by `1: Homer`
9. (Loops) Refactor this loop to use the `range` keyword
10. (Loops) Drop the index variable and refactor this to use the `range` keyword omitting the index. At this point you will output _only_ the names, for example, `Homer` then `Marge`
11. (Maps) Change your list to include the favorite color of each family member by their name (see above). Output it like this: `Marge: red` followed by `Homer: blue`. The ordering doesn't matter. When looping continue to use the `range` keyword.
12. (Functions) Refactor this into a method called `sayFavoriteColors` that outputs the greeting
13. (Maps) Create a new list of family members with their day of the month number that they were born. Then output these, for example: `Homer: 22` followed by `Bart: 5`. (these numbers can be fictional)
14. (Functions) Refactor this into a method called `sayBirthdays`. At this point, your `main` method should be like this:
```golang
func main() {
  sayHi()
  sayFavoriteColors()
  sayBirthdays()
}
```
15. (Maps) Create another list that tracks the weekly allowance of each family member and output this: `Michael: $20.50` - make sure you use the cents place for your values.
16. (Functions) Refactor this into a method `sayAllowance`
17. (Structs) Create a `person` struct that contains the person's name, favorite color, birthday number, and allowance. Refactor various lists into one global list with all of this information, and ensure your methods still work. 
18. (Packages) Migrate your functions to a `greeting` package and call that  package from your `main` package
19. (Sort) Whenever you talk about the family in the above methods, refer to them in alphabetical order (make sure your original list is _not_ in alphabetical order)
20. (Imports, Methods) Whenever you refer to your family by name, refer to them by all-caps, like `HOMER: blue`
21. (Packages) Put the `SayHi` method and associated variable and the family-oriented greetings in different files in the same `greeting` package

# Intermediate

Here we get into files and advanced object functions.

1. (File/IO) Load in the information for the family from the comma-delimited list defined in `simpsons.csv`
2. (File/IO) Load which family is used based on an argument to the program. For example if you run `simpsons.csv` the program should use the simpsons data. If the program has the argument `jacksons.csv` then load the jacksons data.
3. (Channels, Mutex) If the program contains _both_ files (`simpsons.csv jacksons.csv`), load them concurrently into one list. Make sure the list isn't accessed at the same time.
4. (Structs, File/IO, `flag` package) If you pass in `pets pets.csv`, it should output the pets and their favorite toy. For example: `Fido: bone` followed by `Argyle: yarn`
4. (Interfaces) If you pass in `everyone simpsons.csv pets.csv`, it should output the people and pets as animals. For example: `Homer,Marge,Bart,Lisa,Maggie,Fido,Argyle`
5. (Methods) Create `isRich()` method on `person` that will return true if their allowance is greater than `10.50`.
11. (Anonymous functions) create `familyMembers(where func(person member) bool)` method to filter family members
12. (Anonymous functions) if you pass in `rich simpsons.csv`, output family members defined in the CSV that are rich

# Advanced

This extends what you learned above to learn how to create a webserver API and client in golang.

1. Remove method calls from `main()` method
2. (REST Server) Create a `webserver` package, that when calling `http://localhost/sayHi` it returns the results of the `SayHi` method in the `greeting` package
3. Run your webserver with the `server` argument
4. When passing in `webclient`, call the web server at `http://localhost/sayHi` and output the contents of the resulting message.
5. Create an endpoint for `sayFavoriteColors` and call it from `main` when `webclient` flag is given
6. Create an endpoint for `sayBirthdays` and call it from `main` when `webclient` flag is given
5. Create an endpoint for `sayAllowance` and call it from `main` when `webclient` flag is given
6. (JSON Serialization) Create a `http://localhost/familyMembers` endpoint that returns back your family members, deserialized as `person` structs, and output them
7. (PUT API calls) Create an endpoint at `http://localhost/you` where a `PUT` allows you to change the name in memory every time you call `sayHi`. Note that the name is different.
8. (Errors) If someone sets the name to be blank, `setName` method should return an error and you should return an error HTTP response code
9. When starting the server pass in the csv file to load for the family, like: `server simpsons.csv`
10. (Templates) Create an html endpoint at `http://localhost/report` that returns back an HTML formatted list of all family members
