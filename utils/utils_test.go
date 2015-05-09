package utils
import "testing"

func TestSearchBrowserLauncher(t *testing.T) {
    browser := searchBrowserLauncher("darwin")
    if browser != "open" {
        t.Error("On darwin browser is 'open'")
    }

    browser = searchBrowserLauncher("windows")
    if browser != "cmd /c start" {
        t.Error("On windows browser is cmd /c start")
    }
}