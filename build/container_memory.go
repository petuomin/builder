package build

type memoryContainer struct {
	builds map[string]Build
}

func (c memoryContainer) Init(purge bool) error {
	return nil
}

func (c memoryContainer) Close() error {
	return nil
}

func (c memoryContainer) Builds() []string {
	builds := []string{}
	for ID := range c.builds {
		builds = append(builds, ID)
	}
	return builds
}

func (c memoryContainer) Build(ID string) (Build, error) {
	build, ok := c.builds[ID]
	if !ok {
		return nil, ErrNotFound
	}
	return build, nil
}

func (c memoryContainer) New(b Buildable) (Build, error) {
	build, err := New(b)
	if err != nil {
		return nil, err
	}
	c.builds[build.ID()] = build
	return build, nil
}

func (c memoryContainer) AddStage(buildID string, stage Stage) error {
	b, err := c.Build(buildID)
	if err != nil {
		return err
	}
	return b.AddStage(stage)
}

func (c memoryContainer) Output(buildID string, output []byte) error {
	b, err := c.Build(buildID)
	if err != nil {
		return err
	}
	return b.Output(output)
}
