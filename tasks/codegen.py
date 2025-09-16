import os

from invoke.tasks import task
from invoke.context import Context

root_dir = os.path.dirname(os.path.abspath(__file__))

# where the proto code is first generated as we declare packages with v5
v5_dir=os.path.join(root_dir, "v5")

# where all toolchains (protobuf compilers, protobuf plugins,etc) are stored
toolchain_dir=os.path.join(root_dir, "toolchains")

# binaries
toolchain_bin_dir=os.path.join(toolchain_dir, "bin")

# include path for the legacy and main protobuf compilers respectively
toolchain_include_dir=os.path.join(toolchain_dir, "include", "proto")

protoc_version="21.12"
protoc_binary=os.path.join(toolchain_bin_dir, "protoc" + protoc_version)

gogo_tag = "v1.3.2"
gogo_dir=os.path.join(toolchain_dir,  "gogo")
gogo_include = f"{toolchain_dir}/gogo/src:{toolchain_dir}/gogo/src/github.com/gogo/protobuf/protobuf"
gogo_bin = os.path.join(toolchain_bin_dir, f"gogo-bin-{gogo_tag}")
protoc_jsonschema_version="73d5723"

### toolchains is meant to store a cache of all binary dependencies needed to build the agent-payload.
### this rakefile will download those dependencies on the fly if needed.
### the structure of toolchains is as follows:
### /toolchains
###     /toolchains/bin -- contains any binaries used for building
###         /toolchains/bin/protoc21.12 -- the new protobuf compiler
###         /toolchains/bin/protoc-gen-go -- the go code generator for protoc
###         /toolchains/bin/protoc-gen-XXX -- other protobuf generators (vtproto, java, etc.)
###     /toolchains/include -- contains any protobuf library files for common types (like proto.Duration)
###         /toolchains/includes/proto -- protobuf libraries bundled with the new protobuf compiler
### =>  /toolchains/gogo -- temp directory used to build the gogo_faster generator

@task
def clean(ctx: Context):
    ctx.run(f"rm -rf {gogo_dir}")
    if os.path.exists(toolchain_dir):
        ctx.run(f"chmod u+w -R {toolchain_dir}")
    ctx.run(f"rm -rf {toolchain_dir}")


@task(pre=["install_protoc_all"])
def setup_gogo(ctx: Context):
    og_dir = os.getcwd()
    if os.path.exists(gogo_dir):
        with open("#{gogo_dir}/src/github.com/gogo"):
            ctx.run(f"git clone https://github.com/gogo/protobuf.git {gogo_dir}/src/github.com/gogo/protobuf")
    else:
        print(f"gogo already cloned into {gogo_dir}")

    os.chdir(f"{gogo_dir}/src/github.com/gogo/protobuf")
    ctx.run(f"git checkout {gogo_tag}")
    ctx.run(f"PATH=$PATH:/tmp GOBIN={toolchain_bin_dir} GOPATH={gogo_dir} make clean install")
    os.chdir(og_dir)

@task
def install_protoc(ctx: Context):
    if ctx.run(f'bash -c "{protoc_binary} --version"').result != f"libprotoc {protoc_version}":
        ctx.run(f"""
        /bin/bash <<BASH

        set -euo pipefail
        if [ ! -f {protoc_binary} ] ; then
            echo "Downloading protoc {protoc_version}"
            cd /tmp
            if [ "$(uname -s)" = "Darwin" ] ; then
                curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v{protoc_version}/protoc-{protoc_version}-osx-x86_64.zip
            else
                curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v{protoc_version}/protoc-{protoc_version}-linux-x86_64.zip
            fi

            mkdir -p {toolchain_bin_dir}
            rm -rf include/ # make sure there is no stale include
            unzip -o protoc-{protoc_version}*.zip
            mv bin/protoc {protoc_binary}

            mkdir -p {toolchain_include_dir}
            rm -rf {toolchain_include_dir}/**
            mv include/* {toolchain_include_dir}
        fi
BASH
        """)

@task(pre=["install_protoc"])
def install_protoc_all(ctx: Context):
    pass

@task(pre=["install_protoc", "setup_gogo"])
def protoc(ctx: Context):
    ctx.run(f"""
        /bin/bash <<BASH
        set -euo pipefail

        export GO111MODULE=auto

# These .proto can be imported by other repositories (not only generated code). To have a consistent import path, we use the GOPATH and we add a /v5/ folder as declared packaged in `go.mod` is ../v5
        mkdir -p {v5_dir}
        ln -sf {root_dir}/proto {v5_dir}/proto

        echo "Generating logs proto"
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path=$GOPATH/src:{gogo_dir}/src:. --gogofast_out=$GOPATH/src proto/logs/agent_logs_payload.proto

        echo "Generating metrics proto (go)"
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path=$GOPATH/src:{gogo_include}:{toolchain_include_dir}:. --gogofast_out=$GOPATH/src proto/metrics/agent_payload.proto
        echo "done"

        echo "Generating process proto (go)"
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path={toolchain_include_dir}:{gogo_include}:. --gogofaster_out=$GOPATH/src proto/process/*.proto

        GOPATH={toolchain_dir} go install github.com/DataDog/protoc-gen-gostreamer@v0.2.0
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path=$GOPATH/src:{gogo_dir}/src:.  --gostreamer_out=$GOPATH/src proto/process/*.proto
        mv v5/process/proto/process/*.go process

# Install protoc-gen-go
        GOPATH={toolchain_dir} go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
        GOPATH={toolchain_dir} go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema@{protoc_jsonschema_version}

        echo "Generating contlcycle proto"
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path={toolchain_include_dir}:. --go_out=$GOPATH/src proto/contlcycle/contlcycle.proto

        echo "Generating kubernetes autoscaling proto"
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path={toolchain_include_dir}:proto/deps:$GOPATH/src --go_out=$GOPATH/src --jsonschema_out=type_names_with_no_package:jsonschema $GOPATH/src/github.com/DataDog/agent-payload/v5/proto/autoscaling/kubernetes/*.proto

        echo "Generating contimage proto"
        PATH={toolchain_bin_dir}  {protoc_binary} --proto_path={toolchain_include_dir}:. --go_out=$GOPATH/src proto/contimage/contimage.proto

        echo "Generating sbom proto"
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path={toolchain_include_dir}:. --go_out=$GOPATH/src proto/deps/github.com/CycloneDX/specification/schema/bom-1.4.proto
        PATH={toolchain_bin_dir}  {protoc_binary} --proto_path={toolchain_include_dir}:. --go_out=$GOPATH/src proto/sbom/sbom.proto

# Install protoc-gen-go-vtproto
        GOPATH={toolchain_dir} go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.5.0

        echo "Generating CWS Activity Dumps v1"
        PATH={toolchain_bin_dir} {protoc_binary} --proto_path=$GOPATH/src:. \
        --java_out=java \
        --go_out=$GOPATH/src \
        --go-vtproto_out=$GOPATH/src \
        --go-vtproto_opt=features=pool+marshal+unmarshal+size \
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.SecDump \
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.ProcessActivityNode \
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.FileActivityNode \
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.FileInfo \
        --go-vtproto_opt=pool=github.com/DataDog/agent-payload/v5/cws/dumpsv1.ProcessInfo \
        proto/cws/dumpsv1/activity_dump.proto

        rm -f {v5_dir}/proto
        cp -r v5/* .
        rm -rf v5
BASH
    """)

@task(pre=["protoc"])
def all(ctx: Context):
    pass
