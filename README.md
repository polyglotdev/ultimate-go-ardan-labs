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

> "Let's imagine a project that's going to end up with a million lines of code or more. The probability of those projects being successful in the United States these days is very low - well under 50%. That's debatable." - Tom Love (inventor of Objective C)

> "100k lines of code fit inside a box of paper." - Tom Love (inventor of Objective C)

> "One of our many problems with thinking is “cognitive load”: the number of things we can pay attention to at once. The cliche is 7±2, but for many things it is even less. We make progress by making those few things be more powerful." - Alan Kay

> "The hardest bugs are those where your mental model of the situation is just wrong, so you can't see the problem at all." - Brian Kernighan

> "Everyone knows that debugging is twice as hard as writing a program in the first place. So if you're as clever as you can be when you write it, how will you ever debug it?" - Brian Kernighan

> "Debuggers don't remove bugs. They only show them in slow motion." - Unknown

> "Fixing bugs is just a side effect. Debuggers are for exploration." - @Deech (Twitter)

## Mental Model Resources

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

> "Choosing the right limitations for a certain problem domain is often much more powerful than allowing anything." - Jason Moiron

> "The correctness of the implementation is the most important concern, but there is no royal road to correctness. It involves diverse tasks such as thinking of invariants, testing and code reviews. Optimization should be done, but not prematurely." - Al Aho (inventor of AWK)

> "The basic ideas of good style, which are fundamental to write clearly and simply, are just as important now as they were 35 years ago. Simple, straightforward code is just plain easier to work with and less likely to have problems. As programs get bigger and more complicated, it's even more important to have clean, simple code." - Brian Kernighan

> "Problems can usually be solved with simple, mundane solutions. That means there's no glamorous work. You don't get to show off your amazing skills. You just build something that gets the job done and then move on. This approach may not earn you oohs and aahs, but it lets you get on with it." - Jason Fried

**Resources:**

