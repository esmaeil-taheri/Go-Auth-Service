# یکی از ویژگی های کلود اپلیکیشن ها disposable بودن هست.
# باید اپلیکیشن statless باشه

# هر پروسسی توی سیستم یه سری متغییر های محیطی داره
# مثل $GOPATH
# os.Getenv("GOPATH")

# همچنین توی ~/.bashrc OR ~/.zshrc
# برای ست کردن export GAMEAPP_PORT="8099" به صورت موقت
# بعد از ست باید یه ترمینال جدید باز کنیم و از اونجا بخونیم مگر اینکه source ~./bashrc
# echo $GAMEAPP_PORT

# GAMEAPP_HOST="localhost" go run main.go حتی با اجرای باینری

# میتونیم کانفیگ رو از جیسون هم لود کنیم
# توی قسمت متغییر های داکر کامپوز اونایی که باید باشن و سکرت نیستن

# متغییر محلی - فایل - دیفالت به صورت ترکیبی
# 1- Default
# 2- File Read
# 3- Get from env (hightest priority)

# Viper - Koanf - GOdontenv


# Mission: Build a Distinctive, Premium Personal Portfolio for Esmaeil Taheri

## 1. Your Role

Act as a multidisciplinary team consisting of:

- Senior Brand Identity Designer
- Award-winning Digital Art Director
- Senior Product Designer
- UX/UI Designer specializing in developer portfolios
- Senior Frontend Engineer
- Technical Copywriter with expertise in software engineering careers

Your mission is to **design and build a complete, production-ready personal portfolio website** for Esmaeil Taheri, a Backend Engineer specializing in scalable systems, financial infrastructure, distributed architectures, and AI backend services.

This must NOT look like a generic developer portfolio, a modified template, a résumé website, or a typical Tailwind landing page.

Build a memorable, sophisticated personal brand expressed through typography, layout, interactions, visual identity, technical storytelling, and exceptional implementation quality.

The website should feel custom-designed specifically for this engineer and his professional background.

**Do not stop at a design proposal, wireframe, or implementation plan. Implement the actual website.**

---

# 2. Understand the Person Before Designing

## Identity

**Name:** Esmaeil Taheri

**Professional title:** Backend Engineer

**Primary expertise:**
- Python Backend Engineering
- Django / Django REST Framework
- FastAPI
- Software Architecture
- Distributed Systems
- Financial Infrastructure
- Event-Driven Systems
- AI Backend Infrastructure

**Additional language:** Go

**Professional experience:** 4+ years, as described in the supplied résumé.

**Email:** Esi.taheri@yahoo.com

**GitHub:** https://github.com/esmaeil-taheri

**LinkedIn:** https://linkedin.com/in/esmaeil-taheri-developer

## Professional positioning

Esmaeil is not simply an API developer.

His background includes engineering backend systems where correctness, consistency, performance, reliability, and architectural decisions directly affect real products.

Important distinguishing experiences:

1. Engineering ledger-based financial wallets for a digital gold platform serving 18M+ users.
2. Implementing atomic financial transactions in an environment handling 50K+ daily transactions.
3. Developing asynchronous settlement and fault-tolerant transaction processing.
4. Building event-driven AI microservices and usage-based credit infrastructure.
5. Applying Clean and Hexagonal Architecture in actual backend projects.
6. Working with authentication, subscriptions, concurrency, data consistency, queues, distributed workflows, and real-time systems.

These experiences should become the foundation of the personal brand.

## Primary audience

The website should communicate effectively with:

- Engineering Managers
- Senior Backend Engineers
- Software Architects
- Technical Recruiters
- CTOs and Technical Founders
- Companies hiring Backend or Platform Engineers
- International engineering teams

The experience should demonstrate technical credibility before the visitor even opens the résumé.

Avoid unsupported positioning such as Principal Engineer, Staff Engineer, or CTO.

---

