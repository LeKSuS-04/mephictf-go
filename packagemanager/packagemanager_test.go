package packagemanager

import (
	"context"
	"math/rand/v2"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSimpleInstallation(t *testing.T) {
	t.Parallel()

	repo := Repository{
		PackageDependencies: map[Package][]Package{
			{Name: "A", Version: 1}: {{Name: "B", Version: 1}},
			{Name: "B", Version: 1}: {{Name: "C", Version: 1}},
			{Name: "C", Version: 1}: {},
		},
	}
	order, err := GetInstallationOrder(repo, []Package{{Name: "A", Version: 1}})
	require.NoError(t, err)
	require.Equal(t, []Package{
		{Name: "C", Version: 1},
		{Name: "B", Version: 1},
		{Name: "A", Version: 1},
	}, order)
}

func TestInstallationMultipleDeps(t *testing.T) {
	t.Parallel()

	repo := Repository{
		PackageDependencies: map[Package][]Package{
			{Name: "A", Version: 1}: {{Name: "B", Version: 1}, {Name: "C", Version: 1}},
			{Name: "B", Version: 1}: {{Name: "D", Version: 1}},
			{Name: "C", Version: 1}: {{Name: "B", Version: 1}},
			{Name: "D", Version: 1}: {},
		},
	}

	order, err := GetInstallationOrder(repo, []Package{{Name: "A", Version: 1}})
	require.NoError(t, err)
	require.Equal(t, []Package{
		{Name: "D", Version: 1},
		{Name: "B", Version: 1},
		{Name: "C", Version: 1},
		{Name: "A", Version: 1},
	}, order)
}

func TestInstallationMultiplePackages(t *testing.T) {
	t.Parallel()

	repo := Repository{
		PackageDependencies: map[Package][]Package{
			{Name: "A", Version: 1}: {{Name: "B", Version: 1}},
			{Name: "B", Version: 1}: {{Name: "C", Version: 1}},
			{Name: "C", Version: 1}: {},
			{Name: "D", Version: 1}: {{Name: "C", Version: 1}},
		},
	}

	order, err := GetInstallationOrder(repo, []Package{
		{Name: "A", Version: 1},
		{Name: "D", Version: 1},
	})
	require.NoError(t, err)

	pos := make(map[string]int, len(order))
	for i, p := range order {
		require.Equal(t, 1, p.Version)
		pos[p.Name] = i
	}

	validateOrder(t, repo, order)
}

func TestInstallationCircularDependency(t *testing.T) {
	t.Parallel()

	repo := Repository{
		PackageDependencies: map[Package][]Package{
			{Name: "A", Version: 1}: {{Name: "B", Version: 1}},
			{Name: "B", Version: 1}: {{Name: "C", Version: 1}},
			{Name: "C", Version: 1}: {{Name: "D", Version: 1}},
			{Name: "D", Version: 1}: {{Name: "B", Version: 1}},
		},
	}

	_, err := GetInstallationOrder(repo, []Package{{Name: "A", Version: 1}})
	require.ErrorIs(t, ErrCircularDependency, err)
}

func TestInstallationDependencyNotFound(t *testing.T) {
	t.Parallel()

	repo := Repository{
		PackageDependencies: map[Package][]Package{
			{Name: "A", Version: 1}: {{Name: "B", Version: 1}},
		},
	}

	_, err := GetInstallationOrder(repo, []Package{{Name: "A", Version: 1}})
	require.ErrorIs(t, ErrDependencyNotFound, err)
}

func TestVersionConflict(t *testing.T) {
	t.Parallel()

	repo := Repository{
		PackageDependencies: map[Package][]Package{
			{Name: "A", Version: 1}: {{Name: "C", Version: 1}},
			{Name: "B", Version: 1}: {{Name: "C", Version: 2}},
			{Name: "C", Version: 1}: {},
			{Name: "C", Version: 2}: {},
		},
	}

	_, err := GetInstallationOrder(repo, []Package{
		{Name: "A", Version: 1},
		{Name: "B", Version: 1},
	})
	require.ErrorIs(t, ErrVersionConflict, err)
}

func TestInstallationBigRepository(t *testing.T) {
	t.Parallel()

	packageCount := 100_000
	maxDepsPerPackage := 100
	packages := make([]Package, packageCount)
	deps := make(map[Package][]Package, maxDepsPerPackage)
	for i := range packageCount {
		p := Package{
			Name:    "P" + strconv.Itoa(i),
			Version: rand.Int(),
		}
		packages[i] = p

		dependencyCount := rand.IntN(min(maxDepsPerPackage, i+1))
		dependencies := make(map[Package]struct{}, dependencyCount)
		for range dependencyCount {
			r := rand.IntN(packageCount)
			_, ok := dependencies[packages[r]]
			for ok {
				r = (r + 1) % i
				_, ok = dependencies[packages[r]]
			}
			dependencies[packages[r]] = struct{}{}
		}

		deps[p] = make([]Package, 0, len(dependencies))
		for d := range dependencies {
			deps[p] = append(deps[p], d)
		}
	}
	repo := Repository{
		PackageDependencies: deps,
	}

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var result []Package
	var err error
	resultChan := make(chan struct{})
	go func() {
		result, err = GetInstallationOrder(repo, packages)
		close(resultChan)
	}()

	select {
	case <-resultChan:
		break

	case <-ctx.Done():
		t.Error("execution took too long")
	}

	require.NoError(t, err)
	validateOrder(t, repo, result)
}

func validateOrder(t *testing.T, repo Repository, order []Package) {
	pos := make(map[Package]int, len(order))
	for i, pkg := range order {
		pos[pkg] = i
	}

	for pkg, deps := range repo.PackageDependencies {
		require.Contains(t, pos, pkg)
		for _, dep := range deps {
			require.True(t, pos[dep] < pos[pkg])
		}
	}
}
