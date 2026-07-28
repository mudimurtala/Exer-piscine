# TypeScript + React Deep Dive
### Built from freeCodeCamp's "TypeScript in React – Full Tutorial" (Rachel Johnson, Scrimba)

## How to use this document

This is not a summary. It is a full companion to the video, expanded with the
foundational explanations you asked for. Open this file in VS Code, open a
second terminal, scaffold the project exactly as described in Part 1, and
work through each part in order. Type every code snippet yourself. Do not
copy paste. Typing it is what builds the muscle memory that watching alone
never gave you.

Each part below follows the same structure:

- **What is happening in the video** — plain explanation, no jargon left undefined
- **Concept deep dive** — the fundamental React or TypeScript idea behind that
  part, explained from the ground up, with a fresh example that is not from
  the video
- **Code** — the pattern you should type into your own editor
- **Your turn** — a small extra exercise so you are not just transcribing

At the very end you will find a glossary of every important term used in
this document, and a short practice plan for the days after you finish
reading.

---

## Before we start: the four words that keep confusing you

You told me your biggest fear is being handed a snippet and not being able
to tell a prop from a hook from a state from a callback. Let's kill that
confusion right now, before touching any code from the video.

### Component

A component is just a JavaScript function that returns something React can
render on screen (usually JSX). That is the whole definition. Nothing
mystical.

```tsx
function Greeting() {
  return <h1>Hello there</h1>;
}
```

That is a complete, valid React component. It takes no input and always
returns the same JSX.

### Props

Props are the *input* to a component. Exactly like function parameters,
because that is literally what they are under the hood. They are read only
from inside the component. A parent component passes props down to a child.

```tsx
function Greeting(props: { name: string }) {
  return <h1>Hello {props.name}</h1>;
}

// used like this:
<Greeting name="Mudi" />
```

If you see curly braces being read off a function parameter in a component,
you are looking at props. Full stop. That is the tell.

### State

State is data that a component *owns* and that can *change over time*,
where a change should cause the component to re render with the new value.
State is created with the `useState` hook. Unlike props, state is not
passed in from outside, it lives inside the component that declares it.

```tsx
import { useState } from "react";

function Counter() {
  const [count, setCount] = useState(0);
  return <button onClick={() => setCount(count + 1)}>{count}</button>;
}
```

If you see `useState(...)` being called and destructured into an array of
two things, you are looking at state.

### Hook

A hook is simply a special function, always starting with the word `use`,
that lets a function component tap into React features like state,
lifecycle timing, or context. `useState` is a hook. `useEffect` is a hook.
`useRef` is a hook. Props and state are not hooks themselves, they are
things hooks give you access to. So the rule of thumb is: if the function
name starts with `use`, it is a hook. `useState` is the hook. What it
returns (`count` and `setCount`) is the state and the state setter, not the
hook itself.

### Callback

A callback is simply a function that you pass as an argument into another
function or component, to be called later, usually in response to
something happening (a click, a change, an event). It is called a
"callback" because the receiving code calls it back at the right moment.

```tsx
function Button({ onClick }: { onClick: () => void }) {
  return <button onClick={onClick}>Click me</button>;
}

// The arrow function below is the callback, passed in as a prop:
<Button onClick={() => console.log("clicked")} />
```

So notice something important: a callback is very often *also* a prop,
because the most common way to hand a function down into a child component
is to pass it as a prop. That is exactly why these words tangle up for you.
They are not mutually exclusive categories, they describe different things
about the same value. `onClick` above is simultaneously a prop (because it
is passed into the component) and a callback (because of what it is and how
it gets used).

Keep this quick test in your head as you go through the rest of this
document:

| You see this pattern | It is called this |
|---|---|
| A value read from the function's parameter object, coming from a parent | Prop |
| `useState(...)` being called, its two returned values | State |
| Any function whose name starts with `use` | Hook |
| A function value being passed around to be run later, often in response to an event | Callback |

---

## Foundational TypeScript primer

Before the video even starts refactoring anything, you need one solid idea
in your head: TypeScript does not change how your code runs. It only checks,
while you are writing code, that the values flowing through your program
match the shapes you said they would. It disappears completely at build
time, compiling down to plain JavaScript. Its entire job is to catch
mistakes on your screen, in your editor, before you ever run the app.