# 3. Personal Brand Strategy

## Brand name

Use the actual personal name:

**Esmaeil Taheri**

Do not invent a new company-like brand name.

Create a distinctive visual identity around his name, initials, engineering philosophy, and professional personality.

## Brand philosophy

Core idea:

**Engineering Systems That Matter.**

Supporting message:

**Building reliable backends where correctness, architecture, and scale matter.**

Alternative copy may be explored if it is more compelling, but the final positioning must remain credible, technically grounded, and sophisticated.

The brand should communicate:

- Precision
- Reliability
- Systems thinking
- Architectural clarity
- Technical depth
- Calm confidence
- Pragmatism
- Attention to detail

Avoid exaggerated terms like "genius," "visionary," "world-class," or "10x engineer."

## Brand personality

Think of a combination of:

- The discipline of a systems architect
- The precision of financial engineering
- The minimalism of premium digital products
- The clarity of strong technical documentation
- The sophistication of a boutique design studio

The result should feel engineered, not decorated.

## Distinctive visual concept

Develop a cohesive visual language inspired by the structure of software systems:

- Data paths
- Architectural layers
- Connection points
- Transaction records
- Controlled grids
- Structured information
- System boundaries
- Event flows

Use these ideas abstractly and elegantly.

Do not fill the website with literal code snippets, random circuit illustrations, binary numbers, or generic tech graphics.

Consider creating a custom **ET monogram** that subtly references connected components, modular architecture, or controlled data flow.

The monogram must remain recognizable at small sizes and work as a favicon.

Develop a compact brand system consisting of:

- Wordmark treatment
- ET monogram
- Color palette
- Typography pairing
- Spacing system
- Shape language
- Motion principles
- Graphic motifs

Every visual decision should have a coherent reason.

---

# 4. Art Direction

## Design quality

Aim for a premium, editorial, highly polished aesthetic.

The result should be comparable in craftsmanship to thoughtfully designed independent engineering portfolios and premium technology studios.

However, do not copy an existing portfolio's identity or layout.

Create an original composition.

## Visual direction

Preferred overall direction:

**Dark, architectural, minimal, precise, and quietly futuristic.**

Suggested visual foundation:

- Deep graphite or near-black backgrounds
- Soft off-white typography
- Subtle slate surfaces
- Restrained accent colors
- Carefully controlled border contrast
- Generous whitespace
- Strong typographic hierarchy

Explore a distinctive accent color appropriate to the brand.

A muted metallic amber, cool electric blue, or a similarly restrained technical accent may work.

Do not overuse gold merely because of financial-sector experience.

Avoid generic neon purple gradients.

## Typography

Typography is a core brand asset.

Use:

- A refined modern sans-serif for major headings
- A highly readable typeface for body copy
- A restrained monospace font for technical metadata, numbers, and system labels

The hero should be instantly recognizable through excellent typography and composition.

Avoid excessive text gradients, outlined text gimmicks, and oversized decorative punctuation.

## Layout principles

- Editorial composition
- Intentional asymmetry where appropriate
- Consistent visual rhythm
- Strong alignment
- Distinctive section transitions
- Readable content widths
- Meaningful negative space
- Responsive composition, not simple element stacking

Use whitespace as part of the visual identity.

Do not rely on repetitive rounded cards for every section.

## Motion language

Animation should feel like a well-engineered system.

Characteristics:

- Precise
- Smooth
- Subtle
- Intentional
- Responsive
- Never distracting

Use motion to explain relationships and reinforce hierarchy.

Possible ideas:

- Structured reveal animations
- A subtle system-path animation in the hero
- Animated architectural connections
- Carefully choreographed text entrances
- Interactive technical diagrams
- Smooth hover transitions
- Intelligent scroll-based reveals

Avoid excessive parallax, random floating shapes, constant movement, and unnecessary loading animations.

Honor `prefers-reduced-motion`.

---

