# 23-02-2026
* Added Deepsource as a static code analyzer, which can help us analyze software quality with AI.
* Next step is to address issues app.deepsource.com

# 24-02-2026
* Added Continous Intergration for automatic test, so we detect errors early, prevent merging conflicts and make sure tests are always executed.
* We also We cache dependencies to reduces build time, make the pipeline faster and more efficient and also so dependencies wont be fetched again on each run.
* When it came to badge the default GitHub Actions badge only displays the status of our workflow, but to have more flexability for the future, we choose to go with https://shields.io.