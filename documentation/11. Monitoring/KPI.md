# 16-04-2026

## The task

### CPU load on the server

1. <1%

### Total amount of users

2. Checking the highest database userid. **2078**

### Cost of the infrastructure / complete setup (monthly and/or total so far).

3. VM = 50kr

Domain = free, but around 100kr.

### [Optional] Total amount of active users

4. Efter breach, kan vi se hvilke brugere der er aktive, ved at se hvem der har skiftet adgangskode.

```SELECT COUNT(*) FROM users where must_change_password = 0;```

### [Optional] Average amount of searches per day

5. På simulationen kan vi se hvor mange searches der er på specifikke endpoints hver uge.

[Simulationen](http://158.158.49.35:8000/charts/success/weekly)

* Siden 9. april har der været 333 searches.

* Det er 8 dage siden.

* 333 % 8 = 41,625

* Bring them to class next time and be prepared to answer how you calculated these values
