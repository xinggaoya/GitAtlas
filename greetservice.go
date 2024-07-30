package main

import (
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log/slog"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

type GreetService struct{}

func (g *GreetService) InitWordRepo() string {
	dialog := application.OpenFileDialog()
	dialog.SetTitle("选择Git文件夹")
	dialog.CanChooseDirectories(true)
	dialog.CanChooseFiles(false)

	selection, err := dialog.PromptForSingleSelection()
	if err != nil {
		slog.Error("failed to prompt for selection", "error", err)
		return ""
	}
	return selection
}

// InitGitRepository 初始化仓库
func (g *GreetService) InitGitRepository(path string, remoteURL string) {
	if path == "" || remoteURL == "" {
		slog.Error("path or remoteURL is empty")
		return
	}
	err := InitRepo(path)
	if err != nil {
		slog.Error("failed to init repo", "error", err)
		return
	}
	err = SetRemote(path, "origin", remoteURL)
	if err != nil {
		slog.Error("failed to set remote", "error", err)
		return
	}
	slog.Info("init git repository successfully")
}

// GetRemoteURL 获取远程仓库地址
func (g *GreetService) GetRemoteURL(path string) []string {
	if path == "" {
		slog.Error("path is empty")
		return nil
	}
	remoteURLs, err := GetRemoteURL(path)
	if err != nil {
		slog.Error("failed to get remote url", "error", err)
		return nil
	}
	return remoteURLs
}

// GetUpdateFiles 获取变更文件列表
func (g *GreetService) GetUpdateFiles(path string) UpdateFiles {
	if path == "" {
		slog.Error("path is empty")
		return UpdateFiles{}
	}
	files, err := GetUpdateFiles(path)
	if err != nil {
		slog.Error("failed to get update files", "error", err)
		return UpdateFiles{}
	}
	return files
}

// PullChanges 拉取最新代码
func (g *GreetService) PullChanges(path string, remote string, branch string) {
	if path == "" || remote == "" || branch == "" {
		slog.Error("path or remote or branch is empty")
		return
	}
	err := PullChanges(path, remote, branch)
	if err != nil {
		slog.Error("failed to pull changes", "error", err)
		return
	}
	slog.Info("pull changes successfully")
}

// CommitChanges 提交更改
func (g *GreetService) CommitChanges(path string, message string) {
	if path == "" || message == "" {
		slog.Error("path or message is empty")
		return
	}
	err := CommitChanges(path, message)
	if err != nil {
		slog.Error("failed to commit changes", "error", err)
		return
	}
	slog.Info("commit changes successfully")
}

// PushChanges 推送提交
func (g *GreetService) PushChanges(path string, remote string, branch string) {
	if path == "" || remote == "" || branch == "" {
		slog.Error("path or remote or branch is empty")
		return
	}
	err := PushChanges(path, remote, branch)
	if err != nil {
		slog.Error("failed to push changes", "error", err)
		return
	}
	slog.Info("push changes successfully")
}

// InitRepo 初始化Git仓库
func InitRepo(path string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to init repo: %v, output: %s", err, string(output))
	}
	fmt.Println(string(output))
	return nil
}

// 设置远程仓库地址
func SetRemote(path, remoteName, remoteURL string) error {
	cmd := exec.Command("git", "remote", "add", remoteName, remoteURL)
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set remote: %v, output: %s", err, string(output))
	}
	fmt.Println(string(output))
	return nil
}

// 提交更改
func CommitChanges(path, message string) error {
	cmd := exec.Command("git", "add", "--all")
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add all files: %v, output: %s", err, string(output))
	}
	cmd = exec.Command("git", "commit", "-m", message)
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to commit changes: %v, output: %s", err, string(output))
	}
	fmt.Println(string(output))
	return nil
}

// 推送到远程仓库
func PushChanges(path, remote, branch string) error {
	cmd := exec.Command("git", "push", remote, branch)
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to push changes: %v, output: %s", err, string(output))
	}
	fmt.Println(string(output))
	return nil
}

// 从远程仓库拉取最新代码
func PullChanges(path, remote, branch string) error {
	cmd := exec.Command("git", "pull", remote, branch)
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to pull changes: %v, output: %s", err, string(output))
	}
	fmt.Println(string(output))
	return nil
}

// 获取git远程地址
func GetRemoteURL(path string) ([]string, error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to get remote url: %v, output: %s", err)
	}
	// 解析每一行
	var remoteURLs []string
	urlRegex := regexp.MustCompile(`https?://[^\s"]+`)
	urls := urlRegex.FindAllString(string(output), -1)

	for _, rawURL := range urls {
		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			continue
		}
		// 排除重复
		bol := false
		for _, url := range remoteURLs {
			if url == parsedURL.String() {
				bol = true
				break
			}
		}
		if !bol {
			remoteURLs = append(remoteURLs, parsedURL.String())
		}
	}
	return remoteURLs, nil
}

type UpdateFiles struct {
	AddFiles    []string `json:"addFiles"`
	DeleteFiles []string `json:"deleteFiles"`
	ModifyFiles []string `json:"modifyFiles"`
}

// 获取变更文件列表
func GetUpdateFiles(path string) (UpdateFiles, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return UpdateFiles{}, fmt.Errorf("failed to get update files: %v, output: %s", err)
	}
	var updateFiles UpdateFiles
	for _, line := range strings.Split(string(output), "\n") {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, " ")
		if split[0] == "??" {
			updateFiles.AddFiles = append(updateFiles.AddFiles, split[1])
		}
		switch split[1] {
		case "D":
			updateFiles.DeleteFiles = append(updateFiles.DeleteFiles, split[2])
		case "M":
			updateFiles.ModifyFiles = append(updateFiles.ModifyFiles, split[2])
		}
	}
	return updateFiles, nil
}
