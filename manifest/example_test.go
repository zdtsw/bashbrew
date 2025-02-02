package manifest_test

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/docker-library/bashbrew/manifest"
)

func Example() {
	man, err := manifest.Parse(bufio.NewReader(strings.NewReader(`# RFC 2822

	# I LOVE CAKE

Maintainers: InfoSiftr <github@infosiftr.com> (@infosiftr),
             Johan Euphrosine <proppy@google.com> (@proppy)
GitFetch: refs/heads/master
GitRepo: https://github.com/docker-library/golang.git
SharedTags: latest
arm64v8-GitRepo: https://github.com/docker-library/golang.git
Architectures: amd64, amd64


 # hi


 	 # blasphemer


# Go 1.6
Tags: 1.6.1, 1.6, 1
arm64v8-GitRepo: https://github.com/docker-library/golang.git
Directory: 1.6
GitCommit: 0ce80411b9f41e9c3a21fc0a1bffba6ae761825a
Constraints: some-random-build-server


# Go 1.5
Tags: 1.5.3
GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
SharedTags: 1.5.3-debian, 1.5-debian
Directory: 1.5
s390x-GitCommit: b6c460e7cd79b595267870a98013ec3078b490df
i386-GitFetch: refs/heads/i386
ppc64le-Directory: 1.5/ppc64le/


Tags: 1.5
SharedTags: 1.5-debian
GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
Directory: 1.5
s390x-GitCommit: b6c460e7cd79b595267870a98013ec3078b490df
i386-GitFetch: refs/heads/i386
ppc64le-Directory: 1.5/ppc64le

Tags: 1.5-alpine
GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
Directory: 1.5
File: Dockerfile.alpine
s390x-File: Dockerfile.alpine.s390x.bad-boy

SharedTags: raspbian
GitCommit: deadbeefdeadbeefdeadbeefdeadbeefdeadbeef
Tags: raspbian-s390x
Architectures: s390x, i386
File: Dockerfile-raspbian
s390x-File: Dockerfile


`)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("-------------\n2822:\n%s\n", man)

	fmt.Printf("\nShared Tag Groups:\n")
	for _, group := range man.GetSharedTagGroups() {
		fmt.Printf("\n  - %s\n", strings.Join(group.SharedTags, ", "))
		for _, entry := range group.Entries {
			fmt.Printf("    - %s\n", entry.TagsString())
		}
	}
	fmt.Printf("\n")

	// Output:
	// -------------
	// 2822:
	// Maintainers: InfoSiftr <github@infosiftr.com> (@infosiftr), Johan Euphrosine <proppy@google.com> (@proppy)
	// SharedTags: latest
	// GitRepo: https://github.com/docker-library/golang.git
	// arm64v8-GitRepo: https://github.com/docker-library/golang.git
	//
	// Tags: 1.6.1, 1.6, 1
	// GitCommit: 0ce80411b9f41e9c3a21fc0a1bffba6ae761825a
	// Directory: 1.6
	// Constraints: some-random-build-server
	//
	// Tags: 1.5.3, 1.5
	// SharedTags: 1.5.3-debian, 1.5-debian
	// GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
	// Directory: 1.5
	// i386-GitFetch: refs/heads/i386
	// ppc64le-Directory: 1.5/ppc64le
	// s390x-GitCommit: b6c460e7cd79b595267870a98013ec3078b490df
	//
	// Tags: 1.5-alpine
	// GitCommit: d7e2a8d90a9b8f5dfd5bcd428e0c33b68c40cc19
	// Directory: 1.5
	// File: Dockerfile.alpine
	// s390x-File: Dockerfile.alpine.s390x.bad-boy
	//
	// Tags: raspbian-s390x
	// SharedTags: raspbian
	// Architectures: i386, s390x
	// GitCommit: deadbeefdeadbeefdeadbeefdeadbeefdeadbeef
	// File: Dockerfile-raspbian
	// s390x-File: Dockerfile
	//
	// Shared Tag Groups:
	//
	//   - latest
	//     - 1.6.1, 1.6, 1
	//     - 1.5-alpine
	//
	//   - 1.5.3-debian, 1.5-debian
	//     - 1.5.3, 1.5
	//
	//   - raspbian
	//     - raspbian-s390x
}

func ExampleFetch_local() {
	repoName, tagName, man, err := manifest.Fetch("testdata", "bash:4.4")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s:%s\n\n", repoName, tagName)

	fmt.Println(man.GetTag(tagName).ClearDefaults(manifest.DefaultManifestEntry).String())

	// Output:
	// bash:4.4
	//
	// Maintainers: Tianon Gravi <admwiggin@gmail.com> (@tianon)
	// Tags: 4.4.12, 4.4, 4, latest
	// GitRepo: https://github.com/tianon/docker-bash.git
	// GitCommit: 1cbb5cf49b4c53bd5a986abf7a1afeb9a80eac1e
	// Directory: 4.4
}

func ExampleFetch_remote() {
	repoName, tagName, man, err := manifest.Fetch("/home/jsmith/docker/official-images/library", "https://github.com/docker-library/official-images/raw/1a3c4cd6d5cd53bd538a6f56a69f94c5b35325a7/library/bash:4.4")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s:%s\n\n", repoName, tagName)

	fmt.Println(man.GetTag(tagName).ClearDefaults(manifest.DefaultManifestEntry).String())

	// Output:
	// bash:4.4
	//
	// Maintainers: Tianon Gravi <admwiggin@gmail.com> (@tianon)
	// Tags: 4.4.12, 4.4, 4, latest
	// GitRepo: https://github.com/tianon/docker-bash.git
	// GitCommit: 1cbb5cf49b4c53bd5a986abf7a1afeb9a80eac1e
	// Directory: 4.4
}
