# React Cheat Sheet for AppInnoHealth

This file focuses only on React as it appears in this project.

Goal:
- help you understand React from the ground up
- show where each React concept appears in your codebase
- help you explain the code in an interview without sounding memorized

Primary React files to study first:
- [src/App.tsx](src/App.tsx)
- [src/main.tsx](src/main.tsx)
- [src/components/sections/Navbar.tsx](src/components/sections/Navbar.tsx)
- [src/components/sections/BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [src/components/sections/HeroSection.tsx](src/components/sections/HeroSection.tsx)
- [src/components/sections/OurServices.tsx](src/components/sections/OurServices.tsx)
- [src/components/sections/ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx)
- [src/components/sections/TeamCarousel.tsx](src/components/sections/TeamCarousel.tsx)
- [src/components/sections/BlogList.tsx](src/components/sections/BlogList.tsx)
- [src/components/sections/BlogPost.tsx](src/components/sections/BlogPost.tsx)
- [src/components/sections/FloatingBlogButton.tsx](src/components/sections/FloatingBlogButton.tsx)
- [src/components/sections/Footer.tsx](src/components/sections/Footer.tsx)

---

## 1) What React Is

React is a JavaScript library for building user interfaces.

In simple terms:
- You break the page into small reusable pieces called components.
- Each component returns what the UI should look like.
- When data changes, React re-renders the changed parts.

In your project, React is used to build:
- the main page sections
- the blog page
- the modal windows
- the forms
- the floating blog button
- the navigation and footer

If an interviewer asks, “What is React?” a good beginner-friendly answer is:

“React is a component-based UI library. In this project, I used it to split the website into reusable sections like the hero, services, navbar, footer, blog list, and forms. React updates the UI when state changes, such as when a modal opens or a carousel advances.”

---

## 2) The React Mental Model

The easiest way to understand React is this:

1. You describe the UI as components.
2. Components can receive data through props.
3. Components can manage their own data through state.
4. React re-renders when state or props change.
5. Effects run when something outside render needs to happen, like a timer or event listener.

In this project, that mental model appears everywhere:
- [App.tsx](src/App.tsx) composes sections and routes.
- [HeroSection.tsx](src/components/sections/HeroSection.tsx) changes images with state and an effect.
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx) switches between patient and doctor forms.
- [BlogList.tsx](src/components/sections/BlogList.tsx) changes visible posts with state.

---

## 3) Components

A component is a function that returns UI.

Example idea:
- one component for a navbar
- one component for a hero section
- one component for a modal
- one component for a form

In your app:
- [Navbar.tsx](src/components/sections/Navbar.tsx) is a component.
- [HeroSection.tsx](src/components/sections/HeroSection.tsx) is a component.
- [Footer.tsx](src/components/sections/Footer.tsx) is a component.
- [BlogPost.tsx](src/components/sections/BlogPost.tsx) is a component.

### Why components matter
- They make code easier to read.
- They make code reusable.
- They make it easier to change one part without breaking everything else.

### What to say in an interview
“React components help me break the app into focused pieces. Instead of one huge file, I separated the homepage into sections so each part has one job.”

---

## 4) JSX

JSX looks like HTML inside JavaScript/TypeScript.

Example:
- React uses JSX to describe the UI.
- JSX is converted into JavaScript behind the scenes.

In your project, JSX is everywhere:
- the routes in [App.tsx](src/App.tsx)
- the buttons and modals in [Navbar.tsx](src/components/sections/Navbar.tsx)
- the cards in [BlogList.tsx](src/components/sections/BlogList.tsx)

### Important JSX rules
- Use className instead of class.
- Wrap expressions in curly braces.
- Return one parent element from a component.
- Self-close tags like img, input, and br.

### Interview answer
“JSX lets me write UI in a syntax that is easy to read. It still compiles to JavaScript, but it feels close to HTML, which makes component structure more intuitive.”

---

