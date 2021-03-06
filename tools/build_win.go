package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"fmt"
)

const signtool = "C:\\Program Files (x86)\\Windows Kits\\10\\bin\\10.0.18362.0\\x64\\signtool.exe"

func main() {
	cmd := exec.Command("git", "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Remove(filepath.Join("openvpn_win",
		"openvpn-install-2.4.6-I602.exe"))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(filepath.Join("openvpn_win",
		"openvpn-install-2.4.6-I602.exe.asc"))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Remove(filepath.Join("build", "win", "FVPN.exe"))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.RemoveAll(filepath.Join("build", "win"))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	err = os.Chdir("tuntap_win")
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "build", "-v", "-o", "tuntap.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Join("..", "service"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello, world\n")
	cmd = exec.Command("go", "get", "-u", "-f")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "build", "-v", "-ldflags", "-H windowsgui")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Join("..", "client"))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("npm", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(".\\node_modules\\.bin\\electron-rebuild")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(
		".\\node_modules\\.bin\\electron-packager",
		".\\",
		"fvpn",
		"--platform=win32",
		"--arch=x64",
		"--icon=www\\img\\logo.ico",
		"--out=..\\build\\win",
		"--prune",
		"--asar",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Join("..", "build", "win",
		"fvpn-win32-x64"))
	if err != nil {
		panic(err)
	}



	err = os.Chdir(filepath.Join("..", "..", "..",
		"resources_win", "post_install"))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "build", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}


	err = os.Chdir(filepath.Join("..", "pre_uninstall"))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "build", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Join("..", "..", "service_win"))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command(signtool,
		"sign",
		"/a",
		"/n", "FVPN",
		"/tr", "http://timestamp.digicert.com",
		"/td", "sha256",
		"/fd", "sha256",
		"/d", "FVPN",
		"nssm.exe",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(filepath.Join("..", "resources_win"))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("C:\\Program Files (x86)\\Inno Setup 6\\ISCC.exe",
		"setup.iss")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
