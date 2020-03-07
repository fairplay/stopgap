package main

import (
    "fmt"
    "bytes"
    "io/ioutil"
    "math"
    "os"
    "strings"
    "time"
    "github.com/russross/blackfriday"
)

type Config struct {
    Title string
    Description string
    Email string
}

var config Config

func getConfig() {
    var config_md = strings.Split(string(getFile("_sections/config.md")), "\n")
    config.Title = config_md[0]
    config.Description = config_md[1]
    config.Email = config_md[2]
}

func getLayoutStart(title string) string {
    return fmt.Sprintf(string(getFile("_sections/start.md")), title, config.Title, config.Description)
}

func getLayoutEnd() string {
    return fmt.Sprintf(string(getFile("_sections/end.md")), config.Email)
}

func getFile(f string) []byte {
    b, err := ioutil.ReadFile(f)

    if err != nil {
        panic(err)
    }

    return b
}

func getDir(dir string) []os.FileInfo {
    p, err := ioutil.ReadDir(dir)

    if err != nil {
        panic(err)
    }

    return p
}

func writeFile(fileName string, b bytes.Buffer) {
    err := ioutil.WriteFile(fileName+".html", b.Bytes(), 0644)

    if err != nil {
        panic(err)
    }
}

func formatDate(date string) string {
    layout := "2006-01-02"
    t, _ := time.Parse(layout, date)
    return t.Format("Jan 02, 2006")
}

func getPostMeta(fi os.FileInfo) (string, string, string) {
    id := fi.Name()[:len(fi.Name())-3]
    date := fi.Name()[0:10]
    title := strings.Split(string(getFile("_posts/"+fi.Name())), "\n")[0][2:]

    return id, formatDate(date), title
}

func getPageMeta(fi os.FileInfo) (string, string, bool) {
    fileName := fi.Name()
    runes := []rune(fileName)
    if runes[0] == '.' {
        return "", "", true
    }
    id := fi.Name()[:len(fileName)-3]
    title := strings.Split(string(getFile("_pages/"+fi.Name())), "\n")[0][2:]

    return id, title, false
}

func writeIndex() {
    var b bytes.Buffer
    b.WriteString(getLayoutStart(config.Title))
    writePostsSection(&b)
    writePagesSection(&b)
    b.WriteString(getLayoutEnd())
    writeFile("index", b)
}

func writePostsSection(b *bytes.Buffer) {
    b.WriteString("<h2>Posts</h2><nav class=\"posts\"><ul>")

    posts := getDir("_posts")
    limit := int(math.Max(float64(len(posts))-5, 0))

    for i := len(posts) - 1; i >= limit; i-- {
        fileName, date, title := getPostMeta(posts[i])

        b.WriteString("<li><span class=\"date\">" + date +
            "</span><a href=\"posts/" +
            fileName + ".html\">" +
            title + "</a></li>\n")
    }

    b.WriteString("</ul></nav><p class=\"all-posts\"><a href=\"all-posts.html\">All posts</a></p>")
}

func writePagesSection(b *bytes.Buffer) {
    b.WriteString("<h2>Pages</h2><nav class=\"pages\"><ul>")

    pages := getDir("_pages")

    for i := 0; i < len(pages); i++ {
        id, title, skip := getPageMeta(pages[i])
        if skip {
            continue
        }

        b.WriteString("<li><a href=\"pages/" +
            id + ".html\">" +
            title + "</a></li>\n")
    }

    b.WriteString("</ul></nav>")
}

func writePosts() {
    posts := getDir("_posts")

    for i := 0; i < len(posts); i++ {
        id, date, title := getPostMeta(posts[i])

        var b bytes.Buffer
        b.WriteString(getLayoutStart(title + " – " + config.Title))
        b.WriteString("<a class=\"back\" href=\"../index.html\">←</a>")
        b.WriteString("<h2><span class=\"date\">" + date + "</span></h2>")
        b.Write(blackfriday.MarkdownCommon(getFile("_posts/" + posts[i].Name())))
        b.WriteString(getLayoutEnd())

        writeFile("posts/"+id, b)
    }
}

func writePostsPage() {
    posts := getDir("_posts")
    var b bytes.Buffer

    b.WriteString(getLayoutStart("All posts – " + config.Title))
    b.WriteString("<a class=\"back\" href=\"index.html\">←</a>")
    b.WriteString("<h2><span>All posts</span></h2>")
    b.WriteString("<nav class=\"posts\"><ul>")

    for i := len(posts) - 1; i >= 0; i-- {
        id, date, title := getPostMeta(posts[i])

        b.WriteString("<li><span class=\"date\">" + date +
            "</span><a href=\"posts/" +
            id + ".html\">" +
            title + "</a></li>\n")
    }

    b.WriteString(getLayoutEnd())
    writeFile("all-posts", b)
}

func writePages() {
    pages := getDir("_pages")

    for i := 0; i < len(pages); i++ {
        fileName, title, skip := getPageMeta(pages[i])
        if skip {
            continue
        }

        var b bytes.Buffer
        b.WriteString(getLayoutStart(title + " – " + config.Title))
        b.WriteString("<p><a href=\"../index.html\">←</a></p>")
        b.Write(blackfriday.MarkdownCommon(getFile("_pages/" + pages[i].Name())))
        b.WriteString(getLayoutEnd())

        writeFile("pages/"+fileName, b)
    }
}