- [Prototype your design!](https://www.youtube.com/watch?v=vLxX3yZmw5Q)
---
## Rules

- Rules have costs.
- Rules must pull their weight - Don’t be clever (high level).
- Value the standard, don’t idolize it.
- Be consistent!
- Semantics convey ownership.

**Quotes**

> "An architecture isn't a set of pieces, it's a set of rules about what you can expect of them." - Michael Feathers

**Resources:**

- [The Philosophy of Google's C++ Code](https://www.youtube.com/watch?v=NOCElcMcFik)
---
## Senior vs Junior Developers

What is the difference between a Senior and Junior developer?

**Quotes**

> "You are personally responsible for the software you write." - Stephen Bourne (Bourne shell)

> "And the difference between juniors+seniors to those who are in-between, is the confidence to ask "dumb" questions." - Natalie Pistunovich

> "Mistakes are an inevitable consequence of doing something new and, as such, should be seen as valuable; without them, we'd have no originality." - Ed Catmull (President of Pixar)

> "It takes considerable knowledge just to realize the extent of your own ignorance." - Thomas Sowell

> "If you don’t make mistakes, you’re not working on hard enough problems." - Frank Wilczek

> "Don’t cling to a mistake because you spent so much time making it." - Aubrey de Grey
---
## Design Philosophy

You can't look at a piece of code, function or algorithm and determine if it smells good or bad without a design philosophy. These four major categories are the basis for code reviews and should be prioritized in this order: Integrity, Readability, Simplicity and then Performance. You must consciously and with great reason be able to explain the category you are choosing.

---
## Integrity

**_We need to become very serious about reliability._**

There are two driving forces behind integrity:
* Integrity is about every allocation, read and write of memory being accurate, consistent and efficient. The type system is critical to making sure we have this `micro` level of integrity.
* Integrity is about every data transformation being accurate, consistent and efficient. Writing less code and error handling is critical to making sure we have this `macro` level of integrity.

**Write Less Code:**

There have been studies that have researched the number of bugs you can expect to have in your software. The industry average is around 15 to 50 bugs per 1000 lines of code. One simple way to reduce the number of bugs, and increase the integrity of your software, is to write less code.

Bjarne Stroustrup stated that writing more code than you need results in `Ugly`, `Large` and `Slow` code:

- `Ugly`: Leaves places for bugs to hide.
- `Large`: Ensures incomplete tests.
- `Slow`: Encourages the use of shortcuts and dirty tricks.

**Error Handling:**

When error handling is treated as an exception and not part of the main code path, you can expect the majority of your critical failures to be due to error handling.

There was a study that looked at a couple hundred bugs in Cassandra, HBase, HDFS, MapReduce, and Redis. The study identified 48 critical failures that fell into these categories.

- 92%: Failures from bad error handling
  - 35% : Incorrect handling
    - 25% : Simply ignoring an error
    - 8% : Catching the wrong exception
    - 2% : Incomplete TODOs
  - 57% System specific
    - 23% : Easily detectable
    - 34% : Complex bugs
  - 8% : Failures from latent human errors

**Quotes**

> "Failure is expected, failure is not an odd case. Design systems that help you identify failure. Design systems that can recover from failure." - JBD

> "Product excellence is the difference between something that only works under certain conditions, and something that only breaks under certain conditions". - Kelsey Hightower

> "Instability is a drag on innovation." - Yehudah Katz

**Resources:**

- [Software Development for Infrastructure](http://www.stroustrup.com/Software-for-infrastructure.pdf) - Bjarne Stroustrup
- [Normalization of Deviance in Software](http://danluu.com/wat/) - danluu.com
- [Lessons learned from reading postmortems](http://danluu.com/postmortem-lessons/) - danluu.com
- [Technical Debt Quadrant](https://martinfowler.com/bliki/TechnicalDebtQuadrant.html) - Martin Fowler
- [Design Philosophy On Integrity](https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-integrity.html) - William Kennedy
- [Ratio of bugs per line of code](https://www.mayerdan.com/ruby/2012/11/11/bugs-per-line-of-code-ratio) - Dan Mayer
- [Masterminds of Programming](http://dl.acm.org/citation.cfm?id=1592983) - Federico Biancuzzi and Shane Warden
- [Developing Software The Right Way, with Intent and Carefulness](http://ipengineer.net/2017/04/developing-software-the-right-way-with-intent-and-carefulness) - David Gee
- [What bugs live in the Cloud](https://www.usenix.org/system/files/login/articles/login_aug15_08_gunawi.pdf) - usenix.org
---
## Readability

**_We must structure our systems to be more comprehensible._**

This is about writing simple code that is easier to read and understand without the need of mental exhaustion. Just as important, it's about not hiding the cost/impact of the code per line, function, package and the overall ecosystem it runs in.

[Example Readability Issue](https://cpp.sh/?source=%2F%2F+Example+program+to+show+how+copy+elision+can+make+it+hard+to%0A%2F%2F+read+code+and+understand+its+cost.%0A%0A%23include+%3Ciostream%3E%0A%23include+%3Cstring%3E%0A%0Astruct+Foo+%7B%0A++++Foo()+%7B+std%3A%3Acout+%3C%3C+%22Constructed%22+%3C%3C+std%3A%3Aendl%3B+%7D%0A++++Foo(const+Foo+%26)+%7B+std%3A%3Acout+%3C%3C+%22Copy-constructed%22+%3C%3C+std%3A%3Aendl%3B+%7D%0A++++Foo(Foo+%26%26)+%7B+std%3A%3Acout+%3C%3C+%22Move-constructed%22+%3C%3C+std%3A%3Aendl%3B+%7D%0A++++~Foo()+%7B+std%3A%3Acout+%3C%3C+%22Destructed%22+%3C%3C+std%3A%3Aendl%3B+%7D%0A++%0A++++Foo+operator%3D(const+Foo%26)+%7B+std%3A%3Acout+%3C%3C+%22Assignment-overload%22+%3C%3C+std%3A%3Aendl%3B+return+*this%3B+%7D%0A%7D%3B%0A+%0AFoo+f1()+%7B%0A++++return+Foo()%3B%0A%7D%0A+%0Aint+main()+%7B++%0A%0A++++%2F%2F+How+many+objects+are+created+for+this+construction%3F%0A++++Foo+foo1+%3D+f1()%3B%0A++++std%3A%3Acout+%3C%3C+%22**************************%5Cn%22%3B%0A%0A++++%2F%2F+How+many+objects+are+created+for+this+assignment%3F%0A++++foo1+%3D+f1()%3B%0A++++std%3A%3Acout+%3C%3C+%22**************************%5Cn%22%3B%0A%7D)

**Code Must Never Lie**

It doesn't matter how fast the code might be if no one can understand or maintain it moving forward.

**Quotes**

> "This is a cardinal sin amongst programmers. If code looks like it’s doing one thing when it’s actually doing something else, someone down the road will read that code and misunderstand it, and use it or alter it in a way that causes bugs. That someone might be you, even if it was your code in the first place." - Nate Finch

[Code Must Never Lie](https://npf.io/2017/08/lies)

**Average Developer**

You must be aware of who you are on your team. When hiring new people, you must be aware of where the new person falls. Code must be written for the average developer to comprehend. If you are below average for your team, you have the responsibility to work to be average. If you are above average, you have the responsibility to reduce writing clever code and coach/mentor.

**Quotes**

> "Can you explain it to the median user (developer)? as opposed to will the smartest user (developer) figure it out?" - Peter Weinberger (inventor of AWK)

**Real Machine**

In Go, the underlying machine is a real machine, unlike what you would find in Java or C# with their virtual machine layer. The model of computation is that of the computer. Here is the key, Go gives you direct access to the machine while still providing abstraction mechanisms to allow higher-level ideas to be expressed.

**Quotes**

> "Making things easy to do is a false economy. Focus on making things easy to understand and the rest will follow." - Peter Bourgon

> "Reducing complexity is more powerful than hiding it." - Chris Hines
---
## Simplicity

**_We must understand that simplicity is hard to design and complicated to build._**

This is about hiding complexity. A lot of care and design must go into simplicity because this can cause more problems then good. It can create issues with readability and it can cause issues with performance.

**Complexity Sells Better**

Focus on encapsulation and validate that you're not generalizing or even being too concise. You might think you are helping the programmer or the code but validate things are still easy to use, understand, debug and maintain.

**Quotes**

> "Simplicity is a great virtue but it requires hard work to achieve it and education to appreciate it. And to make matters worse: complexity sells better." - Edsger W. Dijkstra

> "Everything should be made as simple as possible, but not simpler." - Albert Einstein

> "You wake up and say, I will be productive, not simple, today." - Dave Cheney

**Encapsulation**

Encapsulation is what we have been trying to figure out as an industry for 40+ years. Go is taking a slightly new approach with the package. Bringing encapsulation up a level and providing richer support at the language level.

**Quotes**

> Paraphrasing: "Encapsulation and the separation of concerns are drivers for designing software. This is largely based on how other industries handle complexity. There seems to be a human pattern of using encapsulation to wrestle complexity to the ground." - Brad Cox (inventor of Objective C)

> "The purpose of abstraction is not to be vague, but to create a new semantic level in which one can be absolutely precise." - Edsger W. Dijkstra

> "A proper abstraction decouples the code so that every change doesn’t echo throughout the entire code base." - Ronna Steinburg

> "A good API is not just easy to use but also hard to misuse." - JBD

> "Computing is all about abstractions. Those below yours are just details. Those above yours are limiting complicated crazy town." - Joe Beda

**Resources:**

- [Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM) - Rob Pike
- [What did Alan Kay mean by, "Lisp is the greatest single programming language ever designed"?](https://www.quora.com/What-did-Alan-Kay-mean-by-Lisp-is-the-greatest-single-programming-language-ever-designed/answer/Alan-Kay-11) - Alan Kay

---

### Performance

**_We must compute less to get the results we need._**

This is about not wasting effort and achieving execution efficiency. Writing code that is mechanically sympathetic with the runtime, operating system and hardware. Achieving performance by writing less and more efficient code, but staying within the idioms and framework of the language.

**Quotes**

> "Programmers waste enormous amounts of time thinking about, or worrying about, the speed of noncritical parts of their programs, and these attempts at efficiency actually have a strong negative impact when debugging and maintenance are considered. We should forget about small efficiencies, say about 97% of the time: premature optimization is the root of all evil. Yet we should not pass up our opportunities in that critical 3%." — Donald E. Knuth

> "I don't trust anything until it runs... In fact, I don't trust anything until it runs twice." - Andrew Gelman (one of the greatest living statisticians at Columbia University).

Rules of Performance:
  - Never guess about performance.
  - Measurements must be relevant.
  - Profile before you decide something is performance critical.
  - Test to know you are correct.

- [Example Benchmark](https://github.com/ardanlabs/gotraining/blob/master/topics/go/testing/benchmarks/basic/basic_test.go)

**Broad Engineering**

Performance is important but it can't be your priority unless the code is not running fast enough. You only know this once you have a working program and you have validated it. We place those who we think know how to write performant code on a pedestal. We need to put those who write code that is optimized for correctness and performs fast enough on those pedestals.

**Quotes**

> "When we're computer programmers we're concentrating on the intricate little fascinating details of programming and we don't take a broad engineering point of view about trying to optimize the total system. You try to optimize the bits and bytes." - Tom Kurtz (inventor of BASIC)

---

## Micro-Optimizations

Micro-Optimizations are about squeezing every ounce of performance as possible. When code is written with this as the priority, it is very difficult to write code that is readable, simple or idiomatic. You are writing clever code that may require the unsafe package or you may need to drop into assembly.

[Example Micro Optimization](https://play.golang.org/p/D_bImirgXL)

---

### Data-Oriented Design

> "Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming." - Rob Pike

**Design Philosophy:**

- If you don't understand the data, you don't understand the problem.
- All problems are unique and specific to the data you are working with.
- Data transformations are at the heart of solving problems. Each function, method and work-flow must focus on implementing the specific data transformations required to solve the problems.
- If your data is changing, your problems are changing. When your problems are changing, the data transformations needs to change with it.
- Uncertainty about the data is not a license to guess but a directive to STOP and learn more.
- Solving problems you don't have, creates more problems you now do.
- If performance matters, you must have mechanical sympathy for how the hardware and operating system work.
- Minimize, simplify and REDUCE the amount of code required to solve each problem. Do less work by not wasting effort.
- Code that can be reasoned about and does not hide execution costs can be better understood, debugged and performance tuned.
- Coupling data together and writing code that produces predictable access patterns to the data will be the most performant.
- Changing data layouts can yield more significant performance improvements than changing just the algorithms.
- Efficiency is obtained through algorithms but performance is obtained through data structures and layouts.

**Resources:**

- [Data-Oriented Design and C++](https://www.youtube.com/watch?v=rX0ItVEVjHc) - Mike Acton
- [Efficiency with Algorithms, Performance with Data Structures](https://www.youtube.com/watch?v=fHNmRkzxHWs) - Chandler Carruth

---

## Interface And Composition Design

**Design Philosophy:**

- Interfaces give programs structure.
- Interfaces encourage design by composition.
- Interfaces enable and enforce clean divisions between components.
  - The standardization of interfaces can set clear and consistent expectations.
- Decoupling means reducing the dependencies between components and the types they use.
  - This leads to correctness, quality and performance.
- Interfaces allow you to group concrete types by what they do.
  - Don't group types by a common DNA but by a common behavior.
  - Everyone can work together when we focus on what we do and not who we are.
- Interfaces help your code decouple itself from change.
  - You must do your best to understand what could change and use interfaces to decouple.
  - Interfaces with more than one method have more than one reason to change.
  - Uncertainty about change is not a license to guess but a directive to STOP and learn more.
- You must distinguish between code that:
  - defends against fraud vs protects against accidents

**Validation:**

Use an interface when:
- users of the API need to provide an implementation detail.
- API’s have multiple implementations they need to maintain internally.
- parts of the API that can change have been identified and require decoupling.

Don't use an interface:
- for the sake of using an interface.
- to generalize an algorithm.
- when users can declare their own interfaces.
- if it's not clear how the interface makes the code better.

**Resources:**

- [Methods, interfaces and Embedding](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy
- [Composition with Go](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html) - William Kennedy
- [Reducing type hierarchies](https://www.ardanlabs.com/blog/2016/10/reducing-type-hierarchies.html) - William Kennedy
- [Interface pollution in Go](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275) - Burcu Dogan
- [Application Focused API Design](https://www.ardanlabs.com/blog/2016/11/application-focused-api-design.html) - William Kennedy
- [Avoid interface pollution](https://www.ardanlabs.com/blog/2016/10/avoid-interface-pollution.html) - William Kennedy
- [Interface Values Are Valueless](https://www.ardanlabs.com/blog/2018/03/interface-values-are-valueless.html) - William Kennedy
- [Interface Semantics](https://www.ardanlabs.com/blog/2017/07/interface-semantics.html) - William Kennedy

---

## Package-Oriented Design

> Package Oriented Design allows a developer to identify where a package belongs inside a Go project and the design guidelines the package must respect. It defines what a Go project is and how a Go project is structured. Finally, it improves communication between team members and promotes clean package design and project architecture that is discussable.

---

## Concurrent Software Design

Concurrency means undefined “out of order” execution. Taking a set of instructions that would otherwise be executed in sequence and finding a way to execute them out of order and still produce the same result. For the problem in front of you, it has to be obvious that out of order execution would add value.

When I say value, I mean add enough of a performance gain for the complexity cost. Depending on your problem, out of order execution may not be possible or even make sense.

It’s also important to understand that [concurrency is not the same as parallelism](https://blog.golang.org/concurrency-is-not-parallelism). Parallelism means executing two or more instructions at the same time. This is a different concept from concurrency. Parallelism is only possible when you have at least 2 operating system (OS) and hardware threads available to you and you have at least 2 Goroutines, each executing instructions independently on each OS/hardware thread.

Both you and the runtime have a responsibility of managing the concurrency of the application. You are responsible for managing these three things when writing concurrent software:

**Design Philosophy:**

- The application must startup and shutdown with integrity.
- Know how and when every goroutine you create terminates.
- All goroutines you create should terminate before main returns.
- Applications should be capable of shutting down on demand, even under load, in a controlled way.
    - You want to stop accepting new requests and finish the requests you have (load shedding).
- Identify and monitor critical points of back pressure that can exist inside your application.
- Channels, mutexes and atomic functions can create back pressure when goroutines are required to wait.
- A little back pressure is good, it means there is a good balance of concerns.
- A lot of back pressure is bad, it means things are imbalanced.
- Back pressure that is imbalanced will cause:
    - Failures inside the software and across the entire platform.
    - Your application to collapse, implode or freeze.
- Measuring back pressure is a way to measure the health of the application.
- Rate limit to prevent overwhelming back pressure inside your application.
- Every system has a breaking point, you must know what it is for your application.
- Applications should reject new requests as early as possible once they are overloaded.
    - Don’t take in more work than you can reasonably work on at a time.
    - Push back when you are at critical mass. Create your own external back pressure.
- Use an external system for rate limiting when it is reasonable and practical.
- Use timeouts to release the back pressure inside your application.
- No request or task is allowed to take forever.
- Identify how long users are willing to wait.
- Higher-level calls should tell lower-level calls how long they have to run.
- At the top level, the user should decide how long they are willing to wait.
- Use the `Context` package.
    - Functions that users wait for should take a `Context`.
        - These functions should select on <-ctx.Done() when they would otherwise block indefinitely.
    - Set a timeout on a `Context` only when you have good reason to expect that a function's execution has a real time limit.
    - Allow the upstream caller to decide when the `Context` should be canceled.
    - Cancel a `Context` whenever the user abandons or explicitly aborts a call.
- Architect applications to:
- Identify problems when they are happening.
- Stop the bleeding.
- Return the system back to a normal state.

Index of the three part series:
1) [Scheduling In Go : Part I - OS Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
2) [Scheduling In Go : Part II - Go Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
3) [Scheduling In Go : Part III - Concurrency](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)

---
## Channel Design

Channels allow goroutines to communicate with each other through the use of signaling semantics. Channels accomplish this signaling through the use of sending/receiving data or by identifying state changes on individual channels. Don't architect software with the idea of channels being a queue, focus on signaling and the semantics that simplify the orchestration required.

**Language Mechanics:**

- Use channels to orchestrate and coordinate goroutines.
  - Focus on the signaling semantics and not the sharing of data.
  - Signaling with data or without data.
  - Question their use for synchronizing access to shared state.
      - > There are cases where channels can be simpler for this but initially question.
- Unbuffered channels:
  - Receive happens before the Send.
  - Benefit: 100% guarantee the signal being sent has been received.
  - Cost: Unknown latency on when the signal will be received.
- Buffered channels:
  - Send happens before the Receive.
  - Benefit: Reduce blocking latency between signaling.
  - Cost: No guarantee when the signal being sent has been received.
      - The larger the buffer, the less guarantee.
      - Buffer of 1 can give you one delayed send of guarantee.
- Closing channels:
  - Close happens before the Receive. (like Buffered)
  - Signaling without data.
  - Perfect for signaling cancellations and deadlines.
- NIL channels:
  - Send and Receive block.
  - Turn off signaling
  - Perfect for rate limiting or short-term stoppages.

**Design Philosophy:**

Depending on the problem you are solving, you may require different channel semantics. Depending on the semantics you need, different architectural choices must be taken.

- If any given Send on a channel `CAN` cause the sending goroutine to block:
  - Be careful with Buffered channels larger than 1.
      - Buffers larger than 1 must have reason/measurements.
  - Must know what happens when the sending goroutine blocks.
- If any given Send on a channel `WON'T` cause the sending goroutine to block:
  - You have the exact number of buffers for each send.
      - Fan Out pattern
  - You have the buffer measured for max capacity.
      - Drop pattern
- Less is more with buffers.
  - Don’t think about performance when thinking about buffers.
  - Buffers can help to reduce blocking latency between signaling.
      - Reducing blocking latency towards zero does not necessarily mean better throughput.
      - If a buffer of one is giving you good enough throughput then keep it.
      - Question buffers that are larger than one and measure for size.
      - Find the smallest buffer possible that provides good enough throughput.

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

> 🧠 You can simply multiply the **number of bytes by 8** **to get the number of bits**.

---

## Struct Types

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#struct-types)

Struct types are a way of creating complex types that group fields of data together. They are a great way of organizing and sharing the different aspects of the data your program consumes.

A computer architecture’s potential performance is determined predominantly by its word length (the number of bits that can be processed per access) and, more importantly, memory size, or the number of words that it can access.

## Notes

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#notes)

- We can use the struct literal form to initialize a value from a struct type.
- The dot (`.`) operator allows us to access individual field values.
- We can create anonymous structs.

## `Struct` Quotes

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#quotes)

> "Implicit conversion of types is the Halloween special of coding. Whoever thought of them deserves their own special hell." - Martin Thompson

## Links

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#links)

[Understanding Type in Go](https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html) - William Kennedy
[Object Oriented Programming in Go](https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html) - William Kennedy
[Padding is hard](https://dave.cheney.net/2015/10/09/padding-is-hard) - Dave Cheney
[Structure Member Alignment, Padding and Data Packing](https://www.geeksforgeeks.org/structure-member-alignment-padding-and-data-packing/)
[The Lost Art of Structure Packing](http://www.catb.org/esr/structure-packing) - Eric S. Raymond

## Code Review

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#code-review)

[Declare, create and initialize struct types](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/example1/example1.go) ([Go Playground](https://play.golang.org/p/djzGT1JtSwy))
[Anonymous struct types](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/example2/example2.go) ([Go Playground](https://play.golang.org/p/09cxjnmfcdC))
[Named vs Unnamed types](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/example3/example3.go) ([Go Playground](https://play.golang.org/p/ky91roJDjir))

## Advanced Code Review

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#advanced-code-review)

[Struct type alignments](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/rAvtS7cgD0z))

## Exercises

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#exercises)

### Exercise 1

[](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/README.md#exercise-1)

**Part A:** Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initialize with values and display each field.

**Part B:** Declare and initialize an anonymous struct type with the same three fields. Display the value.

[Template](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/h-7BEn2U3Rz)) | [Answer](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/struct_types/exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/eT_gLZKeHk-))

## Explain Golang Padding and Alignment

In Golang, padding and alignment are crucial concepts for optimizing memory usage and ensuring efficient access to data structures. These concepts are particularly important when dealing with structs, which are composite data types that group together variables under a single name.

Alignment refers to the way data is arranged and accessed in memory. Each data type in Go has an alignment requirement, which is the boundary on which the data must be placed. This requirement ensures that the CPU can access the data efficiently. The alignment values for common data types are as follows:

- `byte`, `int8`, `uint8`: 1 byte
- `int16`, `uint16`: 2 bytes
- `int32`, `uint32`, `float32`, `complex64`: 4 bytes
- `int64`, `uint64`, `float64`, `complex128`, `string`, `slice`: 8 bytes

The alignment requirement means that the memory address of a variable must be a multiple of its alignment value. For example, an int32 must be placed at an address that is a multiple of 4.

## Padding

Padding is the extra space added between fields in a struct to satisfy alignment requirements. When fields in a struct are not aligned properly, the compiler adds padding bytes to ensure that each field starts at an address that is a multiple of its alignment value. This padding can lead to wasted memory if not managed properly.

```go
type Example struct {
  a int8
  b string
  c int8
  d int32
}
```

In this struct, the fields are not optimally ordered, leading to padding:

- `a` (int8) starts at offset 0
- `b` (string) starts at offset 8 (7 bytes of padding after `a`)
- `c` (int8) starts at offset 24 (16 bytes for `b` and 7 bytes of padding)
- `d` (int32) starts at offset 28 (3 bytes of padding after `c`)

The total size of this struct is 32 bytes, with significant padding added to align the fields properly

## Pointers

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#pointers)

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.

## Pointer Notes

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#notes)

- Use pointers to share data.
- Values in Go are always pass by value.
- "Value of", what's in the box. "Address of" ( **&** ), where is the box.
- The (\*) operator declares a pointer variable and the "Value that the pointer points to".

## Escape Analysis

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#escape-analysis)

- When a value could be referenced after the function that constructs the value returns.
- When the compiler determines a value is too large to fit on the stack.
- When the compiler doesn’t know the size of a value at compile time.
- When a value is decoupled through the use of function or interface values.

## Garbage Collection History

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#garbage-collection-history)

The design of the Go GC has changed over the years:

- Go 1.0, Stop the world mark sweep collector based heavily on tcmalloc.
- Go 1.2, Precise collector, wouldn't mistake big numbers (or big strings of text) for pointers.
- Go 1.3, Fully precise tracking of all stack values.
- Go 1.4, Mark and sweep now parallel, but still stop the world.
- Go 1.5, New GC design, focusing on latency over throughput.
- Go 1.6, GC improvements, handling larger heaps with lower latency.
- Go 1.7, GC improvements, handling larger number of idle goroutines, substantial stack size fluctuation, or large package-level variables.
- Go 1.8, GC improvements, collection pauses should be significantly shorter than they were in Go 1.7, usually under 100 microseconds and often as low as 10 microseconds.
- Go 1.9, Large object allocation performance is significantly improved in applications using large (>50GB) heaps containing many large objects.
- Go 1.10, Many applications should experience significantly lower allocation latency and overall performance overhead when the garbage collector is active.

## Garbage Collection Semantics

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#garbage-collection-semantics)

[Garbage Collection Semantics Part I](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html) - William Kennedy

## Stack vs Heap

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#stack-vs-heap)

> "The stack is for data that needs to persist only for the lifetime of the function that constructs it, and is reclaimed without any cost when the function exits. The heap is for data that needs to persist after the function that constructs it exits, and is reclaimed by a sometimes costly garbage collection." - Ayan George

## Pointer Links

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#links)

### Pointer Mechanics

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#pointer-mechanics)

[Pointers vs. Values](https://golang.org/doc/effective_go.html#pointers_vs_values)
[Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html) - William Kennedy
[Using Pointers In Go](https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html) - William Kennedy
[Understanding Pointers and Memory Allocation](https://www.ardanlabs.com/blog/2013/07/understanding-pointers-and-memory.html) - William Kennedy

### Stacks

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#stacks)

[Contiguous Stack Proposal](https://docs.google.com/document/d/1wAaf1rYoM4S4gtnPh0zOlGzWtrZFQ5suE8qr2sD8uWQ/pub)

### Escape Analysis and Inlining

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#escape-analysis-and-inlining)

[Go Escape Analysis Flaws](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw)
[Compiler Optimizations](https://github.com/golang/go/wiki/CompilerOptimizations)

### Garbage Collection

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#garbage-collection)

[The Garbage Collection Handbook](http://gchandbook.org/)
[GC Pacer Redesign - 2021](https://github.com/golang/proposal/blob/master/design/44167-gc-pacer-redesign.md) - Michael Knyszek
[Tracing Garbage Collection](https://en.wikipedia.org/wiki/Tracing_garbage_collection)
[Go Blog - 1.5 GC](https://blog.golang.org/go15gc)
[Go GC: Solving the Latency Problem](https://www.youtube.com/watch?v=aiv1JOfMjm0&index=16&list=PL2ntRZ1ySWBf-_z-gHCOR2N156Nw930Hm)
[Concurrent garbage collection](http://rubinius.com/2013/06/22/concurrent-garbage-collection)
[Go 1.5 concurrent garbage collector pacing](https://docs.google.com/document/d/1wmjrocXIWTr1JxU-3EQBI6BK6KgtiFArkG47XK73xIQ/edit)
[Eliminating Stack Re-Scanning](https://github.com/golang/proposal/blob/master/design/17503-eliminate-rescan.md)
[Why golang garbage-collector not implement Generational and Compact gc?](https://groups.google.com/forum/m/#!topic/golang-nuts/KJiyv2mV2pU) - Ian Lance Taylor
[Getting to Go: The Journey of Go's Garbage Collector](https://blog.golang.org/ismmkeynote) - Rick Hudson
[Garbage Collection In Go : Part I - Semantics](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html) - William Kennedy
[Garbage Collection In Go : Part II - GC Traces](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html) - William Kennedy
[Garbage Collection In Go : Part III - GC Pacing](https://www.ardanlabs.com/blog/2019/07/garbage-collection-in-go-part3-gcpacing.html) - William Kennedy
[Go memory ballast: How I learnt to stop worrying and love the heap](https://blog.twitch.tv/en/2019/04/10/go-memory-ballast-how-i-learnt-to-stop-worrying-and-love-the-heap-26c2462549a2/) - Ross Engers

### Static Single Assignment Optimizations

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#static-single-assignment-optimizations)

[GopherCon 2015: Ben Johnson - Static Code Analysis Using SSA](https://www.youtube.com/watch?v=D2-gaMvWfQY)
[package ssa](https://godoc.org/golang.org/x/tools/go/ssa)
[Understanding Compiler Optimization](https://www.youtube.com/watch?v=FnGCDLhaxKU)

### Debugging code generation

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#debugging-code-generation)

[Debugging code generation in Go](https://rakyll.org/codegen/) - JBD

## Pointer Code Review

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#code-review)

[Pass by Value](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/example1/example1.go) ([Go Playground](https://play.golang.org/p/9kxh18hd_BT))
[Sharing data I](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/example2/example2.go) ([Go Playground](https://play.golang.org/p/mJz5RINaimn))
[Sharing data II](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/example3/example3.go) ([Go Playground](https://play.golang.org/p/GpmPICMGMre))
[Escape Analysis](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/example4/example4.go) ([Go Playground](https://play.golang.org/p/BCtJrNRJGun))
[Stack grow](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/example5/example5.go) ([Go Playground](https://play.golang.org/p/vBKF2hXvKBb))

### Escape Analysis Flaws

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#escape-analysis-flaws)

[Indirect Assignment](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/flaws/example1/example1_test.go)
[Indirection Execution](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/flaws/example2/example2_test.go)
[Assignment Slices Maps](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/flaws/example3/example3_test.go)
[Indirection Level Interfaces](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/flaws/example4/example4_test.go)
[Unknown](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/flaws/example5/example5_test.go)

## Pointer Exercises

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#exercises)

### Pointer Exercise 1

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#exercise-1)

**Part A** Declare and initialize a variable of type int with the value of 20. Display the _address of_ and _value of_ the variable.

**Part B** Declare and initialize a pointer variable of type int that points to the last variable you just created. Display the _address of_ , _value of_ and the _value that the pointer points to_.

[Template](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/6QYTKWzF8s8)) | [Answer](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/qq5P9gRDHKc))

### Pointer Exercise 2

[](https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers#exercise-2)

Declare a struct type and create a value of this type. Declare a function that can change the value of some field in this struct type. Display the value before and after the call to your function.

[Template](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/exercises/template2/template2.go) ([Go Playground](https://play.golang.org/p/nolKjrgBX-X)) | [Answer](https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/exercises/exercise2/exercise2.go) ([Go Playground](https://play.golang.org/p/i6utWhgDUH4))

## Value vs Pointer Semantics

1. Value Semantics: When you use value semantics, you are **making a copy of the original value** and passing around that copy. Any changes you make to the copy will not affect the original value. This is the default way of passing data in Go. All data types in Go are pass-by-value. This means that when you call a function with a parameter, a new copy of the parameter is created for the function's use.
2. Pointer Semantics: When you use pointer semantics, you are **passing around the memory address** (or pointer) where the value is stored. This means that if you change the value at that memory address, the change will be seen by anyone who has access to that memory address. This is useful when you want to change the value of a variable that was defined outside the function, or when you want to avoid copying large amounts of data.