## 5) Props

Props are inputs passed from one component to another.

Think of props like arguments to a function.

Example from your project:
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx) passes props to [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx) and [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx).
- [BlogPost.tsx](src/components/sections/BlogPost.tsx) receives file, author, and date as props.
- [AboutUsModalContent.tsx](src/components/sections/AboutUsModalContent.tsx) accepts an optional onClose prop.

### Why props are useful
- They make components reusable.
- They let parent components control child behavior.
- They keep data flow clear.

### Example to explain
In the blog post component, the parent passes the markdown file path and metadata, and the child component renders the content.

### Good interview answer
“Props let me make reusable components. For example, BlogPost does not hardcode the article content. It receives the file path, author, and date from BlogList, so the same component can render different posts.”

### Parent and child relationship
- Parent component owns the data.
- Child component receives it.

---

## 6) State

State is data that changes over time inside a component.

React state tells React when to re-render.

You use state in many places:
- [HeroSection.tsx](src/components/sections/HeroSection.tsx) uses current and direction.
- [OurServices.tsx](src/components/sections/OurServices.tsx) uses activeIndex, hoveredIndex, and isMobile.
- [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx) uses activeIndex, isHovered, and showSwipeHint.
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx) uses role, submitted, and submitting.
- [BlogList.tsx](src/components/sections/BlogList.tsx) uses selectedPost and startIndex.
- [Navbar.tsx](src/components/sections/Navbar.tsx) uses isMenuOpen, showTeam, and showAbout.
- [Footer.tsx](src/components/sections/Footer.tsx) uses many modal visibility states.

### Beginner definition of state
State is like memory for a component.
If the state changes, the component updates.

### Why state matters here
Without state:
- the hero would not auto-slide
- the modal would not open and close
- the selected blog post would not change
- the form buttons would not show loading text

### Good interview answer
“State is used for values that change during user interaction or time. For example, BookAppointmentModal stores whether the user chose patient or doctor, and HeroSection stores which image is currently visible.”

---

## 7) Hooks

Hooks are special React functions that let functional components use React features.

The main hooks used in this project are:
- useState
- useEffect
- useRef

### 7.1 useState
useState is used when you want to store changing data in a component.

Example files:
- [Navbar.tsx](src/components/sections/Navbar.tsx)
- [HeroSection.tsx](src/components/sections/HeroSection.tsx)
- [BlogList.tsx](src/components/sections/BlogList.tsx)
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)

Simple mental model:
- useState gives you a value and a setter.
- When you call the setter, React re-renders.

Example idea:
- const [isMenuOpen, setIsMenuOpen] = useState(false)

What it means:
- isMenuOpen stores whether the mobile menu is open.
- setIsMenuOpen changes it.

Interview answer:
“useState is my main tool for UI state. I use it for toggles, counters, selected items, and modal visibility.”

### 7.2 useEffect
useEffect runs side effects.

A side effect is anything outside rendering, such as:
- timers
- event listeners
- fetching data
- manual DOM interaction

Example files:
- [main.tsx](src/main.tsx)
- [HeroSection.tsx](src/components/sections/HeroSection.tsx)
- [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx)
- [GovernanceAccountability.tsx](src/components/sections/GovernanceAccountability.tsx)
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [BlogPost.tsx](src/components/sections/BlogPost.tsx)
- [Footer.tsx](src/components/sections/Footer.tsx)

What useEffect does in your project:
- adds and cleans up resize listeners
- starts and stops timers
- loads markdown content from a file
- delays the app mount loader

Important rule:
- Always clean up timers and listeners when needed.

Example from your code:
- HeroSection uses setInterval and clears it when the component unmounts.
- Footer and Navbar use window resize listeners and remove them later.
- BlogPost fetches markdown file content when the file path changes.

Good interview answer:
“useEffect is used for actions that React should not do during render. In this app, I use it for image rotation timers, resize listeners, and loading markdown files.”

