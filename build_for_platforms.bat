@echo off

cd /d %~dp0
set DIST=%~dp0dist

net.exe session 1>NUL 2>NUL && (
    call:main
) || (
    call:not_admin
)
goto :eof

:not_admin
%1 mshta vbscript:CreateObject("Shell.Application").ShellExecute("cmd.exe","/k %~s0 ::","","runas",1)(window.close)
:pause
goto :eof


:main
call:build
call:test
goto:eof

:build

@echo build start
rmdir %DIST% /S /Q

set CGO_ENABLED=
set GOOS=
set GOARCH=
go build -o %DIST%\default\go_cli_test.exe main.go

set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go  build -o %DIST%\windows\go_cli_test.exe main.go

set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go  build -o %DIST%\mac\go_cli_test main.go

@echo build finish

goto:eof

:test
@echo test start
call:clearupenv

%DIST%\default\go_cli_test.exe
@echo test finish
goto:eof


:clearupenv
choco uninstall python python3 --yes
rmdir /s /q %ChocolateyInstall%

reg delete "HKLM\system\controlset001\control\session manager\environment" /v ChocolateyInstall /f
reg delete "HKLM\system\controlset001\control\session manager\environment" /v ChocolateyToolsLocation /f
reg delete "HKLM\system\controlset001\control\session manager\environment" /v ChocolateyLastPathUpdate /f

refreshenv
goto:eof