# 23-02-2026
* Added Deepsource as a static code analyzer, which can help us analyze software quality with AI.
* Next step is to address issues app.deepsource.com

# 24-02-2026

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