# 5. Information Architecture

Build a polished portfolio containing:

1. Home
2. Work / Case Studies
3. Expertise / Engineering Approach
4. About / Career
5. Contact

Use dedicated routes for detailed case studies if this improves clarity.

Recommended starting route structure:

- `/` — Main portfolio
- `/work` — Selected engineering work
- `/work/[slug]` — Detailed case study
- `/about` — Professional background

Contact can be a section on the homepage rather than a separate page.

Avoid introducing unnecessary pages.

The homepage should be effective for visitors who spend only 30–60 seconds scanning it.

---

# 6. Homepage — Detailed Requirements

## Section A: Navigation

Create a restrained, high-quality navigation bar.

Include:

- ET brand mark
- Esmaeil Taheri
- Work
- Expertise
- About
- Contact

Include one clear primary action, such as:

**Get in Touch**

Keep navigation simple and elegant.

On mobile, navigation must be easy to use with one hand.

## Section B: Hero

The hero is the most important visual element.

It must immediately communicate identity, specialization, and credibility.

Suggested hierarchy:

Small eyebrow:
BACKEND ENGINEER / SYSTEMS & ARCHITECTURE

Prominent identity:
Esmaeil Taheri

Main statement:
**I build systems where correctness matters.**

Supporting copy:
"I engineer reliable backend systems, from financial ledgers and high-volume transactions to event-driven AI infrastructure."

Include concise CTAs:

- Explore My Work
- GitHub / Contact

A subtle custom architectural visualization may complement the typography.

The visualization could represent connected services, transaction flows, or abstract architectural layers.

It must be original, visually elegant, and lightweight.

Do not render a generic animated code terminal.

Do not use stock photographs.

If a professional portrait is not provided, do not invent one.

## Section C: Credibility / Engineering Impact

Show meaningful professional signals.

Potential metrics:

- 4+ years of experience
- 18M+ users served by the digital gold platform
- 50K+ daily financial transactions handled by that platform
- Financial, AI, and EdTech systems

**Critical accuracy requirement:**

18M+ users and 50K+ daily transactions describe the production platform and operating environment, not personal user acquisition or individually measured performance improvements.

The wording must make this distinction clear.

Use excellent typographic number treatment, not generic KPI dashboard cards.

## Section D: Selected Engineering Work

Create a featured work section presenting substantial engineering problems.

Prioritize:

1. Financial Ledger & Transaction Infrastructure — MelliGold
2. Event-Driven AI Services — Sepas Holding
3. EdTech Platform Architecture — Doino
4. Domain-Specific AI Chatbot — DrillAI

Present these as engineering stories.

Each project should communicate:

- What problem existed
- Why it was challenging
- What was engineered
- Which architectural decisions mattered
- Which technologies were used
- What impact or capability resulted

Distinguish professional work from independent projects.

Avoid presenting projects as if they all had the same commercial scale.

Create distinctive previews using architecture-focused custom graphics rather than generic screenshots.

## Section E: Engineering Philosophy

Create a thoughtful section about engineering principles.

Suggested themes:

**Correctness First**

Financial consistency, reliable transactions, concurrency control, and predictable behavior.

**Architecture With Purpose**

Clear boundaries, maintainable systems, explicit dependencies, and pragmatic abstractions.

**Reliability Under Failure**

Asynchronous processing, retries, transaction states, failure handling, and operational stability.

**Performance Through Understanding**

Efficient data access, appropriate caching, bottleneck analysis, and sound system design.

Keep the writing specific rather than motivational.

Do not imply experience or practices beyond the résumé without supporting evidence.

## Section F: Technical Expertise

Do not create a giant wall of colorful technology logos.

Organize capabilities into thoughtful categories.

### Backend Engineering
Python, Go, Django, Django REST Framework, FastAPI, SQLAlchemy, WebSockets, OpenAPI