A "type" is just a description of the shape of a value. Here are the
building blocks you will lean on throughout this document.

**Basic types**

```ts
let age: number = 27;
let name: string = "Mudi";
let isActive: boolean = true;
let tags: string[] = ["react", "typescript"];
```

**Custom types**, built with the `type` keyword, describe the shape of an
object so you do not have to write that shape out inline every time:

```ts
type Developer = {
  name: string;
  yearsOfExperience: number;
  isRemote: boolean;
};

const mudi: Developer = {
  name: "Mudi",
  yearsOfExperience: 1,
  isRemote: true,
};
```

**Function types**, describing parameters and return values:

```ts
function double(value: number): number {
  return value * 2;
}
```

Reading that signature out loud: this function takes one parameter called
`value` which must be a `number`, and it returns a `number`.

**Union types**, meaning a value can be one of several possible types,
joined with a pipe symbol:

```ts
function formatId(id: string | number): string {
  return `ID-${id}`;
}
```

`id` here is allowed to be a `string` OR a `number`, nothing else.

**Generics**, which is the idea that gave you trouble before by your own
account, so let's slow down on it. A generic lets a function or type accept
a *type as if it were a parameter*, so the same reusable code can work
safely with different shapes of data without losing type safety.

```ts
function wrapInArray<T>(value: T): T[] {
  return [value];
}

wrapInArray<string>("hello"); // T becomes string, returns string[]
wrapInArray<number>(5);       // T becomes number, returns number[]
```

`T` is just a placeholder name (you could call it anything, `T` is
convention) that gets filled in with whatever type you specify in the
angle brackets when you call the function. This is exactly the mechanism
`useState<T>()` uses, which is why the video keeps bringing generics back
up. You will see it constantly below.

---

## Part 1 — Project setup with Vite (00:00 to 05:02)

### What is happening

Instead of adding TypeScript to an already existing JavaScript React
project by hand (installing packages, renaming files, writing config), the
video scaffolds a brand new project using Vite with TypeScript built in
from the start. Vite is a build tool, it is the thing that runs your dev
server and bundles your code, and it ships a project generator that can
set everything up for you in one command.

### Concept deep dive: why a fresh TypeScript project instead of converting one

When you convert an existing JavaScript project to TypeScript by hand you
need to install `typescript`, `@types/react`, `@types/react-dom`, create a
`tsconfig.json`, and rename every `.jsx` file to `.tsx`. Vite's scaffolding
tool does all of that for you correctly in seconds, which is why the video
recommends starting fresh and then copying your existing source files
across, rather than editing an existing setup in place.

### Code: create the project

```bash
npm create vite@latest
```

You will be prompted for a project name, then asked to pick React as the
framework, and TypeScript as the variant. Afterward:

```bash
npm install
npm run dev
```

Look inside the generated `package.json`. You will see `typescript`,
`@types/react`, and `@types/react-dom` already listed as dependencies. Open
`tsconfig.json` and you will see it references `tsconfig.app.json` and
`tsconfig.node.json`, which hold the actual compiler settings.

### Your turn

Scaffold this project yourself right now with the exact commands above,
using the name `assembly-endgame-ts`. Open the generated `tsconfig.app.json`
and read through it once, even if half of it does not make sense yet. You
are just building familiarity with where things live.

---

## Part 2 — TypeScript refresher: basic and custom types (05:02 to 11:50)

### What is happening

The video works through a `words.ts` file (an array of strings) and a
`languages.ts` file (an array of objects), typing both. This is where the
two ways of typing an array of objects show up: an inline object type, and
a separately declared custom type.

### Concept deep dive: inline types versus custom types, and why custom types win

```ts
// Inline: works, but gets unreadable fast and cannot be reused
const languages: { name: string; backgroundColor: string; color: string }[] = [
  { name: "TypeScript", backgroundColor: "#3178c6", color: "#fff" },
];

// Custom type: declared once, reused anywhere, easy to read
type Language = {
  name: string;
  backgroundColor: string;
  color: string;
};

const languages: Language[] = [
  { name: "TypeScript", backgroundColor: "#3178c6", color: "#fff" },
];
```

The second version wins for two reasons. First, readability, the intent is
named. Second, and this matters a lot later in this document, once a type
is declared separately you can `export` it and reuse it in other files,
including components that need to know the shape of a prop.

