# MLH Fellowship Interview Cheat Sheet

This guide is tailored to your project:
https://github.com/mudimurtala/AppInnoHealth

Goal: help you explain this project confidently and honestly in a technical interview.

---

## 1) How Your Project Matches MLH Code Sample Requirements

### Representative of your abilities
- You can explain architecture, routing, forms, animations, responsive behavior, and deployment.
- You should be honest that AI helped speed up implementation, but you understand and can defend the final code decisions.

### Existing sample
- This project was built before this application and is already deployed.

### Public on GitHub
- Repository is public.

### Multiple files and substance
- Multi-file React + TypeScript app with many components, forms, modals, and deployment config.

### Real problem
- Real healthcare communication and outreach web app for InnoHealth Africa Technology.

### No notebook
- This is a deployable web app, not a notebook.

### Aligned language
- Primary language: TypeScript.
- Also uses JavaScript tooling files and CSS.

---

## 2) Project Architecture You Must Be Able To Explain

### Entry flow
1. index.html bootstraps app and preloads fonts/images.
2. src/main.tsx mounts React after a minimum loader duration.
3. src/App.tsx defines routes and section composition.

### Routing
- Uses BrowserRouter.
- Routes:
  - / -> Home sections
  - /blog -> Blog list and modal post reader

### Component structure
- src/components/sections/ contains page sections and modals.
- src/components/ui/ contains reusable UI primitives.

### Data and content model
- Most content is local static arrays in components.
- Blog content is Markdown files in public/blog loaded with fetch and rendered with react-markdown.

### Forms
- Two React forms in modal flow:
  - PatientAppointmentForm
  - DoctorRegistrationForm
- One static fallback HTML form in public/book-appointment.html
- Submission target uses Formspree endpoint.

### Styling system
- Tailwind + extensive inline style objects.
- Brand tokens in src/brand.css and Tailwind theme extension.
- Mobile/desktop branching done mostly with runtime screen-width checks and media-query blocks.

### Deployment
- Vite build output to build/ directory.
- Netlify settings in netlify.toml.
- SPA redirect in public/_redirects.

---

## 3) Full Repository Map (What Each Folder/File Is For)

## Top level
- README.md: project overview and setup steps.
- index.html: app shell, preloads, initial loading UI.
- package.json: dependencies and scripts.
- postcss.config.js: PostCSS pipeline.
- tailwind.config.js: Tailwind theme extension and content paths.
- vite.config.ts: Vite plugins, aliases, build/server config.
- netlify.toml: Netlify build/publish settings.
- .gitignore: ignored files.
- .github/copilot-instructions.md: project AI guidance.

