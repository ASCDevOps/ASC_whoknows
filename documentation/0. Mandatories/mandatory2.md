# How are you DevOps?
For at forstå DevOps kan man se det som en sammensætning af Development og Operations. For os handler det derfor både om at udvikle software, sørge for stabil drift og automatisere så meget som muligt, så det hele hænger sammen.
Vi arbejder DevOps-orienteret ved at bruge CI/CD pipelines i GitHub Actions, hvor vores kode automatisk bliver bygget, testet og tjekket med linting ved hver pull request.
Derudover bruger vi pre-commit hooks lokalt til at sikre, at vores kode lever op til de samme krav, inden vi overhovedet pusher.
Vi bruger også Docker til at containerisere vores applikation, så vi får et ensartet miljø, uanset hvor den kører.
Selve applikationen kører på en virtuel maskine, hvilket betyder, at vi selv står for både udvikling og drift.
Derudover har vi implementeret password hashing og arbejdet med databaseændringer, hvilket viser, at vi også tænker sikkerhed ind som en del af vores proces – altså et DevSecOps-perspektiv.