### Your turn

Create `src/words.ts` with an exported array of five programming language
names, typed as `string[]`. Then create `src/languages.ts` with a `Language`
type exactly like above, and an exported array of three objects matching
that type.

---

## Part 3 — TypeScript refresher: functions, and building `getRandomIndex` (08:31 to 14:07)

### What is happening

Two utility functions, `getRandomWord` and `getFarewellText`, get typed.
Along the way the video notices both functions independently calculate a
random index using the same line of code, and refactors that repeated line
out into its own reusable `getRandomIndex` function.

### Concept deep dive: the `any` type is your enemy

When TypeScript cannot figure out what type something is, and you have not
told it, it silently falls back to a type called `any`. `any` turns off
type checking completely for that value, which defeats the entire purpose
of using TypeScript in the first place. Whenever your editor shows a
squiggly warning about an implicit `any`, that is TypeScript telling you
"I have no idea what this is supposed to be, please tell me."

```ts
function getFarewellText(language) {
  // language is implicitly `any` here — TypeScript is not helping you at all
}

function getFarewellText(language: string) {
  // now TypeScript will catch it if you ever call this with a number by mistake
}
```

### Code

```ts
function getRandomIndex(array: string[]): number {
  return Math.floor(Math.random() * array.length);
}

function getRandomWord(words: string[]): string {
  return words[getRandomIndex(words)];
}

function getFarewellText(language: string): string {
  const options: string[] = [
    `Goodbye ${language}`,
    `${language} has left the game`,
  ];
  return options[getRandomIndex(options)];
}
```

### Your turn

Write a third function, `getRandomLanguage(languages: Language[]): Language`,
that reuses `getRandomIndex` to pick a random language object out of the
`languages` array you built in Part 2. This forces you to type a function
whose parameter is an array of a custom type, and whose return value is
that custom type, not a primitive.

---

## Part 4 — Typing `useState` (14:07 to 17:36)

### What is happening

This is the first genuinely React specific typing in the whole video, and
it is the one most worth slowing down on given what you told me about your
struggles with state.

### Concept deep dive: how `useState` infers versus how you explicitly type it

When you give `useState` an initial value, TypeScript looks at that value
and infers the type automatically.

```tsx
const [currentWord, setCurrentWord] = useState("react");
// TypeScript infers currentWord is a string, because "react" is a string
```

This works fine as long as there is a meaningful initial value. The problem
shows up with empty arrays or `null` starting values, because TypeScript
cannot infer anything useful from nothing.

```tsx
const [guessedLetters, setGuessedLetters] = useState([]);
// TypeScript infers this as any[], which is nearly useless
```

This is exactly where the generic syntax from the primer above comes back.
`useState` is written internally to accept a generic type parameter, so you
can tell it explicitly what the state will hold, regardless of what the
initial value looks like:

```tsx
const [guessedLetters, setGuessedLetters] = useState<string[]>([]);
// Now guessedLetters is explicitly string[], and setGuessedLetters will
// reject anything that is not an array of strings.
```

The angle bracket syntax `useState<string[]>` is the generic argument.
Read it as: "this call to useState is being told, explicitly, that T is
string[]." Both the value returned and the setter function that updates it
now respect that type. If you ever try `setGuessedLetters(5)`, TypeScript
will immediately flag it as an error, before you even run the code. That
is the entire value proposition of typed state.

### Code

```tsx
const [currentWord, setCurrentWord] = useState<string>(getRandomWord(words));
const [guessedLetters, setGuessedLetters] = useState<string[]>([]);
```

### Your turn

In your scaffolded project, create a state variable called `score` typed
explicitly as a `number` starting at `0`, and a state variable called
`selectedLanguage` typed explicitly as `Language | null`, starting at
`null`. That second one is a union type combining your custom `Language`
type with `null`, which is a very common real world pattern: "this might
hold a language object, or it might hold nothing yet."

---

## Part 5 — Typing derived values and arrow functions (17:36 to 19:48)

### What is happening

Several plain constants (like `numGuessesLeft`, `isGameOver`) sit directly
underneath the state declarations. These are called derived values because
their value is calculated from state, they are not state themselves, they
do not get their own `useState` call. The video also types a couple of
small inline arrow functions used for checks.

### Concept deep dive: derived values do not need `useState`