### 7.3 useRef
useRef gives you a persistent reference that does not cause re-renders when it changes.

Example files:
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)
- [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx)
- [GovernanceAccountability.tsx](src/components/sections/GovernanceAccountability.tsx)
- [TeamCarousel.tsx](src/components/sections/TeamCarousel.tsx)
- [Partnerships.tsx](src/components/sections/Partnerships.tsx)

What it is used for here:
- getting the form element so FormData can be created
- pointing to scroll containers for carousels
- controlling DOM scrolling behavior

Interview answer:
“I use useRef for things that need direct access to DOM elements, like reading a form or scrolling a container, without forcing re-renders.”

---

## 8) Rendering and Re-rendering

React renders a component when:
- it first appears
- its props change
- its state changes

That is why a state update causes the UI to change.

Examples:
- In [BlogList.tsx](src/components/sections/BlogList.tsx), clicking next changes the visible cards.
- In [Navbar.tsx](src/components/sections/Navbar.tsx), clicking the menu button opens the mobile menu.
- In [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx), choosing patient or doctor changes the visible form.

### Beginner explanation
Render means React calculates what the UI should look like right now.
Re-render means React calculates it again after something changed.

### Interview answer
“React re-renders when state or props change. That is how clicking a button can update the page without manually rewriting the DOM.”

---

## 9) Conditional Rendering

Conditional rendering means showing different UI based on a condition.

You use it heavily in this project.

Examples:
- [App.tsx](src/App.tsx) shows different routes.
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx) shows patient form, doctor form, or success message depending on state.
- [Navbar.tsx](src/components/sections/Navbar.tsx) shows the mobile menu only when open.
- [BlogList.tsx](src/components/sections/BlogList.tsx) shows a modal only when a blog post is selected.
- [Footer.tsx](src/components/sections/Footer.tsx) shows different layouts for mobile and desktop.

### Beginner explanation
If condition is true, show one thing.
If it is false, show another thing.

### Interview answer
“I use conditional rendering to switch UI states. For example, the booking modal shows the role selection screen first, then the patient or doctor form, and finally a success message after submission.”

---

## 10) Lists and Keys

When you render many repeated items in React, you usually use map().

Examples:
- [TeamCarousel.tsx](src/components/sections/TeamCarousel.tsx) maps TEAM_MEMBERS.
- [BlogList.tsx](src/components/sections/BlogList.tsx) maps posts.
- [Partnerships.tsx](src/components/sections/Partnerships.tsx) maps partnerLogos and collaborationCategories.
- [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx) maps impactItems.
- [GovernanceAccountability.tsx](src/components/sections/GovernanceAccountability.tsx) maps trustBadges.
- [AboutUsModalContent.tsx](src/components/sections/AboutUsModalContent.tsx) maps bullets in lists.

### Why keys matter
Keys help React identify list items correctly.

### Good interview answer
“I use map for repeated UI like team cards, partner logos, and blog posts. Keys help React track each item during updates.”

### Important warning
Do not use array index as key if the list can reorder often.
In your code, some static lists use index because the items are stable, but the better pattern is a stable ID when available.

---

## 11) Event Handling

React lets you respond to user actions like clicks and input changes.

Examples:
- onClick on navbar buttons
- onSubmit in forms
- onMouseEnter and onMouseLeave in cards and buttons
- onChange in newsletter input
- onScroll in some containers

Files:
- [Navbar.tsx](src/components/sections/Navbar.tsx)
- [Footer.tsx](src/components/sections/Footer.tsx)
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)
- [BlogList.tsx](src/components/sections/BlogList.tsx)

### React event basics
- You receive a synthetic event.
- You can prevent default behavior with e.preventDefault().
- You can stop bubbling with e.stopPropagation().

### In your project
- Forms use e.preventDefault() so the page does not refresh.
- Modal content uses e.stopPropagation() to prevent overlay close when clicking inside.

