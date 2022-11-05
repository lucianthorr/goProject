# A really simple Go project

I've been working on Go projects for a few years and I have found that they have all started to converge into a single pattern.  I thought it might be cool to boil it down to the simplest possible example and then maybe get some feedback.

## How to run

```
export DB_PASSWORD=123; go run ./cmd/app
```

## The `configs` folder/package

This probably violates more Go project-structure recommandation than I can think up but the `configs/` folder is used to store configuration files and has a really simple `configs` package that will read the files and unmarshall them into each package's Config struct.  Typically, Go packages should be singular but I like plural because it denotes that the folder is storing config files and we don't have to have a `config` package folder and a `configs` folder for config files.  That just seems like too much.

And the `Config` struct is really clean.
```
type Config struct {
	API *api.Config `yaml:"api"`
	DB  *db.Config  `yaml:"db"`
    // ... any other packages in the project
}
```

## Each package

In this example, I'm just using `api` and `db` as example packages.  There's nothing really implemented but you'll see that there's mostly structured the same.

Each package has a .go file with the same name as the package.  This is to draw attention to the fact that it is the point of entry when there might be lots of files in a package.  And inside each is the package's own `Config` struct and a `Client` interface that defines all the externally available functions for that package.  The package will then define a private] implementation of that `Client` called `client`.

Then each package has a `New` function that returns a real implementation of that interface.  Each `New` should take all other dependencies in as Clients from the other packages.  This makes fakes/mocks/stubs super easy during unit testing because you can create your own private `client` for the package you're testing and fake out all the other package's `Client` interfaces.

## Commands

It's pretty common among Go projects, especially that generate more than one binary, to use the `cmd/` folder to store each command.  For this example, I just have one uncreative application called "app."


Here, we start things by creating the single `config` which contains each packages Config.  After that, generate a Client for each package that's needed for that command and wire them together with `New` and then start whatever actual logic needs to be performed.

## Unit tests

In [api_test.go](api/api_test.go), I have an example of using this design to fake the `db` package and test the API's [GetThing](api/api.go#51) endpoint.  It's pretty straightforward to both test that the API request is processed correctly, the DB is receiving the expected input and that the API is returning the expected output.  We can also test that errors are handled correctly at each point along the way.

## Caveats and other notes.

Regarding the API, I have mostly been using [gin](https://github.com/gin-gonic/gin) but I don't want to be prescriptive so I'm keeping it simple with the standard library here.

There's probably a few better ways to handle secrets.  Here