### Databases & Infrastructure
PostgreSQL, MySQL, MongoDB, Redis, Elasticsearch, MinIO

### Messaging & Background Processing
Kafka, Celery, Redis

### Software Architecture
Clean Architecture, Hexagonal Architecture, Microservices, Event-Driven Architecture, CQRS

### AI Systems
RAG, AI API integrations, asynchronous AI workloads

### DevOps & Observability
Docker, Nginx, Linux, Jenkins, GitHub Actions, GitLab CI/CD, Prometheus, Grafana

### Security & Quality
JWT, OAuth2, OTP, TOTP, unit testing, integration testing, TDD

Highlight the strongest demonstrated expertise.

Do not imply equal professional depth in every listed technology.

For example, Go is part of the technical skill set and personal projects, but the résumé does not establish the same production depth in Go as in Python.

## Section G: Professional Journey

Design a sophisticated career timeline.

Include:

**Yara E-Commerce Foundation (MelliGold)**  
Backend Engineer  
May 2024 – May 2026

**Sepas Holding**  
Backend Engineer, Remote  
March 2025 – May 2026

**Sigloyland**  
Backend Engineer  
January 2024 – May 2024

**Freelance**  
Backend Developer  
July 2023 – December 2023

Preserve the actual overlapping dates for MelliGold and Sepas Holding.

Do not invent promotions, managerial positions, or current employment.

The timeline should feel editorial, not like a standard résumé widget.

## Section H: About

Write a concise, human, technically credible introduction.

The personality should come through in the way the engineer approaches problems.

Emphasize:

- Enjoyment of complex backend problems
- Interest in clear architecture
- Attention to data consistency
- Building maintainable systems
- Understanding failure scenarios
- Engineering practical solutions

Do not invent hobbies, education, location, personal stories, or motivations not supported by the source materials.

Where personal facts are missing, keep the copy professionally grounded.

## Section I: Contact

Create a strong closing section.

Suggested heading:

**Let's build something reliable.**

Supporting copy should invite technical conversations, interesting engineering opportunities, or collaboration.

Include:

Email: Esi.taheri@yahoo.com  
GitHub: github.com/esmaeil-taheri  
LinkedIn: linkedin.com/in/esmaeil-taheri-developer

Do not invent additional social accounts.

If implementing a contact form, ensure it actually works. Otherwise use email links.

Use a strong final visual composition, not a generic footer contact card.

---

# 7. Case Study Content — Source of Truth

The following information is taken from the résumé and must be represented accurately.

Do not fabricate business impact, benchmarks, implementation details, or architectural decisions.

## Case Study 1: MelliGold — Financial Infrastructure

**Organization:** Yara E-Commerce Foundation (MelliGold)

**Period:** May 2024 – May 2026

**Industry:** Digital gold trading

**Environment:** Platform serving 18M+ users and handling 50K+ daily transactions.

**Primary technologies:**
Django, Django REST Framework, PostgreSQL, Celery.

**Verified engineering contributions:**

- Led migration from a legacy Django-template monolith to a Django REST Framework API.
- Designed a maintainable four-layer architecture.
- Engineered ledger-based wallets for fiat (IRT) and gold balances.
- Represented balance changes as immutable transaction entries.
- Supported separate asset wallets, treasury management, and history-derived balances.
- Implemented atomic buy/sell workflows with PostgreSQL row-level locking.
- Addressed concurrency and race conditions in financial transactions.
- Developed asynchronous settlement workflows with Celery.
- Implemented pending-transaction processing, price validation, retries, and fault-tolerant background processing.
- Automated KYC, identity verification, fraud checks, and user validation.
- Implemented OTP login and optional TOTP-based 2FA.
- Developed automated withdrawal workflows for digital balances and physical gold delivery.

### Case study storytelling

Explain the engineering challenges around:

- Financial correctness
- Concurrent transactions
- Transaction integrity
- Ledger-based accounting
- Asynchronous settlement
- Failure recovery
- Maintainable API architecture

