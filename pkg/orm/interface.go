package orm


type Repository interface {
}

type getRepository(entity) Repository

type Orm interface {
	// getRepository() — Returns the repository for a given entity
	GetRepository(entity) Repository
	// getRepositories() — Returns all registered repositories
	GetRepositories() []Repository
	// registerRepository() — Registers a repository for a given entity
	RegisterRepository(entity, Repository)
	// registerRepositories() — Registers repositories for a given entity
	RegisterRepositories(entity, ...Repository)
	// unregisterRepository() — Unregisters a repository for a given entity
	UnregisterRepository(entity)
	// unregisterRepositories() — Unregisters repositories for a given entity
	UnregisterRepositories(entity)
	// hasRepository() — Determines whether a repository is registered for a given entity
	HasRepository(entity) bool
	// hasRepositories() — Determines whether repositories are registered for a given entity
}