# TypeScript Cheat Sheet for AppInnoHealth

This sheet is TypeScript-only and beginner-friendly.

Goal:
- make you genuinely confident in TypeScript for your MLH interview
- connect each TypeScript concept to your own project files
- give practical exercises you can do this week

Project context:
- TypeScript is used mainly in .ts and .tsx files under src/
- Your app is React + TypeScript + Vite

---

## 1) What TypeScript Is

TypeScript is JavaScript plus static types.

Simple meaning:
- JavaScript runs at runtime.
- TypeScript helps you catch mistakes before runtime.

Why this helps you:
- fewer bugs
- clearer code contracts
- safer refactors
- better editor support (autocomplete, warnings, jump-to-definition)

In your project, TypeScript is used to type:
- component props
- state values
- refs
- event handlers
- arrays of objects
- reusable utility APIs

Interview answer:
“TypeScript helped me make my React components safer and easier to maintain by defining clear types for props, state, refs, and helper functions.”

---

## 2) Where TypeScript Appears in Your Codebase

Core files to study first:
- [src/main.tsx](src/main.tsx)
- [src/App.tsx](src/App.tsx)
- [src/components/sections/BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [src/components/sections/DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- [src/components/sections/PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)
- [src/components/sections/ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx)
- [src/components/ui/image-with-fallback.tsx](src/components/ui/image-with-fallback.tsx)
- [src/components/ui/button.tsx](src/components/ui/button.tsx)
- [src/components/ui/utils.ts](src/components/ui/utils.ts)
- [vite.config.ts](vite.config.ts)

Important note:
- I did not find a tsconfig file in this repo during search.
- TypeScript still works via Vite defaults, but adding tsconfig is a good improvement for stricter checks.

---

## 3) TypeScript Foundations You Must Know

## 3.1 Primitive types
Examples:
- string
- number
- boolean
- null
- undefined

In your code:
- boolean state: submitting, submitted, isMobile
- number state: activeIndex, current
- string values: URLs, labels, text content

## 3.2 Object types
You use typed object structures often.

Example from [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx):
- impactItems is typed as an array of objects with text and icon.

Why it matters:
- prevents wrong shapes in data arrays
- gives autocomplete while mapping lists

## 3.3 Array types
Example:
- { text: string; icon: LucideIcon }[]

Use this mental model:
- Type[] means array of that type.

## 3.4 Function types
You define callback types in props.

Example from [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx):
- onSuccess: () => void
- setSubmitting: (v: boolean) => void

Why this is excellent:
- parent and child agree on exact function signature
- avoids accidental wrong callback usage

---

## 4) Interfaces and Props (Very Important)

You use interfaces to type component props.

Example from [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx):
- interface BookAppointmentModalProps {
  - open: boolean
  - onClose: () => void
}

Example from [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx):
- interface PatientAppointmentFormProps {
  - onSuccess: () => void
  - submitting: boolean
  - setSubmitting: (v: boolean) => void
  - onBack: () => void
}

Why this is interview gold:
- it proves you understand component contracts
- it shows maintainable React architecture

Interview answer:
“I use interfaces for component props so each component has a clear input contract. That makes reuse safer and reduces integration bugs.”

---

## 5) Union Types

Union type means a value can be one of several options.

Excellent real example in [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx):
- role state type is null | 'patient' | 'doctor'

Why this is powerful:
- stops invalid states
- documents allowed values
- helps conditional rendering stay safe

Beginner explanation:
- without union type, role could be any string
- with union type, role can only be patient, doctor, or null

Interview answer:
“I used a union type for role selection to enforce valid UI states and avoid unexpected values.”

---

## 6) Type Inference vs Explicit Typing

TypeScript can often infer types automatically.

Examples in your code:
- const [submitted, setSubmitted] = useState(false)
  - TypeScript infers boolean

When explicit typing is useful:
- refs
- complex object arrays
- function params in reusable utilities
- unions and optional fields

Balanced approach you already use:
- infer simple things
- explicitly type complex or critical structures

---

## 7) Typing React State

State typing patterns in your code:

## 7.1 Inferred boolean
- useState(false)

## 7.2 Inferred number
- useState(0)

## 7.3 Explicit union
- useState<null | 'patient' | 'doctor'>(null)

Why this matters:
- state type controls what updates are allowed
- catches invalid setState calls during development

Practice thought:
If you accidentally call setRole('admin'), TS should reject it.

---

## 8) Typing Refs

Refs are one of the most useful TS areas in React.

Real examples:
- [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)
- [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx)

Patterns you use:
- useRef<HTMLFormElement>(null)
- useRef<HTMLDivElement>(null)

Why this is good:
- TypeScript knows available DOM properties like scrollLeft, scrollWidth
- prevents calling wrong APIs on wrong element types

Interview answer:
“I type refs explicitly, like HTMLFormElement and HTMLDivElement, so DOM interactions remain safe and discoverable in the editor.”

---

## 9) Event Typing

You use many React events in forms and click handlers.

Example in [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx):
- onSubmit gets event e, then e.preventDefault()

Example in modal components:
- click handlers use stopPropagation

Where explicit typing appears:
- [Footer.tsx](src/components/sections/Footer.tsx) uses React.FormEvent in newsletter submit

Why this matters:
- event type gives access to correct methods and properties
- reduces mistakes with wrong event assumptions

---

## 10) Non-Null Assertion Operator (!)

Real example in [main.tsx](src/main.tsx):
- document.getElementById("root")!

Meaning:
- you are telling TS: this value is not null here

When to use carefully:
- only when you are sure element exists

In this project it is safe because index.html contains root.

Risk:
- if root is missing, app will crash at runtime

Interview answer:
“I used non-null assertion at React root mount because the root element is guaranteed by index.html.”

---

## 11) Type Assertions (as ...)

You use type assertions in places like number input controls.

Example from form files:
- document.getElementById('patient-age-input') as HTMLInputElement

Why used:
- getElementById returns HTMLElement | null
- you need HTMLInputElement methods like stepUp and stepDown

Good practice:
- keep null checks before use (you do this)

Interview answer:
“I used type assertions when DOM APIs return generic HTMLElement, and I narrowed to HTMLInputElement for input-specific methods.”

---

## 12) Optional Props and Optional Values

Example in [AboutUsModalContent.tsx](src/components/sections/AboutUsModalContent.tsx):
- onClose?: () => void

Meaning:
- prop can be passed or omitted

Why useful:
- flexible component reuse

Another pattern in project:
- optional strings for some modal content props

---

## 13) Reusable Type Utilities in Your UI Layer

## 13.1 React.ImgHTMLAttributes
File: [image-with-fallback.tsx](src/components/ui/image-with-fallback.tsx)

You type component props as:
- React.ImgHTMLAttributes<HTMLImageElement>

Why this is great:
- your component accepts normal img props safely
- avoids manually rewriting every prop type

## 13.2 React.ComponentProps
File: [button.tsx](src/components/ui/button.tsx)

You use:
- React.ComponentProps<"button">

Meaning:
- inherit native button prop types automatically

## 13.3 VariantProps from class-variance-authority
File: [button.tsx](src/components/ui/button.tsx)

You combine variant types with native button props.

Interview answer:
“In the UI layer I use utility types like ComponentProps and HTMLAttributes to preserve native element typing while extending behavior.”

---

## 14) Generic Types and Advanced Typing in Your Code

Your [button.tsx](src/components/ui/button.tsx) uses a strong advanced pattern:
- React.forwardRef with generic params

Pattern idea:
- forwardRef<HTMLButtonElement, PropsType>(...)

Why this is strong:
- ref stays typed as button element
- props stay typed and composable

Interview value:
- shows you can work with reusable component libraries in TypeScript

---

## 15) External Library Types

Examples:
- [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx) imports LucideIcon type
- [utils.ts](src/components/ui/utils.ts) imports ClassValue type from clsx

Why this matters:
- you can type project data with library types instead of loose any
- improves API correctness when integrating third-party packages

---

## 16) .ts vs .tsx in This Project

- .ts files: TypeScript without JSX
  - [utils.ts](src/components/ui/utils.ts)
  - [countries.ts](src/components/ui/countries.ts)
  - [nigeriaStates.ts](src/components/ui/nigeriaStates.ts)
  - [vite.config.ts](vite.config.ts)

- .tsx files: TypeScript with JSX
  - most section and UI component files

Interview answer:
“I use .tsx for React components with JSX and .ts for non-UI TypeScript modules.”

---

## 17) TypeScript Patterns You Are Already Using Well

1. Props interfaces per component
2. Callback function typing in props
3. Ref element typing
4. Union state typing for finite UI states
5. Typed arrays of objects for mapped rendering
6. Utility types in reusable UI components
7. Safe null checks before DOM operations

These are strong practical TypeScript practices.

---

## 18) Common TypeScript Risks in This Project (And How To Defend)

## Risk 1: No explicit tsconfig file found
Defense:
- project still compiles with Vite defaults
- next improvement is adding tsconfig with stricter options

## Risk 2: Some assertions with as HTMLInputElement
Defense:
- assertions are limited and guarded with null checks
- future improvement: use refs for those inputs instead of getElementById

## Risk 3: Some inferred any in callback params in large files
Defense:
- core paths are typed; can further tighten event/handler signatures incrementally

Good interview wording:
“I used TypeScript practically and safely, and I can improve strictness further by adding explicit tsconfig and reducing DOM assertions.”

---

## 19) TypeScript Interview Questions and Answers (Project-Specific)

Q1. Why TypeScript instead of JavaScript only?
A. It catches type mistakes early and makes component contracts clearer, especially with many forms and modal flows.

Q2. How do you type component props?
A. I define interfaces per component, for example BookAppointmentModalProps and form props interfaces.

Q3. Give an example of union types in your code.
A. role state in the booking modal uses null | 'patient' | 'doctor'.

Q4. How do you type refs?
A. I use generics like useRef<HTMLFormElement>(null) and useRef<HTMLDivElement>(null).

Q5. Where did you use type assertions?
A. In form stepper controls where getElementById is narrowed to HTMLInputElement for stepUp/stepDown.

Q6. Why use optional props?
A. Some modal content components can be reused with or without a close callback.

Q7. Explain non-null assertion in your project.
A. I used ! on getElementById('root') in main.tsx because index.html guarantees that element.

Q8. What advanced typing do you use?
A. forwardRef generics and utility types like React.ComponentProps and VariantProps in reusable button component.

Q9. How would you improve TS strictness next?
A. Add explicit tsconfig, enable strict options, and replace DOM id assertions with refs.

Q10. What practical value did TS give this project?
A. Safer integration across many components, especially props callbacks and ref-based DOM interactions.

---

## 20) Hands-On Practice Exercises (Beginner to Advanced)

Do these one by one in this project. These are practical and interview-focused.

## Exercise 1: Identify 10 interfaces and explain each
Task:
- Open files in sections and list every interface.
- For each, write what each property means.

Expected files:
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)
- modal content files

## Exercise 2: Strengthen one prop type
Task:
- Pick one component currently using broad typing and tighten it.
- Example: add stricter literal unions for modal titles/categories where useful.

## Exercise 3: Replace one type assertion with refs
Task:
- In patient or doctor form, replace getElementById + as HTMLInputElement with a typed ref.
- Keep behavior same.

Learning outcome:
- safer DOM handling pattern.

## Exercise 4: Create a reusable type alias
Task:
- Add type alias for common callback signature:
  - type VoidCallback = () => void
- Reuse in 2-3 interface files.

Learning outcome:
- consistency in prop typing.

## Exercise 5: Add readonly to data arrays
Task:
- For static lists (like impact items), try readonly typing.

Learning outcome:
- prevent accidental mutation.

## Exercise 6: Add explicit return types to key functions
Task:
- In 3 files, add explicit return types to helper functions and handlers.

Learning outcome:
- clearer function contracts.

## Exercise 7: Build a type-safe helper
Task:
- Create a small helper in a ui or utils file with typed input and output.
- Example: formatDate(date: string): string

## Exercise 8: Create one discriminated union
Task:
- For modal content state, model UI state as discriminated union instead of multiple booleans in one component.

Learning outcome:
- better state correctness.

## Exercise 9: Add a local tsconfig (practice branch)
Task:
- Create tsconfig with strict true and run build.
- Record errors and classify them.

Learning outcome:
- understand strict mode migration.

## Exercise 10: Explain every TS keyword you used
Task:
- Build your own glossary for: interface, type, union, optional, generic, assertion, non-null assertion, utility type.

---

## 21) 7-Day TypeScript Practice Plan

Day 1:
- Learn basic TS types and interfaces
- Review all props interfaces in this project

Day 2:
- Focus on union types, optional props, and function types
- Explain role state in booking modal out loud

Day 3:
- Focus on refs and event typing
- Refactor one getElementById assertion into useRef

Day 4:
- Study utility types in [button.tsx](src/components/ui/button.tsx)
- Understand forwardRef generics

Day 5:
- Create 3 small typed helpers in a practice file
- Add explicit return types

Day 6:
- Mock TypeScript interview using section 19 questions
- Record your answers and tighten weak points

Day 7:
- Do mini refactor with type improvements and summarize tradeoffs

---

## 22) TypeScript Mini Glossary (Say These Clearly in Interview)

interface:
- defines object shape contract

type alias:
- names a type expression, useful for unions and reusable signatures

union type:
- value can be one of multiple types

optional property:
- property may be undefined

generic:
- type parameter for reusable typed logic/components

type assertion:
- tells TS to treat value as a more specific type

non-null assertion (!):
- tells TS value is not null/undefined here

utility type:
- helper type from React/libraries to compose types safely

---

## 23) How To Present Your TypeScript Confidence to MLH

You can say:

“I chose TypeScript because I use it actively in this project for component contracts, state safety, and DOM/ref typing. I can explain practical concepts like interfaces, unions, callback typing, utility types, and ref generics with examples from my codebase.”

That is strong and honest.

---

## 24) Quick Self-Test (Before Interview)

If you can answer yes to most of these, you are ready:
- Can I explain interface vs type?
- Can I explain union type with my booking modal example?
- Can I explain why useRef<HTMLFormElement> is useful?
- Can I explain non-null assertion in main.tsx?
- Can I explain type assertion and its risk?
- Can I explain utility types in button.tsx?
- Can I explain one TypeScript improvement I would make next?

---

## 25) Final Reminder

Do not memorize syntax only.
Focus on this formula:
- What type did I choose?
- Why did I choose it?
- What bug does it prevent?
- How would I improve it?

That is what interviewers care about most.
