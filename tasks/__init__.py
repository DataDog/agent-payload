from invoke.collection import Collection

from tasks import (codegen, go)

ns = Collection()

ns.add_collection(codegen, "codegen")
ns.add_collection(go)