## Public assets and static files
- public/_redirects: SPA route fallback.
- public/book-appointment.html: standalone HTML appointment form.
- public/blog/*.md: blog article Markdown source files.
- public/blog-images/*: blog card image assets.
- public/images/branding/*: logo, iconmark, favicon assets.
- public/images/hero/*: hero carousel images.
- public/images/partners/*: partner logos.
- public/images/team/*: team member photos.

## Source root
- src/main.tsx: React root render + loader timing logic.
- src/App.tsx: Router + section layout.
- src/index.css: Tailwind generated styles + imports.
- src/brand.css: design tokens and reusable brand classes.
- src/Attributions.md: third-party attribution notes.
- src/guidelines/Guidelines.md: template guidelines placeholder.

## Section components (src/components/sections)
- AboutUsModalContent.tsx: About modal long-form content.
- BlogList.tsx: featured blog carousel and post modal trigger.
- BlogPost.tsx: Markdown renderer and typography mapping.
- BookAppointmentModal.tsx: role selection + form flow wrapper.
- ComingSoonModalContent.tsx: generic coming-soon modal.
- ContactModalContent.tsx: contact modal and social links.
- DoctorRegistrationForm.tsx: doctor registration form and Formspree POST.
- FloatingBlogButton.tsx: animated floating CTA to /blog.
- Footer.tsx: footer layout, links, socials, and several modals.
- GetInvolvedModalContent.tsx: get involved modal content.
- GovernanceAccountability.tsx: governance section and badge slider.
- HeroSection.tsx: hero image carousel with AnimatePresence.
- HowWeWork.tsx: hybrid model section and office/trust cards.
- ImpactFocus.tsx: auto-scrolling impact badge strip.
- Navbar.tsx: top nav, mobile menu, and portal modals.
- NavbarMenu.tsx: dropdown action menu and modal triggers.
- OurServices.tsx: 3D-style service carousel.
- Partnerships.tsx: collaboration categories and partner logo scroller.
- PatientAppointmentForm.tsx: patient booking form and Formspree POST.
- PrivacyPolicyModalContent.tsx: privacy bullet modal.
- ProgramsModalContent.tsx: programs modal content.
- TeamCarousel.tsx: horizontal team card carousel.
- TermsModalContent.tsx: terms modal content.
- index.ts: barrel exports for sections.

## UI primitives (src/components/ui)
- button.tsx: cva-powered button variants.
- dropdown-menu.tsx: Radix dropdown wrapper components.
- utils.ts: className merge helper (clsx + tailwind-merge).
- image-with-fallback.tsx: fallback image component.
- countries.ts: country options array.
- nigeriaStates.ts: Nigeria states array.
- CustomSelect.tsx: currently empty file.

---

## 4) Stack-by-Stack: What You Must Understand Before Interview

## React
- Functional components and hooks usage (useState, useEffect, useRef).
- Component composition through section exports.
- Conditional rendering for mobile and modal states.
- Portal rendering with createPortal for modals.
- Event handling and propagation control in overlays.

## TypeScript
- Prop interfaces for components and forms.
- Typed refs like useRef<HTMLFormElement>(null).
- Typed union states, for example role state in appointment modal.
- Types around icon components and utility props.

## Routing
- BrowserRouter/Routes/Route setup.
- Why one route is section-stack and another route is blog page.
- Why static host needs redirect fallback for SPA routes.

## Styling
- Why both Tailwind and inline style objects are used.
- Brand variable system in brand.css.
- Responsive behavior via:
  - runtime isMobile checks
  - CSS media queries
- Tradeoff: inline styles improve control but can reduce consistency.

## Animation
- Framer Motion in hero transitions.
- CSS keyframes in floating button.
- requestAnimationFrame loops in scrolling components.

## Forms + HTTP
- FormData construction and POST with fetch.
- Form submission state handling (submitting, success, error).
- Endpoint management and how to safely replace endpoints.

## Build and deployment
- npm run dev and npm run build.
- Vite bundling and output folder.
- Netlify build config and redirect behavior.

---

## 5) 40 Interview Questions With Strong, Honest Sample Answers

Use these as memory anchors. Do not memorize word-for-word; understand the logic.

## A) Architecture and project choices

Q1. Why did you choose React + TypeScript + Vite?
A. I wanted a fast UI-focused stack. Vite gives fast local startup, React gives component structure, and TypeScript improves maintainability by catching prop and state errors early.

Q2. Why not Next.js?
A. This project is primarily a static marketing and engagement app with client-side interaction. I did not need SSR complexity for this version.

Q3. How is your app structured?
A. App.tsx handles routes and composes section components. Each section is isolated in src/components/sections so UI changes stay localized.

Q4. How do users navigate the app?
A. Primary navigation is section scrolling on home route and a dedicated /blog route for blog posts.

Q5. How do you avoid deeply coupled components?
A. I keep local state within each section and pass explicit props to form/modal components rather than global shared mutable state.

## B) React and state

Q6. Why local state instead of Redux?
A. The state is mostly UI-local and ephemeral: modal open/close, active carousel index, and form submission state. Global state manager would add unnecessary complexity here.

Q7. Explain one useEffect you are confident about.
A. In HeroSection, useEffect creates an interval for slide changes and clears it on unmount. That cleanup prevents timer leaks.

Q8. How do you handle modals?
A. I render modal overlays with createPortal into document.body, so they are not constrained by parent stacking contexts.

Q9. How do you prevent closing modal when clicking inside content?
A. The overlay handles close on click, while content wrapper calls stopPropagation to keep clicks inside from bubbling.

Q10. What is one improvement you would make to modal state?
A. I would centralize modal state logic to reduce duplication between Navbar and Footer.

## C) TypeScript

Q11. Where did TypeScript help you most?
A. It helped define strict props for forms and modals, and typed refs prevented null mistakes when reading form elements.

Q12. Show an example of a useful type pattern.
A. The role state in BookAppointmentModal uses a union type for patient, doctor, or null to keep transitions explicit.

Q13. Any weak TypeScript areas in this project?
A. Some style and event patterns could be tightened further, and there are places where stronger shared types would reduce repeated prop declarations.

## D) Forms and API

Q14. How do form submissions work?
A. On submit, I collect FormData from formRef and POST to Formspree endpoint with Accept application/json. Then I branch on response status.

Q15. How do you communicate submit progress to users?
A. I use submitting boolean to disable button and switch button label text.

Q16. Why use Formspree?
A. It gives a quick reliable form backend for a static frontend deployment.

Q17. What are risks of this approach?
A. Limited server-side customization, potential spam risk, and less control than a custom backend.

Q18. How would you harden form handling?
A. Add better input validation, anti-spam measures, and user-friendly inline error UI instead of alert-only feedback.

## E) Styling and responsive design

Q19. How do you keep brand consistency?
A. I define CSS variables in brand.css and map key colors in Tailwind theme, then reuse those across sections.

Q20. Why mix Tailwind and inline styles?
A. Tailwind is great for quick structure, while inline styles handle custom gradients and dynamic values. Tradeoff is consistency and maintainability.

Q21. How do you handle mobile differences?
A. I use runtime width checks for behavior differences and CSS media queries for style differences.

Q22. What responsive bug do you watch for?
A. Horizontal overflows in carousels and modal viewport clipping on small screens.

Q23. How do you reduce layout shifts?
A. I preload key images/fonts in index.html and constrain image containers.

## F) Animation and interaction

Q24. Explain your hero animation briefly.
A. AnimatePresence transitions between images with directional motion and spring timing. It gives smooth, readable movement without complex state.

Q25. Why use requestAnimationFrame in sliders?
A. It syncs animation with browser paint for smoother motion and better performance than fixed intervals for continuous scrolling.

Q26. What animation tradeoff did you accept?
A. Rich motion improves experience but increases complexity. I kept animations mostly component-local.

## G) Blog system

Q27. How is blog content stored and rendered?
A. Markdown files in public/blog are fetched by path and rendered with react-markdown in BlogPost.

Q28. Why is this useful?
A. Easy content edits without changing component code.

Q29. What are the limits?
A. Not a full CMS workflow; scaling many posts needs better indexing/search and likely a content backend.

## H) Deployment and DevOps

Q30. How does SPA routing work on Netlify?
A. public/_redirects routes all paths to index.html with 200, then React Router handles route selection client-side.

Q31. What does netlify.toml do?
A. Defines build command and publish directory.

Q32. What does vite.config.ts contribute?
A. Plugin setup, alias configuration, build target/outDir, and local dev server behavior.

## I) Quality, risk, and honesty

Q33. What is missing technically?
A. Automated tests, stronger accessibility audits, centralized constants, and observability/monitoring.

Q34. If asked about AI assistance, what do you say?
A. I used AI to speed implementation and explore options, but I reviewed, edited, and integrated the final code. I can explain architecture, tradeoffs, and follow-up improvements.

Q35. If interviewer asks something you do not know?
A. I explain what I do know, state uncertainty clearly, and propose a concrete way I would verify or fix it.

Q36. One design decision you would change now?
A. I would reduce duplicated modal code and move repeated inline style blocks into reusable styled components or theme utilities.

Q37. One performance improvement you would add?
A. Code-split heavy components and lazy-load non-critical modal content.

Q38. One accessibility improvement you would add?
A. Better keyboard focus trapping for modals, stronger focus-visible styles, and automated a11y checks.

Q39. One maintainability improvement you would add?
A. Introduce a constants layer for section data and shared style tokens to avoid duplication.

Q40. Why are you a good candidate based on this project?
A. I can ship a real user-facing TypeScript project end-to-end, from UI architecture and responsive behavior to integration and deployment, and I can discuss tradeoffs honestly.

---

## 6) Risky Questions You Should Prepare For (With Safe Answers)

### Risk: “Did you write all this yourself?”
Safe answer:
- I used AI as a coding assistant for speed and iteration, like many developers use modern tooling.
- I reviewed and adjusted outputs, and I can explain the architecture and decisions in detail.

### Risk: “Why so many inline styles?”
Safe answer:
- It provided precise control for custom visuals and animation states quickly.
- I would refactor repeated style patterns into reusable style utilities/components as the next cleanup step.

### Risk: “Where are your tests?”
Safe answer:
- This version is manually tested, but I agree tests are the next maturity step.
- I would start with Vitest for utilities/components and Playwright for critical flows (forms and modal behavior).

### Risk: “How do you handle accessibility?”
Safe answer:
- I use semantic elements and aria labels in key places, and modal overlays with explicit close controls.
- I would still run axe-based audits and improve focus trap/keyboard behavior.

### Risk: “How scalable is this architecture?”
Safe answer:
- For current app size, local section state is simple and stable.
- As complexity grows, I would introduce context/query layers and split large components.

---

## 7) Exact Files To Revisit Before Interview (High Priority)

Read and be ready to explain these first:
1. src/App.tsx
2. src/main.tsx
3. src/components/sections/Navbar.tsx
4. src/components/sections/Footer.tsx
5. src/components/sections/BookAppointmentModal.tsx
6. src/components/sections/PatientAppointmentForm.tsx
7. src/components/sections/DoctorRegistrationForm.tsx
8. src/components/sections/HeroSection.tsx
9. src/components/sections/OurServices.tsx
10. src/components/sections/ImpactFocus.tsx
11. src/components/sections/BlogList.tsx
12. src/components/sections/BlogPost.tsx
13. tailwind.config.js
14. src/brand.css
15. vite.config.ts
16. netlify.toml
17. public/_redirects
18. index.html

---

## 8) 7-Day Study Plan (Practical and Realistic)

Day 1
- Explain architecture flow out loud from index.html to App routes.
- Draw component tree on paper.

Day 2
- Practice forms and modal flow.
- Be able to explain each form field and submit lifecycle.

Day 3
- Practice animation and carousel logic in Hero, Services, Impact, Team.
- Explain useEffect cleanup and requestAnimationFrame choices.

Day 4
- Deep dive styling system: brand.css + tailwind.config + inline tradeoffs.

Day 5
- Deployment and routing: Vite build, Netlify publish, redirects.
- Practice answering what would break without redirects.

Day 6
- Run mock interview Q and A from section 5.
- Focus on difficult honesty questions (AI usage, tests, scalability).

Day 7
- Final review and concise storytelling:
  - Problem
  - Solution
  - Architecture
  - Tradeoffs
  - Improvements

---

## 9) 60-Second Project Pitch You Can Memorize

I built InnoHealth Africa Technology as a TypeScript React web app to communicate healthcare services, impact, and partnerships while supporting user engagement through appointment and registration forms. I structured it as modular section components with route-based navigation, responsive behavior, and interactive UI patterns like modals and carousels. I used Formspree for backendless form processing, Markdown-powered blog content for easy publishing, and Netlify deployment with SPA redirects. The project demonstrates practical front-end engineering across architecture, UX, integration, and deployment, and I can clearly discuss the tradeoffs and next improvements such as testing, stronger accessibility automation, and state consolidation.

---

## 10) Honesty Script About AI Assistance (Recommended)

I used AI tools to speed up coding and design iteration, especially for boilerplate and UI experimentation. I still reviewed, integrated, and modified the code to fit the project requirements. I understand the architecture and decisions and can explain or refactor the key parts confidently.

This answer is honest and professional.

---

## 11) Final Checklist Before You Submit To MLH

- Repo is public and README is accurate.
- You can explain 18 high-priority files from section 7.
- You can answer at least 25 of the 40 questions confidently.
- You can defend tradeoffs and name concrete improvements.
- You can clearly state what part AI helped with and what you validated.

If you can do this, you are interview-ready.
