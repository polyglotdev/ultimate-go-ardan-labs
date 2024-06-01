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

Software engineering might be the only field that asks you to write before you were taught to read.

**Quotes**

> “If most computer people lack understanding and knowledge, **then what they will select will also be lacking**.” - Alan Kay

> "The software business is one of the few places we teach people to write before we teach them to read." - Tom Love (inventor of Objective C)

> "Code is read many more times than it is written." - Dave Cheney

> "Programming is, among other things, **a kind of writing**. One way to learn writing is to write, **but in all other forms of writing, one also reads**. We read examples both good and bad to facilitate learning. But how many programmers learn to write programs by reading programs?" - Gerald M. Weinberg

> "Skill develops when we produce, not consume." - Katrina Owen


We are reaching a point in the field where we have a lot of legacy code, and that could be a problem.


**Quotes**

> "There are two kinds of software projects: those that fail, and those that turn into legacy horrors." - Peter Weinberger (inventor of AWK)

> "Legacy software is an unappreciated but serious problem. Legacy code may be the downfall of our civilization." - Chuck Moore (inventor of Forth)

> "Few programmers of any experience would contradict the assertion that most programs are modified in their lifetime. Why then do we rarely find a program that contains any evidence of having been written with an eye to subsequent modification." - Gerald M. Weinberg

> "We think awful code is written by awful devs. But in reality, it's written by reasonable devs in awful circumstances." - Sarah Mei

> "There are many reasons why programs are built the way they are, although we may fail to recognize the multiplicity of reasons because we usually look at code from the outside rather than by reading it. When we do read code, we find that some of it gets written because of machine limitations, some because of language limitations, some because of programmer limitations, some because of historical accidents, and some because of specifications—both essential and inessential." - Gerald M. Weinberg
---

## Mental Models

You must constantly make sure your mental model of the code you are writing and maintaining is clear. When you can't remember where a piece of logic is or you can't remember how something works, you’re losing your mental model of the code. This is a clear indication that you need to refactor the code. Focus time on structuring code that provides the best mental model possible and during code reviews validate your mental models are still intact.

How much code do you think you can maintain in your head? I believe asking a single developer to maintain a mental model of more than one ream of copy paper (~10k lines of code) is asking a lot. If you do the math, it takes a team of 100 people to work on a code base that hits a million lines of code. That’s 100 people that need to be coordinated, grouped, tracked and in a constant feedback loop of communication.

**Quotes**

> "Let's imagine a project that's going to end up with a million lines of code or more. The probability of those projects being successful in the United States these days is very low - well under 50%. That's debatable." - Tom Love (inventor of Objective C)_

> "100k lines of code fit inside a box of paper." - Tom Love (inventor of Objective C)_

> "One of our many problems with thinking is “cognitive load”: the number of things we can pay attention to at once. The cliche is 7±2, but for many things it is even less. We make progress by making those few things be more powerful." - Alan Kay

> "The hardest bugs are those where your mental model of the situation is just wrong, so you can't see the problem at all." - Brian Kernighan

> "Everyone knows that debugging is twice as hard as writing a program in the first place. So if you're as clever as you can be when you write it, how will you ever debug it?" - Brian Kernighan

> "Debuggers don't remove bugs. They only show them in slow motion." - Unknown

> "Fixing bugs is just a side effect. Debuggers are for exploration." - @Deech (Twitter)

### Mental Model Resources

- [The Magical Number Seven, Plus or Minus Two](https://en.wikipedia.org/wiki/The_Magical_Number_Seven,_Plus_or_Minus_Two)
- [Psychology of Code Readability](https://medium.com/@egonelbre/psychology-of-code-readability-d23b1ff1258a)
---

## Productivity vs Performance

Productivity and performance both matter, but in the past you couldn’t have both. You needed to choose one over the other. We naturally gravitated to productivity, with the idea or hope that the hardware would resolve our performance problems for free. This movement towards productivity has resulted in the design of programming languages that produce sluggish software that is outpacing the hardware’s ability to make them faster.

By following Go’s idioms and a few guidelines, we can write code that can be reasoned about by average developers. We can write software that simplifies, minimizes and reduces the amount of code we need to write to solve the problems we are working on. We don’t have to choose productivity over performance or performance over productivity anymore. We can have both.

**Quotes**

> "The hope is that the progress in hardware will cure all software ills. However, a critical observer may observe that software manages to outgrow hardware in size and sluggishness. Other observers had noted this for some time before, indeed the trend was becoming obvious as early as 1987." - Niklaus Wirth

> "The most amazing achievement of the computer software industry is its continuing cancellation of the steady and staggering gains made by the computer hardware industry." - Henry Petroski (2015)

> "The hardware folks will not put more cores into their hardware if the software isn’t going to use them, so, it is this balancing act of each other staring at each other, and we are hoping that Go is going to break through on the software side.” - Rick Hudson (2015)

> "C is the best balance I've ever seen between power and expressiveness. You can do almost anything you want to do by programming fairly straightforwardly and you will have a very good mental model of what's going to happen on the machine; you can predict reasonably well how quickly it's going to run, you understand what's going on .... - Brian Kernighan (2000)

> "The trend in programming language design has been to create languages that enhance software reliability and programmer productivity. What we should do is develop languages alongside sound software engineering practices so the task of developing reliable programs is distributed throughout the software lifecycle, especially into the early phases of system design." - Al Aho (2009)
---
## Correctness vs Performance

You want to write code that is optimized for correctness. Don't make coding decisions based on what you think might perform better. You must benchmark or profile to know if code is not fast enough. Then and only then should you optimize for performance. This can't be done until you have something working.

Improvement comes from writing code and thinking about the code you write. Then refactoring the code to make it better. This requires the help of other people to also read the code you are writing. Prototype ideas first to validate them. Try different approaches or ask others to attempt a solution. Then compare what you have learned.

Too many developers are not prototyping their ideas first before writing production code. It’s through prototyping that you can validate your thoughts, ideas and designs. This is the time when you can break down walls and figure out how things work. Prototype in the concrete and consider contracts after you have a working prototype.

Refactoring must become part of the development cycle. Refactoring is the process of improving the code from the things that you learn on a daily basis. Without time to refactor, code will become impossible to manage and maintain over time. This creates the legacy issues we are seeing today.

**Quotes**

> "Make it correct, make it clear, make it concise, make it fast. In that order." - Wes Dyer

> "Good engineering is less about finding the "perfect" solution and more about understanding the tradeoffs and being able to explain them." - JBD

> "Choosing the right limitations for a certain problem domain is often much more powerful than allowing anything." - Jason Moiron_

> "The correctness of the implementation is the most important concern, but there is no royal road to correctness. It involves diverse tasks such as thinking of invariants, testing and code reviews. Optimization should be done, but not prematurely." - Al Aho (inventor of AWK)

> "The basic ideas of good style, which are fundamental to write clearly and simply, are just as important now as they were 35 years ago. Simple, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have clean, simple code." - Brian Kernighan

> "Problems can usually be solved with simple, mundane solutions. That means there's no glamorous work. You don't get to show off your amazing skills. You just build something that gets the job done and then move on. This approach may not earn you oohs and aahs, but it lets you get on with it." - Jason Fried

**Resources:**

- [Prototype your design!](https://www.youtube.com/watch?v=vLxX3yZmw5Q) - Robert Griesemer

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