func createConfig() {
    if _, err := os.Stat("_sections/config.md"); os.IsNotExist(err) {
        err := ioutil.WriteFile(
            "_sections/config.md",
            []byte("Title\nDescription\nemail@email.com"), 0644)

        if err != nil {
            panic(err)
        }
    }
}

func createCss() {
    if _, err := os.Stat("css/styles.css"); os.IsNotExist(err) {
        err := ioutil.WriteFile(
            "css/styles.css",
            []byte(`body {
    font-family: 'IBM Plex Sans', sans-serif;
    line-height: 1.875;
    font-weight: 400;
    font-size: 16px;
    background-color: #fdffff;
    color: #0f3763;
}

h1, h2, h3 {
    font-weight: 300;
}

h1 {
    font-size: 26.79296875px;
    margin-top: 24px;
}

h2 {
    font-size: 22.5625px;
    margin-top: 40px;
}

h3 {
    font-size: 19px;
    margin-top: 40px;
}

h1, h2, h3 {
    color: #003880;
}

@media (max-width: 1023.98px) {
    .container {
        margin: 28px auto 48px auto;
    }
}

@media (min-width: 1024px) {
    .container {
        margin: 48px auto;
    }
}

.container {
    max-width: 896px;
    padding: 0 16px;
}

nav ul {
    list-style-type: none;
    padding: 0;
}

nav li {
    margin-bottom: 8px;
}

nav li .date {
    display: inline-block;
    width: 104px;
}

.all-posts {
    font-size: 13.47368421052632px;
    margin: 16px 0 32px 0;
}

a {
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

a {
    color: #006aee;
}

pre {
    overflow: auto;
    padding: 0.25rem 0.75rem;
    margin-bottom: 32px;
}

pre {
    background-color: #f4f7ff;
}

code {
    font-size: 0.875em;
    font-family: 'IBM Plex Mono', monospace;
}

table {
    border-collapse: collapse;
    width: 100%;
    margin-bottom: 32px;
}

@media (max-width: 1023.98px) {
    table {
        font-size: 14px;
    }
}

@media (min-width: 1024px) {
    table {
        font-size: 15px;
    }
}

tr {
    border-bottom: 0.5px solid #bdc5d8;
}

th {
    text-align: left;
    font-weight: 500;
    padding: 12px;
    white-space: nowrap;
}

td {
    padding: 12px;
    white-space: nowrap;
}

.footer {
    text-align: right;
    margin-top: 1px;
    padding-top: 1px;
    color: #aaa;
    font-size: 20px;
    font-weight: 300;
}

a.back {
    margin-right: 12px;
    text-decoration: none;
}`), 0644)

        if err != nil {
            panic(err)
        }
    }
}

func createLayoutStart() {
    if _, err := os.Stat("_sections/start.md"); os.IsNotExist(err) {
        err := ioutil.WriteFile(
            "_sections/start.md",
            []byte(`<!DOCTYPE html>
    <html>
        <head>
            <meta charset="utf-8">
            <meta name="viewport" content="width=device-width, initial-scale=1">
            <link href="https://fonts.googleapis.com/css?family=IBM+Plex+Sans:300,400,400i,600&display=swap&subset=cyrillic-ext" rel="stylesheet">
            <link href="https://fonts.googleapis.com/css?family=IBM+Plex+Mono:400" rel="stylesheet">
            <link type="text/css" rel="stylesheet" href="/css/styles.css" media="all">
            <title>%s</title>
        </head>
        <body>
            <div class="container">
                <div class="header">
                    <h1>%s</h1>
                    <div>%s</div>
                </div>`), 0644)

        if err != nil {
            panic(err)
        }
    }
}

func createLayoutEnd() {
    if _, err := os.Stat("_sections/end.md"); os.IsNotExist(err) {
        err := ioutil.WriteFile(
            "_sections/end.md",
            []byte(`<div class="footer">%s
                    </div>
            </div>
        </body>
    </html>`), 0644)

        if err != nil {
            panic(err)
        }
    }
}

func createFilesAndDirs() {
    os.MkdirAll("_sections", 0755)
    os.MkdirAll("_posts", 0755)
    os.MkdirAll("_pages", 0755)
    os.MkdirAll("css", 0755)

    createConfig()
    createCss()
    createLayoutStart()
    createLayoutEnd()

    if _, err := os.Stat("posts"); os.IsNotExist(err) {
        err := ioutil.WriteFile(
            "_posts/"+time.Now().Format("2006-01-02")+"-initial-post.md",
            []byte("# Initial post\n\nThis is the initial post."),
            0644)

        if err != nil {
            panic(err)
        }
    }

    if _, err := os.Stat("pages"); os.IsNotExist(err) {
        err := ioutil.WriteFile(
            "_pages/about.md",
            []byte("# About\n\nThis is the about page."),
            0644)

        if err != nil {
            panic(err)
        }
    }

    os.MkdirAll("posts", 0755)
    os.MkdirAll("pages", 0755)
}

func main() {
    createFilesAndDirs()
    getConfig()
    writeIndex()
    writePosts()
    writePostsPage()
    writePages()
}
