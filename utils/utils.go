package utils
import (
    "os/exec"
    "os"
    "runtime"
    "errors"
    "strings"
)

func BrowserLauncher() ([]string, error) {
    browser := os.Getenv("BROWSER")
    if browser == "" {
        browser = searchBrowserLauncher(runtime.GOOS)
    }

    if browser == "" {
        return nil, errors.New("Please set $BROWSER to a web launcher")
    }

    return strings.Split(browser, " "), nil
}

func searchBrowserLauncher(goos string) string {
    var browser string
    switch goos {
        case "darwin":
            browser = "open"
        case "windows":
            browser = "cmd /c start"
        default:
            candidates := []string{"xdg-open", "cygstart", "x-www-browser", "firefox", "opera", "mozilla", "netscape"}
            for _, b := range candidates {
                path, err := exec.LookPath(b)
                if err == nil {
                    browser = path
                    break
                }
            }
    }
    return browser
}