Where useful, illustrate general architectural concepts with clearly labeled conceptual diagrams.

Do not invent the actual proprietary database schema or internal architecture.

Avoid exposing confidential infrastructure information.

## Case Study 2: Sepas Holding — AI Backend Infrastructure

**Period:** March 2025 – May 2026

**Product:** AI content generation platform

**Primary technology:** FastAPI

**Verified engineering contributions:**

- Built authentication, content generation, and subscription services.
- Applied Clean Architecture in a Dockerized microservices environment.
- Integrated OpenAI, Replicate, and custom content-generation workflows through a consistent API.
- Supported text, image, and video generation.
- Implemented Kafka-based event-driven credit processing.
- Ensured credit deductions occurred only after successful content generation.
- Built asynchronous Celery and Redis workflows with retry handling.
- Used PostgreSQL, MongoDB, and MinIO for different storage responsibilities.

### Case study storytelling

Explain:

- The complexity of unifying AI providers
- Long-running asynchronous generation jobs
- Success-dependent credit charging
- Event-driven coordination
- Failure and retry handling
- Service boundaries
- Structured and unstructured data storage

Do not invent throughput, latency, availability, revenue, or cost-saving statistics.

## Case Study 3: Doino — EdTech SaaS

**Type:** Project

**Description:** A platform connecting high-school students with counselors for curriculum planning, performance tracking, and analytics.

**Technologies:**

Python 3.12, FastAPI, PostgreSQL, SQLAlchemy 2.0, asyncpg, Redis, Docker, Jenkins.

**Verified features:**

- Clean / Hexagonal Architecture
- Ports and Adapters
- Explicit Dependency Injection
- Fully asynchronous backend stack
- Redis caching and OTP rate limiting
- Scheduling conflict detection
- Conflict prevention across hourly, volumetric, and video tasks
- Individual and cohort scheduling support
- Tier-based subscriptions
- Capacity-based pricing
- Automated invoicing
- Zibal payment integration
- Passwordless OTP authentication
- JWT access and refresh tokens
- Zero-downtime Docker and Jenkins CI/CD

### Case study storytelling

Emphasize the real-world domain complexity of scheduling, subscription management, authentication, architectural boundaries, and deployment.

Create a conceptual architecture diagram that can be verified against the description.

## Case Study 4: DrillAI — RAG Chatbot

**Type:** Project

**Technologies:** FastAPI, WebSockets, RAG.

**Verified contributions:**

- Developed the chatbot backend for DrillAI.
- Built a RAG pipeline using oil and drilling industry data.
- Used WebSockets for real-time conversational communication.

Highlight domain-specific retrieval and real-time backend communication.

Do not invent retrieval accuracy, model providers, vector database choices, or response-time results.

## Additional Projects

### Quiz Game Backend

Go, Echo, MySQL, Clean Architecture.

Includes layered architecture, dependency inversion, JWT access/refresh authentication, and bcrypt.

### TODO CLI Application

Go, Ports and Adapters, a single dependency-injection composition root, and interchangeable file/in-memory storage.

These should appear in a compact additional-project section.

Do not give small projects the same visual prominence as the major engineering case studies.

---

# 8. Signature Interactive Experience

Create one distinctive, technically meaningful interaction that makes this portfolio memorable.

Recommended concept:

**Architecture Explorer**

An elegant interactive visualization of selected engineering systems.

Allow visitors to explore three conceptual views:

### Financial Ledger Flow

Illustrate a conceptual flow involving:

API request → transaction validation → database transaction → ledger entry → balance derived from history.

### Event-Driven AI Generation

Illustrate a conceptual flow involving:

API → asynchronous generation workflow → provider integration → success event → credit deduction.

### Modular EdTech Architecture

Illustrate:

Presentation / API → Application Use Cases → Domain → Infrastructure Adapters.

