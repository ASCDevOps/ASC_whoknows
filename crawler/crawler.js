import { JSDOM } from 'jsdom';
import { URL } from 'url';
import fs from 'fs';

const visitedUrls = new Set();
const queue = ['https://en.wikipedia.org/wiki/Elasticsearch'];

const MAX_PAGES = 100;
const DELAY_MS = 1000;

const OUTPUT_FILE = './data/pages.jsonl';

function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function savePage(pageData) {
    fs.appendFileSync(
        OUTPUT_FILE,
        JSON.stringify(pageData) + '\n',
        'utf8'
    );
}

async function crawl() {

    while (queue.length > 0 && visitedUrls.size < MAX_PAGES) {

        const currentUrl = queue.shift();

        try {

            const cleanUrl = new URL(currentUrl);
            cleanUrl.hash = '';

            if (visitedUrls.has(cleanUrl.href)) {
                continue;
            }

            visitedUrls.add(cleanUrl.href);

            console.log(
                `Crawling (${visitedUrls.size}/${MAX_PAGES}): ${cleanUrl.href}`
            );

            const response = await fetch(cleanUrl.href, {
                headers: {
                    'User-Agent': 'MyWikiCrawler/1.0'
                }
            });

            if (!response.ok) {
                console.error(`Failed: ${response.status}`);
                continue;
            }

            const html = await response.text();

            const dom = new JSDOM(html);
            const document = dom.window.document;

            // Title
            const title =
                document.querySelector('#firstHeading')?.textContent?.trim() ||
                'No title';

            // Main content
            const contentElements = document.querySelectorAll(
                '.mw-parser-output p, .mw-parser-output h1, .mw-parser-output h2, .mw-parser-output h3'
            );

            const mainContent = Array.from(contentElements)
                .map(el => el.textContent.trim())
                .filter(text => text.length > 0)
                .join(' ')
                .replace(/\s+/g, ' ')
                .trim();

            // Structured document
            const pageData = {
                url: cleanUrl.href,
                title,
                content: mainContent,
                crawledAt: new Date().toISOString()
            };

            // Save locally
            savePage(pageData);

            console.log(`Indexed + Saved: ${cleanUrl.href}`);
            console.log(mainContent.substring(0, 300), '...\n');

            // Find new links
            const links = Array.from(document.querySelectorAll('a'))
                .map(link => link.getAttribute('href'))
                .filter(href =>
                    href &&
                    href.startsWith('/wiki/') &&
                    !href.includes(':')
                );

            for (const href of links) {

                const resolvedUrl =
                    new URL(href, cleanUrl.origin).href.split('#')[0];

                if (
                    !visitedUrls.has(resolvedUrl) &&
                    !queue.includes(resolvedUrl)
                ) {
                    queue.push(resolvedUrl);
                }
            }

            await delay(DELAY_MS);

        } catch (error) {
            console.error(
                `Error crawling ${currentUrl}:`,
                error.message
            );
        }
    }

    console.log(`Finished. Indexed ${visitedUrls.size} pages.`);
}

crawl();