Interview answer:
“I use React event handlers to control user interactions like opening menus, submitting forms, and closing modals. I stop event propagation when I want the click to affect only the intended element.”

---

## 12) Forms in React

Forms are a major React concept in your project.

You have two main React forms:
- [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)

### How your forms work
1. User fills the form.
2. onSubmit runs.
3. The code prevents default browser submit.
4. FormData is created from the form element.
5. fetch sends data to Formspree.
6. Button text changes while submitting.
7. On success, the modal shows a success message.

### Beginner explanation
React forms are usually controlled or uncontrolled.

In your project, these forms are mostly uncontrolled because you read values from the form element using FormData.

That means:
- the browser manages the inputs
- React does not store each keystroke in state
- the form is easier to build for a simple submission workflow

### Why this is okay here
For a public intake form, this is simpler and easier to maintain than wiring every field to React state.

### Interview answer
“These forms use FormData rather than full controlled state. That keeps the code simpler because I only need the final values at submit time.”

### What to know if asked about controlled vs uncontrolled forms
Controlled form:
- React state stores each input value.
- every keystroke updates state.

Uncontrolled form:
- the DOM stores input value.
- React reads it at submit time.

Your current implementation is closer to uncontrolled.

---

## 13) Form Submission State

In your forms, you store submission status in state.

Example:
- submitting tells the UI whether a request is in progress.
- submitted tells the modal whether to show success.

Why this matters:
- prevents double submit
- gives feedback to the user
- improves user experience

Files:
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- [PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)

Interview answer:
“I keep a separate submitting flag so users know the form is being processed and I can disable the button during the request.”

---

## 14) Portals

A portal lets you render a component outside its normal parent DOM tree.