This trips a lot of learners up. A common beginner mistake is to make
*everything* a piece of state. But if a value can always be recalculated
from existing state, it should just be a plain constant recalculated on
every render, not its own state.

```tsx
const [wrongGuessCount, setWrongGuessCount] = useState(0);

// This is a DERIVED VALUE, not state. It is recalculated every render
// from wrongGuessCount, so it does not need its own useState call.
const isGameLost = wrongGuessCount >= 8;
```

TypeScript usually infers these correctly on its own (`isGameLost` above is
inferred as `boolean`), but you can still type them explicitly for clarity:

```tsx
const isGameLost: boolean = wrongGuessCount >= 8;
```

### Code: typing a small arrow function

```ts
const isLetterGuessed = (letter: string): boolean =>
  guessedLetters.includes(letter);
```

### Your turn

Write a derived value `remainingGuesses: number` calculated from a
`maxGuesses` constant and your `wrongGuessCount` state. Then write a typed
arrow function `isGameWon(word: string, guessed: string[]): boolean` that
checks whether every letter in `word` appears in `guessed`.

---

## Part 6 — Typing the two `App.tsx` functions (19:48 to 23:10)

### What is happening

`startNewGame` (no parameters, returns nothing) and `addGuessedLetter` (one
parameter, updates state through a callback) get fully typed.

### Concept deep dive: the `void` return type

`void` means "this function does not return a usable value." You will see
it constantly on functions whose entire job is to run side effects, like
updating state, rather than computing and returning something.

```ts
function startNewGame(): void {
  setCurrentWord(getRandomWord(words));
  setGuessedLetters([]);
}
```

### Concept deep dive: the state setter callback form

State setter functions like `setGuessedLetters` can be called two ways.
Either you pass the new value directly, or you pass a callback function
that receives the *previous* state and returns the new state. The callback
form is safer when the new state depends on the old state, because it
avoids stale closures.

```tsx
function addGuessedLetter(letter: string): void {
  setGuessedLetters((prevLetters: string[]): string[] =>
    prevLetters.includes(letter) ? prevLetters : [...prevLetters, letter]
  );
}
```

TypeScript is usually smart enough to infer that `prevLetters` is
`string[]` on its own, because it already knows `guessedLetters` is
`string[]` from Part 4. Typing it explicitly anyway, like above, is good
practice while you are still building the habit.

### Your turn

Write `resetScore(): void` that sets your `score` state back to `0`. Then
write `incrementScore(points: number): void` that adds `points` to the
current score using the callback setter form.

---

## Part 7 — Splitting into components (23:10 to 25:12)

### What is happening

No new typing concept here, this is a structural pivot. The single
`App.tsx` file gets broken apart into separate component files (`Header`,
`GameStatus`, `LanguageChips`, `WordLetters`, `AriaLiveStatus`, `Keyboard`,
`NewGameButton`, `ConfettiContainer`), each receiving data through props
from `App.tsx`. This is the setup for everything that follows.

### Concept deep dive: why break a UI into components at all

Splitting UI into components is not about TypeScript, it is a core React
practice, done for reusability and readability. Each component should
ideally do one visual job. `App.tsx` becomes the "parent" that holds state
and hands pieces of it down as props to each "child" component that needs
to display or interact with that piece.

### Your turn

Sketch (on paper or in a scratch file, does not need to run) which pieces
of state from your Assembly Endgame clone would need to be passed down to
a `Keyboard` component versus a `WordDisplay` component. This is a design
exercise, not a coding one, and it is exactly the kind of thinking
interviewers ask about.

---

## Part 8 — Typing React components with `JSX.Element` (25:12 to 27:48)

### What is happening

The video imports the `JSX` type from React and uses `JSX.Element` to
explicitly type what a component function returns.

### Concept deep dive: why type a component's return value at all

By default, TypeScript infers the return type of a component from what it
actually returns. If a component always returns JSX, TypeScript infers
`JSX.Element` on its own, and you technically do not need to write it out.
But writing it explicitly is a safety net, most valuable on a large
codebase or team, because it stops you from accidentally returning
something that is not renderable JSX (a stray `true`, a `string` on its
own outside of JSX, and so on) without TypeScript complaining loudly at the
point of the mistake.

