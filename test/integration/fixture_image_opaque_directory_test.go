package integration

import (
	"testing"

	"github.com/khulnasoft-labs/goscope/pkg/file"
	"github.com/khulnasoft-labs/goscope/pkg/filetree"
	"github.com/khulnasoft-labs/goscope/pkg/imagetest"
)

func TestImage_SquashedTree_OpaqueDirectoryExistsInFileCatalog(t *testing.T) {
	image := imagetest.GetFixtureImage(t, "docker-archive", "image-opaque-directory")

	tree := image.SquashedTree()
	path := "/usr/lib/jvm"
	_, ref, err := tree.File(file.Path(path), filetree.FollowBasenameLinks)
	if err != nil {
		t.Fatalf("unable to get file=%q : %+v", path, err)
	}

	_, err = image.FileCatalog.Get(*ref.Reference)
	if err != nil {
		t.Fatal(err)
	}
}
