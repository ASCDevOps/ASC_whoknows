# How are you DevOps?


## Chris
* To understand DevOps, it can be seen as a combination of Development and Operations. For us, this means both developing software, ensuring stable operation, and automating as much as possible so everything works together.

* We work in a DevOps-oriented way by using CI/CD pipelines in GitHub Actions, where our code is automatically built, tested, and checked with linting on every pull request.

* In addition, we use pre-commit hooks locally to ensure that our code meets the same standards before we even push it.

* We also use Docker to containerize our application, ensuring a consistent environment regardless of where it runs.

* The application itself runs on a virtual machine, which means we are responsible for both development and operations.

* Additionally, we have implemented password hashing and worked with database changes, showing that we also consider security as part of our process – in other words, a DevSecOps perspective.

## Asger

* All 4 psycological safeties are integrated in our team.

1. *Inclusion Safety: Feeling safe to be your authentic self.*
2. *Learner Safety: Feeling safe to learn through questions and mistakes.*
3. *Contributor Safety: Feeling safe to contribute ideas.*
4. *Challenger Safety: Feeling safe to challenge the status quo.*

* End-to-End responsibility, all team members have a stake in both Dev and Ops
*You build it, you run it!*

* Continous learning and improvement, all team members have been through a learning experience, and continoues to learn.

* Culture of collaboration and shared responsibility. We collaborate on features, ask for help, and have a shared responsibility of the code that is written.

* Reduce WIP, we have been getting better, but you can always be better.

* Automatize as much as possible, we are still missing things as Deployment. It will come in the near future (09-04-2026)

### The Three Ways

* **Flow** - Small batch sizes(Small commits), visible work (PR's,), software quality tools (Superlinter, Deepsource), CI/CD/CD.
* **Feedback** - If the simulation fails, we take action, fast. We have monitors setup, to take action even faster.
* **Continual Learning and Experimentation** - We are in a learning process. We are doing new tasks each week from the course.

* We should be better at Lean manufacturing and its value stream. Right now we are pretty slow at getting features out to the customers.
* Why could this be?:

1. Private matters in the team
2. Too big tasks? Everyone should participate, by breaking down the tasks.
3. Easter-vacation, sickness

### Monitoring Realization

* With monitoring via Prometheus and visualization in Grafana, we gained insight into how our system is actually used in practice.
* We became aware of how many things can be monitored – from the VM’s health to what users are searching for.

* Since WhoKnows is a website where search is central, it became clear how important it is to track what is being searched, how often searches occur, and whether errors happen during searches. At the same time, it became important to understand the user experience: can users log in, and what might be going wrong?

* Initially, we used Postman for monitoring, where we received email alerts on failures. Later, we switched to Discord, where we receive instant notifications if something fails – especially if our VM goes down.

* In addition, we monitor our endpoints to understand how our API is used in practice.

* We also became aware of the importance of our VM’s uptime, as it is crucial for the availability of the application.

* Realization – Through monitoring, we became aware of what is actually happening in our system, such as how our endpoints are used and where errors may occur.

* Fix – Based on this, we improved our monitoring by setting up Prometheus and Grafana to better collect and visualize data. At the same time, we moved from Postman to Discord notifications and increased our focus on the most frequently used endpoints.

* Overall, we learned that monitoring is not only about operations, but also about understanding users and improving the system based on their behavior.

* In short: We went from simply observing the system to actively improving it based on data.

### Software Quality

####  DeepSource Issues

**Which ones did you fix? Why?**

- Removed unused method receivers, doesn't make sense to have it if not used.
- Removed unused code, same as other.

**Which ones did you ignore? Why?**

- Removing log.fatal, which will run if database isn't setup correctly.

Log.fatal will run os.exit which will terminate the program, apparently its bad practice to use, but it makes sense for us because we dont want to run the program if there is no database connected.

- The documentation of an exported type should start with the type’s name

I feel like this is something each team should choose for themselves.
Instead of the type's name we have earlier chosen to write the endpoint which uses the type.

**Conclusion: Do you agree with the findings?**

- DeepSource has some uses, it found some possibilities of bugs, but it primarily cleaned up the code.
- Some of its "issues" was kind of intrusive and shouldn't necessarily be "fixed"
