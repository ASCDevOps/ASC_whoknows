### Crawler

#### Initialize Crawler

Make data directory.

```bash
cd crawler
```

```bash
mkdir data
```

Install jsdom.

```bash
npm install jsdom
```
---

#### Use Crawler

First cd into crawler.

```bash
cd crawler
```

To use the crawler.

```bash
node crawler.js
```

---

#### SCP Pages into vm

```bash
cd crawler
```

SCP from local to vm.

```bash
scp data/pages.jsonl appuser@your-vm-ip:/home/appuser/pages.jsonl
```

SSH into the vm and run this.

```bash
python3 - <<'EOF'
import json

with open('/home/appuser/pages.jsonl') as f, open('/home/appuser/pages.sql', 'w') as out:
    for line in f:
        d = json.loads(line.strip())
        title   = d['title'].replace("'", "''")
        url     = d['url'].replace("'", "''")
        content = d['content'].replace("'", "''")
        crawled = d['crawledAt']
        out.write(f"INSERT INTO pages (title, url, language, last_updated, content) VALUES ('{title}', '{url}', 'en', '{crawled}', '{content}') ON CONFLICT (title) DO NOTHING;\n")
print("Done writing pages.sql")
EOF
```

Pipe it into the database.

```bash
docker exec -i compose-db-1 psql -U whoknows -d whoknows < /home/appuser/pages.sql
```

---
