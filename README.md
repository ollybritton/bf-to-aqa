# `bf-to-aqa`
- [`bf-to-aqa`](#bf-to-aqa)
    - [Usage](#usage)
    - [Limitations](#limitations)
    - [In the Future?](#in-the-future)
    - [Examples](#examples)
      - [Hello, World!](#hello-world)
      - [Factorial](#factorial)
      - [Fibonacci](#fibonacci)


`bf-to-aqa` is a small project for converting the [Brainf***](https://en.wikipedia.org/wiki/Brainfuck) esoteric programming language into [AQA assembly language](https://filestore.aqa.org.uk/resources/computing/AQA-75162-75172-ALI.PDF) which can then be run by a simulator like the [one by Peter Higginson](https://www.peterhigginson.co.uk/AQA/).


### Usage
```
cat code.bf | bf-to-aqa
```

To install:

```
go get -u github.com/ollybritton/bf-to-aqa
```

### Limitations
* The assembly output isn't nearly as efficient as assembly written by hand.
* Too many instructions means that [Peter Higginson's AQA CPU](https://www.peterhigginson.co.uk/AQA/) won't work. Any program over 200 instructions long won't work and since instructions and the "tape" share the same memory, there won't be enough space left over anyway.

### In the Future?
* Write a new AQA simulator in Go, something like `go-aqa-asm`, similar to [`go-lmc`](https://github.com/ollybritton/go-lmc) so that bigger programs can work.
* Use the new capacity along with an existing project like [`ELVM`](https://github.com/shinh/elvm) or [`c2bf`](https://github.com/arthaud/c2bf) to allow compilation from C to the AQA assembly language without actually writing a compiler.

### Examples
#### Hello, World!
> `>++++++++[-<+++++++++>]<.>>+>-[+]++>++>+++[>[->+++<<+++>]<<]>-----.>->`  
> `+++..+++.>-.<<+[>[+>+]>>]<--------------.>>.+++.------.--------.>+.>+.`

File: [`examples/hello-world.bf`](./examples/hello-world.bf)

```
$ cat examples/hello-world.bf | bf-to-aqa

defineRegisters:
	mov r0,#0
	mov r1,#body
run:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#8
	str r0,[r1]
LS9:
	cmp r0,#0
	beq LE22
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#9
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	b LS9
LE22:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	out r0,7
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
LS30:
	cmp r0,#0
	beq LE32
	add r0,r0,#1
	str r0,[r1]
	b LS30
LE32:
	add r0,r0,#2
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#2
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#3
	str r0,[r1]
LS42:
	cmp r0,#0
	beq LE59
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS44:
	cmp r0,#0
	beq LE56
	sub r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#3
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#3
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	b LS44
LE56:
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
	b LS42
LE59:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#5
	out r0,7
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#3
	str r0,[r1]
	out r0,7
	out r0,7
	add r0,r0,#3
	str r0,[r1]
	out r0,7
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
	out r0,7
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
LS86:
	cmp r0,#0
	beq LE95
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS88:
	cmp r0,#0
	beq LE92
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	b LS88
LE92:
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	b LS86
LE95:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#14
	out r0,7
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	out r0,7
	add r0,r0,#3
	str r0,[r1]
	out r0,7
	sub r0,r0,#6
	out r0,7
	sub r0,r0,#8
	out r0,7
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	out r0,7
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	out r0,7
	HALT
body: dat 0
```

#### Factorial
> `+++++`
> `[->+<]>[[>+>+<<-]>>-<[-<+>]>[-<+>]<]<[->+<]<[[->>[-<+>>+<]>[-<+>]<<<]>>[-]<<<]>>[-<<+>>]<<`

File: [`examples/factorial.bf`](examples/factorial.bf)

```
$ cat examples/factorial.bf | bf-to-aqa | xsel -b -i

defineRegisters:
	mov r0,#0
	mov r1,#body
run:
	add r0,r0,#5
	str r0,[r1]
LS6:
	cmp r0,#0
	beq LE11
	sub r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	b LS6
LE11:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS13:
	cmp r0,#0
	beq LE41
LS14:
	cmp r0,#0
	beq LE22
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
	sub r0,r0,#1
	b LS14
LE22:
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS27:
	cmp r0,#0
	beq LE32
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	b LS27
LE32:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS34:
	cmp r0,#0
	beq LE39
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	b LS34
LE39:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	b LS13
LE41:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS43:
	cmp r0,#0
	beq LE48
	sub r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	b LS43
LE48:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS50:
	cmp r0,#0
	beq LE83
LS51:
	cmp r0,#0
	beq LE74
	sub r0,r0,#1
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
LS55:
	cmp r0,#0
	beq LE63
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	b LS55
LE63:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS65:
	cmp r0,#0
	beq LE70
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	b LS65
LE70:
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	b LS51
LE74:
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
LS77:
	cmp r0,#0
	beq LE79
	sub r0,r0,#1
	b LS77
LE79:
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	b LS50
LE83:
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
LS86:
	cmp r0,#0
	beq LE93
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	b LS86
LE93:
	HALT
body: dat 0
```

#### Fibonacci
> `+++++++++++`
> `>+>>>>++++++++++++++++++++++++++++++++++++++++++++`
> `>++++++++++++++++++++++++++++++++<<<<<<[>[>>>>>>+>`
> `+<<<<<<<-]>>>>>>>[<<<<<<<+>>>>>>>-]<[>++++++++++[-`
> `<-[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<[>>>+<<<`
> `-]>>[-]]<<]>>>[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]`
> `>[<<+>>[-]]<<<<<<<]>>>>>[+++++++++++++++++++++++++`
> `+++++++++++++++++++++++.[-]]++++++++++<[->-<]>++++`
> `++++++++++++++++++++++++++++++++++++++++++++.[-]<<`
> `<<<<<<<<<<[>>>+>+<<<<-]>>>>[<<<<+>>>>-]<-[>>.>.<<<`
> `[-]]<<[>>+>+<<<-]>>>[<<<+>>>-]<<[<+>-]>[<+>-]<<<-]`

File: [`examples/fibonacci.bf`](examples/fibonacci.bf)
Note: this one won't actually work in the Peter Higginson AQA assembly simulator because there's just too many instructions (451 lines long).

```
$ cat examples/fibonacci.bf | bf-to-aqa

defineRegisters:
	mov r0,#0
	mov r1,#body
run:
	add r0,r0,#11
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#4
	ldr r0,[r1]
	add r0,r0,#44
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#32
	str r0,[r1]
	sub r1,r1,#6
	ldr r0,[r1]
LS102:
	cmp r0,#0
	beq LE520
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS104:
	cmp r0,#0
	beq LE123
	str r0,[r1]
	add r1,r1,#6
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#7
	ldr r0,[r1]
	sub r0,r0,#1
	b LS104
LE123:
	str r0,[r1]
	add r1,r1,#7
	ldr r0,[r1]
LS131:
	cmp r0,#0
	beq LE148
	str r0,[r1]
	sub r1,r1,#7
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#7
	ldr r0,[r1]
	sub r0,r0,#1
	b LS131
LE148:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS150:
	cmp r0,#0
	beq LE285
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#10
	str r0,[r1]
LS162:
	cmp r0,#0
	beq LE226
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
LS167:
	cmp r0,#0
	beq LE177
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS167
LE177:
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
LS181:
	cmp r0,#0
	beq LE190
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS181
LE190:
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS193:
	cmp r0,#0
	beq LE202
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS195:
	cmp r0,#0
	beq LE197
	sub r0,r0,#1
	b LS195
LE197:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS199:
	cmp r0,#0
	beq LE201
	sub r0,r0,#1
	b LS199
LE201:
	b LS193
LE202:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS204:
	cmp r0,#0
	beq LE223
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
LS207:
	cmp r0,#0
	beq LE217
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS207
LE217:
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
LS220:
	cmp r0,#0
	beq LE222
	sub r0,r0,#1
	b LS220
LE222:
	b LS204
LE223:
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
	b LS162
LE226:
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
LS230:
	cmp r0,#0
	beq LE240
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS230
LE240:
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
LS244:
	cmp r0,#0
	beq LE253
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS244
LE253:
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS256:
	cmp r0,#0
	beq LE265
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS258:
	cmp r0,#0
	beq LE260
	sub r0,r0,#1
	b LS258
LE260:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS262:
	cmp r0,#0
	beq LE264
	sub r0,r0,#1
	b LS262
LE264:
	b LS256
LE265:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS268:
	cmp r0,#0
	beq LE277
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
LS274:
	cmp r0,#0
	beq LE276
	sub r0,r0,#1
	b LS274
LE276:
	b LS268
LE277:
	str r0,[r1]
	sub r1,r1,#7
	ldr r0,[r1]
	b LS150
LE285:
	str r0,[r1]
	add r1,r1,#5
	ldr r0,[r1]
LS291:
	cmp r0,#0
	beq LE345
	add r0,r0,#48
	str r0,[r1]
	out r0,7
LS342:
	cmp r0,#0
	beq LE344
	sub r0,r0,#1
	b LS342
LE344:
	b LS291
LE345:
	add r0,r0,#10
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
LS357:
	cmp r0,#0
	beq LE362
	sub r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	b LS357
LE362:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#48
	str r0,[r1]
	out r0,7
LS414:
	cmp r0,#0
	beq LE416
	sub r0,r0,#1
	b LS414
LE416:
	str r0,[r1]
	sub r1,r1,#12
	ldr r0,[r1]
LS430:
	cmp r0,#0
	beq LE442
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#4
	ldr r0,[r1]
	sub r0,r0,#1
	b LS430
LE442:
	str r0,[r1]
	add r1,r1,#4
	ldr r0,[r1]
LS447:
	cmp r0,#0
	beq LE458
	str r0,[r1]
	sub r1,r1,#4
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#4
	ldr r0,[r1]
	sub r0,r0,#1
	b LS447
LE458:
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
LS461:
	cmp r0,#0
	beq LE474
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	out r0,7
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	out r0,7
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
LS471:
	cmp r0,#0
	beq LE473
	sub r0,r0,#1
	b LS471
LE473:
	b LS461
LE474:
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
LS477:
	cmp r0,#0
	beq LE487
	str r0,[r1]
	add r1,r1,#2
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS477
LE487:
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
LS491:
	cmp r0,#0
	beq LE500
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS491
LE500:
	str r0,[r1]
	sub r1,r1,#2
	ldr r0,[r1]
LS503:
	cmp r0,#0
	beq LE508
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
	b LS503
LE508:
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
LS510:
	cmp r0,#0
	beq LE515
	str r0,[r1]
	sub r1,r1,#1
	ldr r0,[r1]
	add r0,r0,#1
	str r0,[r1]
	add r1,r1,#1
	ldr r0,[r1]
	sub r0,r0,#1
	b LS510
LE515:
	str r0,[r1]
	sub r1,r1,#3
	ldr r0,[r1]
	sub r0,r0,#1
	b LS102
LE520:
	HALT
body: dat 0
```