The visualization should:

- Be attractive and understandable
- Respond to user input
- Respect accessibility requirements
- Work on mobile
- Avoid excessive animation
- Have meaningful explanatory captions
- Be clearly labeled as conceptual where the exact production architecture is not documented

It must complement the portfolio rather than dominate the entire experience.

If this idea becomes unnecessarily complex or harms performance, implement a simpler and more polished version.

Prioritize craftsmanship over complexity.

---

# 9. Content & Copywriting Standards

All content should be written in professional, natural English suitable for international engineering audiences.

Writing principles:

- Technically accurate
- Concise
- Confident
- Human
- Specific
- No corporate buzzwords
- No exaggerated achievements
- No generic motivational paragraphs
- No artificial-sounding marketing language

Prefer statements like:

"Designed a ledger-based wallet architecture with immutable transaction records."

Avoid statements like:

"Revolutionized the financial industry with groundbreaking technology."

Prioritize engineering problems, solutions, and decisions over lists of tools.

Do not fabricate:

- Project URLs
- Repository URLs
- Testimonials
- Client feedback
- Awards
- Certifications
- University degrees
- Employment status
- Company logos
- Personal photographs
- Performance improvements
- Additional years of experience

If the résumé is attached, use it as the final factual source of truth.

If additional implementation details would improve a case study, identify them as missing content rather than inventing them.

---

# 10. Technical Implementation

Build the portfolio using a modern, maintainable frontend stack.

Recommended:

- Next.js with App Router
- TypeScript
- Tailwind CSS
- Framer Motion / Motion where justified
- Lucide icons
- CSS and SVG for custom technical visualizations

Use the existing project dependencies if an appropriate application already exists.

Do not introduce unnecessary libraries, state management systems, or a backend solely for a static portfolio.

Prefer Server Components by default and Client Components only where interaction is required.

## Architecture requirements

Use a clean project structure with separation between:

- Pages and routes
- Reusable UI components
- Visual system components
- Portfolio content
- Case-study data
- Animation utilities where actually needed
- Shared styles and design tokens

Keep the content easily editable.

Store project information, experience, expertise, and navigation data in clear structured modules or MDX if appropriate.

Avoid hardcoding the same data across multiple components.

Avoid overengineering.

Do not build a CMS or an API layer unless there is a genuine requirement.

## Responsive design

The portfolio must work beautifully on:

- Small mobile screens
- Standard smartphones
- Tablets
- Laptops
- Large desktop monitors

Design mobile deliberately.

Do not merely shrink the desktop design.

Important mobile considerations:

- Hero typography
- Navigation
- Work presentation
- Architecture visualizations
- Readable metrics
- Touch interactions
- Comfortable spacing
- Accessible button sizes
- No horizontal overflow

## Performance

Aim for:

- Lighthouse Performance: 90+
- Accessibility: 95+
- Best Practices: 95+
- SEO: 95+

These are targets to validate, not scores to claim without testing.

Optimize:

- Font loading
- JavaScript bundle size
- Images
- Animation performance
- Core Web Vitals
- Server rendering
- Caching behavior where appropriate

Avoid unnecessary WebGL, canvas animation libraries, or heavy 3D frameworks.

## Accessibility

Implement:

- Semantic HTML
- Proper heading hierarchy
- Keyboard-accessible navigation
- Visible focus states
- Accessible contrast
- Screen-reader descriptions for meaningful visualizations
- Reduced-motion support
- Touch-friendly interactions

## SEO & discoverability

Include:

- Relevant metadata
- Open Graph metadata
- Twitter/X sharing metadata
- Sitemap
- Robots configuration
- Canonical URL support when a domain is available
- Structured data appropriate to a personal professional profile
- A distinctive social sharing image
- ET favicon

Never invent a personal domain. Use configuration for an eventual production URL.

## Resume availability

If the original résumé PDF exists in the project, make it available through a working résumé download link.