```tsx
import type { JSX } from "react";

function Header(): JSX.Element {
  return (
    <header>
      <h1>Assembly Endgame</h1>
      <p>Guess the word before Assembly takes over</p>
    </header>
  );
}
```

### Concept deep dive: `React.FC` versus `JSX.Element`, and why this document avoids `React.FC`

You will see a lot of older React and TypeScript code use `React.FC` (or
`React.FunctionComponent`) to type a whole component:

```tsx
const Header: React.FC = () => {
  return <header>Assembly Endgame</header>;
};
```

`React.FC` automatically types the return value as `JSX.Element` for you
and also automatically includes a `children` prop even when the component
never uses `children`. That last part is the actual problem, it silently
widens your prop type with something you did not ask for. Because of that,
current TypeScript and React convention (and this document, and the video)
prefers typing the return value directly with `JSX.Element` and typing
props separately and explicitly, giving you full control over exactly what
each component accepts.

### Your turn

Write a component called `Footer` that takes no props and returns a
`JSX.Element` containing a copyright line. Type its return value
explicitly, the way shown above, even though TypeScript could infer it.

---

## Part 9 — Union return types: `JSX.Element | null` (27:48 to 29:24)

### What is happening

`ConfettiContainer` sometimes returns actual confetti JSX, and sometimes
returns `null` (React's way of saying "render nothing"), depending on
whether the game has been won. This means its return type cannot honestly
be `JSX.Element` alone.

### Concept deep dive: conditionally rendering nothing

Returning `null` from a component is completely normal in React, it simply
means "render nothing here." When a component's return type can genuinely
be more than one shape, a union type is the honest way to describe that,
exactly like the union type primer earlier in this document.

```tsx
import type { JSX } from "react";

function ConfettiContainer({ isGameWon }: { isGameWon: boolean }): JSX.Element | null {
  if (isGameWon) {
    return <Confetti />;
  }
  return null;
}
```

### Your turn

Write a component `ScoreBadge({ score }: { score: number })` that returns a
`JSX.Element` showing the score if it is greater than `0`, and returns
`null` otherwise. Type the return value as `JSX.Element | null`.

---

## Part 10 — Typing component props, inline (29:24 to 32:46)

### What is happening

`ConfettiContainer` gets its single prop, `isGameWon`, typed inline
directly on the function signature. `GameStatus` then gets five props
typed the same inline way, which starts to look messy with that many
props, setting up the next part.

### Concept deep dive: the syntax for typing props inline

Props in a function component are just a single object parameter. You type
that object the same way you would type any object parameter in a plain
function.

```tsx
function ConfettiContainer({ isGameWon }: { isGameWon: boolean }): JSX.Element | null {
  // ...
}
```

Reading this: the component takes one parameter (destructured into
`isGameWon`), and that parameter's type is an object with a single field
`isGameWon` of type `boolean`.

With multiple props, you keep chaining fields inside that same object type:

```tsx
function GameStatus({
  isGameWon,
  isGameLost,
  isGameOver,
  isLastGuessIncorrect,
  wrongGuessCount,
}: {
  isGameWon: boolean;
  isGameLost: boolean;
  isGameOver: boolean;
  isLastGuessIncorrect: boolean;
  wrongGuessCount: number;
}): JSX.Element {
  // ...
}
```

### Your turn

Write `PlayerCard` accepting two inline props: `name` (`string`) and
`isOnline` (`boolean`). Return a `JSX.Element` that renders the name and a
colored dot depending on `isOnline`.

---

## Part 11 — Custom prop types (32:46 to 34:34)

### What is happening

The five inline props on `GameStatus` get pulled out into a separately
declared `GameStatusProps` type, following the naming convention
`ComponentName` + `Props`. This is the same "inline versus custom type"
lesson from Part 2, now applied specifically to props, and it is the
pattern used for every remaining component in the video.

### Code

```tsx
import type { JSX } from "react";

type GameStatusProps = {
  isGameWon: boolean;
  isGameLost: boolean;
  isGameOver: boolean;
  isLastGuessIncorrect: boolean;
  wrongGuessCount: number;
};

function GameStatus({
  isGameWon,
  isGameLost,
  isGameOver,
  isLastGuessIncorrect,
  wrongGuessCount,
}: GameStatusProps): JSX.Element {
  // ...
}
```

From this point forward in this document, and in your own projects, always
reach for this custom `ComponentNameProps` pattern rather than inline
typing the moment a component has more than one or two props.

### Your turn

Refactor the `PlayerCard` component from Part 10 to use a named
`PlayerCardProps` type instead of inline typing.

---

## Part 12 — Practice component: `AriaLiveStatus` (34:34 to 36:53)

### What is happening

A self contained challenge component with four props (`currentWord`,
`lastGuessedLetter`, `guessedLetters`, `numGuessesLeft`), plus a small
inline callback inside the JSX that also needs typing. This exists purely
so you practice everything from Parts 8 through 11 in one go, without new
concepts.

### Code

```tsx
import type { JSX } from "react";

type AriaLiveStatusProps = {
  currentWord: string;
  lastGuessedLetter: string;
  guessedLetters: string[];
  numGuessesLeft: number;
};

function AriaLiveStatus({
  currentWord,
  lastGuessedLetter,
  guessedLetters,
  numGuessesLeft,
}: AriaLiveStatusProps): JSX.Element {
  return (
    <p className="sr-only" role="status">
      Current word: {currentWord.split("").map((letter: string): string =>
        guessedLetters.includes(letter) ? letter : "blank"
      )}
      . {numGuessesLeft} guesses left. Last guess was {lastGuessedLetter}.
    </p>
  );
}
```

### Your turn

Build this component yourself from scratch in your project before looking
at the snippet above, then compare. That struggle is where the learning
actually happens, not in the reading.

---

## Part 13 — Exporting and importing custom types (36:53 to 39:15)

### What is happening

`LanguageChips` needs the `Language` type that lives in `languages.ts`
(from Part 2). The video exports that type from its file and imports it
into the component file, so the same type definition is shared and never
duplicated.

### Concept deep dive: one type, many files

This is the payoff for putting custom types in their own declarations
instead of inlining them everywhere. A type is just a piece of code like
any other, and it can be exported and imported exactly like a function or
a variable.

```ts
// languages.ts
export type Language = {
  name: string;
  backgroundColor: string;
  color: string;
};
```

```tsx
// LanguageChips.tsx
import type { Language } from "./languages";

type LanguageChipsProps = {
  languages: Language[];
  wrongGuessCount: number;
};

function LanguageChips({ languages, wrongGuessCount }: LanguageChipsProps) {
  // ...
}
```

Notice the `import type` syntax rather than a plain `import`. This tells
TypeScript (and your build tool) that you are only importing a type, not
actual runtime code, which lets the compiler strip it out completely when
it generates the final JavaScript, since types do not exist anymore at
runtime.

### Your turn

Move your `Developer` type from the very first TypeScript primer section
into its own file, `developer.ts`, export it, and import it into a new
component `DeveloperCard` that accepts a `developer: Developer` prop.

---

## Part 14 — Mapping typed data and the `Omit` utility type (39:15 to 42:14)

### What is happening

`LanguageChips` maps over the `languages` array to render one chip per
language, and needs its callback parameters, its callback's return value,
and the final array all typed. Along the way, a `styles` object needs a
type that is almost identical to `Language` except it should not include
the `name` field, which introduces the `Omit` utility type.

### Concept deep dive: utility types

TypeScript ships a handful of built in "utility types" that transform an
existing type into a new one without you retyping it from scratch. `Omit`
is one of the most useful: it takes a type and a list of keys to remove.

```ts
type Language = {
  name: string;
  backgroundColor: string;
  color: string;
};

type LanguageStyles = Omit<Language, "name">;
// LanguageStyles is now equivalent to: { backgroundColor: string; color: string }
```

Read `Omit<Language, "name">` as: "take everything from `Language`, except
the `name` field." Two other utility types worth knowing exist for the
opposite situation, when you want only some fields rather than excluding
one: `Pick<Language, "name">` would give you just `{ name: string }`.

### Concept deep dive: typing `.map()` correctly

```tsx
const languageElements: JSX.Element[] = languages.map(
  (lang: Language, index: number): JSX.Element => {
    const isLanguageLost = index < wrongGuessCount;
    const styles: LanguageStyles = {
      backgroundColor: lang.backgroundColor,
      color: lang.color,
    };
    return (
      <span key={lang.name} style={styles} className={isLanguageLost ? "lost" : ""}>
        {lang.name}
      </span>
    );
  }
);
```

Notice the pattern: the callback parameter (`lang`) is typed as a single
`Language`, the callback's return value is typed as a single `JSX.Element`
because that is what gets produced on each loop, and the final variable
holding the whole mapped result is typed as an array of that,
`JSX.Element[]`, because `.map()` always returns an array.

### Your turn

Using your `languages.ts` array from Part 2, write a `.map()` that returns
an array of plain strings, each formatted as `"LanguageName is a language"`.
Type the callback parameter, the callback return value, and the final
result explicitly, following the exact pattern above.

---

## Part 15 — Practice component: `WordLetters` (42:14 to 44:49)

### What is happening

A second full practice component, three props (`currentWord`,
`guessedLetters`, `isGameLost`), a `.map()` inside, and two internal
boolean and string variables to type. Same idea as Part 12, pure
repetition to build the habit.

### Code

```tsx
import type { JSX } from "react";

type WordLettersProps = {
  currentWord: string;
  guessedLetters: string[];
  isGameLost: boolean;
};

function WordLetters({ currentWord, guessedLetters, isGameLost }: WordLettersProps): JSX.Element {
  return (
    <div>
      {currentWord.split("").map((letter: string, index: number): JSX.Element => {
        const shouldRevealLetter: boolean = guessedLetters.includes(letter) || isGameLost;
        const letterClassName: string = shouldRevealLetter ? "revealed" : "hidden";
        return (
          <span key={index} className={letterClassName}>
            {shouldRevealLetter ? letter : ""}
          </span>
        );
      })}
    </div>
  );
}
```

### Your turn

Build it yourself first, from the props list alone, before checking the
snippet.

---

## Part 16 — Typing function props, callbacks passed down (44:49 to 47:33)

### What is happening

`NewGameButton` receives `startNewGame`, the actual `void` returning
function from `App.tsx`, as a prop. This is the moment your two confusing
concepts collide directly: a callback being passed down as a prop, and you
now need to type that.

### Concept deep dive: the syntax for a function typed as a prop

The pattern is: prop name, then a colon, then arrow function syntax
describing parameters and return type, exactly like typing a standalone
function, just written as the value of an object field instead.

```ts
type NewGameButtonProps = {
  isGameOver: boolean;
  startNewGame: () => void;
};
```

Read `startNewGame: () => void` as: "the `startNewGame` prop must be a
function that takes no parameters and returns nothing." If the function
took a parameter and returned something, it would look like this instead:

```ts
type SearchBoxProps = {
  onSearch: (query: string) => number;
};
```

That says: `onSearch` must be a function taking one `string` parameter and
returning a `number`.

### Code

```tsx
import type { JSX } from "react";

type NewGameButtonProps = {
  isGameOver: boolean;
  startNewGame: () => void;
};

function NewGameButton({ isGameOver, startNewGame }: NewGameButtonProps): JSX.Element | null {
  if (!isGameOver) return null;
  return <button onClick={startNewGame}>New Game</button>;
}
```

### Your turn

Write a `LikeButton` component with props `likeCount: number` and
`onLike: () => void`. Then write a second version, `RatingWidget`, whose
`onRate` prop is a function taking a `rating: number` parameter and
returning `void`. This second one is the shape you will run into constantly
in real apps, forms and interactive widgets that report a value back up to
a parent.

---

## Part 17 — The final challenge component: `Keyboard` (47:33 to 50:23)

### What is happening

This is the capstone, one component pulling together everything from every
part above: a typed return value, five typed props (including a callback
prop that itself takes a parameter), a `.map()` with typed callback
parameters and return value, and several internal boolean and string
variables.

### Code

```tsx
import type { JSX } from "react";

type KeyboardProps = {
  alphabet: string;
  guessedLetters: string[];
  currentWord: string;
  isGameOver: boolean;
  addGuessedLetter: (letter: string) => void;
};

function Keyboard({
  alphabet,
  guessedLetters,
  currentWord,
  isGameOver,
  addGuessedLetter,
}: KeyboardProps): JSX.Element {
  const keyboardElements: JSX.Element[] = alphabet.split("").map((letter: string): JSX.Element => {
    const isGuessed: boolean = guessedLetters.includes(letter);
    const isCorrect: boolean = isGuessed && currentWord.includes(letter);
    const isWrong: boolean = isGuessed && !currentWord.includes(letter);
    const className: string = isCorrect ? "correct" : isWrong ? "wrong" : "";

    return (
      <button
        key={letter}
        className={className}
        disabled={isGameOver || isGuessed}
        onClick={() => addGuessedLetter(letter)}
      >
        {letter}
      </button>
    );
  });

  return <div className="keyboard">{keyboardElements}</div>;
}
```

Walk through this out loud to yourself before moving on: `alphabet` is a
`string` prop, `guessedLetters` is a `string[]` prop, `currentWord` is a
`string` prop, `isGameOver` is a `boolean` prop, and `addGuessedLetter` is
a callback prop, a function taking one `letter: string` parameter and
returning `void`. Inside, `.map()` produces one `JSX.Element` per letter of
the alphabet, and the whole result, `keyboardElements`, is typed as
`JSX.Element[]`.

### Your turn

This is your real final challenge, matching the video's own final
challenge. Build the entire `Keyboard` component from just the props list
above, with the file closed, no peeking at the snippet until you are done
or genuinely stuck for more than ten minutes.

---

## What you actually built

By the end you will have a fully typed clone of Assembly Endgame using:

- Basic and custom TypeScript types
- Typed functions, parameters, and return values
- Generics, specifically through `useState<T>()`
- Union types, for values or return types that can be more than one shape
- The `JSX.Element` type for component return values
- Inline versus custom prop types, and the `ComponentNameProps` convention
- Exporting and importing shared types across files
- The `Omit` utility type
- Typing `.map()` callbacks correctly
- Typing callback props, including ones that take parameters

That list above is a genuinely solid, honest answer to the question "what
is your TypeScript and React experience," the kind of answer that holds up
under a follow up question in an interview, because you built every piece
of it with your own hands rather than watching it go by.

---

## Glossary

**Component** — a function that returns JSX, the basic building block of a
React UI.

**Props** — the input data passed into a component from its parent,
read only from inside the component.

**State** — data owned by a component that can change over time and
triggers a re render when it does, created with `useState`.

**Hook** — a special function starting with `use` that gives a function
component access to React features like state or side effects.

**Callback** — a function passed as a value to be called later, usually in
response to an event; very often also a prop.

**Type** — a description of the shape a value is allowed to take.

**Interface versus type** — this video only used `type`, which is fine and
common. `interface` is a close alternative mostly used for object shapes
that might be extended later; for the purposes of this document treat them
as interchangeable.

**Generic** — a placeholder for a type, filled in at the point of use, like
a parameter but for types instead of values. Written in angle brackets,
e.g. `useState<string>()`.

**Union type** — a type meaning "this value can be one of several listed
types," written with a pipe, e.g. `string | null`.

**Utility type** — a built in TypeScript helper that transforms one type
into another, e.g. `Omit<T, "field">`.

**`any`** — the type TypeScript falls back to when it cannot infer
anything, and which turns off type checking for that value. Treat every
`any` warning as something to fix.

**`void`** — the return type for a function that does not return a usable
value, typically one whose job is a side effect like updating state.

**`JSX.Element`** — the type representing a piece of renderable JSX,
commonly used as a component's return type.

**Inference** — TypeScript automatically figuring out a type from context,
without you writing it explicitly.

---

## A practice plan for the days after you finish this

Given how you described your own pattern (finishing something, then coming
back weeks later and it feeling brand new again), here is what I would
suggest instead of moving straight to a new tutorial.

1. **Day 1 to 2**: Rebuild the entire Assembly Endgame project from this
   document alone, with the video closed the whole time. Use this file as
   your only reference.
2. **Day 3**: Without any reference at all, from memory, write out the
   `KeyboardProps` type, a `useState<string[]>` declaration, and one
   `.map()` with fully typed callback parameters and return value, on a
   blank page or blank file. This is your self test.
3. **Day 4**: Take your own InnoHealth codebase and pick one component that
   is still plain JavaScript or loosely typed, and add proper prop types
   and state types to it, using this document's patterns as your
   reference.
4. **Day 5**: Explain out loud, to yourself or to someone else, the
   difference between a prop, a hook, state, and a callback, using your
   own examples, not the ones from this document. If you can do that
   fluently, this concept is genuinely yours now, not borrowed.

Whenever you are ready for the next video or tutorial, bring me the
transcript again and we will build the next one of these together the same
way.
