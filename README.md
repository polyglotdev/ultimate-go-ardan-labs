# Ultimate Go

## [Prepare your 🧠](https://github.com/ardanlabs/gotraining/tree/master/topics/go#prepare-your-mind)

> If you do not agree with the idioms and patterns of a given language, you should not be using that language because you are just going to be frustrated. - William Kennedy

I want to take the above statement and expound on it. If you are not willing to let Go be Go, you will inevitably find yourself pissed off. And I think that same goes for language that you are using. What my humble suggestion is that you see the beauty in all programming languages, yes, even the ones that you do not agree with. The language you choose should be the one that is the right tool for the job. Remembering that will make you a better engineer.

**Somewhere Along The Line**

- We became impressed with programs that contain large amounts of code.
- We strived to create large abstractions in our code base.
- We forgot that the hardware is the platform.
- We lost the understanding that every decision comes with a cost.

**These Days Are Gone**

- We can throw more hardware at the problem.
- We can throw more developers at the problem.

**Open Your Mind**

- Technology changes quickly but people's minds change slowly.
- Easy to adopt new technology but hard to adopt new ways of thinking.

**Interesting Questions - What do they mean to you?**

- Is it a good program?
- Is it an efficient program?
- Is it correct?
- Was it done on time?
- What did it cost?

**Aspire To**

- Be a champion for quality, efficiency and simplicity.
- Have a point of view.
- Value introspection and self-review.

---
## Polymorphism - 4.2.1

> "Polymorphism means that you write a certain program and it behaves differently depending on the data that it operates on."-
Tom Kurtz (inventor of BASIC)

```go
// reader is an interface that defines the act of reading data.
type reader interface {
  read(b []byte) (int, error)
}

// file defines a system file.
type file struct {
  name string
}
```

`var r reader` declares a variable r of type reader. The interface type reader is a contract that states that any concrete type that implements the read method is also of type reader. The reason this is a code smell is because interface types do not define concrete data, it defines a method set of behavior therefore they are valueless.

API design needs to be first simple to understand, not simple to do however. Lets take a look at a poorly implemented reader interface.

```go
// reader is an interface that defines the act of reading data.
type reader interface {
  read(b int) ([]byte, error)
}
```

Here on the field on reader we have a method read that takes an integer and returns a slice of bytes and an error. This is a poor design because the read method is not idiomatic to the io.Reader interface. Not to mention what its really saying is give me a number of bytes and I will return you slice of bytes. There is no guarantee that the entirety of what you passing is read because I am only going to read up to what you handed me. The io.Reader interface is defined as:

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

On the bad method in `reader` every time that its called there has to be an allocation on the heap. The first allocation comes from the construction of the slice and the length is set by `b`. The return at some point will have to be returned back up the call stack and when you do the pointer to the back array will no longer be on frame anymore.

We have to be consistent with our data semantics, but we must also be concerned with the impact the design is going to effect the users program. Designing API that are sympathetic with things like the garbage collector.

> Remember reader is a an interface type, which makes it a valueless type. And yes, you can declare a variable of the interface type `r` does not exist because we can neither read nor write to it. It is a valueless type.