If not, do not create a broken download button.

---

# 11. Design Anti-Patterns to Avoid

Do NOT create:

- A generic developer template
- A standard hero with floating technology icons
- A fake terminal interface
- An animated typing-effect introduction
- A random purple-to-blue gradient background
- Dozens of identical skill cards
- A large unstructured technology-logo cloud
- A generic particle background
- A dashboard disguised as a portfolio
- Decorative charts without real meaning
- Unnecessary glassmorphism
- Stock developer illustrations
- Overly complicated navigation
- Excessive motion
- An unreadable dark interface
- A résumé copied into a webpage with no storytelling

Do not equate technical complexity with design quality.

Do not sacrifice usability for novelty.

---

# 12. Expected User Experience

A recruiter should understand within approximately 10 seconds:

1. Who Esmaeil Taheri is.
2. What kind of engineering work he specializes in.
3. Why his experience is credible.
4. Where to inspect his work.
5. How to contact him.

An engineering manager should be able to discover deeper technical substance through the case studies.

A senior engineer should recognize that the content reflects actual backend architecture and implementation challenges rather than superficial buzzwords.

The visual experience must support these goals.

---

# 13. Implementation Workflow

Follow this workflow autonomously.

## Phase 1 — Discovery

Inspect the existing repository, if any.

Read the supplied résumé and available project context.

Identify the strongest evidence-based positioning.

Define the content structure.

## Phase 2 — Art Direction

Establish:

- Visual identity
- Typography
- Color tokens
- Layout rules
- Interaction principles
- Reusable component patterns

Choose one coherent design direction.

Do not produce multiple unrelated concepts or ask for visual preferences unless absolutely necessary.

## Phase 3 — Build

Implement:

- Complete homepage
- Work overview
- Detailed case-study pages
- About section/page
- Contact section
- Responsive navigation
- Brand identity elements
- Custom technical illustrations
- Meaningful interactions
- SEO metadata

Use real résumé content.

No lorem ipsum or meaningless placeholder sections.

## Phase 4 — Refine

Polish:

- Typography
- Alignment
- Spacing
- Component consistency
- Motion timing
- Mobile layouts
- Hover and focus states
- Loading and navigation behavior
- Visual hierarchy

Pay special attention to the first viewport and the featured case studies.

## Phase 5 — Validate

Run the appropriate checks available in the repository.

At minimum:

- Type checking
- Linting
- Production build
- Navigation and link validation
- Responsive layout review
- Accessibility review
- Reduced-motion review

If browser or screenshot tools are available, inspect the rendered website at mobile and desktop widths and fix visible issues.

Do not claim tests passed unless they actually ran.

---

# 14. Deliverables

Provide an implemented, usable website in the repository, not just recommendations.

The result should include:

1. Complete source code
2. Cohesive custom personal brand
3. Responsive pages and components
4. Actual résumé-based content
5. Detailed engineering case studies
6. Custom architecture visuals
7. Functional navigation and contact links
8. SEO configuration
9. README with installation and deployment steps
10. Clear locations for editing content
11. A brief explanation of the design direction
12. A list of any details that still require verification

At the end, explain what was implemented, what checks were performed, and which optional improvements remain.

Do not claim deployment was completed unless deployment actually occurred.

---

# 15. Final Quality Standard

The finished portfolio must feel like a deliberately crafted digital identity for a serious backend engineer.

It should be:

**Distinctive enough to remember.**

**Minimal enough to remain timeless.**

**Technical enough to establish credibility.**

**Human enough to build a personal connection.**

**Fast enough to feel professionally engineered.**

**Clear enough to help turn visitors into real professional opportunities.**

The final experience should communicate one central impression:

**Esmaeil Taheri builds reliable systems, thinks carefully about architecture, and understands the engineering consequences of his decisions.**

Make the design and implementation demonstrate that philosophy.

Build the complete portfolio now.