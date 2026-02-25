# 23-02-2026
* Added Deepsource as a static code analyzer, which can help us analyze software quality with AI.
* Next step is to address issues app.deepsource.com

# 24-02-2026
* Added Continous Intergration for automatic test, so we detect errors early, prevent merging conflicts and make sure tests are always executed.
* I have chosen to stay at checkout/@v4 and setup-go/@v5 because they are more stable and less bugs compared to v6
* We also cache dependencies to reduces build time, make the pipeline faster and more efficient and also so dependencies wont be fetched again on each run.
* When it came to badge the default GitHub Actions badge only displays the status of our workflow, but to have more flexability for the future, we choose to go with https://shields.io.

## DeepSource Issues
**Which ones did you fix? Why?**
* Removed unused method receivers, doesn't make sense to have it if not used.
* Removed unused code, same as other.

**Which ones did you ignore? Why?**
* Removing log.fatal, which will run if database isn't setup correctly.

Log.fatal will run os.exit which will terminate the program, apparently its bad practice to use, but it makes sense for us because we dont want to run the program if there is no database connected.

* The documentation of an exported type should start with the type’s name

I feel like this is something each team should choose for themselves.
Instead of the type's name we have earlier chosen to write the endpoint which uses the type.


**Conclusion: Do you agree with the findings?**
* DeepSource has some uses, it found some possibilites of bugs, but it primarily cleaned up the code.
* Some of its "issues" was kind of intrusive and shouldn't necessarily be "fixed"
