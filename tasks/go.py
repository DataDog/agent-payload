from invoke.context import Context
from invoke.exceptions import Exit
from invoke.tasks import task

from codegen import clean

@task(pre=[clean])
def test(ctx: Context):
    cmd = "go list ./... | grep -v vendor | xargs go test -v "
    ctx.run(cmd)

@task
def fuzz(ctx: Context):
    cmd = "go list ./... | grep -v vendor | xargs -n 1 go test -v -fuzztime=20s -fuzz '.*' "
    ctx.run(cmd)

@task
def deps(ctx: Context):
    ctx.run("go mod tidy")

@task
def gimme(ctx: Context):
    go_version = "1.21.9"

    if ctx.run("which gimme"):
        ctx.run(f"gimme {go_version}")
        results = ctx.run(f". ~/.gimme/envs/go{go_version}.env > /dev/null; echo $GOOS; echo $GOARCH; echo $GOROOT; echo $PATH").stdout.split("\n")
        env = {
            "GOOS": results[0],
            "GOARCH": results[1],
            "GOROOT": results[2],
            "PATH": results[3],
        }
        print(f"The code generated with versions other than the pinned version {go_version} might cause conflict")
    elif ctx.run("which go"):
        print("using local go toolchain, install gimme to use pinned go version")
    else:
        raise Exit("You need either gimme or go in your path.")