In your project, portals are used for modals:
- [Navbar.tsx](src/components/sections/Navbar.tsx)
- [NavbarMenu.tsx](src/components/sections/NavbarMenu.tsx)
- [BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- [Footer.tsx](src/components/sections/Footer.tsx)

### Why portals are useful
- modals can escape parent overflow
- modals can sit above everything else
- they help avoid z-index stacking problems

### Beginner explanation
If a modal is nested inside a section with overflow or stacking issues, it might get clipped. A portal moves it to document.body so it behaves like a top-level overlay.

### Interview answer
“I use createPortal for modal dialogs so they render outside the component hierarchy and avoid layout and stacking context problems.”

---

## 15) useEffect Cleanup

This is important to understand.

Whenever you set a timer or event listener inside useEffect, you should usually clean it up.

Examples:
- [HeroSection.tsx](src/components/sections/HeroSection.tsx) clears the interval.
- [Footer.tsx](src/components/sections/Footer.tsx) removes resize listeners.
- [Navbar.tsx](src/components/sections/Navbar.tsx) removes overlay event listeners.
- [ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx) cancels animation frame.
- [GovernanceAccountability.tsx](src/components/sections/GovernanceAccountability.tsx) cancels animation frame.

### Why cleanup matters
Without cleanup:
- timers keep running after component unmount
- event listeners can leak memory
- unexpected behavior may happen later

### Interview answer
“I always clean up timers and event listeners in useEffect. That prevents memory leaks and keeps the UI predictable.”

---

## 16) Refs and the DOM

React usually manages the UI for you, but sometimes you need direct DOM access.

That is where refs help.

Examples:
- formRef in the two forms
- scrollRef in carousel components

### What refs are doing here
- read form contents at submit time
- scroll horizontally with buttons or animation

### Interview answer
“I use refs for DOM tasks that do not belong in state, like reading form fields or scrolling a container.”

---

## 17) Routing and React

React itself does not provide routing. Your project uses react-router-dom.

In [App.tsx](src/App.tsx):
- BrowserRouter wraps the app
- Routes and Route define pages

Current routes:
- / -> homepage sections
- /blog -> blog page

Why routing is useful:
- lets you treat different views as pages
- keeps the homepage and blog separate
- supports deep linking

### Important interview point
Because this is a client-side app, the server must support unknown routes. That is why you also need the redirect rule in [public/_redirects](public/_redirects).

---

## 18) React Composition

Composition means building a page by combining smaller components.

In your app:
- App composes the whole page.
- Navbar composes menu and modal triggers.
- BookAppointmentModal composes patient and doctor forms.
- Footer composes multiple modal contents.

### Why composition is good
- easier to maintain
- easier to reuse
- easier to test

### Interview answer
“I use composition heavily. The app is not one giant component; it is a set of focused sections that are combined in App.tsx.”

---

## 19) Where React Lives in Your Project

If someone asks “Show me React in your code,” point to these files and what they prove.

### App shell
- [src/App.tsx](src/App.tsx): routes and page composition
- [src/main.tsx](src/main.tsx): mounting the React app

### State and effects
- [src/components/sections/HeroSection.tsx](src/components/sections/HeroSection.tsx): timer-driven image carousel
- [src/components/sections/ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx): auto scroll and hover state
- [src/components/sections/GovernanceAccountability.tsx](src/components/sections/GovernanceAccountability.tsx): mobile auto-scroll animation

### User interactions
- [src/components/sections/Navbar.tsx](src/components/sections/Navbar.tsx): menu and modal interaction
- [src/components/sections/BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx): multi-step UI flow
- [src/components/sections/BlogList.tsx](src/components/sections/BlogList.tsx): selected post modal

### Forms
- [src/components/sections/PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)
- [src/components/sections/DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)

### Reusable pieces
- [src/components/sections/index.ts](src/components/sections/index.ts)
- [src/components/ui/button.tsx](src/components/ui/button.tsx)
- [src/components/ui/dropdown-menu.tsx](src/components/ui/dropdown-menu.tsx)
- [src/components/ui/image-with-fallback.tsx](src/components/ui/image-with-fallback.tsx)

---

## 20) Beginner React Terms You Must Know

### Component
A reusable UI block.

### Props
Input passed into a component.

### State
Changing memory inside a component.

### Hook
A React function like useState or useEffect.

### Re-render
React drawing the component again after data changes.

### Effect
Code that runs outside normal rendering, like timers or fetches.

### Ref
A stable pointer to a DOM element or mutable value.

### Portal
Rendering a component outside its parent DOM hierarchy.

### Conditional rendering
Showing different things depending on a condition.

### Controlled form
Input values are stored in React state.

### Uncontrolled form
Input values live in the DOM until read.

### Key
A stable identifier for list rendering.

---

## 21) Common React Interview Questions You Should Be Ready For

Q1. What is a component?
A. A reusable UI unit that returns what should be displayed.

Q2. What is state?
A. Data that changes over time inside a component and triggers re-rendering.

Q3. What is a prop?
A. Data passed from parent to child.

Q4. Why use useEffect?
A. To handle side effects like timers, fetching, and event listeners.

Q5. Why use useRef?
A. To access DOM elements or store values without re-rendering.

Q6. Why use portals?
A. To render overlays and modals outside parent layout constraints.

Q7. Why use map with keys?
A. To render repeated items efficiently and help React track them.

Q8. What is the difference between controlled and uncontrolled forms?
A. Controlled forms store input values in React state. Uncontrolled forms read values directly from the DOM when needed.

Q9. Why do you clean up effects?
A. To prevent leaks and unexpected repeated behavior.

Q10. Why split the app into many components?
A. Better maintainability, reuse, and readability.

---

## 22) React Concepts in This Project, by File

### [src/App.tsx](src/App.tsx)
- router setup
- app composition
- component tree structure

### [src/main.tsx](src/main.tsx)
- root mount
- initial loader delay
- ReactDOM createRoot

### [src/components/sections/Navbar.tsx](src/components/sections/Navbar.tsx)
- menu state
- modal state
- portal rendering
- scroll-to-section behavior

### [src/components/sections/BookAppointmentModal.tsx](src/components/sections/BookAppointmentModal.tsx)
- multi-step form flow
- modal state reset
- conditional rendering
- portal overlay

### [src/components/sections/PatientAppointmentForm.tsx](src/components/sections/PatientAppointmentForm.tsx)
- form submission
- useRef form access
- fetch POST
- submit loading state

### [src/components/sections/DoctorRegistrationForm.tsx](src/components/sections/DoctorRegistrationForm.tsx)
- same as above with a different form layout and fields

### [src/components/sections/HeroSection.tsx](src/components/sections/HeroSection.tsx)
- image index state
- interval effect
- animated transitions

### [src/components/sections/OurServices.tsx](src/components/sections/OurServices.tsx)
- active card state
- mobile/desktop branch
- card transforms

### [src/components/sections/ImpactFocus.tsx](src/components/sections/ImpactFocus.tsx)
- auto-scroll loop
- hover pausing
- scroll indicators

### [src/components/sections/GovernanceAccountability.tsx](src/components/sections/GovernanceAccountability.tsx)
- mobile auto-scroll
- trust badges

### [src/components/sections/TeamCarousel.tsx](src/components/sections/TeamCarousel.tsx)
- mapped data list
- horizontal carousel
- card layout

### [src/components/sections/BlogList.tsx](src/components/sections/BlogList.tsx)
- visible posts state
- blog modal state
- pagination dots

### [src/components/sections/BlogPost.tsx](src/components/sections/BlogPost.tsx)
- fetch markdown content
- render markdown to JSX
- content typography mapping

### [src/components/sections/FloatingBlogButton.tsx](src/components/sections/FloatingBlogButton.tsx)
- useLocation to detect route
- expanded state
- click handling
- timeout cleanup

### [src/components/sections/Footer.tsx](src/components/sections/Footer.tsx)
- multiple modal states
- responsive layout
- portal-based modal contents

---

## 23) How To Explain Your React Knowledge Simply

If you need a short explanation in an interview:

“I understand React as a way to build UI from reusable components. In this project, I used state for changing UI, props for passing data, effects for timers and loading content, refs for forms and scroll containers, and portals for modal overlays. The app is built as a set of modular sections, which makes it easier to maintain and expand.”

---

## 24) Things You Should Practice Saying Out Loud

Practice these until they feel natural:
- What is a component?
- What is state?
- What is a prop?
- What does useEffect do?
- Why did you use a portal for modals?
- How do your forms work?
- Why did you choose local state instead of Redux?
- How does your carousel work?
- Why does the blog page use Markdown files?
- What would you improve next?

---

## 25) A Good Beginner Learning Order for React

Learn React in this order:
1. Components
2. JSX
3. Props
4. State
5. Conditional rendering
6. Lists and keys
7. Event handling
8. useEffect
9. useRef
10. Forms
11. Portals
12. Composition
13. Basic performance and cleanup

If you learn in this order, the project will make more sense.

---

## 26) A Simple Summary of Your React Project

Your app is not using React in a complicated enterprise way.
It uses React in a practical front-end way:
- build a page from sections
- manage UI state locally
- show modals and forms
- animate small interactions
- fetch content from Markdown files
- keep the app deployable as a static site

That is exactly the kind of project you can defend in an interview if you understand the basics deeply.

---

## 27) Next Step After This File

When you are comfortable with this React sheet, the next file should be the TypeScript sheet.

Suggested order:
1. React
2. TypeScript
3. CSS and Tailwind
4. Routing
5. Forms
6. Animation
7. Deployment
8. JavaScript tooling

---

## 28) Final Reminder

Do not try to memorize the entire code.
Instead, learn:
- what each React concept means
- where it appears in your app
- why you used it
- what tradeoff it brings
- how you would improve it next

That is enough to answer confidently.
