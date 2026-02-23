# Create a dependency graph
![Dependency Graph](<../1. Understanding system/dependencygraph1.svg>)

# Problems with the codebase
[Problems with the codebase](<../1. Understanding system/Problems_with_the_codebase.md>)

# Generate your own OpenAPI specification
[OpenApi Specification](openapi.yaml)

Made in .yml which is originally used by Swagger 2.0 but is still supported by OpenAPI 3.0 (JSON)

# Choose a branching strategy
[Blog used for inspiration](https://blog.prateekjain.dev/the-ultimate-guide-to-git-branching-strategies-6324f1aceac2)

## Github Flow
**Workflow:**
1. Create **Feature Branches** from default branch. 
2. Push commits to the **feature branch**.
3. Open a **pull request** for team and AI code reviews.
4. Once approved, merge into ``` main ```
5. Deploying immediately is encouraged **(Continous Deployment)**

This means that everything in ``` main ``` branch should be production ready.

**Pros:**
1. The branching strategy is simple, which is important for a small team.
2. Strong pipelines is plays well with Github Flow.
3. Encourages Continous Deployment => makes it DevOps friendly.
4. Avoids long-lived branches => DevOps Friendly


**Cons:**
1. Highly dependent on a robust pipeline & testing so we ensure no bugs goes unnoticed into ``` main ```
=> Until we have a strong pipeline (automated test, linting, static analysis) code reviews and testing rely heavily on manual effort, which takes time.

**Conclusion:**
<br>
Github Flow seems to align the best with our needs for a suitable branching strategy, due to its simplicity, DevOps friendliness and pipeline focus.

Although it requires manual review early (before pipelines), this can be viewed as a pro. Through manual code reviews we seek to strengthen **culture of collaboration** and **shared responsibilty**.

Before pipelines, the manual code reviews will also support **transparency, visibility and knowledge sharing**.

Strategy is enforced by applying branch protection rules on default branch and making it mandatory to have atleast 2 reviewers on a PR.